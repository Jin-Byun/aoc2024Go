package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day10Pt1(lines []string, n int) {
	res := 0
	for r := range n {
		for c, b := range lines[r] {
			if b == '0' {
				dst := map[string]struct{}{}
				countTrail1(lines, dst, r, c, n)
				res += len(dst)
			}
		}
	}
	fmt.Println("part 1: ", res)
}


func day10Pt2(res int) {
	fmt.Println("part 2: ", res)
}

func countTrail1(m []string, h map[string]struct{}, r, c, n int) {
	if m[r][c] == '9' {
		h[utils.GetCoord(c, r)] = struct{}{}
		return
	}
	for i := range 4 {
		dy, dx := r + utils.MoveCardinal[i], c + utils.MoveCardinal[i+1]
		if utils.IdxInValid(dy, dx, n) { continue }
		if m[dy][dx] != m[r][c] + 1 { continue }
		countTrail1(m, h, dy, dx, n)
	}
}

func countTrail2(m []string, acc *int, r, c, n int) {
	if m[r][c] == '9' {
		*acc++
		return
	}
	for i := range 4 {
		dy, dx := r + utils.MoveCardinal[i], c + utils.MoveCardinal[i+1]
		if utils.IdxInValid(dy, dx, n) { continue }
		if m[dy][dx] != m[r][c] + 1 { continue }
		countTrail2(m, acc, dy, dx, n)
	}
}

func Day10() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input10.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	n := len(lines)
	res := 0
	for r := range n {
		for c, b := range lines[r] {
			if b == '0' {
				countTrail2(lines, &res, r, c, n)
			}
		}
	}
	day10Pt1(lines, n)
	day10Pt2(res)
}