package palsvm

import (
	"crypto/rand"
	"encoding/base64"
)

// RandomSecret generate random string
func RandomSecret(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
