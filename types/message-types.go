package types

import (
	"net"

	"github.com/google/uuid"
)

type Message struct {
	Type              MessageType
	ControlUnitName   string
	RemoteProcedureID RemoteProcID
	RequestID         uuid.UUID
	ReturnCode        ReturnType
	Address           net.UDPAddr
	Content           string
}

type ReturnType uint8

const (
	NONE     ReturnType = 0
	ACCEPTED ReturnType = 1
	DECLINED ReturnType = 2
	INFO     ReturnType = 3
	STRING   ReturnType = 4
	ERROR    ReturnType = 5
)

type ReturnTuple struct {
	Content string
	Code    ReturnType
}

type MessageType uint8

const (
	Request  MessageType = 0
	Response MessageType = 1
)

// Remote Procedure ID
type RemoteProcID uint8

const (
	None            RemoteProcID = 0
	GetString       RemoteProcID = 1
	GetConfirmation RemoteProcID = 2
	Info            RemoteProcID = 3
)

type RequestMsg struct {
	RemoteProcedureID RemoteProcID
	Content           string
}

type RequestStatus uint8

const (
	Free    RequestStatus = 0
	Pending RequestStatus = 1
)
