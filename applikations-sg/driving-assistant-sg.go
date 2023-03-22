package applikationssg

import (
	"fmt"
	"math/rand"
	"time"

	commtypes "github.com/cricton/comm-types"
)

// Create sg struct using composition
type Assistantsg struct {
	Name        string
	ControlUnit *ControlUnit
}

// SG mainloop; Waits random amount of milliseconds and sends a random message to the HMI-controller
func (sg Assistantsg) Mainloop() {
	fmt.Println("Starting Airback Steuergerät")
	for {

		message := sg.constructRandomMessage()

		sg.ControlUnit.sendMessage(message)
		sg.ControlUnit.receiveMessage()

		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	}
}

// TODO proper message ID
func (sg Assistantsg) constructRandomMessage() commtypes.Message {

	message := commtypes.Message{
		MsgID:    sg.ControlUnit.messageID,
		SenderID: int16(sg.ControlUnit.clientID),
		Content:  "test2",
	}

	return message
}
