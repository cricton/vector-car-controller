package applikationssg

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	commmiddleware "github.com/cricton/comm-middleware"
	"github.com/cricton/types"
)

// Create cu struct using composition
type ControlUnit struct {
	Name         string
	ID           uint8
	LocalAddress net.UDPAddr
	HMIAddress   net.UDPAddr
	Middleware   *commmiddleware.Middleware
	requests     []types.RequestMsg
}

// SG mainloop; Waits random amount of seconds and sends a random message to the HMI-controller
func (cu ControlUnit) Mainloop() {
	fmt.Println("Starting ", cu.Name)
	time.Sleep(time.Duration(rand.Intn(10)+3) * time.Second)
	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	//send registration message to HMI
	cu.Register()

	for {

		request := cu.constructRandomMessage()
		cu.Middleware.SendMessage(request, cu.HMIAddress)
		response := cu.Middleware.ReceiveMessage()
		fmt.Println("Received response: ", response)
		time.Sleep(time.Duration(rand.Intn(20)+3) * time.Second)
	}
}

func (cu ControlUnit) Register() {
	registerMessage := types.Message{
		SgID:    cu.ID,
		RpID:    types.Register,
		Content: cu.LocalAddress.String(),
	}
	cu.Middleware.SendMessage(registerMessage, cu.HMIAddress)

	response := cu.Middleware.ReceiveMessage()

	if response.ReturnCode != types.ACCEPTED {
		log.Fatal("Could not register Control Unit ", cu.ID)
	}
}

func (cu ControlUnit) constructRandomMessage() types.Message {

	if len(cu.requests) <= 0 {
		log.Fatal("no messages to be sent")
	}
	request := cu.requests[rand.Intn(len(cu.requests))]

	requestMessage := types.Message{
		SgID:    cu.ID,
		RpID:    request.RpID,
		Content: request.Content,
	}

	return requestMessage
}

func (cu *ControlUnit) AddRequest(request types.RequestMsg) {
	cu.requests = append(cu.requests, request)
}

func (cu ControlUnit) GetRequests() []types.RequestMsg {
	return cu.requests
}
