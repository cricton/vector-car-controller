package main

import (
	"net"

	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
	hmisg "github.com/cricton/hmi-sg"
)

func main() {

	//start HMI-SG
	//start Applikations-SG

	//TODO read this from config file?
	HMIAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
	SG0Addr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081}
	SG1Addr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8082}

	sg1 := &applikationssg.Airbacksg{
		ID:           0,
		LocalAddress: SG0Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	go sg1.Mainloop()

	hmi := hmisg.HMI{
		LocalAddress: HMIAddr,
		SGAddresses:  [16]net.UDPAddr{SG0Addr, SG1Addr},
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	hmi.HMI_main_loop()
	//go sg1.Mainloop()

	//hmi.HMI_main_loop()

}
