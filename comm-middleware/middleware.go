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
	fmt.Println("Listening to port: ", address.Port)

	listener, err := net.ListenUDP("udp", &address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// create a temp buffer
	recBuffer := make([]byte, 512)
	for {

		_, _, err := listener.ReadFromUDP(recBuffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(recBuffer)

		tmpstruct := new(types.Message)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)

		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		middleware.IncomingChannel <- *tmpstruct

	}
	//if received send to local channel
	//middleware.incomingChannel <- received
}

func (middleware Middleware) SendMessage(message types.Message, destinationAddress net.UDPAddr) {
	//Connect to address and send message
	c, err := net.Dial("udp", destinationAddress.String())
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

func (middleware Middleware) ReceiveMessage() types.Message {
	response := <-middleware.IncomingChannel

	return response
}
