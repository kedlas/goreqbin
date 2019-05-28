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

const timeFormat = "2006-01-02 15:04:05"

func (cf *ConsoleFormatter) Format(msg servers.Msg) string {
	switch msg.Type() {
	case servers.HTTPRequest:
		r, ok := msg.Data().(*http.Request)
		if !ok {
			return fmt.Sprintf("Unable to format HTTP request received at %s", msg.Time().Format(timeFormat))
		}

		return cf.formatHTTP(r)
	case servers.UDPRequest:
		r, ok := msg.Data().(string)
		if !ok {
			return fmt.Sprintf("Unable to format UDP request received at %s", msg.Time().Format(timeFormat))
		}

		return cf.formatUDP(r)
	}

	return "Unknown message type"
}

func (cf *ConsoleFormatter) formatHTTP(r *http.Request) string {
	// request basic info
	rLine := fmt.Sprintf("%-10s %-7s %s", r.Proto, r.Method, r.URL.Path)

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

func (cf *ConsoleFormatter) formatUDP(r string) string {
	return fmt.Sprintf("%-17s %s \n\n", "UDP", r)
}
