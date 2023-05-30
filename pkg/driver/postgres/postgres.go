package postgres

import (
	"github.com/agile-app/flexdb/internal/profile"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Offer() (metaSource *gorm.DB, dataSources []*gorm.DB, err error) {
	metaSource, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  profile.P.MetaDSN, // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,              // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return
}
