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

	middleware := &commmiddleware.Middleware{CurrentMsgID: 0}
	hmi := hmisg.HMI{}

	//create channel for sg-middleware interaction
	sg1.ControlUnit.CreateChannel(middleware)

	hmi.CreateChannel(middleware)

	go middleware.Mainloop()
	go hmi.HMI_main_loop()
	sg1.Mainloop()

}
