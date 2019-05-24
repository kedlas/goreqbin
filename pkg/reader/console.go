package reader

import (
	"fmt"
	"net/http"
	"strings"

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
	// request basic info
	rLine := fmt.Sprintf("%-7s %s %s", r.Method, r.URL.Path, r.Proto)

	// headers info
	var hdrs []string
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			hdrs = append(hdrs, fmt.Sprintf("%v='%v'", name, h))
		}
	}
	hLine := fmt.Sprintf("          Headers: %v", hdrs)

	// request body
	bLine := fmt.Sprintf("          Body: %s", r.Context().Value("body"))

	return fmt.Sprintf("%s \n %s \n %s \n\n", rLine, hLine, bLine)
}
