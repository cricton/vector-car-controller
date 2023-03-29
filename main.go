package main

import (
	applicationcu "github.com/cricton/application-cu"
	commmiddleware "github.com/cricton/comm-middleware"
	hmicu "github.com/cricton/hmi-cu"
	"github.com/cricton/types"
)

func main() {

	//---------------------------Setup variables----------------------------------------//

	cu0 := &applicationcu.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           types.CU0ID,
		LocalAddress: types.CU0Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	cu0.AddRequest(types.RequestMsg{RemoteProcedureID: types.GetConfirmation, Content: "Idle too long, disable Airbag?"})
	cu0.AddRequest(types.RequestMsg{RemoteProcedureID: types.Info, Content: "Airbag malfunctioning."})
	cu0.AddRequest(types.RequestMsg{RemoteProcedureID: types.GetConfirmation, Content: "Fault detected. Restart Airbag system?"})

	cu1 := &applicationcu.ControlUnit{
		Name:         "Infotainment Control Unit",
		ID:           types.CU1ID,
		LocalAddress: types.CU1Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	cu1.AddRequest(types.RequestMsg{RemoteProcedureID: types.GetConfirmation, Content: "New shows available. Update Infotainment system?"})
	cu1.AddRequest(types.RequestMsg{RemoteProcedureID: types.Info, Content: "Download completed."})
	cu1.AddRequest(types.RequestMsg{RemoteProcedureID: types.GetConfirmation, Content: "Out of disk space. Archive unused files?"})

	cu2 := &applicationcu.ControlUnit{
		Name:         "Navigation Control Unit",
		ID:           types.CU2ID,
		LocalAddress: types.CU2Addr,
		HMIAddress:   types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	cu2.AddRequest(types.RequestMsg{RemoteProcedureID: types.GetString, Content: "Enter new destination..."})
	cu2.AddRequest(types.RequestMsg{RemoteProcedureID: types.GetString, Content: "Enter home address..."})
	cu2.AddRequest(types.RequestMsg{RemoteProcedureID: types.Info, Content: "GPS signal lost."})

	hmi := hmicu.HMI{
		LocalAddress: types.HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	//---------------------------Start Main Loops---------------------------------//

	go cu0.Mainloop()
	go cu1.Mainloop()
	go cu2.Mainloop()

	hmi.Mainloop()

}
