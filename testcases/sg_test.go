package testcases

import (
	"math/rand"
	"net"
	"testing"

	"fyne.io/fyne/v2/test"
	applicationcu "github.com/cricton/application-cu"
	commmiddleware "github.com/cricton/comm-middleware"
	hmicu "github.com/cricton/hmi-cu"
	types "github.com/cricton/types"
)

func TestIllegalRemoteProcedureID(t *testing.T) {

	//---------------------------Setup variables---------------------------------//

	ip := net.ParseIP("127.0.0.1")
	HMIAddr := net.UDPAddr{IP: ip, Port: 8082}
	CUAddr := net.UDPAddr{IP: ip, Port: 8083}

	cu := &applicationcu.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CUAddr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}
	request := types.Message{
		Type:              types.Request,
		RemoteProcedureID: 255, //non existant RemoteProcedureID
		Content:           "Idle too long. Deactivate Airbag?",
	}

	hmi := hmicu.HMI{
		LocalAddress: HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	//--Start UDP Servers--//
	go cu.Middleware.StartUDPServer(CUAddr)
	go hmi.Middleware.StartUDPServer(HMIAddr)

	receivedAtHmi := hmi.ReceiveMessage()
	hmi.SendResponse(receivedAtHmi)

	go cu.Middleware.SendMessage(request, cu.HMIAddress)
	receivedAtHMI := hmi.ReceiveMessage()
	go hmi.SendResponse(receivedAtHMI)

	receivedAtSG := cu.Middleware.ReceiveMessage()

	//---------------------------Check results---------------------------------//
	if len(receivedAtSG.Content) > 0 {
		t.Errorf("Reiceved content = %s; wanted \"\"", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != types.ERROR {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, types.ERROR)
	}

	if receivedAtSG.ControlUnitName != cu.Name {
		t.Errorf("Client ID = %s; want %s", receivedAtSG.ControlUnitName, cu.Name)
	}

	if receivedAtSG.RemoteProcedureID != types.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RemoteProcedureID, types.NONE)
	}

}

func TestGetString(t *testing.T) {

	//---------------------------Setup variables----------------------------------------//

	ip := net.ParseIP("127.0.0.1")
	HMIAddr := net.UDPAddr{IP: ip, Port: 8080}
	CUAddr := net.UDPAddr{IP: ip, Port: 8081}

	cu := &applicationcu.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CUAddr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	message := types.Message{
		Type:              types.Request,
		RemoteProcedureID: types.GetString,
		Content:           "Hello ...",
	}
	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	hmi := hmicu.HMI{
		LocalAddress: HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	hmi.PrepareGUI()

	go hmi.Commloop()

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

	if receivedAtSG.ReturnCode != types.STRING {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, types.ERROR)
	}

	if receivedAtSG.ControlUnitName != cu.Name {
		t.Errorf("Client ID = %s; want %s", receivedAtSG.ControlUnitName, cu.Name)
	}

	if receivedAtSG.RemoteProcedureID != types.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RemoteProcedureID, types.NONE)
	}
}

func TestGetConfirmation(t *testing.T) {

	//---------------------------Setup variables----------------------------------------//

	ip := net.ParseIP("127.0.0.1")
	HMIAddr := net.UDPAddr{IP: ip, Port: 8084}
	CUAddr := net.UDPAddr{IP: ip, Port: 8085}

	cu := &applicationcu.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CUAddr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	message := types.Message{
		Type:              types.Request,
		RemoteProcedureID: types.GetConfirmation,
		Content:           "Hello?",
	}

	hmi := hmicu.HMI{
		LocalAddress: HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	go cu.Middleware.StartUDPServer(cu.LocalAddress)
	hmi.PrepareGUI()

	go hmi.Commloop()

	cu.Middleware.SendMessage(message, cu.HMIAddress)

	var accepted types.ReturnType

	if rand.Intn(2) == 0 {
		accepted = types.ACCEPTED
	} else {
		accepted = types.DECLINED
	}
	hmi.GUIconnector.ResponseChannel <- types.ReturnTuple{Content: "", Code: accepted}

	receivedAtSG := cu.Middleware.ReceiveMessage()

	//---------------------------Check results------------------------------------------//

	if len(receivedAtSG.Content) > 0 {
		t.Errorf("Reiceved content = %s; wanted \"\"", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != accepted {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, accepted)
	}

	if receivedAtSG.ControlUnitName != cu.Name {
		t.Errorf("Client ID = %s; want %s", receivedAtSG.ControlUnitName, cu.Name)
	}

	if receivedAtSG.RemoteProcedureID != types.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RemoteProcedureID, types.NONE)
	}

}

func TestGetInfo(t *testing.T) {

	//---------------------------Setup variables----------------------------------------//

	ip := net.ParseIP("127.0.0.1")
	HMIAddr := net.UDPAddr{IP: ip, Port: 8086}
	CUAddr := net.UDPAddr{IP: ip, Port: 8087}

	cu := &applicationcu.ControlUnit{
		Name:         "Airbag Control Unit",
		ID:           0,
		LocalAddress: CUAddr,
		HMIAddress:   HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	message := types.Message{
		Type:              types.Request,
		RemoteProcedureID: types.Info,
		Content:           "Hello!",
	}

	hmi := hmicu.HMI{
		LocalAddress: HMIAddr,
		Middleware:   &commmiddleware.Middleware{IncomingChannel: make(chan types.Message)},
	}

	//---------------------------Run necessary commands---------------------------------//

	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	hmi.PrepareGUI()

	go hmi.Commloop()

	cu.Middleware.SendMessage(message, cu.HMIAddress)

	hmi.GUIconnector.ResponseChannel <- types.ReturnTuple{Content: "", Code: types.INFO}

	receivedAtSG := cu.Middleware.ReceiveMessage()

	//---------------------------Check results------------------------------------------//

	if len(receivedAtSG.Content) != 0 {
		t.Errorf("Reiceved content = %s; wanted \"\"", receivedAtSG.Content)
	}

	if receivedAtSG.ReturnCode != types.INFO {
		t.Errorf("Return code = %d; want %d", receivedAtSG.ReturnCode, types.INFO)
	}

	if receivedAtSG.ControlUnitName != cu.Name {
		t.Errorf("Client ID = %s; want %s", receivedAtSG.ControlUnitName, cu.Name)
	}

	if receivedAtSG.RemoteProcedureID != types.None {
		t.Errorf("Remote procedure= %d; want %d", receivedAtSG.RemoteProcedureID, types.NONE)
	}
}
