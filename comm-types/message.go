package commtypes

import graphicinterface "github.com/cricton/graphic-interface"

type Message struct {
	Type       MessageType
	SgID       uint8
	MsgID      uint16
	RpID       RemoteProcID
	ReturnCode graphicinterface.ReturnType
	Content    string
}

const (
	MHIID = 0
)

type MessageType uint8

const (
	Request  MessageType = 0
	Response MessageType = 1
)

// Remote Procedure ID
type RemoteProcID uint8

const (
	GetString       RemoteProcID = 0
	GetConfirmation RemoteProcID = 1
	Info            RemoteProcID = 2
)
