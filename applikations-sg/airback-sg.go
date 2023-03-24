package applikationssg

import (
	"fmt"
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
	fmt.Println("Starting Airback SG")
	for {

		message := sg.constructRandomMessage()

		sg.ControlUnit.sendMessage(message)
		sg.ControlUnit.receiveMessage()

		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)

	}
}

func (sg Airbacksg) constructRandomMessage() commtypes.Message {

	message := commtypes.Message{
		Type:  commtypes.Request,
		MsgID: sg.ControlUnit.messageID,
		SgID:  sg.ControlUnit.clientID,
	}

	message.RpID = commtypes.Info
	message.Content = "Idle too long. Deactivate Airbag?"

	return message
}
