package environments

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	ENV := os.Getenv("GO_ENV")

	switch ENV {
	case "development":
		_ = godotenv.Load(".env.development")
		break
	default:
		_ = godotenv.Load(".env.development")
		break
	}
}
