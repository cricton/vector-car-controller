package applikationssg

import (
	"fmt"
	"math/rand"
	"time"

	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

// Create sg struct using composition
type Airbacksg struct {
	Name         string
	LocalAddress string
	HMIAddress   string
	Middleware   *commmiddleware.Middleware
}

// SG mainloop; Waits random amount of milliseconds and sends a random message to the HMI-controller
func (sg Airbacksg) Mainloop() {
	fmt.Println("Starting Airback SG")
	time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
	go sg.Middleware.StartTCPServer(sg.LocalAddress)

	for {

		request := sg.constructRandomRequest()
		sg.Middleware.SendMessage(request, sg.HMIAddress)
		response := sg.Middleware.ReceiveMessage()
		fmt.Println(response)
		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
	}
}

func (sg Airbacksg) constructRandomRequest() commtypes.Message {

	request := commtypes.Message{
		RpID:    commtypes.ProcIDs[rand.Intn(3)],
		Content: "Idle too long. Deactivate Airbag?",
	}

	return request
}
