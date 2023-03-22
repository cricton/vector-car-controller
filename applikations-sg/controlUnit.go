package applikationssg

import (
	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

type ControlUnit struct {
	Channel chan commtypes.Message
}

func (controlUnit ControlUnit) sendMessage(message commtypes.Message) {
	controlUnit.Channel <- message
}

func (controlUnit ControlUnit) receiveMessage() commtypes.Message {

	message := <-controlUnit.Channel

	return message
}

// creates a new channel and adds it to the Middleware and the controlUnit
func (controlUnit *ControlUnit) CreateChannel(commmiddleware commmiddleware.Middleware) {
	channel := make(chan commtypes.Message)
	commmiddleware.RegisterChannel(channel)
	controlUnit.Channel = channel
}
