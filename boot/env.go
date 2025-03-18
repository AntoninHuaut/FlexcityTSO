package boot

import (
	"FlexcityTest/domain"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	HttpPort          int
	SelectedAlgorithm domain.AlgorithmType
)

func LoadEnvironments() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}

	SelectedAlgorithm = loadEnum("ALGORITHM", domain.AvailableAlgorithms)
	HttpPort = loadInt("HTTP_PORT")
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

func loadEnum[T ~string](env string, enumList []T) T {
	val := loadString(env)
	for _, enum := range enumList {
		if enum == T(val) {
			return enum
		}
	}
	log.Fatalf("Failed to load env '%s': invalid value", env)
	return ""
}
