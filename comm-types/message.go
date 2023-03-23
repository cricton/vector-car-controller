package commtypes

type Message struct {
	Type    messageType
	MsgID   int16
	SgID    int16
	Content string
}

const (
	MHIID = 0
)

type messageType uint8

const (
	Request  messageType = 0
	Response messageType = 1
)
