package middlewares

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func DotEnvVariable(key string) string {
	if flag.Lookup("test.v") == nil {
		// normal run
		err := godotenv.Load("../.env")
		if err != nil {
			err = godotenv.Load(".env")
			if err != nil {
				log.Fatalf("Error loading .env file")
				os.Exit(1)
			}
		}
		return os.Getenv(key)
	} else {
		// run under go test
		path, _ := os.Getwd()
		err := godotenv.Load(strings.Split(path, "src")[0] + "/" + ".env")
		fmt.Println(path)
		if err != nil {
			log.Fatalf("Error loading .env.test file")
		}
		return os.Getenv(key)
	}
}
