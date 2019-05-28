package servers

import (
	"github.com/sirupsen/logrus"

	"goreqbin/pkg/config"
)

type Server interface {
	Start()
	Stop()
}

type Servers struct {
	cfg     *config.Configuration
	log     *logrus.Logger
	msgs    chan Msg
	httpSrv *HTTP
	udpSrv  *UDP
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
		s.httpSrv = NewHTTPServer(&s.cfg.HTTP, s.log, s.msgs)
		s.httpSrv.Start()
		s.log.Infoln("HTTP server started", s.cfg.HTTP.Port)
	}

	if s.cfg.UDP.Enabled {
		s.udpSrv = NewUDPServer(&s.cfg.UDP, s.log, s.msgs)
		s.udpSrv.Start()
		s.log.Infoln("UDP server started, port: ", s.cfg.UDP.Port)
	}

	s.log.Infoln("All servers started")
}

func (s *Servers) Stop() {
	s.log.Infoln("Shutting down all servers...")

	if s.httpSrv != nil {
		s.httpSrv.Stop()
		s.log.Infoln("HTTP server stopped")
	}

	if s.udpSrv != nil {
		s.udpSrv.Stop()
		s.log.Infoln("UDP server stopped")
	}

	s.log.Infoln("All servers gracefully stopped")
}
