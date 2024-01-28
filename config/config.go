package config

import (
	"github.com/naoina/toml"
	"os"
)

type EthNode struct {
	URL   string
	Chain string
}

type Config struct {
	Rpc struct {
		Url       string
		PasetoKey string
	}

	Server struct {
		Port string
	}

	Redis struct {
		Url      string
		DB       int
		User     string
		Password string
	}

	Node map[string]EthNode
}

func NewConfig(path string) *Config {
	c := new(Config)

	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err = toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
