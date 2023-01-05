package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DsnTemplate string = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"

var (
	username = "root"
	password = "users_root_password"
	hostname = "users_db"
	dbName   = "users_db"
)

func ConectToDatabase() *sql.DB {
	dsn := buildDSN()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Panicf("error on database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Panicf("error during database ping: %v", err)
	}

	return db
}

func buildDSN() string {
	return fmt.Sprintf(DsnTemplate, username, password, hostname, dbName)
}
