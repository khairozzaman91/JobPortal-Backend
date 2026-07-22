package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func load() string {
	return "host=localhost port=5432 user=postgres password=2173 dbname=jobportal sslmode=disable"
}

func GetConnect() (*sqlx.DB, error) {
	dbSource := load()

	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
    fmt.Println("Database Connected Successfully")
	return db, nil
}
