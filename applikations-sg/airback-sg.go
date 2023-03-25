package applikationssg

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

// Create sg struct using composition
type Airbacksg struct {
	Name         string
	ID           uint8
	LocalAddress net.UDPAddr
	HMIAddress   net.UDPAddr
	Middleware   *commmiddleware.Middleware
}

// SG mainloop; Waits random amount of milliseconds and sends a random message to the HMI-controller
func (sg Airbacksg) Mainloop() {
	fmt.Println("Starting Airback SG")
	time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
	go sg.Middleware.StartUDPServer(sg.LocalAddress)

	for {

		request := sg.constructRandomRequest()
		sg.Middleware.SendMessage(request, sg.HMIAddress)
		response := sg.Middleware.ReceiveMessage()
		fmt.Println("Received response: ", response)
		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
	}
}

func (controlUnit ControlUnit) constructMessage(request commtypes.RequestMsg) commtypes.Message {
	message := commtypes.Message{
		Type: commtypes.Request,
		SgID: controlUnit.ClientID,
	}

	message.RpID = request.RpID
	message.Content = request.Content

	return message
}
func (sg Airbacksg) constructRandomRequest() commtypes.Message {

	request := commtypes.Message{
		SgID:    sg.ID,
		RpID:    commtypes.ProcIDs[rand.Intn(3)],
		Content: "Idle too long. Deactivate Airbag?",
	}

	return request
}
