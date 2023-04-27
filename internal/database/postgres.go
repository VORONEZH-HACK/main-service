package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgreSQl struct {
	Requests map[string]string
	Options  string
}

func (db *PostgreSQl) Prepare(pathToRequestFiles string, request string) error {
	dat, err := os.ReadFile(pathToRequestFiles + request + ".sql")
	if err != nil {
		return err
	}
	db.Requests[request] = string(dat)

	return nil
}

func (db *PostgreSQl) Get(request string) string {
	return db.Requests[request]
}

func (db *PostgreSQl) Open() (*sql.DB, error) {
	dbconn, err := sql.Open("postgres", db.Options)
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}

var PostgresQLDB PostgreSQl

func (db *PostgreSQl) Init(options DatabaseOptions) error {
	db.Requests = make(map[string]string)
	db.Options = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		options.Host, options.Port, options.User, options.Password, options.Database)
	conn, err := db.Open()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		return err
	}

	return nil
}
