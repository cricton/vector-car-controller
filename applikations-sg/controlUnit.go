package applikationssg

import (
	"fmt"

	commtypes "github.com/cricton/comm-types"
)

type ControlUnit struct {
	Channel   chan commtypes.Message
	clientID  uint8
	messageID uint16
}

func (controlUnit ControlUnit) constructMessage(request commtypes.RequestMsg) commtypes.Message {
	message := commtypes.Message{
		Type:  commtypes.Request,
		MsgID: controlUnit.messageID,
		SgID:  controlUnit.clientID,
	}

	message.RpID = request.RpID
	message.Content = request.Content

	fmt.Printf("Sending mesage: %#v\n", message)
	return message
}

func (controlUnit *ControlUnit) receiveMessage() commtypes.Message {

	message := <-controlUnit.Channel
	fmt.Printf("Received message: %#v\n", message)

	controlUnit.messageID = uint16(message.MsgID) + 1
	return message
}

func (controlUnit ControlUnit) GetClientID() uint8 {
	return controlUnit.clientID
}
