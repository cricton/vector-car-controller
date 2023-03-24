package applikationssg

import (
	"fmt"

	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

type ControlUnit struct {
	Channel   chan commtypes.Message
	clientID  uint8
	messageID uint16
}

func (controlUnit ControlUnit) sendMessage(request commtypes.RequestMsg) {
	message := commtypes.Message{
		Type:  commtypes.Request,
		MsgID: controlUnit.messageID,
		SgID:  controlUnit.clientID,
	}

	message.RpID = request.RpID
	message.Content = request.Content

	fmt.Printf("Sending mesage: %#v\n", message)
	controlUnit.Channel <- message
}

func (controlUnit *ControlUnit) receiveMessage() commtypes.Message {

	message := <-controlUnit.Channel
	fmt.Printf("Received message: %#v\n", message)

	controlUnit.messageID = uint16(message.MsgID) + 1
	return message
}

// creates a new channel and adds it to the Middleware and the controlUnit
func (controlUnit *ControlUnit) RegisterClient(commmiddleware *commmiddleware.Middleware) {
	channel := make(chan commtypes.Message)
	clientID := commmiddleware.RegisterClient(channel)
	controlUnit.clientID = uint8(clientID)
	controlUnit.Channel = channel
}
