package commtypes

import graphicinterface "github.com/cricton/graphic-interface"

type Message struct {
	Type       messageType
	MsgID      int16
	SgID       int16
	RpID       RemoteProcID
	Content    string
	ReturnCode graphicinterface.ReturnType
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
	GetString         RemoteProcID = 0
	GetConfirmation   RemoteProcID = 1
	GetButtonResponse RemoteProcID = 2
)
