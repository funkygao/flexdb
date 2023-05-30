package mysql

import (
	"time"

	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/pkg/driver/mysql/callbacks"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

// Offer offers the singleton meta DB and all data DBs defined in the configuration.
func Offer() (metaSource *gorm.DB, dataSources []*gorm.DB, err error) {
	config := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger: newLogger(ormlog.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      ormlog.Info,
			Colorful:      false,
		}),
	}

	// meta
	metaSource, err = gorm.Open(mysql.Open(profile.P.MetaDSN), config)
	if err != nil {
		return
	}
	// setup the meta conn pool
	dbx, err := metaSource.DB()
	if err != nil {
		return nil, nil, err
	}
	dbx.SetMaxIdleConns(5)
	dbx.SetMaxOpenConns(100)

	if profile.Debug() {
		metaSource = metaSource.Debug()
	}

	// intercept all queries of meta
	metaSource.Callback().Query().
		Before("gorm:query").
		Register("flexdb:explain_query", callbacks.ExplainQuery)

	// data
	dataDSNs := profile.P.DataDSNs
	dataSources = make([]*gorm.DB, len(dataDSNs))
	for i, dsn := range dataDSNs {
		dataSources[i], err = gorm.Open(mysql.Open(dsn), config)
		if err != nil {
			return nil, nil, err
		}

		// setup conn pool
		dbx, err := dataSources[i].DB()
		if err != nil {
			return nil, nil, err
		}
		dbx.SetMaxIdleConns(5)
		dbx.SetMaxOpenConns(100)

		// intercept all queries
		dataSources[i].Callback().Query().
			Before("gorm:query").
			Register("flexdb:explain_query", callbacks.ExplainQuery)

		if profile.Debug() {
			dataSources[i] = dataSources[i].Debug()
		}
	}

	return
}
