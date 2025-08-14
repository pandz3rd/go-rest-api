package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // always call this import to register driver
	"go-rest-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:masuk();@tcp(127.0.0.1:3306)/learn?parseTime=true")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
