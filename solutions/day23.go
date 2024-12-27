package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day23Pt1(res int) {
	fmt.Println("part 1:", res)
}

func day23Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day23() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input23.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	day23Pt1(len(lines))
	day23Pt2(len(lines))
}