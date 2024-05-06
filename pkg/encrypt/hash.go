package encrypt

import (
	"crypto/sha256"
	"fmt"
)

func HashValue(val, secret string) string {
	hasher := sha256.New()
	hasher.Write([]byte(val))
	hashedValue := hasher.Sum([]byte(secret))

	return fmt.Sprintf("%x", hashedValue)
}
