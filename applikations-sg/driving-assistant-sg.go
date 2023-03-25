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
	fmt.Println("Starting Driving Assistant SG")
	for {

		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)

	}
}

// TODO proper message ID
func (sg Assistantsg) constructRandomRequest() commtypes.RequestMsg {

	request := commtypes.RequestMsg{

		Content: "Activate cruise control?",
		RpID:    commtypes.GetConfirmation,
	}

	return request
}
