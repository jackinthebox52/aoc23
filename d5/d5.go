package d5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackinthebox52/aoc23/utils"
)

type Range struct {
	dst_start int
	src_start int
	length    int
}

type Map []Range

func Solve1() int {
	lines, err := utils.ReadLines(5, "input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s, m := parseInput(lines)
	return calculateSmallestLocation(s, m)
}

func Solve2() int {
	lines, err := utils.ReadLines(5, "input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s, m := parseInput(lines)
	//Part 2
	seedPairs := [][]int{}
	for i := 0; i < len(s); i += 2 {
		seedPairs = append(seedPairs, s[i:i+2])
	}

	length := 0
	for _, p := range seedPairs {
		length += p[1]
	}
	s = make([]int, length)
	j := 0
	for _, p := range seedPairs {
		for i := p[0]; i < p[0]+p[1]; i++ {
			s[j] = i
			j++
		}
	}
	return calculateSmallestLocation(s, m)

}

func parseInput(lines []string) ([]int, []Map) {
	seeds := func() []int {
		ss := strings.Fields(lines[0][6:])
		i := []int{}
		for _, s := range ss {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				continue
			}
			i = append(i, n)
		}
		return i
	}()

	lines = lines[2:]

	maps := []Map{}
	mi := -1
	for _, line := range lines {
		if strings.Contains(line, "map:") {
			mi += 1
			maps = append(maps, Map{})
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		ds, _ := strconv.Atoi(fields[0])
		ss, _ := strconv.Atoi(fields[1])
		l, _ := strconv.Atoi(fields[2])
		maps[mi] = append(maps[mi], Range{
			dst_start: ds,
			src_start: ss,
			length:    l,
		})
	}

	return seeds, maps
}

func calculateSmallestLocation(seeds []int, maps []Map) int {
	smallest := -1
	for i := range seeds {
	Maps:
		for _, m := range maps {
			for _, r := range m {
				if seeds[i] >= r.src_start && seeds[i] <= r.src_start+r.length-1 {
					seeds[i] = seeds[i] - r.src_start + r.dst_start
					continue Maps
				}
			}
		}
		if smallest == -1 || seeds[i] < smallest {
			smallest = seeds[i]
		}
	}

	return smallest
}
