package random

import "math/rand"

func NewString(length int) string {
	return randomStringFromCharset(alphanumericalCharacters, length)
}

const alphanumericalCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomStringFromCharset(charset string, length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
