package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, def string) string {
	if val := os.Getenv(key); len(val) > 0 {
		return val
	}
	return def
}

func SetupEnvFile() {
	env := os.Getenv("GO_ENV")
	if env != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}
