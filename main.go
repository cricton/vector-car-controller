package main

import (
	"net"

	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	hmisg "github.com/cricton/hmi-sg"
	"github.com/cricton/types"
)

func main() {

	//start HMI-SG
	//start Applikations-SG

	//TODO read this from config file?
	HMIAddr := net.UDPAddr{Port: 8080}
	SG0Addr := net.UDPAddr{Port: 8081}
	SG1Addr := net.UDPAddr{Port: 8082}
	SG2Addr := net.UDPAddr{Port: 8083}

	sg0 := &applikationssg.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: SG0Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg0.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "Idle too long, disable Airbag?"})

	sg1 := &applikationssg.ControlUnit{
		Name:         "Infotainment Control Unit",
		ID:           1,
		LocalAddress: SG1Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg1.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "Is Ryan the best?"})

	sg2 := &applikationssg.ControlUnit{
		Name:         "Infotainment Control Unit",
		ID:           2,
		LocalAddress: SG2Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg2.AddRequest(types.RequestMsg{RpID: types.GetString, Content: "Enter new destination..."})

	go sg0.Mainloop()
	go sg1.Mainloop()
	go sg2.Mainloop()

	hmi := hmisg.HMI{
		LocalAddress: HMIAddr,
		SGAddresses:  [16]net.UDPAddr{SG0Addr, SG1Addr, SG2Addr},
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	hmi.HMI_main_loop()

}
