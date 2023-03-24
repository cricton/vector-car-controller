package commtypes

import graphicinterface "github.com/cricton/graphic-interface"

type Message struct {
	Type       RequestResponse
	SgID       uint8
	MsgID      uint16
	RpID       RemoteProcID
	ReturnCode graphicinterface.ReturnType
	Content    string
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
