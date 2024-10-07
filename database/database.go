package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDb() {
	// Capture connection properties
	config := mysql.Config{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "192.168.1.200:3306",
		DBName: "db",
		Params: map[string]string{"parseTime": "true"},
	}

	// Get a database handle
	var err error
	DB, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}

func CloseDb() {
	DB.Close()
}
