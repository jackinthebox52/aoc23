package utils

import (
	"fmt"
	"os"
	"time"
)

// Returns the contents of the input file for the given day, as a list of bytes ([]byte).
func ReadFileRaw(day int, file string) ([]byte, error) {
	path := fmt.Sprintf("d%d/%v.txt", day, file)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func FunctionTimer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
