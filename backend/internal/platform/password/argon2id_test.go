package password

import (
	"strings"
	"testing"
)

func TestHashAndVerify(t *testing.T) {
	plainPassword := "BelajarGo123!"

	encodedPassword, err := Hash(plainPassword)
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	if strings.Contains(encodedPassword, plainPassword) {
		t.Fatal("hash contains plain password")
	}

	matched, err := Verify(plainPassword, encodedPassword)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}

	if !matched {
		t.Fatal("Verify() matched = false, want true")
	}
}

func TestVerifyRejectsWrongPassword(t *testing.T) {
	encodedPassword, err := Hash("PasswordBenar123!")
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	matched, err := Verify("PasswordSalah123!", encodedPassword)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}

	if matched {
		t.Fatal("Verify() matched = true, want false")
	}
}

func TestHashUsesUniqueSalt(t *testing.T) {
	firstHash, err := Hash("PasswordSama123!")
	if err != nil {
		t.Fatalf("first Hash() error = %v", err)
	}

	secondHash, err := Hash("PasswordSama123!")
	if err != nil {
		t.Fatalf("second Hash() error = %v", err)
	}

	if firstHash == secondHash {
		t.Fatal("two hashes are identical")
	}
}
