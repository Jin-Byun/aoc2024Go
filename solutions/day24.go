package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day24Pt1(res int) {
	fmt.Println("part 1:", res)
}

func day24Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day24() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input24.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	day24Pt1(len(lines))
	day24Pt2(len(lines))
}