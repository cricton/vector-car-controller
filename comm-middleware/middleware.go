package commmiddleware

import (
	"fmt"

	hmisg "github.com/cricton/hmi-sg"
)

type Middleware struct {
	HMI *hmisg.HMI
}

func (message Message) SendMessage() int {
	fmt.Println("Sent a message")
	return 1
}
