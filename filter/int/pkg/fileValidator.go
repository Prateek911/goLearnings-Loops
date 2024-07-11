package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadFile() (*os.File, error) {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading the env file")
	}

	filepath := os.Getenv("FILE_PATH")
	if filepath == "" {
		return nil, fmt.Errorf("failed to get file from env")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s", err)
	}

	return file, nil

}
