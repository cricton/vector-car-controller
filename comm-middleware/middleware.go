package commmiddleware

import (
	"fmt"
	"time"

	commtypes "github.com/cricton/comm-types"
)

type Middleware struct {
	CurrentMsgID int
	Channels     []chan commtypes.Message
	HMIChannel   chan commtypes.Message
}

func SendMessage(commtypes.Message) int {
	fmt.Println("Sent a message")
	return 1
}

// Fetches an unused messageID and updates the value
func (middleware *Middleware) GetAndIncMsgID() int {
	msgID := middleware.CurrentMsgID
	middleware.CurrentMsgID = middleware.CurrentMsgID + 1
	return msgID
}

func (middleware *Middleware) RegisterClient(channel chan commtypes.Message) int {
	middleware.Channels = append(middleware.Channels, channel)

	//get position of channel in array to use as ID
	clientID := len(middleware.Channels)
	return clientID
}

// register HMI module
func (middleware *Middleware) RegisterHMI(channel chan commtypes.Message) {
	middleware.HMIChannel = channel
}

// SG mainloop; Waits random amount of milliseconds and sends a random message to the HMI-controller
func (middleware Middleware) Mainloop() {
	fmt.Println("Starting up middleware...")
	fmt.Println(middleware.Channels)
	for {

		time.Sleep(time.Millisecond)
		//check all channels for incoming messages

		for _, channel := range middleware.Channels {
			select {
			//if a message is received it gets passed to the HMI and waits for a response
			case message := <-channel:
				fmt.Println("Received a message!!")
				middleware.HMIChannel <- message
				message = <-middleware.HMIChannel
				channel <- message
			default:
			}
		}

	}
}
