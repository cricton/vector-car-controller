package hmisg

import (
	"bufio"
	"fmt"
	"os"

	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

type HMI struct {
	Channel chan commtypes.Message
}

func (hmi HMI) HMI_main_loop() int {
	fmt.Println("Starting HMI module...")
	for {

		msg := <-hmi.Channel
		response := handleMessage(msg)
		hmi.Channel <- response
	}

}

// creates a new channel and adds it to the Middleware and the controlUnit
func (hmi *HMI) CreateChannel(commmiddleware *commmiddleware.Middleware) {
	channel := make(chan commtypes.Message)
	commmiddleware.RegisterHMI(channel)
	hmi.Channel = channel
}

// read contents, get user input, create new message
func handleMessage(request commtypes.Message) commtypes.Message {
	fmt.Println("Received message:")
	fmt.Println(request)

	response := commtypes.Message{
		Type:    commtypes.Response,
		MsgID:   request.MsgID + 1,
		SgID:    request.SgID,
		Content: "get user input"}

	fmt.Println("Sending message:")
	fmt.Println(response)
	return response
}

func getUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something...")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
