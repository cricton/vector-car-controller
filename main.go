package main

import (
	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	hmisg "github.com/cricton/hmi-sg"
	"github.com/cricton/types"
)

func main() {

	//---------------------------Setup variables----------------------------------------//

	sg0 := &applikationssg.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: types.CU0Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg0.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "Idle too long, disable Airbag?"})

	sg1 := &applikationssg.ControlUnit{
		Name:         "Infotainment Control Unit",
		ID:           1,
		LocalAddress: types.CU1Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg1.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "Is Ryan the best?"})

	sg2 := &applikationssg.ControlUnit{
		Name:         "Infotainment Control Unit",
		ID:           2,
		LocalAddress: types.CU2Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg2.AddRequest(types.RequestMsg{RpID: types.GetString, Content: "Enter new destination..."})

	hmi := hmisg.HMI{
		LocalAddress: types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	//---------------------------Start Main Loops---------------------------------//

	go sg0.Mainloop()
	//go sg1.Mainloop()
	//go sg2.Mainloop()

	hmi.HMI_main_loop()

}
