package env

import (
	"os"

	"github.com/joho/godotenv"
)

var AdminPassword string
var AllowOrigin string
var Port string

func init() {
	_ = godotenv.Load()
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	AllowOrigin = os.Getenv("ALLOW_ORIGIN")
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
}
