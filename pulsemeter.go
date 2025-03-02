package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/warthog618/go-gpiocdev"
)

var pins = []int{10, 11, 12, 13}

var pulses_meter1 = 0
var pulses_meter2 = 0
var pulses_meter3 = 0
var pulses_meter4 = 0

func pulseEvent(evt gpiocdev.LineEvent) {
	//fmt.Println(evt)
	if evt.Offset == pins[0] {
		pulses_meter1++
	}
}

func savePulsesToDB(sensor1 int, sensor2 int, sensor3 int, sensor4 int) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO pulses (meter1, meter2, meter3, meter4) VALUES (?, ?, ?, ?)",
		sensor1,
		sensor2,
		sensor3,
		sensor4,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func savePulses() {
	for {
		time.Sleep(10 * time.Second)

		go func() {
			fmt.Println(pulses_meter1, pulses_meter2, pulses_meter3, pulses_meter4)

			savePulsesToDB(pulses_meter1, pulses_meter2, pulses_meter3, pulses_meter4)

			pulses_meter1 = 0
			pulses_meter2 = 0
			pulses_meter3 = 0
			pulses_meter4 = 0
		}()
	}
}

func CountPluses(c *gpiocdev.Chip) error {
	// chip := "gpiochip0"
	//pins := []int{10, 11, 12, 13}

	ll, err := c.RequestLines(pins, gpiocdev.WithPullUp,
		gpiocdev.AsInput, gpiocdev.WithFallingEdge, gpiocdev.WithEventHandler(pulseEvent))
	defer ll.Close()
	if err != nil {
		return err
	}
	go savePulses()
	for {
		time.Sleep(60 * time.Second)
	}
}
