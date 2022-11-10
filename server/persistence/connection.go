package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	Name string
	Host string
	User string
	Pwd  string
	Port string
}

type Connection struct {
	conn *pgx.Conn
}

func NewConnection(config *DbConfig) (*Connection, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Pwd, config.Host, config.Port, config.Name)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://./sqlModel", "postgres", driver)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return &Connection{conn}, nil
}
