package main

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = bcrypt.DefaultCost

func encodePassword(plain string, cost int) (string, error) {
	if plain == "" {
		return "", errors.New("plain text must not be empty")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
	if err != nil {
		return "", fmt.Errorf("failed to generate bcrypt hash: %w", err)
	}

	return string(hashed), nil
}

func run(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: password-encoder <plain-text>")
	}

	hashed, err := encodePassword(args[1], DefaultCost)
	if err != nil {
		return err
	}

	fmt.Println(hashed)
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
