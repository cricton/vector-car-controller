package applikationssg

import (
	commtypes "github.com/cricton/comm-types"
)

type ControlUnit struct {
	ClientID uint8
}

func (controlUnit ControlUnit) ConstructMessage(request commtypes.RequestMsg) commtypes.Message {
	message := commtypes.Message{
		Type: commtypes.Request,
		SgID: controlUnit.ClientID,
	}

	message.RpID = request.RpID
	message.Content = request.Content

	return message
}
