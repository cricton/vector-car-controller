package commmiddleware

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"

	commtypes "github.com/cricton/comm-types"
)

type Middleware struct {
	IncomingChannel chan commtypes.Message
}

func (middleware *Middleware) StartTCPServer(address string) {
	fmt.Println("Listening to address: ", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// create a temp buffer
	recBuffer := make([]byte, 512)
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = connection.Read(recBuffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(recBuffer)

		tmpstruct := new(commtypes.Message)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)

		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		// lets print out!
		fmt.Println(tmpstruct) // reflects.TypeOf(tmpstruct) == Message{}
		middleware.IncomingChannel <- *tmpstruct

	}
	//if received send to local channel
	//middleware.incomingChannel <- received
}

func (middleware Middleware) SendMessage(message commtypes.Message, destinationAddress string) {
	//Connect to address and send message
	fmt.Println("Dialing...")
	c, err := net.Dial("tcp", destinationAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Serialize struct to send over TCP
	var byteBuffer bytes.Buffer
	enc := gob.NewEncoder(&byteBuffer)
	err = enc.Encode(message)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	bytes := byteBuffer.Bytes()

	c.Write(bytes)

	c.Close()

}

func (middleware Middleware) ReceiveMessage() commtypes.Message {
	response := <-middleware.IncomingChannel

	return response
}
