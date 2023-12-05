package d4

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/jackinthebox52/aoc23/utils"
)

//	https://adventofcode.com/2023/day/4  --- Day 4: Scratchcards ---

func Solve1() int {
	points := 0

	b, err := utils.ReadInputRaw(4)
	if err != nil {
		fmt.Println(errors.New("couldn't read input file"))
		panic(err)
	}

	line_list := bytes.Split(b, []byte("\n"))
	for _, line := range line_list {
		line = line[7:]
		a := bytes.Split(line, []byte(" | "))
		win_list := strings.Split(string(a[0]), " ")[1:]
		my_list := strings.Split(string(a[1]), " ")
		count := 0
		for _, win := range win_list {
			for _, my := range my_list {
				if win == my {
					if my == "" {
						continue
					}
					if count == 0 {
						count = 1
					} else {
						count = count * 2
					}
				}
			}
		}
		points += count
		//fmt.Printf("Card no. %d is worth %d points.\n", card+1, count)

	}
	return points
}
