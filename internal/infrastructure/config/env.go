package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	DynamoRegion string
	DynamoEndpoint string
}

var (
	instance *Env
	once sync.Once
)

func Get() *Env {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic("failed to load env variables")
		}

		instance = &Env{
			DynamoEndpoint: os.Getenv("DYNAMO_DB_ENDPOINT"),
			DynamoRegion: os.Getenv("DYNAMO_DB_REGION"),
		}
	})
	
	return instance
}