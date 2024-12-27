package solutions

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+	+---+---+
// | 4 | 5 | 6 |	| ^3 | A | <--- start here x2
// +---+---+---+	+---+---+---+
// | 1 | 2 | 3 |	| <2 | v1 | >0 |
// +---+---+---+	+---+---+---+
//     | 0 | A | <--- start here
//     +---+---+

var numpad []string = []string{"789", "456", "123", " 0A"}
var dpad []string = []string{" ^A", "<v>"}

type CodePhase struct {
	code string
	phase int
}

var numCache map[[4]int][]string
func getNumPaths(prev, curr [2]int) []string {
	if v, exists := numCache[[4]int{prev[0], prev[1], curr[0], curr[1]}]; exists {
		return v
	}
	paths := toPaths(numpad, prev, curr)
	dirs := []string{}
	for _, path := range paths {
		dirs = append(dirs, toDirection(path))
	}
	numCache[[4]int{prev[0], prev[1], curr[0], curr[1]}] = dirs
	return dirs
}

var dCache map[[4]int][]string
func getDPaths(prev, curr [2]int) []string {
	if v, exists := dCache[[4]int{prev[0], prev[1], curr[0], curr[1]}]; exists {
		return v
	}
	paths := toPaths(dpad, prev, curr)
	dirs := []string{}
	for _, path := range paths {
		dirs = append(dirs, toDirection(path))
	}
	numCache[[4]int{prev[0], prev[1], curr[0], curr[1]}] = dirs
	return dirs
}

func toDirection(p [][2]int) string {
	var sb strings.Builder
	for i := 1; i < len(p); i++ {
		r1, c1 := p[i-1][0], p[i-1][1]
		r2, c2 := p[i][0], p[i][1]
		if r1 == r2 {
			if c2 > c1 {
				sb.WriteByte('>')
			} else {
				sb.WriteByte('<')
			}
		} else {
			if r2 > r1 {
				sb.WriteByte('v')
			} else {
				sb.WriteByte('^')
			}
		}
	}
	sb.WriteByte('A')
	return sb.String()
}

func toPaths(grid []string, start, end [2]int) [][][2]int {
	h, w := len(grid), len(grid[0])
	dist := make([][]int, h)
	pred := make([][][][2]int, h)
	for r := range h {
		dist[r] = make([]int, w)
		pred[r] = make([][][2]int, w)
		for c := range w {
			dist[r][c] = math.MaxInt32
		}
	}
	q := [][2]int{start}
	dist[start[0]][start[1]] = 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		r, c := curr[0], curr[1]
		d := dist[r][c]
		for direction := range 4 {
			nr, nc := r + utils.MoveCardinal[direction], c + utils.MoveCardinal[direction+1]
			if utils.IdxInValid2(nr, nc, h, w) || grid[nr][nc] == ' ' { continue }
			nd := d+1
			if nd < dist[nr][nc] {
				dist[nr][nc] = nd
				pred[nr][nc] = [][2]int{curr}
				q = append(q, [2]int{nr, nc})
			} else if nd == dist[nr][nc] {
				pred[nr][nc] = append(pred[nr][nc], curr)
			}
		}
	}
	if dist[end[0]][end[1]] == math.MaxInt32 { return [][][2]int{} }
	var buildPath func(pos [2]int) [][][2]int
	buildPath = func(pos [2]int) [][][2]int {
		if pos == start { return [][][2]int{{start}} }
		path := [][][2]int{}
		for _, prePos := range pred[pos[0]][pos[1]] {
			for _, p := range buildPath(prePos) {
				path = append(path, append(p, pos))
			}
		}
		return path
	}
	return buildPath(end)
}

func day21Pt1(lines []string) {
	numPos := map[rune][2]int{}
	for r := range numpad {
		for c, b := range numpad[r] {
			numPos[b] = [2]int{r, c}
		}
	}
	dPos := map[rune][2]int{}
	for r := range dpad {
		for c, b := range dpad[r] {
			dPos[b] = [2]int{r, c}
		}
	}

	seqCache := map[CodePhase]int{}
	lastChar := [3]rune{'A', 'A', 'A'}
	var seq func(string, int) int
	seq = func(s string, phase int) int {
		if v, exists := seqCache[CodePhase{s, phase}]; exists {
			return v
		}
		totalLength := 0
		for _, c := range s {
			var prev, curr [2]int
			var paths []string
			if phase == 0 {
				prev, curr = numPos[lastChar[phase]], numPos[c]
				paths = getNumPaths(prev, curr)
			} else {
				prev, curr = dPos[lastChar[phase]], dPos[c]
				paths = getDPaths(prev, curr)
			}
			lastChar[phase] = c
			minP := math.MaxInt32
			for _, p := range paths {
				tmp := len(p)
				if phase < 2 {
					tmp = seq(p, phase+1)
				}
				minP = min(minP, tmp)
			}
			totalLength += minP
		}
		seqCache[CodePhase{s, phase}] = totalLength
		return totalLength
	}
	total := 0
	for i, l := range lines {
		numeric, _ := strconv.Atoi(l[:3])
		fmt.Println(i, l, numeric)
		total += numeric * seq(l, 0)
	}
	fmt.Println("part 1:", total)
}

func day21Pt2(lines []string) {
	numPos := map[rune][2]int{}
	for r := range numpad {
		for c, b := range numpad[r] {
			numPos[b] = [2]int{r, c}
		}
	}
	dPos := map[rune][2]int{}
	for r := range dpad {
		for c, b := range dpad[r] {
			dPos[b] = [2]int{r, c}
		}
	}

	seqCache := map[CodePhase]int{}
	lastChar := [26]rune{}
	for i := range lastChar {
		lastChar[i] = 'A'
	}
	var seq func(string, int) int
	seq = func(s string, phase int) int {
		if v, exists := seqCache[CodePhase{s, phase}]; exists {
			return v
		}
		totalLength := 0
		for _, c := range s {
			var prev, curr [2]int
			var paths []string
			if phase == 0 {
				prev, curr = numPos[lastChar[phase]], numPos[c]
				paths = getNumPaths(prev, curr)
			} else {
				prev, curr = dPos[lastChar[phase]], dPos[c]
				paths = getDPaths(prev, curr)
			}
			lastChar[phase] = c
			minP := math.MaxInt64
			for _, p := range paths {
				tmp := len(p)
				if phase < 25 {
					tmp = seq(p, phase+1)
				}
				minP = min(minP, tmp)
			}
			totalLength += minP
		}
		seqCache[CodePhase{s, phase}] = totalLength
		return totalLength
	}
	total := 0
	for i, l := range lines {
		numeric, _ := strconv.Atoi(l[:3])
		fmt.Println(i, l, numeric)
		total += numeric * seq(l, 0)
	}
	fmt.Println("part 2:", total)
}

func Day21() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input21.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	numCache = map[[4]int][]string{}
	dCache = map[[4]int][]string{}
	day21Pt1(lines)
	day21Pt2(lines)
}