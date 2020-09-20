package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

//NewMongoDB NewMongoDB
func NewPostgresDB() (*Postgres, error) {
	//config := DBConfig{}

	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"postgresdb", "5432", "postgres", "mypassword", "company_manager")

	//PsqlInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
	//	config.GetUser(), config.GetPassword(), config.GetHost(), config.GetDBName(), config.GetPort())
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{Db: db}, nil
}
