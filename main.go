package main

import (
	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	hmisg "github.com/cricton/hmi-sg"
)

func main() {

	//start HMI-SG
	//start Applikations-SG

	//TODO read this from config file?
	sg1 := &applikationssg.Airbacksg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	sg2 := &applikationssg.Assistantsg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	middleware := &commmiddleware.Middleware{}
	hmi := hmisg.HMI{}

	//create channel for sg-middleware interaction
	sg1.ControlUnit.CreateChannel(middleware)
	sg2.ControlUnit.CreateChannel(middleware)

	hmi.CreateChannel(middleware)

	go middleware.Mainloop()
	go hmi.HMI_main_loop()
	go sg1.Mainloop()
	sg2.Mainloop()

}
