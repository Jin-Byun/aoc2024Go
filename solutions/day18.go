package solutions

import (
	"aoc2024/utils"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Path struct {
	Pos [2]int
	Priority int
}

type PathHeap []Path

func (q PathHeap) Len() int { return len(q) }
func (q PathHeap) Less(i, j int) bool { return q[i].Priority < q[j].Priority }
func (q PathHeap) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *PathHeap) Push(x any) {
	*q = append(*q, x.(Path) )
}
func (q *PathHeap) Pop() any {
	old := *q
	item := old[len(old)-1]
	*q = old[:len(old)-1]
	return item
}

func day18Pt1(res int) {
	fmt.Println("part 1:", res)
}

func day18Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day18() {
	// f, err := os.Open("test.txt")
	// n := 7
	n := 71
	f, err := os.Open("inputs/input18.txt")
	// bytes := 12
	// bytes := 1024
	utils.HandleErr(err)
	defer f.Close()
	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	start := [2]int{0,0}
	goal := [2]int{n-1,n-1}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pos := strings.Split(scanner.Text(), ",")
		x := utils.StrToI(pos[0])
		y := utils.StrToI(pos[1])
		grid[y][x] = '#'
		q := &PathHeap{Path{start, 0}}
		heap.Init(q)
		cost := map[[2]int]int{start: 0}
		for q.Len() > 0 {
			curr := heap.Pop(q).(Path)
			x, y := curr.Pos[0], curr.Pos[1]
			if curr.Pos == goal {
				break
			}
			for d := range 4 {
				dy, dx := y + utils.MoveCardinal[d], x + utils.MoveCardinal[d+1]
				if utils.IdxInValid(dx, dy, n) || grid[dy][dx] == '#' { continue }
				score := cost[curr.Pos] + 1
				next := [2]int{dx, dy}
				if v, exists := cost[next]; !exists || score < v {
					cost[next] = score
					heap.Push(q, Path{next, score + heuristic(goal, next)})
				}
			}
		}
		if _, ok := cost[goal]; !ok {
			fmt.Println(x, y)
			break
		}
	}
	// fmt.Println(cost)
	// day18Pt1(cost[goal])
	// day18Pt2(len(lines))
}