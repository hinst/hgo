package hgo

import "crypto/rand"
import "encoding/base64"

func MakeRandomString(length int) string {
	data := make([]byte, length)
	_, readResult := rand.Read(data)
	AssertResult(readResult)
	text := base64.StdEncoding.EncodeToString(data)
	return text
}
