package applicationcu

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
	Name             string
	ID               uint8
	LocalAddress     net.UDPAddr
	HMIAddress       net.UDPAddr
	Middleware       *commmiddleware.Middleware
	requests         []types.RequestMsg
	pendingResponses [8]types.RequestStatus
}

func (cu *ControlUnit) getNewRequestID() uint8 {
	var requestID uint8
	for k, v := range cu.pendingResponses {
		if v == types.Free {
			requestID = uint8(k)
			break
		}
	}

	cu.pendingResponses[requestID] = types.Pending
	fmt.Println(cu.pendingResponses)
	return requestID
}

func (cu *ControlUnit) clearRequestID(requestID uint8) {
	cu.pendingResponses[requestID] = types.Free
}

func (cu *ControlUnit) sendMessagePeriodically(minTime int, maxTime int) {
	for {
		request := cu.getRandomRequest()
		cu.sendMessage(request)
		time.Sleep(time.Duration(rand.Intn(maxTime)+minTime) * time.Second)

	}

}

func (cu *ControlUnit) sendMessage(request types.RequestMsg) {

	message := types.Message{
		Type:              types.Request,
		ControlUnitID:     cu.ID,
		RequestID:         cu.getNewRequestID(),
		Content:           request.Content,
		RemoteProcedureID: request.RemoteProcedureID,
	}
	cu.Middleware.SendMessage(message, cu.HMIAddress)
}

func (cu *ControlUnit) receiveMessageAsync() types.Message {
	message := cu.Middleware.ReceiveMessageAsync()
	if (message == types.Message{}) {
		return message
	}

	cu.clearRequestID(message.RequestID)
	return message
}

func (cu ControlUnit) Register() {
	//wait for HMI UDP Server to boot up
	time.Sleep(time.Second)

	registerMessage := types.Message{
		ControlUnitID:     cu.ID,
		RemoteProcedureID: types.Register,
		Content:           cu.LocalAddress.String(),
	}
	cu.Middleware.SendMessage(registerMessage, cu.HMIAddress)

	response := cu.Middleware.ReceiveMessage()

	if response.ReturnCode != types.ACCEPTED {
		log.Fatal("Could not register Control Unit ", cu.ID)
	}
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

	//send registration message to HMI
	cu.Register()

	go cu.sendMessagePeriodically(10, 10)

	for {

		response := cu.receiveMessageAsync()
		if (response != types.Message{}) {
			fmt.Println("Received response: ", response)
		}

	}
}
