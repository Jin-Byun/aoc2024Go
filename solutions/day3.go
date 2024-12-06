package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func day3Pt1(data string) {
	res := 0
	addMul(&res, data)
	fmt.Println("part 1: ", res)
}

func addMul(res *int, s string) {
	re := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
	pairs := re.FindAllStringSubmatch(s, -1)
	for _, p := range pairs {
		numbers := strings.Split(p[1], ",")
		*res += utils.StrToI(numbers[0]) * utils.StrToI(numbers[1])
	}
}

func day3Pt2(data string) {
	start := strings.Index(data, "don't()")
	res := 0
	addMul(&res, data[:start])
	re := regexp.MustCompile(`do\(\).*?don't\(\)`)
	enabledLines := re.FindAllString(data, -1)
	for _, line := range enabledLines {
		addMul(&res, line)
	}
	fmt.Println("part 2: ", res)
}

func Day3() {
	f, err := os.ReadFile("day3/input.txt")
	utils.HandleErr(err)
	data := string(f)
	data = strings.ReplaceAll(data, "\n", "")
	day3Pt1(data)
	day3Pt2(data)
}