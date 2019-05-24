package reader

import (
	"fmt"
	"net/http"

	"goreqbin/pkg/servers"
)

type Formatter interface {
	Format(msg servers.Msg) string
}

type ConsoleFormatter struct {
	R *http.Request
}

func (cf *ConsoleFormatter) Format(msg servers.Msg) string {
	r, ok := msg.Data().(*http.Request)
	if ok {
		return cf.formatHttpRequest(r)
	}

	return "Unknown message data"
}

func (cf *ConsoleFormatter) formatHttpRequest(r *http.Request) string {
	return fmt.Sprintf("%s %-7v %s", r.Proto, r.Method, r.URL.Path)
}
