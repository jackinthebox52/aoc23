package main

import (
	"fmt"

	"github.com/jackinthebox52/aoc23/d1"
	"github.com/jackinthebox52/aoc23/d4"
)

func SolveDay1() {
	fmt.Println("Day 1, pt 1 result: ", d1.Solve1())
	fmt.Println("Day 1, pt 2 result: ", d1.Solve2())
}

func main() {
	fmt.Printf("Day 4, pt 1 result: %d\n", d4.Solve1())
}
