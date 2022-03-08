package db

import (
	"database/sql"
	"fmt"

	"github.com/Baraha/clickHouse_event_server.git/pkg/errors"
	"github.com/labstack/gommon/log"
)

type SQLRepo struct {
	DB *sql.DB
}

func NewSQLRepository(dataName string) (*SQLRepo, error) {
	db, err := sql.Open("clickhouse", dataName)
	if err != nil {
		log.Error("Could not get a handle to the database")
		panic(err)
	}
	rep := &SQLRepo{
		DB: db,
	}
	if !rep.available() {
		return nil, errors.NewDBNotAvailable("Database not available")
	}
	return rep, nil
}

func (rep *SQLRepo) available() bool {
	err := rep.DB.Ping()
	if err != nil {
		log.Printf("error : %v", err)
		return false
	}
	return true
}

func (db *SQLRepo) Find(col, table string) sql.Result {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	result, err := tx.Exec(fmt.Sprintf("select %v from %v", col, table))

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result %v", result)

	if tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return result
}

func Insert() {

}
