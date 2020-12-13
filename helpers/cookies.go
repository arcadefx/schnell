package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// GetCookie returns a base64 encoded string, idea here not fully implemented
func GetCookie(data string, token string, browserID string) string {
	// note: this secretKey should be from a config or environment variable
	secretKey := "OreosAndMilk"

	key := []byte(secretKey)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(data + token + browserID))
	base64Data := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return base64Data
}
