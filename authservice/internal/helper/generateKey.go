package helper

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateSecureKey() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal("Failed to generate secure API key")
	}
	return hex.EncodeToString(bytes)
}
