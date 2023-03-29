package hmicu

import (
	"fmt"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	commmiddleware "github.com/cricton/comm-middleware"
	"github.com/cricton/types"
)

type HMI struct {
	LocalAddress net.UDPAddr
	Middleware   *commmiddleware.Middleware
	GUIconnector GUI
}

func (hmi *HMI) PrepareGUI() fyne.Window {
	application := app.New()
	mainWindow := application.NewWindow("HMI Module")
	hmi.GUIconnector = GUI{MainWindow: mainWindow}
	hmi.GUIconnector.ResponseChannel = make(chan types.ReturnTuple)
	hmi.GUIconnector.SetupGUI()

	return mainWindow
}

func (hmi *HMI) Mainloop() {
	fmt.Println("Starting HMI module...")

	mainWindow := hmi.PrepareGUI()

	//Start communication coroutine
	go hmi.Commloop()

	//Start GUI loop
	mainWindow.ShowAndRun()

}

// Waits for a message to arrive from the middleware udp server, blocking call
func (hmi HMI) ReceiveMessage() types.Message {

	message := hmi.Middleware.ReceiveMessage()

	return message
}

// Reads a message from the channel, processes it and sends a response
func (hmi *HMI) SendResponse(request types.Message) {

	response := hmi.handleMessage(request)
	hmi.Middleware.SendMessage(response, request.Address)

}

// Starts a udp server and waits for incoming messages
func (hmi *HMI) Commloop() int {

	//Start local server to listen to incoming messages
	go hmi.Middleware.StartUDPServer(hmi.LocalAddress)

	for {
		message := hmi.ReceiveMessage()

		hmi.SendResponse(message)
	}
}

// Read contents, get user input, create new message
func (hmi *HMI) handleMessage(request types.Message) types.Message {

	var returned types.ReturnTuple

	//switch depending on remote procedure ID
	switch request.RemoteProcedureID {
	case types.Info:
		hmi.GUIconnector.ShowInfo(request.ControlUnitName, request.Content)
		returned = hmi.GUIconnector.AwaitResponse()

	case types.GetString:
		hmi.GUIconnector.GetString(request.ControlUnitName, request.Content)
		returned = hmi.GUIconnector.AwaitResponse()

	case types.GetConfirmation:
		hmi.GUIconnector.GetConfirmation(request.ControlUnitName, request.Content)
		returned = hmi.GUIconnector.AwaitResponse()

	default:
		//Respond with error code in case procedure ID does not exist
		returned = types.ReturnTuple{Content: "", Code: types.ERROR}
	}

	//construct response message
	response := types.Message{
		Type:            types.Response,
		RequestID:       request.RequestID,
		ControlUnitName: request.ControlUnitName,
		Content:         returned.Content,
		ReturnCode:      returned.Code,
		Address:         hmi.LocalAddress}

	return response
}
