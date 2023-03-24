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

		request := sg.constructRandomRequest()

		sg.ControlUnit.sendMessage(request)
		sg.ControlUnit.receiveMessage()

		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)

	}
}

func (sg Airbacksg) constructRandomRequest() commtypes.RequestMsg {

	request := commtypes.RequestMsg{
		RpID:    commtypes.ProcIDs[rand.Intn(4)+1],
		Content: "Idle too long. Deactivate Airbag?",
	}

	return request
}

func (sg Airbacksg) SendMessage(request commtypes.RequestMsg) {
	sg.ControlUnit.sendMessage(request)
}

func (sg Airbacksg) ReceiveMessage() commtypes.Message {

	return sg.ControlUnit.receiveMessage()
}
