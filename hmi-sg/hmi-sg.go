package hmisg

import (
	"fmt"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	commmiddleware "github.com/cricton/comm-middleware"
	graphicinterface "github.com/cricton/graphic-interface"
	"github.com/cricton/types"
)

type HMI struct {
	LocalAddress net.UDPAddr
	SGAddresses  [16]net.UDPAddr
	Channel      chan types.Message
	Middleware   *commmiddleware.Middleware
	GUIconnector graphicinterface.GUI
}

func (hmi *HMI) PrepareGUI() fyne.Window {
	application := app.New()
	mainWindow := application.NewWindow("MHI Module")
	hmi.GUIconnector = graphicinterface.GUI{MainWindow: mainWindow}
	hmi.GUIconnector.ResponseChannel = make(chan types.ReturnTuple)
	hmi.GUIconnector.SetupGUI()

	return mainWindow
}

func (hmi HMI) HMI_main_loop() {
	fmt.Println("Starting HMI module...")

	mainWindow := hmi.PrepareGUI()

	//Start communication coroutine
	go hmi.HMI_comm_loop()

	//Start GUI loop
	mainWindow.ShowAndRun()

}

// Waits for a message to arrive from the middleware
func (hmi HMI) ReceiveMessage() types.Message {

	message := hmi.Middleware.ReceiveMessage()

	return message
}

// reads a message from the channel, processes it and sends a response
func (hmi HMI) SendResponse(message types.Message, address net.UDPAddr) {

	response := hmi.handleMessage(message)
	hmi.Middleware.SendMessage(response, address)

}

// reads a message from the channel, processes it and sends a response
func (hmi HMI) HMI_comm_loop() int {

	//Start local server to listen to incoming messages
	go hmi.Middleware.StartUDPServer(hmi.LocalAddress)

	for {
		message := hmi.ReceiveMessage()

		hmi.SendResponse(message, hmi.SGAddresses[message.SgID])
	}
}

// read contents, get user input, create new message
func (hmi HMI) handleMessage(request types.Message) types.Message {

	var returned types.ReturnTuple

	//switch depending on remote procedure ID
	switch request.RpID {
	case types.Info:
		hmi.GUIconnector.ShowInfo(request.Content)
		returned = hmi.GUIconnector.AwaitResponse()

	case types.GetString:
		hmi.GUIconnector.GetString(request.Content)
		returned = hmi.GUIconnector.AwaitResponse()

	case types.GetConfirmation:
		hmi.GUIconnector.GetConfirmation(request.Content)
		returned = hmi.GUIconnector.AwaitResponse()

	default:
		//Respond with error code in case procedure ID does not exist
		returned = types.ReturnTuple{Content: "", Code: types.ERROR}
	}

	//construct response message
	response := types.Message{
		Type:       types.Response,
		SgID:       request.SgID,
		Content:    returned.Content,
		ReturnCode: returned.Code}

	return response
}
