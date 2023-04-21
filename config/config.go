package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	SimpleLineSecret string
	SimpleLineToken  string
	ChatGPTToken     string
}

var config *Config

func Load() *Config {

	err := godotenv.Load("dev.env")
	if err == nil {
		log.Infoln("Load dev.env file for local dev")
	}

	if config == nil {
		if os.Getenv("SIMPLE_LINE_SECRET") == "" { //other env value might not set as well
			log.Fatalln("SIMPLE_LINE_SECRET is not set:")
		}

		config = &Config{

			SimpleLineSecret: os.Getenv("SIMPLE_LINE_SECRET"),
			SimpleLineToken:  os.Getenv("SIMPLE_LINE_TOKEN"),
			ChatGPTToken:     os.Getenv("CHATGPT_TOKEN"),
		}
	}
	return config
}
