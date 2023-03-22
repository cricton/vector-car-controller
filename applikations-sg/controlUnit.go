package applikationssg

import (
	"fmt"

	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

type ControlUnit struct {
	Channel   chan commtypes.Message
	clientID  int16
	messageID int16
}

func (controlUnit ControlUnit) sendMessage(message commtypes.Message) {
	fmt.Println("Sending a message:")
	controlUnit.Channel <- message
}

func (controlUnit *ControlUnit) receiveMessage() commtypes.Message {

	message := <-controlUnit.Channel
	controlUnit.messageID = int16(message.MsgID) + 1
	return message
}

// creates a new channel and adds it to the Middleware and the controlUnit
func (controlUnit *ControlUnit) CreateChannel(commmiddleware *commmiddleware.Middleware) {
	channel := make(chan commtypes.Message)
	clientID := commmiddleware.RegisterChannel(channel)
	controlUnit.clientID = int16(clientID)
	controlUnit.Channel = channel
}
