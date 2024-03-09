package configs

type Config struct {
	ServerConfig ServerConfig
}

type ServerConfig struct {
	BindAddress string `env:"BIND_ADDRESS" envDefault:":8080"`
}
