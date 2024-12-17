package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day16Pt1(res int) {
	fmt.Println("part 1:", res)
}

func day16Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day16() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input16.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	day16Pt1(len(lines))
	day16Pt2(len(lines))
}