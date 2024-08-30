package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	HttpPort int `env:"HTTP_PORT" env-default:"8080"`
}

func LoadConfig() Config {
	var cfg Config
	_, err := os.Stat(".env")
	if err == nil {
		err = cleanenv.ReadConfig(".env", &cfg)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
