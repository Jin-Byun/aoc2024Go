package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

var N6 int = 130

func day6Pt1(res int) {
	fmt.Println("part 1: ", res)
}

func day6Pt2(res int) {
	fmt.Println("part 2: ", res)
}

func startingPos(grid [][]byte, target byte) (int, int) {
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == target { 
				return c, r
			}
		}
	}
	return -1, -1
}

func getState(x, y, d int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, d)
}
func getPos(coord string) (int, int) {
	pos := strings.Split(coord, ",")
	return utils.StrToI(pos[0]), utils.StrToI(pos[1])
}

func travel(hash map[string]bool, grid [][]byte, x, y int) {
	d := 0
	y--
	for !utils.IdxInValid(y, x, N6) {
		switch grid[y][x] {
		case '#':
			x, y = x - utils.MoveCardinal[d], y + utils.MoveCardinal[d + 1]
			d = (d + 1) % 4
			x, y = x + utils.MoveCardinal[d], y - utils.MoveCardinal[d + 1]
		default:
			hash[utils.GetCoord(x, y)] = true
			x, y = x + utils.MoveCardinal[d], y - utils.MoveCardinal[d + 1]
		}
	}
}

func simulation(grid [][]byte, x, y int) int {
	d := 0
	visited := map[string]struct{}{getState(x, y, d): {}}
	for {
		x, y = x + utils.MoveCardinal[d], y - utils.MoveCardinal[d + 1]
		if utils.IdxInValid(y, x, N6) { return 0 }
		if grid[y][x] == '#' {
			x, y = x - utils.MoveCardinal[d], y + utils.MoveCardinal[d + 1]
			d = (d + 1) % 4
		}
		newState := getState(x, y, d)
		if _, exists := visited[newState]; exists {
			return 1
		}
		visited[newState] = struct{}{}
	}
}

func simulateBlock(hash map[string]bool, grid [][]byte, x, y int, acc *int) {
	for k := range hash {
		c, r := getPos(k)
		grid[r][c] = '#'
		*acc += simulation(grid, x, y)
		grid[r][c] = '.'
	}
}

func Day6() {
	f, err := os.ReadFile("inputs/input6.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	grid := make([][]byte, N6)
	for i := range N6 {
		grid[i] = []byte(lines[i])
	}
	x, y := startingPos(grid, '^')
	hash := map[string]bool{}
	travel(hash, grid, x, y)
	res := 0
	simulateBlock(hash, grid, x, y, &res)

	// Me trying to solve it like a node-graph problem
	// for c, v := range hash {
	// 	for i := range 4 {
	// 		if (v >> i) & 1 == 0 { continue }
	// 		target := (i+3) % 4
	// 		dx, dy := getPos(c, i)
	// 		for !utils.IdxInValid(dx, dy, N6) && grid[dy][dx] != '#' {
	// 			if (hash[utils.GetCoord(dx, dy)] >> target) & 1 == 1 && order[utils.GetCoord(dx, dy)] > order[c] {
	// 				dx, dy = dx + utils.MoveCardinal[target], dy - utils.MoveCardinal[target+1]
	// 				if grid[dy][dx] != '#' {
	// 					obs[utils.GetCoord(dx, dy)] = true
	// 				}
	// 				dx, dy = dx - utils.MoveCardinal[target], dy + utils.MoveCardinal[target+1]
	// 			}
	// 			dx, dy = dx - utils.MoveCardinal[i], dy + utils.MoveCardinal[i+1]
	// 		}
	// 	}
	// }
	day6Pt1(len(hash) + 1)
	day6Pt2(res)
}