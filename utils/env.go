package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Dburl string
}

var Envs = initEnvs()

func initEnvs() *DBConfig {
	godotenv.Load()

	cnfg := &DBConfig{
		Dburl: os.Getenv("DB_URL"),
	}

	return cnfg
}
