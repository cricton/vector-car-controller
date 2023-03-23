package commtypes

type Message struct {
	Type       messageType
	MsgID      int16
	SgID       int16
	RpID       RemoteProcID
	Content    string
	ReturnCode uint16
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
	GetDestination    RemoteProcID = 1
	GetUserResponse   RemoteProcID = 2
	GetButtonResponse RemoteProcID = 3
)
