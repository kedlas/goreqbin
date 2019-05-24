package main

import (
	"os"
	"os/signal"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"

	"goreqbin/pkg/config"
	"goreqbin/pkg/reader"
	"goreqbin/pkg/servers"
)

func main() {
	cfg := &config.Configuration{}

	// assert all required config values existence
	if err := configor.Load(cfg); err != nil {
		panic(err)
	}

	log := logrus.New()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// channel where servers pass the received data
	msgs := make(chan servers.Msg)

	// run the messages reader
	rdr := reader.NewReader(msgs, &reader.ConsoleFormatter{}, log)
	rdr.Read()

	// run the servers that publish messages
	srvs := servers.NewServers(cfg, log, msgs)
	srvs.Start()

	<-stop

	srvs.Stop()
}
