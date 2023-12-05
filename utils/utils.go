package utils

import (
	"fmt"
	"os"
)

// Returns the contents of the input file for the given day, as a list of bytes ([]byte).
func ReadInputRaw(day int) ([]byte, error) {
	path := fmt.Sprintf("d%d/input.txt", day)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ReadTestRaw(day int) ([]byte, error) {
	path := fmt.Sprintf("d%d/test.txt", day)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}
