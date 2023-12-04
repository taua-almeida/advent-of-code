package filehandler

import (
	"log"
	"os"
)

func ReadTxtFile(name string) []byte {
	fileBytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return fileBytes
}
