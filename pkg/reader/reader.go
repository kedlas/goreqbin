package reader

import (
	"github.com/sirupsen/logrus"

	"goreqbin/pkg/servers"
)

type Reader struct {
	msgs chan servers.Msg
	f    Formatter
	log  *logrus.Logger
}

func NewReader(msgs chan servers.Msg, f Formatter, log *logrus.Logger) *Reader {
	return &Reader{msgs: msgs, f: f, log: log}
}

func (r *Reader) Read() {
	select {
	case msg := <-r.msgs:
		r.log.Println(r.f.Format(msg))
	default:
		r.log.Println("Nothing received")
	}
}