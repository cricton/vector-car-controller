package graphicinterface

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

var keyboardMapping1 = [...]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
var keyboardMapping2 = [...]string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
var keyboardMapping3 = [...]string{"y", "x", "c", "v", "b", "n", "m", "-"}
