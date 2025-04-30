package auth

import (
	"testing"
)

func TestHash(t *testing.T) {
	password := "BadPass123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash the password: %v", err)
	}

	if hash == "" {
		t.Fatal("Hash is empty")
	}

	if hash == password {
		t.Fatal("Hash is the same as password, hashing failed")
	}
}

func TestCheckHash(t *testing.T) {
	password := "right_password"
	wrongPass := "wrong_password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash the password: %v", err)
	}

	err = CheckPasswordHash(hash, password)
	if err != nil {
		t.Fatalf("Failed to verify the correct password: %v", err)
	}

	err = CheckPasswordHash(hash, wrongPass)
	if err == nil {
		t.Fatal("successfully failed with wrong password")
	}
}
