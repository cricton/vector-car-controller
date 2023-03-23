package hmisg

import (
	"bufio"
	"fmt"
	"os"

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

	//Start communication coroutine
	go hmi.HMI_comm_loop()

	application := app.New()
	mainWindow := application.NewWindow("MHI Module")
	gui := graphicinterface.GUI{MainWindow: mainWindow}
	gui.SetupGUI()

	mainWindow.ShowAndRun()

}

func (hmi HMI) HMI_comm_loop() int {

	for {

		msg := <-hmi.Channel
		response := handleMessage(msg)
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
func handleMessage(request commtypes.Message) commtypes.Message {
	//fmt.Println("Received message:")
	//fmt.Println(request)

	response := commtypes.Message{
		Type:    commtypes.Response,
		MsgID:   request.MsgID + 1,
		SgID:    request.SgID,
		Content: request.Content}

	//fmt.Println("Sending message:")
	//fmt.Println(response)
	return response
}

func getUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something...")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
