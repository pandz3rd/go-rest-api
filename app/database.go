package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // always call this import to register driver
	"go-rest-api/helper"
	"strconv"
	"time"
)

func NewDB() *sql.DB {
	config := GetConfig()
	username := config.GetString("database.username")
	password := config.GetString("database.password")
	host := config.GetString("database.host")
	port := config.GetInt("database.port")
	name := config.GetString("database.name")

	dataSource := username + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + name + "?parseTime=true"

	db, err := sql.Open("mysql", dataSource)
	helper.PanicIfError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
