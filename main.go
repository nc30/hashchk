package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Augument Error. param is should hash and password string\n")
		fmt.Fprintf(os.Stderr, "ex) hashchk '$2a$10$B.5V561rIKXKG4T1R11Z6.3d355ys4p.ap1ChUIgDzNVz8nx8hJ7i' test \n")
		os.Exit(1)
	}

	hash := os.Args[1]
	password := os.Args[2]

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "ok\n")

	os.Exit(0)
}
