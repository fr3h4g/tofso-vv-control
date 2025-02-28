package dht

import (
	"fmt"
	"strconv"
	"time"

	"github.com/warthog618/go-gpiocdev"
)

var lastTimestamp time.Duration
var bytes = ""

func eh(evt gpiocdev.LineEvent) {
	wt := evt.Timestamp.Microseconds() - lastTimestamp.Microseconds()
	if evt.Type == 2 && wt < 80 {
		if wt > 40 {
			bytes += "1"
		} else {
			bytes += "0"
		}
	}
	lastTimestamp = evt.Timestamp
}

func GetHumidTemp(pin int) (float32, float32, error) {
	bytes = ""
	chip := "gpiochip0"
	l, err := gpiocdev.RequestLine(chip, pin, gpiocdev.WithPullUp,
		gpiocdev.WithBothEdges,
		gpiocdev.WithEventHandler(eh))
	if err != nil {
		return float32(0), float32(0), err
	}
	defer l.Close()

	l.Reconfigure(gpiocdev.AsOutput(1))
	time.Sleep(time.Duration(80) * time.Microsecond)
	l.SetValue(0)
	time.Sleep(time.Duration(80) * time.Microsecond)
	l.SetValue(1)
	l.Reconfigure(gpiocdev.AsInput, gpiocdev.WithBothEdges)

	time.Sleep(time.Duration(2) * time.Second)
	pos := 0
	curByte := 0
	binaryStr := ""

	var nums []int

	for _, char := range bytes {
		binaryStr = binaryStr + string(char)
		pos++
		if pos == 8 {
			curByte++
			if curByte < 5 {
				numx, _ := strconv.ParseInt(binaryStr, 2, 64)
				nums = append(nums, int(numx))
			}
			binaryStr = ""
			pos = 0
		}
	}
	if len(nums) != 4 {
		return float32(0), float32(0), fmt.Errorf("no humid/temp received from sensor")
	}
	return float32(nums[0]) + float32(nums[1])/10, float32(nums[2]) + float32(nums[3])/10, nil
}
