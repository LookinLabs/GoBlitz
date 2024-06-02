package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func generateAPIKey() (string, error) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}

func main() {
	apiKey, err := generateAPIKey()
	if err != nil {
		log.Fatalf("Failed to generate API key: %v", err)
	}

	log.Println(apiKey)
}
