package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	Port string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	fileName := "./config/config.json"
	err := gonfig.GetConf(fileName, &configuration)

	if err != nil {
		log.Fatal(err)
	}

	return configuration
}
