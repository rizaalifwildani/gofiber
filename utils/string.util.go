package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID(text string) uuid.UUID {
	parsedUUID, parsedErr := uuid.Parse(text)
	if parsedErr != nil {
		return uuid.UUID{}
	}
	return parsedUUID
}

func GenerateSlug(text string) string {
	return strings.ReplaceAll(strings.ToLower(text), " ", "-")
}

func GeneratePassword(text string) (string, error) {
	// Create a new SHA-512 hash
	hasher := sha512.New()

	// Write the combined string to the hash
	_, err := hasher.Write([]byte(text))
	if err != nil {
		return "", err
	}

	// Get the final hash and convert it to a hexadecimal string
	hashedBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashedBytes)

	return hashString, nil
}
