package applikationssg

import (
	"fmt"

	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

type ControlUnit struct {
	Channel chan commtypes.Message
	ID      int
}

func (controlUnit ControlUnit) sendMessage(message commtypes.Message) {
	fmt.Println("Sending a message:")
	controlUnit.Channel <- message
}

func (controlUnit ControlUnit) receiveMessage() commtypes.Message {

	message := <-controlUnit.Channel

	return message
}

// creates a new channel and adds it to the Middleware and the controlUnit
func (controlUnit *ControlUnit) CreateChannel(commmiddleware *commmiddleware.Middleware) {
	channel := make(chan commtypes.Message)
	clientID := commmiddleware.RegisterChannel(channel)
	controlUnit.ID = clientID
	controlUnit.Channel = channel
}
