package config

import "os"

type Config struct {
	MongoUri string
}

func LoadConig() *Config {
	return &Config{
		MongoUri: os.Getenv("MONGO_URI"),
	}
}