package main

import (
	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
	hmisg "github.com/cricton/hmi-sg"
)

const (
	HMIAddr string = "127.0.0.1:8080"
	SG1Addr string = "127.0.0.1:8081"
)

func main() {

	//start HMI-SG
	//start Applikations-SG

	//TODO read this from config file?
	sg1 := &applikationssg.Airbacksg{
		LocalAddress: SG1Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	go sg1.Mainloop()

	hmi := hmisg.HMI{
		LocalAddress: HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	hmi.HMI_main_loop()
	//go sg1.Mainloop()

	//hmi.HMI_main_loop()

}
