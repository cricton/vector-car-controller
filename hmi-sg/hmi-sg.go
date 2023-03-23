package hmisg

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
	graphicinterface "github.com/cricton/graphic-interface"
)

type HMI struct {
	Channel      chan commtypes.Message
	GUIconnector graphicinterface.GUI
}

func (hmi HMI) HMI_main_loop() {
	fmt.Println("Starting HMI module...")

	application := app.New()
	mainWindow := application.NewWindow("MHI Module")
	hmi.GUIconnector = graphicinterface.GUI{MainWindow: mainWindow}
	hmi.GUIconnector.ResponseChannel = make(chan graphicinterface.ReturnTuple)
	hmi.GUIconnector.SetupGUI()

	//Start communication coroutine
	go hmi.HMI_comm_loop()

	//Start GUI loop
	mainWindow.ShowAndRun()

}

// reads a message from the channel, processes it and sends a response
func (hmi HMI) HMI_comm_loop() int {

	for {
		msg := <-hmi.Channel
		response := hmi.handleMessage(msg)
		hmi.Channel <- response
	}

}

// creates a new channel and adds it to the Middleware and the controlUnit
func (hmi *HMI) RegisterHMI(commmiddleware *commmiddleware.Middleware) {
	channel := make(chan commtypes.Message)
	commmiddleware.RegisterHMI(channel)
	hmi.Channel = channel
}

// read contents, get user input, create new message
func (hmi HMI) handleMessage(request commtypes.Message) commtypes.Message {
	var returned graphicinterface.ReturnTuple

	//switch depending on remote procedure ID
	switch request.RpID {
	case commtypes.GetButtonResponse:

	case commtypes.GetString:
		hmi.GUIconnector.GetString(request.Content)

		returned = hmi.GUIconnector.AwaitResponse()
	case commtypes.GetConfirmation:
		hmi.GUIconnector.GetConfirmation(request.Content)

		returned = hmi.GUIconnector.AwaitResponse()
	default:
		//Respond with error code in case procedure ID does not exist
		returned = graphicinterface.ReturnTuple{Content: "", Code: graphicinterface.ERROR}
	}

	//construct response message
	response := commtypes.Message{
		Type:       commtypes.Response,
		MsgID:      request.MsgID + 1,
		SgID:       request.SgID,
		Content:    returned.Content,
		ReturnCode: returned.Code}

	return response
}
