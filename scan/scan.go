package scan

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// String reads a single line from stdin and returns it as a string.
//
// Returns io.EOF when input is closed.
func String() (string, error) {
	op := "Scan.String"

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}
		return "", io.EOF
	}

	return scanner.Text(), nil
}

// Int reads a single line from stdin and returns it as an int.
//
// Returns io.EOF when input is closed.
func Int() (int, error) {
	op := "Scan.Int"

	str, err := String()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	result, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to convert into int: %w", op, err)
	}

	return result, nil
}

// Float reads a single line from stdin and returns it as a float64.
//
// Returns io.EOF when input is closed.
func Float() (float64, error) {
	op := "Scan.Float"

	str, err := String()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to convert into float: %w", op, err)
	}

	return result, nil
}
