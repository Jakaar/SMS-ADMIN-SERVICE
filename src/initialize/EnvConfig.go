package initialize

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Env(val string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("ENV does not exist")
	}
	val = os.Getenv(val)
	return val
}
