package applicationcu

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	commmiddleware "github.com/cricton/comm-middleware"
	"github.com/cricton/types"
	"github.com/google/uuid"
)

// Create cu struct using composition
type ControlUnit struct {
	Name           string
	ID             uint8
	LocalAddress   net.UDPAddr
	HMIAddress     net.UDPAddr
	Middleware     *commmiddleware.Middleware
	requests       []types.RequestMsg
	activeRequests uint8
}

func (cu *ControlUnit) sendMessagePeriodically(minTime int, maxTime int) {
	for {
		if cu.activeRequests < types.MaxQueuedResponses {
			request := cu.getRandomRequest()
			cu.sendMessage(request)
		} else {
			fmt.Println("Max concurrent messages sent")
		}
		time.Sleep(time.Duration(rand.Intn(maxTime)+minTime) * time.Second)
	}

}

func (cu *ControlUnit) sendMessage(request types.RequestMsg) {

	message := types.Message{
		Type:              types.Request,
		ControlUnitName:   cu.Name,
		RequestID:         uuid.New(),
		Content:           request.Content,
		RemoteProcedureID: request.RemoteProcedureID,
		Address:           cu.LocalAddress,
	}

	cu.activeRequests += 1
	cu.Middleware.SendMessage(message, cu.HMIAddress)
}

func (cu *ControlUnit) receiveMessageAsync() (types.Message, bool) {
	message, received := cu.Middleware.ReceiveMessageAsync()

	if received {
		cu.activeRequests -= 1
		return message, true
	}

	return message, false
}

func (cu ControlUnit) getRandomRequest() types.RequestMsg {

	if len(cu.requests) <= 0 {
		log.Fatal("no messages to be sent")
	}
	request := cu.requests[rand.Intn(len(cu.requests))]

	return request
}

func (cu *ControlUnit) AddRequest(request types.RequestMsg) {
	cu.requests = append(cu.requests, request)
}

func (cu ControlUnit) GetRequests() []types.RequestMsg {
	return cu.requests
}

// SG mainloop; Waits random amount of seconds and sends a random message to the HMI-controller
func (cu *ControlUnit) Mainloop() {
	fmt.Println("Starting ", cu.Name)
	go cu.Middleware.StartUDPServer(cu.LocalAddress)

	go cu.sendMessagePeriodically(10, 25)

	for {
		response, received := cu.receiveMessageAsync()
		if received {
			fmt.Println("Received response: ", response)
		}
	}
}
