package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashText возвращает SHA-256 хеш от строки текста
func HashText(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
