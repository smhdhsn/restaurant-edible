package random

import (
	"math/rand"
	"time"
)

// init will be called when this package is imported.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateDateBetween generates a random date between two given years.
func GenerateDateBetween(from, to int) time.Time {
	f := time.Date(from, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	t := time.Date(to, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	s := rand.Int63n(t-f) + f
	return time.Unix(s, 0)
}

// GenerateString generates a random string with given length.
func GenerateString(length int) string {
	a := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, length)

	for i := range b {
		b[i] = a[rand.Intn(len(a))]
	}

	return string(b)
}

// GenerateUint generates a random uint between two given values.
func GenerateUint32(from, to int) uint32 {
	s := rand.Intn(from-to) + from
	return uint32(s)
}

// GenerateBool generates a random bool.
func GenerateBool() bool {
	r := rand.Intn(3)
	return r%2 == 0
}
