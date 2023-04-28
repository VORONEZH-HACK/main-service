package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresQL struct {
	Requests map[string]string
	Options  string
}

func (db *PostgresQL) Prepare(pathToRequestFiles string, request string) error {
	dat, err := os.ReadFile(pathToRequestFiles + request + ".sql")
	if err != nil {
		return err
	}
	db.Requests[request] = string(dat)

	return nil
}

func (db *PostgresQL) Get(request string) string {
	return db.Requests[request]
}

func (db *PostgresQL) Open() (*sql.DB, error) {
	dbconn, err := sql.Open("postgres", db.Options)
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}

var PostgresQLDB PostgresQL

func (db *PostgresQL) Init(options DatabaseOptions) error {
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
