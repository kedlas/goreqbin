package servers

import "time"

type Msg interface {
	Type() MsgType
	Time() time.Time
	Data() interface{}
}

type MsgType int

const (
	HTTPRequest MsgType = iota
	UDPRequest
)

type Message struct {
	msgT MsgType
	tim  time.Time
	data interface{}
}

func (msg *Message) Type() MsgType {
	return msg.msgT
}

func (msg *Message) Time() time.Time {
	return msg.tim
}

func (msg *Message) Data() interface{} {
	return msg.data
}
