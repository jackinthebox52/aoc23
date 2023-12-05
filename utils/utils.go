package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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

func ReadLines(day int, file string) ([]string, error) {
	lines, err := ReadFileRaw(day, file)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	buf.Write(lines)
	var res []string
	for _, s := range strings.Split(buf.String(), "\n") {
		res = append(res, s)
	}
	return res, nil
}

func ReadString(day int, file string) (string, error) {
	lines, err := ReadFileRaw(5, file)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	buf.Write(lines)
	return buf.String(), nil
}

func FunctionTimer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
