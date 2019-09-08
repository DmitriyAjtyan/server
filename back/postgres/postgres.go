package postgres

import (
	"database/sql"
	"fmt"
	"server/back/config"

	"github.com/pkg/errors"
)

// DB is a variable contain connection data for postgresql
var DB *sql.DB

// Connect is a function for connecting to postgresql database
func Connect() (err error) {
	DB, err := sql.Open("postgres", fmt.Sprintf(config.ConfigData.Postgres.Host,
		config.ConfigData.Postgres.User,
		config.ConfigData.Postgres.DBName,
		config.ConfigData.Postgres.Password,
		config.ConfigData.Postgres.Port,
		"disable"))
	if err != nil {
		return errors.Wrap(err, "sql open connect")
	}

	if err = DB.Ping(); err != nil {
		return errors.Wrap(err, "db ping")
	}
	return nil
}
