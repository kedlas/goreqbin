package servers

import "time"

// Msg represents the message that is received by some server
type Msg interface {
	Type() MsgType
	Time() time.Time
	Data() interface{}
}

// MsgType contains information about the message type
type MsgType int

const (
	// HTTPRequest type stands for request received via HTTP protocol
	HTTPRequest MsgType = iota
	// UDPRequest type stands for request received via UDP protocol
	UDPRequest
)

// Message is the structure that contains information about received message
type Message struct {
	msgT MsgType
	tim  time.Time
	data interface{}
}

// Type returns the type of the message
func (msg *Message) Type() MsgType {
	return msg.msgT
}

// Time returns the time that message was received
func (msg *Message) Time() time.Time {
	return msg.tim
}

// Data returns the content of the received request
func (msg *Message) Data() interface{} {
	return msg.data
}
