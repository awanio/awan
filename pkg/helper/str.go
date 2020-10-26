package helper

import (
	"crypto/rand"
	"fmt"
	"path/filepath"
	"runtime"
)

var basepath string

// FromBasepath return to main dir
func FromBasepath(file string) string {
	return fmt.Sprintf("%s/../../%s", basepath, file)
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(b)
}

// GenerateRandomString func
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// GenerateRandomNumber func
func GenerateRandomNumber(n int) (string, error) {
	const letters = "0123456789"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// GenerateRandomBytes func
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
