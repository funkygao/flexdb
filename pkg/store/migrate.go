package store

import (
	"database/sql"
	"fmt"
	"strings"
)

// MigrateDB reset database.
func MigrateDB(db *sql.DB) {
	for _, schema := range []string{"template/meta.sql", "template/data.sql"} {
		b, _ := Asset(schema)
		for _, s := range strings.Split(string(b), ";") {
			if strings.TrimSpace(s) == "" {
				continue
			}

			if _, err := db.Exec(s); err != nil {
				fmt.Printf("Warn: %v, %s\n", err, s)
			}
		}
	}

}
