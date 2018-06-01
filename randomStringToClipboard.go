package main

import (
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()")

func main() {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 32)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	clipboard.WriteAll(string(b))
}