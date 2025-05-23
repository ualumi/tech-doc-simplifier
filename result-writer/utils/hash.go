package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashText(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
