package servers

import (
	"net"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"goreqbin/pkg/config"
)

type UDP struct {
	cfg  *config.UDP
	log  *logrus.Logger
	conn *net.UDPConn
	msgs chan Msg
}

func NewUDPServer(cfg *config.UDP, log *logrus.Logger, msgs chan Msg) *UDP {
	return &UDP{cfg: cfg, log: log, msgs: msgs}
}

func (u *UDP) Start() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: u.cfg.Port,
		IP:   net.ParseIP("0.0.0.0"),
	})

	if err != nil {
		u.log.Fatalf("Could not start UDP server. Reason: %s", err)
	}

	u.conn = conn

	go func() {
		for {
			message := make([]byte, 20)
			rlen, _, err := conn.ReadFromUDP(message[:])
			if err != nil {
				u.log.Errorln("Error reading udp message", err)
			}

			data := strings.TrimSpace(string(message[:rlen]))

			u.msgs <- &Message{msgT: UDPRequest, tim: time.Now().UTC(), data: data}
		}
	}()
}

func (u *UDP) Stop() {
	if u.conn != nil {
		err := u.conn.Close()
		if err != nil {
			panic(err)
		}
	}
}
