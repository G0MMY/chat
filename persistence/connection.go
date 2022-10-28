package persistence

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type connection struct {
	conn *pgx.Conn
}

func NewConnection(dbName, dbHost, dbUser, dbPwd string, dbPort int) (*connection, error) {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", dbUser, dbPwd, dbHost, dbPort, dbName)

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &connection{conn}, nil
}
