package db

import (
	"clincker/utils"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func Connect() *sql.DB {
	if db == nil {
		start()
	}

	return db
}

func start() {
	fmt.Printf(os.Getenv("DB_USER"))
	connection := mysql.Config{
		// User:   os.Getenv("DB_USER"),
		// Passwd: os.Getenv("DB_PASSWORD"),
		User:   os.Getenv("DB_ROOT_USER"),
		Passwd: os.Getenv("DB_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr: fmt.Sprintf(
			"%s:%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: os.Getenv("GIN_MODE") == "debug",
	}

	var exception error

	db, exception = sql.Open("mysql", connection.FormatDSN())
	utils.Log().Exception(exception)
	utils.Log().Exception(db.Ping())
}
