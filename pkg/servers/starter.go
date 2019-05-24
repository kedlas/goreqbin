package servers

import (
	"github.com/sirupsen/logrus"

	"goreqbin/pkg/config"
)

type Servers struct {
	cfg     *config.Configuration
	log     *logrus.Logger
	msgs    chan Msg
	httpSrv *HttpServer
}

func NewServers(
	cfg *config.Configuration,
	log *logrus.Logger,
	msgs chan Msg,
) *Servers {
	return &Servers{cfg: cfg, log: log, msgs: msgs}
}

func (s *Servers) Start() {
	if s.cfg.HTTP.Enabled {
		s.httpSrv = NewHttpServer(s.cfg, s.log, s.msgs)
		s.httpSrv.Start()
		s.log.Infoln("HTTP server started")
	}

	s.log.Infoln("All servers started")
}

func (s *Servers) Stop() {
	s.log.Infoln("Shutting down all servers...")

	if s.httpSrv != nil {
		s.httpSrv.Stop()
		s.log.Infoln("HTTP server stopped")
	}

	s.log.Infoln("All servers gracefully stopped")
}
