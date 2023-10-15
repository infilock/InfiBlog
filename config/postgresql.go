package config

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (c *Psql) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Pass, c.Name)
}

func (c *Psql) URL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.Driver, c.User, c.Pass, c.Host, c.Port, c.Name)
}

func GetDBConfig() (*Psql, error) {
	return &Psql{
		Driver:       "postgres",
		MigrationDir: PsqlMigrationDir.Get(),
		Host:         PsqlHost.Get(),
		Port:         PsqlPort.Get(),
		User:         PsqlUser.Get(),
		Pass:         PsqlPass.Get(),
		Name:         PsqlDB.Get(),
	}, nil
}
