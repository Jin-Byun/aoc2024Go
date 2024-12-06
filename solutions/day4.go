package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

var dir1 [5]int = [5]int{0, 1, 0, -1, 0}
var dir2 [5]int = [5]int{-1, 1, 1, -1, -1}

var mas [3]byte = [3]byte{'M', 'A', 'S'};

func idxInValid(r, c int) bool { return r < 0 || c < 0 || r == 140 || c == 140 }

func checkDirection1(r, c, i int, puzzle []string) int {
	for j := range 3 {
		r, c = r+dir1[i], c+dir1[i+1]
		if idxInValid(r, c) { return 0 }
		if puzzle[r][c] != mas[j] { return 0 }
	}
	return 1
}
func checkDirection2(r, c, i int, puzzle []string) int {
	for j := range 3 {
		r, c = r+dir2[i], c+dir2[i+1]
		if idxInValid(r, c) { return 0 }
		if puzzle[r][c] != mas[j] { return 0 }
	}
	return 1
}

func day4Pt1(lines []string) {
	res := 0

	for r := range lines {
		for c := range lines {
			if lines[r][c] != 'X' { continue }
			for i := range 4 {
				res += checkDirection1(r, c, i, lines)
				res += checkDirection2(r, c, i, lines)
			}
		}
	}

	fmt.Println("part 1: ", res)
}

func crossCheck(r, c int, puzzle []string) int {
	if idxInValid(r+1, c+1) || idxInValid(r-1, c-1) { return 0 }
	switch puzzle[r-1][c-1] {
		case 'M':
			if puzzle[r+1][c+1] != 'S' { return 0 }
		case 'S':
			if puzzle[r+1][c+1] != 'M' { return 0 }
		default:
			return 0
	}
	switch puzzle[r+1][c-1] {
		case 'M':
			if puzzle[r-1][c+1] != 'S' { return 0 }
		case 'S':
			if puzzle[r-1][c+1] != 'M' { return 0 }
		default:
			return 0
	}
	return 1
}

func day4Pt2(lines []string) {
	res := 0
	
	for r := range lines {
		for c := range lines {
			if lines[r][c] != 'A' { continue }
			res += crossCheck(r, c, lines)
		}
	}
	
	fmt.Println("part 2: ", res)
}

func Day4() {
	f, err := os.ReadFile("day4/input.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	day4Pt1(lines)
	day4Pt2(lines)
}