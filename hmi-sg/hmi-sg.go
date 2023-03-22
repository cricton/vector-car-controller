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

func HMI_main_loop() int {

	for {
		msg_channel := make(chan commtypes.Message)
		msg := <-msg_channel
		handleMessage(msg)
	}

}

// creates a new channel and adds it to the Middleware and the controlUnit
func (hmi *HMI) CreateChannel(commmiddleware *commmiddleware.Middleware) {
	channel := make(chan commtypes.Message)
	hmi.Channel = channel
}

func handleMessage(message commtypes.Message) commtypes.Message {

	return message
}

func getUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something...")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
