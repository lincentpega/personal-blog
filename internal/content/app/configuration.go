package app

import (
	"os"
)

type Configuration struct {
	MongoUrl string
}

func LoadConfiguration() *Configuration {
	return &Configuration{MongoUrl: os.Getenv("MONGO_URL")}
}
