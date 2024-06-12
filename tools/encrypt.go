package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	flag.Parse()

	fmt.Println("word: " + *wordPtr) //nolint:forbidigo

	hashedString, _ := bcrypt.GenerateFromPassword([]byte(*wordPtr), bcrypt.DefaultCost)

	fmt.Println("hash: " + string(hashedString)) //nolint:forbidigo
}
