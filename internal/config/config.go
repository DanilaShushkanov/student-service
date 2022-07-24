package config

type Config struct {
	DB       string `env:"DB"`
	User     string `env:"USER_DB"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	GRPCAddr string `env:"GRPC_ADDR"`
	HTTPAddr string `env:"HTTP_ADDR"`
}
