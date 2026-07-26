package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khairozzaman91/JobPortal-Backend/config"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.Config) string {

	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s  dbname=%s",
	     cnf.DB.User,
		 cnf.DB.Password,
		 cnf.DB.Host,
		 cnf.DB.Port,
		 cnf.DB.Name,
	)

	if !cnf.DB.EnableSSLMODE{
		connString +=" sslmode=disable"
	}
	
	return connString
}

func NewConnection(cnf *config.Config) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)

	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Database Connected Successfully")
	return db, nil
}
