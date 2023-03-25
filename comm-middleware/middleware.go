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

func (middleware *Middleware) StartUDPServer(address net.UDPAddr) {
	fmt.Println("Listening to address: ", address)

	listener, err := net.ListenUDP("udp", &address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// create a temp buffer
	recBuffer := make([]byte, 512)
	for {

		_, remoteaddr, err := listener.ReadFromUDP(recBuffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(remoteaddr)

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(recBuffer)

		tmpstruct := new(commtypes.Message)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)

		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		middleware.IncomingChannel <- *tmpstruct

	}
	//if received send to local channel
	//middleware.incomingChannel <- received
}

func (middleware Middleware) SendMessage(message commtypes.Message, destinationAddress net.UDPAddr) {
	//Connect to address and send message
	fmt.Println("Dialing...")
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

func (middleware Middleware) ReceiveMessage() commtypes.Message {
	response := <-middleware.IncomingChannel

	return response
}
