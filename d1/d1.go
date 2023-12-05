package d1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Solve1() int {
	file, err := os.Open("d1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(scanner.Text(), -1)
		result := numbers[0] + numbers[len(numbers)-1]
		cplus, err := strconv.Atoi(result)
		count += cplus
		if err != nil {
			log.Fatal(err)
		}
		//println(scanner.Text() + " " + result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count
}

var _IntStringMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func Solve2() int {
	file, err := os.Open("d1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		re := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
		numbers := re.FindAllString(scanner.Text(), -1)

		n1 := numbers[0]
		n2 := numbers[len(numbers)-1]
		var i1 int
		var i2 int
		var str_error error
		if i1, str_error = strconv.Atoi(n1); str_error != nil {
			i1 = _IntStringMap[n1]
		}
		if i2, str_error = strconv.Atoi(n2); str_error != nil {
			i2 = _IntStringMap[n2]
		}
		result, err := strconv.Atoi(fmt.Sprintf("%d%d", i1, i2))
		if err != nil {
			panic(err)
		}
		count += result
		fmt.Println(numbers)
		println(scanner.Text() + " " + fmt.Sprint(result))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count
}
