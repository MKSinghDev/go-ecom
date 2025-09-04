package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be not empty")
	}

	if hash == "password" {
		t.Errorf("expected hash to be different from password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, []byte("password")) {
		t.Errorf("expected password to match hash")
	}

	if ComparePassword(hash, []byte("not-password")) {
		t.Errorf("expected password to not match hash")
	}
}
