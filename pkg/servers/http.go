package servers

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"goreqbin/pkg/config"
)

type Server interface {
	Start()
	Stop()
}

type HttpServer struct {
	cfg  *config.Configuration
	log  *logrus.Logger
	srv  *http.Server
	msgs chan Msg
}

func NewHttpServer(cfg *config.Configuration, log *logrus.Logger, msgs chan Msg) *HttpServer {
	return &HttpServer{cfg: cfg, log: log, msgs: msgs}
}

func (h *HttpServer) Start() {
	h.srv = &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req := r

		// we must read body before sending response as it get closed then
		b, err := ioutil.ReadAll(r.Body)
		if err == nil {
			ctx := context.WithValue(r.Context(), "body", string(b))
			req = r.WithContext(ctx)
		}

		// send response
		_, err = io.WriteString(w, "ok\n")
		if err != nil {
			h.log.Errorln(err)
		}

		// pass received message to be processed
		h.msgs <- &Message{t: time.Now().UTC(), data: req}
	})

	go func() {
		if err := h.srv.ListenAndServe(); err != http.ErrServerClosed {
			h.log.Fatalf("HTTPListenAndServe(): %s", err)
		}
	}()
}

func (h *HttpServer) Stop() {
	if h.srv == nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := h.srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
