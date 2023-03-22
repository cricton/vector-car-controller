package commtypes

type Message struct {
	MsgID      int16
	SenderID   int16
	ReceiverID int16
	Content    string
}

const (
	MHIID = 0
)
