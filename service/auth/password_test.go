package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing passworf: %v", err)
	}
	if hash == "" {
		t.Error("expected hash is empty")
	}
	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing passworf: %v", err)
	}
	if !ComparePassword(hash, []byte("password")) {
		t.Error("expected password does not match hash")
	}
	if ComparePassword(hash, []byte("notpassword")) {
		t.Error("expected password does not match hash")
	}
}
