package d2

import (
	"regexp"
	"slices"
	"strconv"

	"github.com/jackinthebox52/aoc23/utils"
)

func Solve1and2() (int, int) {
	input, _ := utils.ReadLines(2, "test")
	re := regexp.MustCompile(`(\d+) (\w+)`)
	p1, p2 := 0, 0
	for i, s := range input {
		min := map[string]int{}
		for _, c := range re.FindAllStringSubmatch(s, -1) {
			n, _ := strconv.Atoi(c[1])
			min[c[2]] = slices.Max([]int{min[c[2]], n})
		}

		if min["red"] <= 12 && min["green"] <= 13 && min["blue"] <= 14 {
			p1 += i + 1
		}
		p2 += min["red"] * min["green"] * min["blue"]
	}
	return p1, p2
}
