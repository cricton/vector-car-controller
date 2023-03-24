package applikationssg

import (
	"fmt"
	"math/rand"
	"time"

	commtypes "github.com/cricton/comm-types"
)

// Create sg struct using composition
type Navigationsg struct {
	Name        string
	ControlUnit *ControlUnit
}

// SG mainloop; Waits random amount of milliseconds and sends a random message to the HMI-controller
func (sg Navigationsg) Mainloop() {
	fmt.Println("Starting Driving Assistant SG")
	for {

		request := sg.constructRandomMessage()

		sg.ControlUnit.sendMessage(request)
		sg.ControlUnit.receiveMessage()

		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)

	}
}

// TODO proper message ID
func (sg Navigationsg) constructRandomMessage() commtypes.RequestMsg {

	request := commtypes.RequestMsg{
		Content: "You seem to be lost. Start navigation?",
		RpID:    commtypes.GetConfirmation,
	}

	return request
}
