package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func generateAPIKey() (string, error) {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}

func main() {
	apiKey, err := generateAPIKey()
	if err != nil {
		log.Fatalf("Failed to generate API key: %v", err)
	}

	log.Println("Output:" + apiKey)
}
