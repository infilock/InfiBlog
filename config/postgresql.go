package config

import (
	"context"
	"database/sql"
	"fmt"
)

func (c *Psql) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Pass, c.Name)
}

func (c *Psql) URI() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.Driver, c.User, c.Pass, c.Host, c.Port, c.Name)
}

func GetDBConfig() (*Psql, error) {
	return &Psql{
		Driver: "postgres",
		Host:   "127.0.0.1",
		Port:   "5432",
		User:   "infiblog",
		Pass:   "infiblogpw",
		Name:   "infiblogdb",
	}, nil
}

// ConnectionToPSQL get information config postgresql from vault secret dynamic.
func ConnectionToPSQL(cfg *Psql) *sql.DB {
	db, err := sql.Open(cfg.Driver, cfg.DSN())
	if err != nil {
		panic(err)
	}

	if errPing := db.PingContext(context.Background()); errPing != nil {
		panic(errPing)
	}

	return db
}
