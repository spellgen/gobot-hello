package wemos_server

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"log"
)

func main() {
	firmataAdaptor := firmata.NewTCPAdaptor("192.168.1.210:3030")
	fa := firmata.NewAdaptor()

	led := gpio.NewLedDriver(firmataAdaptor, "2")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	log.Printf("Name=%s", robot.Name)

	robot.Start()
}
