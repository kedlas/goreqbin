package reader

import (
	"github.com/sirupsen/logrus"

	"goreqbin/pkg/servers"
)

// Reader represents the struct that is able to read the messages from given channel
type Reader struct {
	msgs chan servers.Msg
	f    Formatter
	log  *logrus.Logger
}

// NewReader creates new Reader instance
func NewReader(msgs chan servers.Msg, f Formatter, log *logrus.Logger) *Reader {
	return &Reader{msgs: msgs, f: f, log: log}
}

// Read starts reading the reader's channel and prints the formatted received messages
func (r *Reader) Read() {
	go func() {
		for msg := range r.msgs {
			r.log.Println(r.f.Format(msg))
		}
	}()
}
