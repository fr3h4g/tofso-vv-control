package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func saveTempHumidToDB(temp float32, humid float32) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO temperature (temperature, humidity) VALUES (?, ?)",
		temp,
		humid,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}
