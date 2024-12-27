package solutions

import (
	"aoc2024/utils"
	"fmt"
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

// Find shortest key press you make and multiply by numeric value
const numA = 10
const dA = 4
var zero rune = '0'
var Numpad [11][2]int8 = [11][2]int8{
	{1,3}, {0,2}, {1,2}, {2,2}, {0,1}, {1,1}, {2,1}, {0,0}, {1,0}, {2,0}, {2,3},
}
var Dpad [5][2]int8 = [5][2]int8{{2,1},{1,1},{0,1},{1,0},{2,0}}

var NumGap [2]int8 = [2]int8{0, 3}
var DGap [2]int8 = [2]int8{2, 0}

func updateY(matrix *[]byte, dy int8) {
	for y := dy; y > 0; y-- { // going down
		*matrix = append(*matrix, 'v')
	}
	for y := dy; y < 0; y++ { // going down
		*matrix = append(*matrix, '^')
	}
}

func updateX(matrix *[]byte, dx int8) {
	for x := dx; x > 0; x-- { // going down
		*matrix = append(*matrix, '>')
	}
	for x := dx; x < 0; x++ { // going down
		*matrix = append(*matrix, '<')
	}
}

func numpadMatrix() [11][11][]byte {
	adjMatrix := [11][11][]byte{}
	for i := range 11 {
		adjMatrix[i][i] = []byte{'A'}
		x1, y1 := Numpad[i][0], Numpad[i][1]
		for j := range 11 {
			if j == i || adjMatrix[i][j] != nil { continue }
			adjMatrix[i][j] = make([]byte, 0)
			adjMatrix[j][i] = make([]byte, 0)
			x2, y2 := Numpad[j][0], Numpad[j][1]
			dx, dy := x2-x1, y2-y1
			if dx < 0 {
				updateY(&adjMatrix[i][j], dy)
				updateX(&adjMatrix[i][j], dx)
				updateX(&adjMatrix[j][i], -dx)
				updateY(&adjMatrix[j][i], -dy)
			} else {
				updateX(&adjMatrix[i][j], dx)
				updateY(&adjMatrix[i][j], dy)
				updateY(&adjMatrix[j][i], -dy)
				updateX(&adjMatrix[j][i], -dx)
			}
			adjMatrix[i][j] = append(adjMatrix[i][j], 'A')
			adjMatrix[j][i] = append(adjMatrix[j][i], 'A')
		}
	}
	return adjMatrix
}

func dpadMatrix() [5][5][]byte {
	adjMatrix := [5][5][]byte{}
	for i := range 5 {
		adjMatrix[i][i] = []byte{'A'}
		x1, y1 := Dpad[i][0], Dpad[i][1]
		for j := range 5 {
			if j == i || adjMatrix[i][j] != nil { continue }
			adjMatrix[i][j] = make([]byte, 0)
			adjMatrix[j][i] = make([]byte, 0)
			x2, y2 := Dpad[j][0], Dpad[j][1]
			dx, dy := x2-x1, y2-y1
			if dx < 0 {
				updateY(&adjMatrix[i][j], dy)
				updateX(&adjMatrix[i][j], dx)
				updateX(&adjMatrix[j][i], -dx)
				updateY(&adjMatrix[j][i], -dy)
			} else {
				updateX(&adjMatrix[i][j], dx)
				updateY(&adjMatrix[i][j], dy)
				updateY(&adjMatrix[j][i], -dy)
				updateX(&adjMatrix[j][i], -dx)
			}
			adjMatrix[i][j] = append(adjMatrix[i][j], 'A')
			adjMatrix[j][i] = append(adjMatrix[j][i], 'A')
		}
	}
	return adjMatrix
}

func day21Pt1(lines []string) {
	nums := make([]int, len(lines))
	start := numA
	nMat := numpadMatrix()
	var sb strings.Builder
	// fmt.Println(lines[4])
	for i, l := range lines {
		v, _ := strconv.Atoi(l[:len(l)-1])
		nums[i] = v
		sb.Reset()
		for _, r := range l {
			dest := min(int(r-zero), numA)
			sb.Write(nMat[start][dest])
			start = dest
		}
		lines[i] = sb.String()
	}
	start = dA
	dMat := dpadMatrix()
	for range 2 {
		for i, l := range lines {
			sb.Reset()
			for _, r := range l {
				var dest int
				switch r {
					case '>': dest = 0
					case 'v': dest = 1
					case '<': dest = 2
					case '^': dest = 3
					case 'A': dest = 4
					}				
					sb.Write(dMat[start][dest])
					start = dest
				}
				lines[i] = sb.String()
			}
		}
	res := 0
	fmt.Println(lines[4])
	for i := range lines {
		// fmt.Println(len(lines[i]), nums[i])
		res += len(lines[i]) * nums[i]
	}
	fmt.Println("part 1:", res)
}

func day21Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day21() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input21.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	day21Pt1(lines)
	day21Pt2(len(lines))
}