package main

import (
	"fmt"
	"log"

	"github.com/jackinthebox52/aoc23/d1"
	"github.com/jackinthebox52/aoc23/d2"
	"github.com/jackinthebox52/aoc23/d4"
	"github.com/jackinthebox52/aoc23/d5"
)

func SolveDay1() {
	fmt.Println("Day 1, pt 1 result: ", d1.Solve1())
	fmt.Println("Day 1, pt 2 result: ", d1.Solve2())
}

func SolveDay4() {
	fmt.Printf("Day 4, pt 1 result: %d\n", d4.Solve1())
}

func SolveDay5() {
	log.Printf("Day 5, pt 1 result: %d\n", d5.Solve1())
	log.Printf("Day 5, pt 2 result: %d\n", d5.Solve2())
}

func SolveDay2() {
	p1, p2 := d2.Solve1and2()
	fmt.Printf("Day 2, pt 1 result: %d, pt 2 result: %d\n", p1, p2)
}

func main() {
	SolveDay2()
}
