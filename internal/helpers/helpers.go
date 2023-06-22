package helpers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func GenerateHmacSha1(data []byte, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(data)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func VerifyDigest(body []byte, digest, secretKey string) bool {
	hmacDigest := GenerateHmacSha1(body, secretKey)
	return hmac.Equal([]byte(digest), []byte(hmacDigest))
}
