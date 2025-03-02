package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/warthog618/go-gpiocdev"
)

func saveTankLevelToDB(sensor1 int, sensor2 int, sensor3 int) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO tanklevel (sensor1, sensor2, sensor3) VALUES (?, ?, ?)",
		sensor1,
		sensor2,
		sensor3,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetTankLevel(c *gpiocdev.Chip) error {
	// chip := "gpiochip0"
	pins := []int{20, 21, 22}

	ll, err := c.RequestLines(pins, gpiocdev.WithPullUp,
		gpiocdev.AsInput)
	defer ll.Close()
	if err != nil {
		return err
	}
	rr := []int{0, 0, 0}
	ll.Values(rr)
	fmt.Println(rr)
	go saveTankLevelToDB(rr[0], rr[1], rr[2])
	return nil
}
