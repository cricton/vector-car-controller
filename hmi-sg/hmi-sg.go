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
	hmi.GUIconnector.ResponseChannel = make(chan graphicinterface.ReturnType)
	hmi.GUIconnector.SetupGUI()

	//Start communication coroutine
	go hmi.HMI_comm_loop()

	//Start GUI loop
	mainWindow.ShowAndRun()

}

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

	hmi.GUIconnector.AddRequest(request.Content)

	returned := hmi.GUIconnector.AwaitResponse()

	response := commtypes.Message{
		Type:       commtypes.Response,
		MsgID:      request.MsgID + 1,
		SgID:       request.SgID,
		Content:    returned.Content,
		ReturnCode: returned.Code}

	return response
}

/*func getUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something...")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}*/
