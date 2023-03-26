package types

type Message struct {
	Type       RequestResponse
	SgID       uint8
	RpID       RemoteProcID
	ReturnCode ReturnType
	Content    string
}

type ReturnType uint8

const (
	NONE     ReturnType = 0
	ACCEPTED ReturnType = 1
	DECLINED ReturnType = 2
	INFO     ReturnType = 3
	STRING   ReturnType = 4
	REGISTER ReturnType = 5
	ERROR    ReturnType = 6
)

type ReturnTuple struct {
	Content string
	Code    ReturnType
}

type RequestMsg struct {
	RpID    RemoteProcID
	Content string
}

type RequestResponse uint8

const (
	Request  RequestResponse = 0
	Response RequestResponse = 1
)

// Remote Procedure ID
type RemoteProcID uint8

const (
	None            RemoteProcID = 0
	GetString       RemoteProcID = 1
	GetConfirmation RemoteProcID = 2
	Info            RemoteProcID = 3
)

var ProcIDs = [...]RemoteProcID{GetString, GetConfirmation, Info}

func (message Message) ToByte() []byte {

	return nil
}
