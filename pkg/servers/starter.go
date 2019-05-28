package servers

import (
	"github.com/sirupsen/logrus"

	"goreqbin/pkg/config"
)

// Server represents the web server that is able to receive messages. May use HTTP, UDP or other protocols
type Server interface {
	// Start starts the server
	Start()
	// Stop stops the server
	Stop()
}

// Servers is wrapper for all servers the app provides
type Servers struct {
	cfg     *config.Configuration
	log     *logrus.Logger
	msgs    chan Msg
	httpSrv *HTTP
	udpSrv  *UDP
}

// NewServers creates the new Servers struct
func NewServers(
	cfg *config.Configuration,
	log *logrus.Logger,
	msgs chan Msg,
) *Servers {
	return &Servers{cfg: cfg, log: log, msgs: msgs}
}

// Start starts all servers enabled by configuration
func (s *Servers) Start() {
	if s.cfg.HTTP.Enabled {
		s.httpSrv = NewHTTP(&s.cfg.HTTP, s.log, s.msgs)
		s.httpSrv.Start()
		s.log.Infoln("HTTP server started, port: ", s.cfg.HTTP.Port)
	}

	if s.cfg.UDP.Enabled {
		s.udpSrv = NewUDP(&s.cfg.UDP, s.log, s.msgs)
		s.udpSrv.Start()
		s.log.Infoln("UDP server started, port: ", s.cfg.UDP.Port)
	}

	s.log.Infoln("All servers started")
}

// Stop terminates all servers that had been started
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
