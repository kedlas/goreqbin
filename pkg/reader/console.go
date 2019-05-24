package reader

import (
	"fmt"
	"net/http"
)

type Formatter interface {
	Format(s interface{}) string
}

type ConsoleFormatter struct {
	R *http.Request
}

func (cf *ConsoleFormatter) Format(s interface{}) string {
	r, ok := s.(*http.Request)
	if ok {
		return cf.formatHttpRequest(r)
	}

	return "Unknown subject format"
}

func (cf *ConsoleFormatter) formatHttpRequest(r *http.Request) string {
	return fmt.Sprintf("%s %-7v %s", r.Proto, r.Method, r.URL.Path)
}
