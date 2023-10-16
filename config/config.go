package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uchupx/geek-garden-test/pkg/helper"
)

type Config struct {
	Port     string
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Redis struct {
		Host     string
		Database int64
		Password string
	}
}

var conf *Config
var path *string

func initialize() {
	conf = &Config{}

	if err := godotenv.Load(helper.StringDefault(path, ".env")); err != nil {
		log.Fatalf("Error loading .env file")
	}

	conf.Database.Host = os.Getenv("DATABASE_HOST")
	conf.Database.Port = os.Getenv("DATABASE_PORT")
	conf.Database.User = os.Getenv("DATABASE_USERNAME")
	conf.Database.Password = os.Getenv("DATABASE_PASSWORD")
	conf.Database.Name = os.Getenv("DATABASE_NAME")

	conf.Redis.Host = os.Getenv("REDIS_HOST")
	conf.Redis.Database = helper.StringToInt64(os.Getenv("REDIS_DATABASE"), 0)
	conf.Redis.Password = os.Getenv("REDIS_PASSWORD")
}

func GetConfig() *Config {
	if conf == nil {
		initialize()
	}

	return conf
}

func SetPath(val string) {
	path = &val
}
