package config

// HTTP represents configuration options for HTTP server
type HTTP struct {
	Enabled bool `env:"HTTP_ENABLED" default:"true"`
	Port    int  `env:"HTTP_PORT" default:"8080"`
}

// UDP represents configuration options for UDP server
type UDP struct {
	Enabled bool `env:"UDP_ENABLED" default:"true"`
	Port    int  `env:"UDP_PORT" default:"9090"`
}

// Configuration wraps the whole app config
type Configuration = struct {
	HTTP
	UDP
}
