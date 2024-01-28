package config

type Config struct {
	Rpc struct {
		Url       string
		PasetoKey string
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	return c
}
