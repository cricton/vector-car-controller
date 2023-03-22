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
func (sg Airbacksg) Mainloop() {

	for {

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		message := sg.constructRandomMessage()

		sg.ControlUnit.sendMessage(message)
		sg.ControlUnit.receiveMessage()

	}
}

// TODO proper message ID
func (sg Airbacksg) constructRandomMessage() commtypes.Message {
	message := commtypes.Message{
		MsgID:      1,
		SenderID:   int16(sg.ControlUnit.ID),
		ReceiverID: commtypes.MHIID,
		Content:    "test",
	}

	return message
}
