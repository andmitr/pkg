package random

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"
)

// MaxInt returns a random integer between 0 and math.MaxInt - 1.
func MaxInt() (int, error) {
	return Int(math.MaxInt - 1)
}

// Int returns a random integer between 0 and max (inclusive).
func Int(max int) (int, error) {
	if max < 0 {
		return 0, fmt.Errorf("random.Int: max must be non-negative")
	}
	random, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("random.Int: %w", err)
	}
	return int(random.Int64()), nil
}

// String generates a random string of the specified length.
//
// If specialChars is true, includes special characters in the output.
func String(length int, specialChars bool) (string, error) {
	op := "random.String"
	if length <= 0 {
		return "", fmt.Errorf("%s: length must be positive", op)
	}
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsExt := "!@#$%^&*()_+-=[]{}|;:'\"<>,.?/\\~`"
	buffer := new(strings.Builder)
	if specialChars {
		letters += charsExt
	}

	for range length {
		random, err := Int(len(letters) - 1)
		if err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}
		buffer.WriteByte(letters[random])
	}

	return buffer.String(), nil
}
