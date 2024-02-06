package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Returns a random int64 number between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0 -> max-min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] // gerneate a random position from 0 to k 0 1
		sb.WriteByte(c)             // write the byte to the string builder
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "TWD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
