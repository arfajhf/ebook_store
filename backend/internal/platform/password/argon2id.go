package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	memory      uint32 = 19 * 1024
	iterations  uint32 = 2
	parallelism uint8  = 1
	saltLength         = 16
	keyLength   uint32 = 32
)

var ErrInvalidHash = errors.New("invalid password hash")

func Hash(plainPassword string) (string, error) {
	salt := make([]byte, saltLength)

	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate password salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(plainPassword),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength,
	)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	encodedPassword := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		iterations,
		parallelism,
		encodedSalt,
		encodedHash,
	)

	return encodedPassword, nil
}

func Verify(
	plainPassword string,
	encodedPassword string,
) (bool, error) {
	parts := strings.Split(encodedPassword, "$")

	if len(parts) != 6 || parts[1] != "argon2id" {
		return false, ErrInvalidHash
	}

	var version int

	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, ErrInvalidHash
	}

	if version != argon2.Version {
		return false, ErrInvalidHash
	}

	var parsedMemory uint32
	var parsedIterations uint32
	var parsedParallelism uint8

	if _, err := fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&parsedMemory,
		&parsedIterations,
		&parsedParallelism,
	); err != nil {
		return false, ErrInvalidHash
	}

	if parsedMemory == 0 ||
		parsedIterations == 0 ||
		parsedParallelism == 0 {
		return false, ErrInvalidHash
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, ErrInvalidHash
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil || len(expectedHash) == 0 {
		return false, ErrInvalidHash
	}

	actualHash := argon2.IDKey(
		[]byte(plainPassword),
		salt,
		parsedIterations,
		parsedMemory,
		parsedParallelism,
		uint32(len(expectedHash)),
	)

	matched := subtle.ConstantTimeCompare(
		expectedHash,
		actualHash,
	) == 1

	return matched, nil
}
