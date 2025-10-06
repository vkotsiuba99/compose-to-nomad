package utils

import (
	"log"
	"os"
)

func CreateOutputDirectory(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}
}
