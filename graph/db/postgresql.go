package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB
var Tx *sqlx.Tx

func InitDB() {
	db, err := sqlx.Connect("postgres", "user=postgres password=123304050 port=5432 dbname=blogdb_v1 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(500)

	Db = db
}
