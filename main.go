package main

import (
	"fmt"
	"os"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Augument Error. param is should single string\n")
		os.Exit(1)
	}

	hash := os.Args[1]
	password := os.Args[2]

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "ok\n", err.Error())

	os.Exit(0)
}
