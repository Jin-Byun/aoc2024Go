package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func setupVisited(n int) [][]bool {
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, n)
	}
	return visited
}

func day12Pt1(lines []string) {
	n := len(lines)
	visited := setupVisited(n)

	var dfs func(int, int, byte) (int, int)
	dfs = func(y, x int, target byte) (int, int) {
		if utils.IdxInValid(x, y, n) || lines[y][x] != target {
			return 0, 1
		}
		if visited[y][x] { return 0, 0 }
		visited[y][x] = true
		area, perimeter := 1, 0
		for i := range 4 {
			dA, dP := dfs(y + utils.MoveCardinal[i], x + utils.MoveCardinal[i+1], target)
			area += dA
			perimeter += dP
		}
		return area, perimeter
	}

	res := 0
	for r := range n {
		for c := range n {
			if visited[r][c] { continue }
			a, p := dfs(r, c, lines[r][c])
			res += a * p
		}
	}

	fmt.Println("part 1:", res)
}

func day12Pt2(lines []string) {
	n := len(lines)
	visited := setupVisited(n)

	fence := [][3]int{}
	var dfs func(int, int, int, byte) int
	dfs = func(y, x, direction int, target byte) int {
		if utils.IdxInValid(x, y, n) || lines[y][x] != target {
			fence = append(fence, [3]int{x, y, direction})
			return 0
		}
		if visited[y][x] { return 0 }
		visited[y][x] = true
		area := 1
		for i := range 4 {
			area += dfs(y + utils.MoveCardinal[i], x + utils.MoveCardinal[i+1], i, target)
		}
		return area
	}

	sortFence := func(a, b [3]int) int {
		if a[2] != b[2] { return a[2] - b[2] }
		if a[2] & 1 == 1 {
			if a[1] != b[1] { return a[1] - b[1] }
			return a[0] - b[0]
		}
		if a[0] != b[0] { return a[0] - b[0] }
		return a[1] - b[1]
	}
	countSide := func() int {
		sides := 1
		for i := 1; i < len(fence); i++ {
			prev, curr := fence[i-1], fence[i]

			if prev[2] != curr[2] {
				sides++
			} else if curr[2] & 1 == 0 {
				if curr[0] != prev[0] || curr[1] - prev[1] > 1 {
					sides++
				}
			} else if curr[1] != prev[1] || curr[0] - prev[0] > 1 {
				sides++
			}
		}
		return sides
	}

	res := 0
	for r := range n {
		for c := range n {
			if visited[r][c] { continue }
			area := dfs(r, c, 0, lines[r][c])
			slices.SortFunc(fence, sortFence)
			sides := countSide()
			res += area * sides
			fence = nil
	}
	}

	fmt.Println("part 2:", res)
}

func Day12() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input12.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	
	day12Pt1(lines)
	day12Pt2(lines)
}