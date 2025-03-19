package boot

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	httpPort int
)

func LoadEnvironments() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}

	httpPort = loadInt("HTTP_PORT")
}

func loadString(env string) string {
	val, ok := os.LookupEnv(env)
	if !ok {
		log.Fatalf("Failed to load env '%s'", env)
		return ""
	}
	return val
}

func loadInt(env string) int {
	val, err := strconv.Atoi(loadString(env))
	if err != nil {
		log.Fatalf("Failed to load env '%s': %s", env, err)
		return 0
	}
	return val
}
