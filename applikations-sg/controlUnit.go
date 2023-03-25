package applikationssg

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	commmiddleware "github.com/cricton/comm-middleware"
	commtypes "github.com/cricton/comm-types"
)

// Create cu struct using composition
type ControlUnit struct {
	Name         string
	ID           uint8
	LocalAddress net.UDPAddr
	HMIAddress   net.UDPAddr
	Middleware   *commmiddleware.Middleware
	requests     []commtypes.RequestMsg
}

// SG mainloop; Waits random amount of milliseconds and sends a random message to the HMI-controller
func (cu ControlUnit) Mainloop() {
	fmt.Println("Starting Airback SG")
	time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	for {

		request := cu.constructRandomMessage()
		cu.Middleware.SendMessage(request, cu.HMIAddress)
		response := cu.Middleware.ReceiveMessage()
		fmt.Println("Received response: ", response)
		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
	}
}

func (cu ControlUnit) constructRandomMessage() commtypes.Message {

	if len(cu.requests) <= 0 {
		requestMessage := commtypes.Message{
			SgID:    cu.ID,
			RpID:    0,
			Content: "",
		}
		return requestMessage
	}
	request := cu.requests[rand.Intn(len(cu.requests))]

	requestMessage := commtypes.Message{
		SgID:    cu.ID,
		RpID:    request.RpID,
		Content: request.Content,
	}

	return requestMessage
}

func (cu *ControlUnit) AddRequest(request commtypes.RequestMsg) {
	cu.requests = append(cu.requests, request)
}

func (cu ControlUnit) GetRequest() []commtypes.RequestMsg {
	return cu.requests
}
