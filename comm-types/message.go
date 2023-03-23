package commtypes

type Message struct {
	Type       messageType
	MsgID      int16
	SgID       int16
	RpID       RemoteProcID
	Content    string
	ReturnCode uint8
}

const (
	MHIID = 0
)

type messageType uint8

const (
	Request  messageType = 0
	Response messageType = 1
)

// Remote Procedure ID
type RemoteProcID uint8

const (
	GetString         RemoteProcID = 1
	GetConfirmation   RemoteProcID = 2
	GetButtonResponse RemoteProcID = 3
)
