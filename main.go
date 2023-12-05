package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackinthebox52/aoc23/d1"
)

// Returns the contents of the input file for the given day
func readInputRaw(day int) (string, error) {
	filePath := filepath.Join("inputs", fmt.Sprintf("%d.txt", day))
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func main() {
	//mt.Printf("Day 1, pt 1 result: %d\n", d1.Solve1())
	fmt.Printf("Day 1, pt 2 result: %d\n", d1.Solve2())
}
