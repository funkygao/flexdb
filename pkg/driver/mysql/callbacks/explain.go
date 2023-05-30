package callbacks

import (
	"database/sql"
	"strconv"

	"github.com/agile-app/flexdb/internal/profile"
	"github.com/funkygao/go-metrics"
	"github.com/funkygao/log4go"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

// https://gorm.io/docs/write_plugins.html

var (
	scanRowsMetric = metrics.NewRegisteredHistogram("db.scan.rows", nil, metrics.NewExpDecaySample(1028, 0.015))
	warnCounter    = metrics.NewRegisteredCounter("db.warn", nil)
)

type explainResultLine struct {
	table  string
	typ    string
	key    string
	keyLen string
	ref    string
	rows   int
	extra  string
}

// ExplainQuery is a gorm Callback that explain query before query executed.
func ExplainQuery(db *gorm.DB) {
	if !profile.Debug() {
		return
	}

	// as ExplainQuery is called before query, db.Statement.SQL is empty
	// we need to build query SQL
	// if ExplainQuery called after query, db.Statement.SQL will be ready
	if db.Statement.SQL.String() == "" {
		callbacks.BuildQuerySQL(db)
	}

	// TODO conditional explain

	dbx, err := db.DB()
	if err != nil {
		log4go.Error(err)
		return
	}

	explainSQL := "EXPLAIN " + db.Statement.SQL.String()
	rows, err := dbx.Query(explainSQL, db.Statement.Vars...)
	if err != nil {
		// ignored
		log4go.Error("%s %v", explainSQL, err)
		return
	}

	// parse explain result
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	args := make([]interface{}, len(values))
	for i := range values {
		args[i] = &values[i]
	}

	//  id | select_type | table | partitions | type | possible_keys | key | key_len | ref | rows | filtered | Extra |
	//  id | select_type | table |              type | possible_keys | key | key_len | ref | rows |            Extra |
	var (
		exp_table, exp_type, exp_key, exp_keyLen, exp_ref, exp_extra string
		explainRows                                                  int
	)
	for rows.Next() {
		if err = rows.Scan(args...); err != nil {
			log4go.Error("%s %v", explainSQL, err)
			rows.Close()
			return
		}

		exp_table = string(values[2])
		switch len(values) {
		case 10:
			exp_type = string(values[3])
			exp_key = string(values[5])
			exp_keyLen = string(values[6])
			exp_ref = string(values[7])
			explainRows, _ = strconv.Atoi(string(values[8]))
			exp_extra = string(values[9])
		case 12:
			exp_type = string(values[4])
			exp_key = string(values[6])
			exp_keyLen = string(values[7])
			exp_ref = string(values[8])
			explainRows, _ = strconv.Atoi(string(values[9]))
			exp_extra = string(values[11])
		default:
			log4go.Warn("Unknown explain fmt: %s", explainSQL)
			rows.Close()
			return
		}

		expl := explainResultLine{
			table:  exp_table,
			typ:    exp_type,
			key:    exp_key,
			keyLen: exp_keyLen,
			ref:    exp_ref,
			rows:   explainRows,
			extra:  exp_extra,
		}

		scanRowsMetric.Update(int64(expl.rows))

		if expl.rows > profile.P.WarnScannedRowsThreshold ||
			expl.key == "" || // not using index
			(expl.extra != "" && expl.extra != "Using where") || // Using where is fine; bad is: Using filesort,Using temporary
			expl.typ == "ALL" { // full table scan
			warnCounter.Inc(1)
			log4go.Warn("%s, %v", explainSQL, expl)
		}
	}
	rows.Close()
}
