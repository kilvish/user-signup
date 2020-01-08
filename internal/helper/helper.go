package helper

import (
	"crypto/sha1"
	"encoding/hex"
)

//GetHash to get the hash of a string
func GetHash(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}
