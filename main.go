package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789!@#$%*")

func main() {
	hexFlag := flag.Bool("hex", false, "Should the random value be valid hexadecimal?")
	flag.Parse()
	if *hexFlag {
		chars = []rune("abcdef0123456789")
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 32)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	err := clipboard.WriteAll(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
