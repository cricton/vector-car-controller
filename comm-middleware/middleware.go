package commmiddleware

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"github.com/cricton/types"
)

type Middleware struct {
	IncomingChannel chan types.Message
}

func (middleware *Middleware) StartUDPServer(address net.UDPAddr) {

	listener, err := net.ListenUDP("udp", &address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// create a temp buffer
	receiveBuffer := make([]byte, 512)
	for {

		_, _, err := listener.ReadFromUDP(receiveBuffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(receiveBuffer)

		messageStruct := new(types.Message)

		// creates a decoder object
		gobDecoder := gob.NewDecoder(tmpbuff)

		// decodes buffer and unmarshals it into a Message struct
		gobDecoder.Decode(messageStruct)

		// pass received message to internal channel
		middleware.IncomingChannel <- *messageStruct

	}
}

func (middleware Middleware) SendMessage(message types.Message, destinationAddress net.UDPAddr) {

	//Connect to address and send message
	c, err := net.Dial("udp", destinationAddress.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	//Serialize struct to send over UDP
	var byteBuffer bytes.Buffer
	gobEncoder := gob.NewEncoder(&byteBuffer)
	err = gobEncoder.Encode(message)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	bytes := byteBuffer.Bytes()

	c.Write(bytes)

	c.Close()

}

// Blocking call to receive the next incoming message
func (middleware Middleware) ReceiveMessage() types.Message {
	response := <-middleware.IncomingChannel

	return response
}

// Non-blocking call to receive the next incoming message
func (middleware Middleware) ReceiveMessageAsync() types.Message {

	select {
	case response := <-middleware.IncomingChannel:
		return response
	default:
		// receiving from middleware.IncomingChannel would block
	}

	//Return empty message
	return types.Message{}
}
