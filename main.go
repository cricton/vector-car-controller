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
		ID:           types.CU0ID,
		LocalAddress: types.CU0Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg0.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "Idle too long, disable Airbag?"})
	sg0.AddRequest(types.RequestMsg{RpID: types.Info, Content: "Airbag malfunctioning."})
	sg0.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "Fault detected. Restart Airbag system?"})

	sg1 := &applikationssg.ControlUnit{
		Name:         "Infotainment Control Unit",
		ID:           types.CU1ID,
		LocalAddress: types.CU1Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg1.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "New shows available. Update Infotainment system?"})
	sg1.AddRequest(types.RequestMsg{RpID: types.Info, Content: "Download completed."})
	sg1.AddRequest(types.RequestMsg{RpID: types.GetConfirmation, Content: "Out of disk space. Archive unused files?"})

	sg2 := &applikationssg.ControlUnit{
		Name:         "Navigation Control Unit",
		ID:           types.CU2ID,
		LocalAddress: types.CU2Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	sg2.AddRequest(types.RequestMsg{RpID: types.GetString, Content: "Enter new destination..."})
	sg2.AddRequest(types.RequestMsg{RpID: types.GetString, Content: "Enter home address..."})
	sg2.AddRequest(types.RequestMsg{RpID: types.Info, Content: "GPS signal lost."})

	hmi := hmisg.HMI{
		LocalAddress: types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	//---------------------------Start Main Loops---------------------------------//

	go sg0.Mainloop()
	go sg1.Mainloop()
	go sg2.Mainloop()

	hmi.HMI_main_loop()

}
