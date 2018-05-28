package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"log"
)

func main() {
	firmataAdaptor := firmata.NewTCPAdaptor("192.168.1.210:3030")
	led := gpio.NewLedDriver(firmataAdaptor, "2")

	work := func() {
		ticker1 := time.NewTicker(950 * time.Millisecond)
		time.Sleep(500 * time.Millisecond)
		ticker2 := time.NewTicker(1000 * time.Millisecond)
		for {
			select {
			case <-ticker1.C:
				led.On()
			case <-ticker2.C:
				led.Off()
			}
		}
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	log.Printf("Name=%s", robot.Name)

	robot.Start()
}
