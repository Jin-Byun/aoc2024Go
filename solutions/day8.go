package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day8Pt1(antennas map[rune][][2]int, n int) {
	antinode := map[string]struct{}{}

	for _, list := range antennas {
		if len(list) < 2 { continue }
		for i := 0; i < len(list) - 1; i++ {
			p1 := list[i]
			for j := i + 1; j < len(list); j++ {
				dx, dy := p1[0] - list[j][0], p1[1] - list[j][1]
				if !utils.IdxInValid(p1[0]+dx, p1[1]+dy, n) { antinode[getCoord(p1[0]+dx, p1[1]+dy)] = struct{}{} }
				if !utils.IdxInValid(list[j][0]-dx, list[j][1]-dy, n) { antinode[getCoord(list[j][0]-dx, list[j][1]-dy)] = struct{}{} }
			}
		}
	}

	fmt.Println("part 1:", len(antinode))
}

func day8Pt2(antennas map[rune][][2]int, n int) {
	antinode := map[string]struct{}{}
	for _, list := range antennas {
		if len(list) < 2 { continue }
		antinode[getCoord(list[0][0], list[0][1])] = struct{}{}
		for i := 0; i < len(list) - 1; i++ {
			antinode[getCoord(list[i+1][0], list[i+1][1])] = struct{}{}
			for j := i + 1; j < len(list); j++ {
				x1, y1 := list[i][0], list[i][1]
				x2, y2 := list[j][0], list[j][1]
				dx, dy := x1 - x2, y1 - y2
				for !utils.IdxInValid(x1+dx, y1+dy, n) {
					x1, y1 = x1+dx, y1+dy
					antinode[getCoord(x1, y1)] = struct{}{}
				}
				for !utils.IdxInValid(x2-dx, y2-dy, n) {
					x2, y2 = x2-dx, y2-dy
					antinode[getCoord(x2, y2)] = struct{}{}
				}
			}
		}
	}
	fmt.Println("part 2:", len(antinode))
}

func Day8() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input8.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	n := len(lines)
	antennas := map[rune][][2]int{}
	for y := range lines {
		for x, r := range lines[y] {
			if r == '.' { continue }
			if antennas[r] == nil { antennas[r] = [][2]int{} }
			antennas[r] = append(antennas[r], [2]int{x, y} )
		}
	}
	
	
	
	day8Pt1(antennas, n)
	day8Pt2(antennas, n)
}