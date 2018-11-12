package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	// set default database
	var cfg = Config()
	var err error
	db, err = sql.Open("mysql", cfg.DB.Addr)
	if err != nil {
		panic(err)
	}

	// set default max open conns and max idle conns
	db.SetMaxOpenConns(cfg.DB.Max)
	db.SetMaxIdleConns(cfg.DB.Idle)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err)
	}
}

func GetDB() *sql.DB {
	initDB()
	return db
}
