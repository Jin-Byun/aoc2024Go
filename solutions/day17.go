package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day17Pt1(res int) {
	fmt.Println("part 1:", res)
}

func day17Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day17() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input17.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	day17Pt1(len(lines))
	day17Pt2(len(lines))
}