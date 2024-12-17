package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func day15Pt1(mapAndMovement []string, tmp []string) {
	grid := make([][]byte, len(tmp)-2)
	for i := 1; i < len(tmp)-1; i++ {
		grid[i-1] = []byte(tmp[i][1:len(tmp[i])-1])
	}
	n := len(grid)
	x, y := startingPos(grid, '@')
	grid[y][x] = '.'
	for _, r := range mapAndMovement[1] {
		d := -1
		switch r {
		case '>':
			d = 0
		case 'v':
			d = 1
		case '<':
			d = 2
		case '^':
			d = 3
		}
		if d == -1 { continue }
		dy, dx := utils.MoveCardinal[d], utils.MoveCardinal[d+1]
		if utils.IdxInValid(x+dx, y+dy, n) || grid[y+dy][x+dx] == '#' { continue }
		i := 1
		for grid[y+dy*i][x+dx*i] == 'O' {
			i++
			if utils.IdxInValid(x+dx*i, y+dy*i, n) || grid[y+dy*i][x+dx*i] == '#' {
				i = 0
				break
			}
		}
		if i > 1 {
			grid[y+dy*i][x+dx*i] = 'O'
			i = 1
		}
		y, x = y+dy*i, x+dx*i
		grid[y][x] = '.'
	}

	sum := 0
	for r := range grid {
		for c, v := range grid[r] {
			if v != 'O' {continue}
			sum += (r+1)*100 + c+1
		}
	}
	fmt.Println("part 1:", sum)
}

func day15Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day15() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input15.txt")
	utils.HandleErr(err)
	mapAndMovement := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	tmp := strings.Fields(mapAndMovement[0])
	day15Pt1(mapAndMovement, tmp)

	n := len(tmp)
	grid := make([][]byte, n)
	x, y := 0, 0
	// part 2
	for i, row := range tmp {
		grid[i] = make([]byte, n*2)
		for c := range row {
			next1, next2 := row[c], row[c]
			switch next1 {
			case '@':
				next1, next2 = '.', '.'
				x, y = c*2, i
			case 'O':
				next1, next2 = '[', ']'
			}
			grid[i][c*2], grid[i][c*2+1] = next1, next2
		}
	}
	var check func(x, y, dy int) bool
	check = func(x, y, dy int) bool {
		if utils.IdxInValid2(x, y, n*2, n) || grid[y][x] == '#' { return false }
		if grid[y][x] == '.' { return true }
		var next, otherNext bool
		sideX := x
		if grid[y][x] == '[' {
			sideX++
			} else {
			sideX--
			}
		next = check(x, y+dy, dy)
		otherNext = check(sideX, y+dy, dy)
		return next && otherNext
	}
	var moveBox func(x, y, dy int)
	moveBox = func(x, y, dy int) {
		if grid[y][x] == '.' { return }
		sideX := x
		this := grid[y][x]
		var other byte
		if this == '[' {
			other = ']'
			sideX++
			} else {
			other = '['
			sideX--
			}
		moveBox(x, y+dy, dy)
		moveBox(sideX, y+dy, dy)
		grid[y+dy][x] = this
		grid[y+dy][sideX] = other
		grid[y][sideX] = '.'
	}
	for _, r := range mapAndMovement[1] {
		d := -1
		switch r {
		case '>':
			d = 0
		case 'v':
			d = 1
		case '<':
			d = 2
		case '^':
			d = 3
		}
		if d == -1 { continue }
		dy, dx := utils.MoveCardinal[d], utils.MoveCardinal[d+1]
		if utils.IdxInValid2(x+dx, y+dy, n*2, n) || grid[y+dy][x+dx] == '#' { continue }
		if grid[y+dy][x+dx] == '.' {
			x, y = x+dx, y+dy
			continue
		}
		// horizontal
		if d & 1 == 0 {
			i := 1
			for grid[y][x+dx*i] == grid[y][x+dx] {
				i += 2
				if utils.IdxInValid2(x+dx*i, y, n*2, n) || grid[y][x+dx*i] == '#' {
					i = 0
					break
				}
			}
			if i == 0 { continue }
			for i > 0 {
				grid[y][x+dx*i] = grid[y][x+dx*(i-1)]
				i--
			}
			x = x+dx
			continue
		}
		fmt.Println(x, y)
		if check(x, y+dy, dy) {
			moveBox(x, y+dy, dy)
			y = y+dy
			grid[y][x] = '.'
		}
	}
	for _, l := range grid {
    fmt.Println(string(l))
	}

	sum := 0
	for r := range grid {
		for c, v := range grid[r] {
			if v != '[' {continue}
			sum += r*100 + c
		}
	}
	day15Pt2(sum)
}