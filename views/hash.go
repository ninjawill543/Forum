package forum

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(password string) string {
	hash := sha1.New()
	hashInBytes := hash.Sum([]byte(password))[:20]
	return hex.EncodeToString(hashInBytes)
	//encoding passwords in sha1
}
