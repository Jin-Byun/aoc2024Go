package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func handleErr(e error) { if e != nil { panic(e) } }

func sToI(s string) int {
	v, _ := strconv.Atoi(s);
	return v
}

func pt1(data string) {
	res := 0
	addMul(&res, data)
	fmt.Println("part 1: ", res)
}

func addMul(res *int, s string) {
	re := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
	pairs := re.FindAllStringSubmatch(s, -1)
	for _, p := range pairs {
		numbers := strings.Split(p[1], ",")
		*res += sToI(numbers[0]) * sToI(numbers[1])
	}
}

func pt2(data string) {
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

func main() {
	f, err := os.ReadFile("day3/input.txt")
	handleErr(err)
	data := string(f)
	data = strings.ReplaceAll(data, "\n", "")
	pt1(data)
	pt2(data)
}