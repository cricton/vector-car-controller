package types

import "net"

var HMIAddr = net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}
var CU0Addr = net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8081}
var CU1Addr = net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8082}
var CU2Addr = net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8083}

const (
	CU0ID = 0
	CU1ID = 1
	CU2ID = 2
)
