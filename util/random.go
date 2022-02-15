package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// generate random number between min and max
func RandomInit(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate random string of a length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// generate random owner with n character
func RandomOwner(n int) string {
	return RandomString(n)
}

// generate random money from 0 to 1000
func RandomMoney() int64 {
	return RandomInit(0, 1000)
}

// generate random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	k := len(currencies)
	return currencies[rand.Intn(k)]
}
