package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fr3h4g/tofso-vv-control/internal/dht"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/go-sql-driver/mysql"
)

// checkEnv checks if all required environment variables are set.
func checkEnv() error {
	envs := []string{
		"MYSQL_DSN",
	}

	for _, env := range envs {
		if os.Getenv(env) == "" {
			return fmt.Errorf("%s is required", env)
		}
	}
	return nil
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func sqlMigrate() error {
	var db *sql.DB

	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		fmt.Println("Error connection to database:", err)
		return err
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		// Codecov ignore
		return err
	}

	return nil
}

func main() {
	log.Println("starting tofso-vv-control")

	godotenv.Load()

	err := checkEnv()
	if err != nil {
		panic(err)
	}

	err = sqlMigrate()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			humid, temp, err := dht.GetHumidTemp(17)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			log.Printf("humidity %.1f%%, temperature: %.1f°C\n", humid, temp)
			saveToDB(temp, humid)
			time.Sleep(time.Duration(60) * time.Second)
		}
	}()

	log.Println("startup complete, running...")

	for {
		time.Sleep(60 * time.Second)
	}
}
