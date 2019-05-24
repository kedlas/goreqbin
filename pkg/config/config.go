package config

type Configuration = struct {
	HTTP struct {
		Enabled bool `env:"HTTP_ENABLED" default:"true"`
		Port    int  `env:"HTTP_PORT" default:"8080"`
	}

	UDP struct {
		Enabled bool `env:"UDP_ENABLED" default:"true"`
		Port    int  `env:"UDP_PORT" default:"9090"`
	}
}
