package servers

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"goreqbin/pkg/config"
)

// HTTP represents the HTTP server
type HTTP struct {
	cfg  *config.HTTP
	log  *logrus.Logger
	srv  *http.Server
	msgs chan Msg
}

type ctxKey string

const (
	Body ctxKey = "body"
)

//NewHTTP creates new HTTP server instance
func NewHTTP(cfg *config.HTTP, log *logrus.Logger, msgs chan Msg) *HTTP {
	return &HTTP{cfg: cfg, log: log, msgs: msgs}
}

// Start starts the server
func (h *HTTP) Start() {
	h.srv = &http.Server{Addr: fmt.Sprintf(":%d", h.cfg.Port)}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req := r

		// we must read body before sending response as it get closed then
		b, err := ioutil.ReadAll(r.Body)
		if err == nil {
			ctx := context.WithValue(r.Context(), Body, string(b))
			req = r.WithContext(ctx)
		}

		// send response
		_, err = io.WriteString(w, "ok\n")
		if err != nil {
			h.log.Errorln("Unable to send http response", err)
		}

		// pass received message to be processed
		h.msgs <- &Message{msgT: HTTPRequest, tim: time.Now().UTC(), data: req}
	})

	go func() {
		if err := h.srv.ListenAndServe(); err != http.ErrServerClosed {
			h.log.Fatalf("HTTPListenAndServe(): %s", err)
		}
	}()
}

// Stop terminates the server
func (h *HTTP) Stop() {
	if h.srv == nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := h.srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
