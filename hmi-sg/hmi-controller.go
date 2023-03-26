package hmisg

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	commmiddleware "github.com/cricton/comm-middleware"
	"github.com/cricton/types"
)

type HMI struct {
	LocalAddress net.UDPAddr
	cuAddresses  [16]net.UDPAddr
	Middleware   *commmiddleware.Middleware
	GUIconnector GUI
}

func (hmi *HMI) PrepareGUI() fyne.Window {
	application := app.New()
	mainWindow := application.NewWindow("MHI Module")
	hmi.GUIconnector = GUI{MainWindow: mainWindow}
	hmi.GUIconnector.ResponseChannel = make(chan types.ReturnTuple)
	hmi.GUIconnector.SetupGUI()

	return mainWindow
}

func (hmi *HMI) HMI_main_loop() {
	fmt.Println("Starting HMI module...")

	mainWindow := hmi.PrepareGUI()

	//Start communication coroutine
	go hmi.HMI_comm_loop()

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
	sgAddress := hmi.cuAddresses[request.SgID]
	hmi.Middleware.SendMessage(response, sgAddress)

}

// Starts a udp server and waits for incoming messages
func (hmi *HMI) HMI_comm_loop() int {

	//Start local server to listen to incoming messages
	go hmi.Middleware.StartUDPServer(hmi.LocalAddress)

	for {
		message := hmi.ReceiveMessage()

		hmi.SendResponse(message)
	}
}

// Reads the cu address from the content field and adds it to the local cuAddresses variable
func (hmi *HMI) RegisterCU(request types.Message) types.ReturnTuple {
	addressAndPort := strings.Split(request.Content, ":")

	port, err := strconv.Atoi(addressAndPort[1])
	if err != nil {
		log.Fatal("Illegal address port")
	}

	ip := net.ParseIP(addressAndPort[0])

	hmi.cuAddresses[request.SgID] = net.UDPAddr{IP: ip, Port: port}

	return types.ReturnTuple{Content: "", Code: types.ACCEPTED}
}

// Read contents, get user input, create new message
func (hmi *HMI) handleMessage(request types.Message) types.Message {

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

	case types.Register:
		returned = hmi.RegisterCU(request)

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

func (hmi HMI) GetcuAddresses() [16]net.UDPAddr {
	return hmi.cuAddresses
}
