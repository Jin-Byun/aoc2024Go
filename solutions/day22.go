package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day22Pt1(res int) {
	fmt.Println("part 1:", res)
}

func day22Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day22() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input22.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	day22Pt1(len(lines))
	day22Pt2(len(lines))
}