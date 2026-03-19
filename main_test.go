package main

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestEncodePassword_ReturnsBcryptHash(t *testing.T) {
	plain := "password"

	hashed, err := encodePassword(plain, bcrypt.MinCost)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if hashed == plain {
		t.Fatal("hashed password must not equal plain text")
	}

	if !strings.HasPrefix(hashed, "$2") {
		t.Fatalf("expected bcrypt hash, got: %s", hashed)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)); err != nil {
		t.Fatalf("hash does not match plain text: %v", err)
	}
}

func TestEncodePassword_EmptyStringReturnsError(t *testing.T) {
	_, err := encodePassword("", bcrypt.MinCost)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRun_WithoutArgumentReturnsUsageError(t *testing.T) {
	err := run([]string{"password-encoder"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "usage:") {
		t.Fatalf("expected usage error, got: %v", err)
	}
}
