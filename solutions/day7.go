package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func moveDigitUp(acc, i int) int {
	for range i {
		acc *= 10
	}
	return acc 
}

func day7Pt1(s []string, exp, acc, i int) bool {
	if i == len(s) {
		return acc == exp
	}
	v := utils.StrToI(s[i])
	return day7Pt1(s, exp, acc + v, i+1) || day7Pt1(s, exp, acc * v, i+1)
}

func day7Pt2(s []string, exp, acc, i int) bool {
	if i == len(s) {
		return acc == exp
	}
	v := utils.StrToI(s[i])
	return day7Pt2(s, exp, acc + v, i+1) || day7Pt2(s, exp, acc * v, i+1) || day7Pt2(s, exp, moveDigitUp(acc, len(s[i])) + v, i+1)
}

func Day7() {
	f, err := os.ReadFile("inputs/input7.txt")
	utils.HandleErr(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	sum1 := 0
	sum2 := 0
	for _, l := range lines {
		calibration := strings.Split(l, ": ")
		expected, testValues := utils.StrToI(calibration[0]), strings.Fields(calibration[1])
		if day7Pt1(testValues, expected, 0, 0) {
			sum1 += expected
		} else if day7Pt2(testValues, expected, 0, 0) {
			sum2 += expected
		}
	}
	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum1 + sum2)
}