package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day20Pt1(lines []string, start, goal [2]int) {
	res := 0
	var timeLimit int
	var traverse func(int, int, [2]int)
	traverse = func(time, direction int, pos [2]int) {
		if pos == goal {
			res++
			return
		}
		x, y := pos[0], pos[1]
		if time == timeLimit || lines[y][x] == '#' {
			return
		}
		for i := range 4 {
			if i == 2 { continue }
			d := (i+direction) % 4
			dx, dy := x+utils.MoveCardinal[d+1], y+utils.MoveCardinal[d]
			traverse(time+1, d, [2]int{dx,dy})
		}
	}
	var traverseWithCheat func(t,c,d int, pos [2]int)
	traverseWithCheat = func(time, cheat, direction int, pos [2]int) {
		if time == timeLimit { return }
		x, y := pos[0], pos[1]
		if utils.IdxInValid(x, y, len(lines)) { return }
		if lines[y][x] == '#' {
			cheat--
		}
		if cheat == 0 { return }
		if cheat < 2 && lines[y][x] != '#' {
				traverse(time, direction, pos)
				return
		}
		for i := range 4 {
			if i == 2 { continue }
			d := (direction + i) % 4
			dx, dy := x+utils.MoveCardinal[d+1], y+utils.MoveCardinal[d]
			traverseWithCheat(time+1, cheat, d, [2]int{dx,dy})
		}
	}
	timeLimit = aStarPath(start, goal, lines, len(lines)) - 100
	traverseWithCheat(0, 2, 0, start)
	fmt.Println("part 1:", res)
}

func day20Pt2(lines []string, start, goal [2]int) {
	path := getPath(start, goal, lines)
	res := 0
	for i := range path {
		for j := i+1; j < len(path); j++ {
			md := heuristic(path[i], path[j])
			if md <= 20 && j-i-md >= 100 {
				res++
			}
		}
	}
	fmt.Println("part 2:", res)
}

func Day20() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input20.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	start, goal := startAndGoal(lines)
	day20Pt1(lines, start, goal)
	day20Pt2(lines, start, goal)
}

func getPath[V string | []byte](pos, goal [2]int, grid []V) [][2]int {
	path := [][2]int{pos}
	prev := [2]int{-2, -2}
	for pos != goal {
		for d := range 4 {
			np := [2]int{pos[0] + utils.MoveCardinal[d+1], pos[1]+utils.MoveCardinal[d]}
			if np != prev && grid[np[1]][np[0]] != '#' {
				prev = pos
				pos = np
				path = append(path, pos)
				break
			}
		}
	}
	return path
}