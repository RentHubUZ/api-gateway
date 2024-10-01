package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE          string
	ACCOMMODATION_SERVICE string
	ACTION_BOARD          string
	API_GATEWAY           string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":1234"))
	config.ACCOMMODATION_SERVICE = cast.ToString(Coalesce("ACCOMMODATION_SERVICE", "nimadurGo11"))
	config.ACTION_BOARD = cast.ToString(Coalesce("ACTION_BOARD", ":9876"))
	config.API_GATEWAY = cast.ToString(Coalesce("API_GATEWAY", ":50052"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
