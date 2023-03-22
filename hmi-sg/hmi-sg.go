package hmisg

import (
	"bufio"
	"fmt"
	"os"

	commtypes "github.com/cricton/comm-types"
)

type HMI struct {
}

func HMI_main_loop() int {

	for {
		msg_channel := make(chan commtypes.Message)
		msg := <-msg_channel
		handleMessage(msg)
	}

}
func handleMessage(message commtypes.Message) commtypes.Message {

	return message
}

func getUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something...")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
