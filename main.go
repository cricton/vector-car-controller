package main

import (
	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
)

func main() {

	//start HMI-SG
	//start Applikations-SG

	//TODO read this from config file?
	sg1 := &applikationssg.Airbacksg{
		ControlUnit: &applikationssg.ControlUnit{},
	}
	middleware := &commmiddleware.Middleware{CurrentMsgID: 0}

	//create channel for sg-middleware interaction
	sg1.ControlUnit.CreateChannel(*middleware)

	//hmi := hmisg.HMI{}

}
