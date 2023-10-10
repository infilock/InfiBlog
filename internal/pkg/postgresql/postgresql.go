package postgresql

import (
	"context"
	"github.com/nasermirzaei89/env"

	//cfgVault "crawler/internal/pkg/vault"
	"database/sql"
	"fmt"
)

func GetPostgresqlConfig() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.MustGetString("PGSQL_HOST"),
		env.MustGetString("PGSQL_PORT"),
		env.MustGetString("PGSQL_USERNAME"),
		env.MustGetString("PGSQL_PASSWORD"),
		env.MustGetString("PGSQL_DB_NAME"),
	)

	//Open opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	//db.SetConnMaxIdleTime()
	//db.SetConnMaxLifetime()
	//db.SetMaxIdleConns()
	//db.SetMaxOpenConns()

	//Ping verifies a connection to the database is still alive, establishing a connection if necessary.
	err = db.PingContext(context.Background())
	if err != nil {
		panic(err)
	}

	return db
}
