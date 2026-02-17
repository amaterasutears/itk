package config

type Server struct {
	Port               int `env:"SERVER_PORT" env-default:"8080"`
	ShutdownTimeoutSec int `env:"SERVER_SHUTDOWN_TIMEOUT_SEC" env-default:"10"`
}
