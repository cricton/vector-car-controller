package testcases

import (
	"math/rand"
	"testing"

	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
	graphicinterface "github.com/cricton/graphic-interface"
	hmisg "github.com/cricton/hmi-sg"
)

func TestHMIReceive(t *testing.T) {

	sg := &applikationssg.Airbacksg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	request := commtypes.RequestMsg{
		RpID:    commtypes.ProcIDs[rand.Intn(3)],
		Content: "Idle too long. Deactivate Airbag?",
	}

	//start necessary infrastructure
	middleware := &commmiddleware.Middleware{}
	hmi := hmisg.HMI{}

	//create channel for sg-middleware interaction
	sg.ControlUnit.RegisterClient(middleware)

	hmi.RegisterHMI(middleware)

	go middleware.Mainloop()

	go sg.SendMessage(request)

	received := hmi.ReceiveMessage()

	receivedRequest := commtypes.RequestMsg{
		RpID:    received.RpID,
		Content: received.Content,
	}

	if receivedRequest != request {
		t.Errorf("Reiceved request = %#v; wanted %#v", receivedRequest, request)
	}

	if received.ReturnCode != graphicinterface.NONE {
		t.Errorf("Return code = %d; want %d", received.ReturnCode, graphicinterface.NONE)
	}

	if received.SgID != sg.ControlUnit.GetClientID() {
		t.Errorf("Client ID = %d; want %d", received.SgID, sg.ControlUnit.GetClientID())
	}

}

func TestSGReceive(t *testing.T) {

	sg := &applikationssg.Airbacksg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	request := commtypes.RequestMsg{
		RpID:    255, //non existant RpID
		Content: "Idle too long. Deactivate Airbag?",
	}

	//start necessary infrastructure
	middleware := &commmiddleware.Middleware{}
	hmi := hmisg.HMI{}

	//create channel for sg-middleware interaction
	sg.ControlUnit.RegisterClient(middleware)

	hmi.RegisterHMI(middleware)

	go middleware.Mainloop()

	go sg.SendMessage(request)

	receivedAtHMI := hmi.ReceiveMessage()
	go hmi.SendMessage(receivedAtHMI)

	receivedAtSG := sg.ReceiveMessage()

	if len(receivedAtSG.Content) > 0 {
		t.Errorf("Reiceved content = %s; wanted \"\"", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != graphicinterface.ERROR {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, graphicinterface.ERROR)
	}

	if receivedAtSG.SgID != sg.ControlUnit.GetClientID() {
		t.Errorf("Client ID = %d; want %d", receivedAtSG.SgID, sg.ControlUnit.GetClientID())
	}

	if receivedAtSG.MsgID != (receivedAtHMI.MsgID + 1) {
		t.Errorf("MsgID= %d; want %d", receivedAtSG.MsgID, receivedAtHMI.MsgID)
	}

	if receivedAtSG.RpID != commtypes.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RpID, graphicinterface.NONE)
	}

}
