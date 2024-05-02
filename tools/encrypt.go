package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	flag.Parse()

	fmt.Println("word:", *wordPtr)

	hashedString, _ := bcrypt.GenerateFromPassword([]byte(*wordPtr), 8)

	fmt.Println("hash:", string(hashedString))

}
