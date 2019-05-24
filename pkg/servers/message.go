package servers

import "time"

type Msg interface {
	Timestamp() time.Time
	Data() interface{}
}

type Message struct {
	t    time.Time
	data interface{}
}

func (msg *Message) Timestamp() time.Time {
	return msg.t
}

func (msg *Message) Data() interface{} {
	return msg.data
}
