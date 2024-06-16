package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		cs := alphabet[rand.Intn(k)]
		sb.WriteByte(cs)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6) + "gmail.com"
}

func RandomPassWordHash() string {
	return RandomString(10)
}
