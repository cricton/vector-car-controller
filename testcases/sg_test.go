package testcases

import (
	"math/rand"
	"net"
	"testing"

	"fyne.io/fyne/v2/test"
	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
	graphicinterface "github.com/cricton/graphic-interface"
	hmisg "github.com/cricton/hmi-sg"
)

func TestIllegalRpID(t *testing.T) {

	//---------------------------Setup variables---------------------------------//

	HMIAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8082}
	CU0Addr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8083}

	cu := &applikationssg.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CU0Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}
	request := commtypes.Message{
		Type:    commtypes.Request,
		RpID:    255, //non existant RpID
		Content: "Idle too long. Deactivate Airbag?",
	}

	hmi := hmisg.HMI{
		LocalAddress: HMIAddr,
		SGAddresses:  [16]net.UDPAddr{CU0Addr},
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	go cu.Middleware.StartUDPServer(CU0Addr)
	go hmi.Middleware.StartUDPServer(HMIAddr)

	go cu.Middleware.SendMessage(request, cu.HMIAddress)

	receivedAtHMI := hmi.ReceiveMessage()
	go hmi.SendResponse(receivedAtHMI, hmi.SGAddresses[0])

	receivedAtSG := cu.Middleware.ReceiveMessage()

	//---------------------------Check results---------------------------------//
	if len(receivedAtSG.Content) > 0 {
		t.Errorf("Reiceved content = %s; wanted \"\"", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != graphicinterface.ERROR {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, graphicinterface.ERROR)
	}

	if receivedAtSG.SgID != cu.ID {
		t.Errorf("Client ID = %d; want %d", receivedAtSG.SgID, cu.ID)
	}

	if receivedAtSG.RpID != commtypes.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RpID, graphicinterface.NONE)
	}

}

func TestGetString(t *testing.T) {

	//---------------------------Setup variables----------------------------------------//

	HMIAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
	CU0Addr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081}

	cu := &applikationssg.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CU0Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	message := commtypes.Message{
		Type:    commtypes.Request,
		RpID:    commtypes.GetString,
		Content: "Hello ...",
	}
	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	hmi := hmisg.HMI{
		LocalAddress: HMIAddr,
		SGAddresses:  [16]net.UDPAddr{CU0Addr},
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	hmi.PrepareGUI()

	go hmi.HMI_comm_loop()

	cu.Middleware.SendMessage(message, cu.HMIAddress)

	enterButton := hmi.GUIconnector.GetEnterButton()

	testInput := "Hello World!"
	hmi.GUIconnector.UserEntry.WriteString(testInput)
	test.Tap(enterButton)

	receivedAtSG := cu.Middleware.ReceiveMessage()

	//---------------------------Check results------------------------------------------//

	if receivedAtSG.Content != testInput {
		t.Errorf("Reiceved content = %s; wanted Hello World!", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != graphicinterface.STRING {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, graphicinterface.ERROR)
	}

	if receivedAtSG.SgID != cu.ID {
		t.Errorf("Client ID = %d; want %d", receivedAtSG.SgID, cu.ID)
	}

	if receivedAtSG.RpID != commtypes.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RpID, graphicinterface.NONE)
	}
}

func TestGetConfirmation(t *testing.T) {

	//---------------------------Setup variables----------------------------------------//

	HMIAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8084}
	CU0Addr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8085}

	cu := &applikationssg.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CU0Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	message := commtypes.Message{
		Type:    commtypes.Request,
		RpID:    commtypes.GetConfirmation,
		Content: "Hello?",
	}
	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	hmi := hmisg.HMI{
		LocalAddress: HMIAddr,
		SGAddresses:  [16]net.UDPAddr{CU0Addr},
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	hmi.PrepareGUI()

	go hmi.HMI_comm_loop()

	cu.Middleware.SendMessage(message, cu.HMIAddress)

	var accepted graphicinterface.ReturnType

	if rand.Intn(2) == 0 {
		accepted = graphicinterface.ACCEPTED
	} else {
		accepted = graphicinterface.DECLINED
	}
	hmi.GUIconnector.ResponseChannel <- graphicinterface.ReturnTuple{Content: "", Code: accepted}

	receivedAtSG := cu.Middleware.ReceiveMessage()

	//---------------------------Check results------------------------------------------//

	if len(receivedAtSG.Content) > 0 {
		t.Errorf("Reiceved content = %s; wanted \"\"", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != accepted {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, accepted)
	}

	if receivedAtSG.SgID != cu.ID {
		t.Errorf("Client ID = %d; want %d", receivedAtSG.SgID, cu.ID)
	}

	if receivedAtSG.RpID != commtypes.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RpID, graphicinterface.NONE)
	}

}

func TestGetInfo(t *testing.T) {

	//---------------------------Setup variables----------------------------------------//

	HMIAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8086}
	CU0Addr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8087}

	cu := &applikationssg.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CU0Addr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	message := commtypes.Message{
		Type:    commtypes.Request,
		RpID:    commtypes.Info,
		Content: "Hello!",
	}
	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	hmi := hmisg.HMI{
		LocalAddress: HMIAddr,
		SGAddresses:  [16]net.UDPAddr{CU0Addr},
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan commtypes.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	hmi.PrepareGUI()

	go hmi.HMI_comm_loop()

	cu.Middleware.SendMessage(message, cu.HMIAddress)

	hmi.GUIconnector.ResponseChannel <- graphicinterface.ReturnTuple{Content: "", Code: graphicinterface.INFO}

	receivedAtSG := cu.Middleware.ReceiveMessage()

	//---------------------------Check results------------------------------------------//

	if len(receivedAtSG.Content) != 0 {
		t.Errorf("Reiceved content = %s; wanted \"\"", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != graphicinterface.INFO {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, graphicinterface.INFO)
	}

	if receivedAtSG.SgID != cu.ID {
		t.Errorf("Client ID = %d; want %d", receivedAtSG.SgID, cu.ID)
	}

	if receivedAtSG.RpID != commtypes.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RpID, graphicinterface.NONE)
	}
}
