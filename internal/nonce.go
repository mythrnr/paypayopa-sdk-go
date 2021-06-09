package internal

import (
	"math/rand"
	"time"
)

// nonce generates a random string with a specified number of characters.
//
// nonce は指定された文字数のランダムな文字列を生成する.
func nonce(n uint) string {
	if n == 0 {
		return ""
	}

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 0, n)

	for i := uint(0); i < n; i++ {
		b = append(b, charset[seed.Intn(len(charset))])
	}

	return string(b)
}
