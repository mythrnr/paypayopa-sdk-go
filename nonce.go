package paypayopa

import (
	"math/rand"
	"time"
)

// Nonce generates a random string with a specified number of characters.
//
// Nonce は指定された文字数のランダムな文字列を生成する.
func Nonce(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		b = append(b, charset[seed.Intn(len(charset))])
	}

	return string(b)
}
