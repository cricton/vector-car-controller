package applikationssg

import (
	"math/rand"
	"time"

	commtypes "github.com/cricton/comm-types"
)

// Create sg struct using composition
type Airbacksg struct {
	Name        string
	ControlUnit *ControlUnit
}

// SG mainloop; Waits random amount of milliseconds and sends a random message to the HMI-controller
func (controlUnit ControlUnit) Mainloop() {

	for {

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		message := constructRandomMessage()
		controlUnit.sendMessage(message)
		controlUnit.receiveMessage()

	}
}

func constructRandomMessage() commtypes.Message {
	message := commtypes.Message{MsgID: 1}
	return message
}
