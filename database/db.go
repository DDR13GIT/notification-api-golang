package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	connectionString := "notifuser:admin123@tcp(localhost:3306)/notification_db"
	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err.Error())
	}

	fmt.Println("Connected to the database")
}