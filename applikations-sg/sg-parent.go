package applikationssg

import commmiddleware "github.com/cricton/comm-middleware"

type ControlUnit struct {
	Name string
}

func (controlUnit ControlUnit) GetName() string {
	return controlUnit.Name
}

func (controlUnit ControlUnit) ContactDriver() commmiddleware.Message {
	message := commmiddleware.Message{}

	message.SendMessage()

	return message
}
