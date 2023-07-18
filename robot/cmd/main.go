package main

import (
	"log"

	"github.com/paulheg/kit_plm_cad/robot/internal/network"
	"github.com/paulheg/kit_plm_cad/robot/internal/parser"
	"github.com/paulheg/kit_plm_cad/robot/internal/robot"
)

func main() {

	robot, err := robot.New(robot.JoystickFixFlutterData(robot.JoystickPolarStrategy(2)))
	if err != nil {
		log.Fatal(err)
	}
	defer robot.Dispose()

	net := network.New(network.Configuration{
		Host:    "robots.hegenberg.dev",
		Secure:  true,
		Channel: "ws/robot/Cerberus",
	})

	net.Run(func(message string) {
		parser.Parse(message, robot)
		// parser.Parse(message, &robot.LoggerStrategy{})
	})

}
