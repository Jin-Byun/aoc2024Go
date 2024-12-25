package solutions

import (
	"aoc2024/utils"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type State struct {
	Pos [2]int
	Score, Direction int
	path map[[2]int]int
}

type PQ []State

func (q PQ) Len() int { return len(q) }
func (q PQ) Less(i, j int) bool { return q[i].Score < q[j].Score }
func (q PQ) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *PQ) Push(x any) {
	*q = append(*q, x.(State) )
}
func (q *PQ) Pop() any {
	old := *q
	item := old[len(old)-1]
	*q = old[:len(old)-1]
	return item
}

func Day16() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input16.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	start, goal := startAndGoal(lines)
	q := &PQ{State{start, 0, 0, map[[2]int]int{}}}
	heap.Init(q)
	// var goalState State
	// cost := map[[2]int]int{start: 0}
	// cameFrom := map[[2]int][2]int{start: {-1, -1}}
	visit := map[[3]int]struct{}{}
	for q.Len() > 0 {
		curr := heap.Pop(q).(State)
		x, y := curr.Pos[0], curr.Pos[1]
		if _, ok := visit[[3]int{x, y, curr.Direction}]; ok { continue }
		curr.path[curr.Pos] = curr.Score
		if curr.Pos == goal {
			fmt.Println(curr.Score, len(curr.path))
			continue
		}
		// 	if curr.Pos == goal {
		// 	goalState = curr
		// 	break
		// }
		for i := range 4 {
			if i == 2 { continue }
			d := (curr.Direction + i) % 4
			dy, dx := y + utils.MoveCardinal[d], x + utils.MoveCardinal[d+1]
			if lines[dy][dx] == '#' { continue }
			if _, ok := visit[[3]int{dx, dy, d}]; ok { continue }
			score := curr.Score + 1
			// score := cost[curr.Pos] + 1
			if d != curr.Direction {
				score += 1000
			}
			next := [2]int{dx, dy}
			// if v, exists := cost[next]; !exists || score < v {
				// cost[next] = score
				heap.Push(q, State{next, score, d, copyMap(curr.path)})
				// cameFrom[next] = curr.Pos
				// heap.Push(q, State{next, score + heuristic(goal, next), d})
			// }
		}
		visit[[3]int{x, y, curr.Direction}] = struct{}{}
	}
	// fmt.Println(goalState.Score, len(goalState.path))
	// fmt.Println(goalState.Score, getUniqueCount(lines, goalState.path, start, goal))
}

func startAndGoal(lines []string) ([2]int, [2]int) {
	start := [2]int{0,0}
	goal := [2]int{0,0}
	for r := range lines {
		for c, b := range lines[r] {
			if b == 'S' {
				start[0] = c
				start[1] = r
			}
			if b == 'E' {
				goal[0] = c
				goal[1] = r
			}
		}
	}
	return start, goal
}

func heuristic(a, b [2]int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	if x < 0 { x = -x }
	if y < 0 { y = -y }
	return x+y
}

func copyMap[V bool | int](path map[[2]int]V) map[[2]int]V {
	new := make(map[[2]int]V, len(path))
	for key, value := range path {
		new[key] = value
	}
	return new
}

// func getUniqueCount(matrix []string, path map[[2]int]int, start, goal [2]int) int {
// 	fmt.Println(start, goal)
// 	q := &PQ{State{start, 0, 0, map[[2]int]int{}}}
// 	heap.Init(q)

// 	visited := make(map[[3]int]struct{})
// 	newSafeCoordinates := make(map[[2]int]struct{})

// 	for q.Len() > 0 {
// 		curr := heap.Pop(q).(State)
// 		x, y := curr.Pos[0], curr.Pos[1]
// 		if score, ok := path[curr.Pos]; ok && score == curr.Score {
// 			for point := range curr.path {
// 				if _, ok := path[point]; !ok {
// 					newSafeCoordinates[point] = struct{}{}
// 				}
// 			}
// 		}

// 		if _, ok := visited[[3]int{x, y, curr.Direction}]; ok {
// 			continue
// 		}

// 		curr.path[curr.Pos] = curr.Score

// 		if curr.Pos == goal {
// 			continue
// 		}

// 		for i := range 4 {
// 			if i == 2 { continue }
// 			d := (curr.Direction + i) % 4
// 			dy, dx := y + utils.MoveCardinal[d], x + utils.MoveCardinal[d+1]
// 			if matrix[dy][dx] == '#' { continue }
// 			if _, ok := visited[[3]int{dx, dy, d}]; ok { continue }
// 			score := curr.Score + 1
// 			if d != curr.Direction {
// 				score += 1000
// 			}
// 			next := [2]int{dx, dy}
// 			heap.Push(q, State{next, score, d, copyMap(curr.path)})
// 		}

// 		visited[[3]int{x, y, curr.Direction}] = struct{}{}
// 	}
// 	return len(path) + len(newSafeCoordinates)
// }

// What was used to find part 2 (check why above didn't work)
// import (
// 	"aoc2024/utils"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// const (
// 	dot  = '.'
// 	end  = 'E'
// 	wall = '#'
// )

// type direction struct{ row, col int }
// type position struct{ row, col int }

// var (
// 	east       = direction{col: 1}
// 	south      = direction{row: 1}
// 	west       = direction{col: -1}
// 	north      = direction{row: -1}
// )

// func (d direction) turnright() direction {
// 	switch d {
// 	case east:
// 		return south
// 	case south:
// 		return west
// 	case west:
// 		return north
// 	case north:
// 		return east
// 	default:
// 		log.Fatalf("unknown direction %v", d)
// 		return d
// 	}
// }

// func (d direction) turnleft() direction {
// 	switch d {
// 	case east:
// 		return north
// 	case north:
// 		return west
// 	case west:
// 		return south
// 	case south:
// 		return east
// 	default:
// 		log.Fatalf("unknown direction %v", d)
// 		return d
// 	}
// }

// func (p position) move(dir direction) position {
// 	return position{row: p.row + dir.row, col: p.col + dir.col}
// }

// type state struct {
// 	pos position
// 	dir direction
// }

// func (s state) possible() (straight, left, right state) {
// 	straight = state{pos: s.pos.move(s.dir), dir: s.dir}
// 	left = state{pos: s.pos, dir: s.dir.turnleft()}
// 	right = state{pos: s.pos, dir: s.dir.turnright()}
// 	return
// }

// type provenance struct {
// 	cost    int
// 	parents []state
// }

// func (p *provenance) maybeAdd(parent state, cost int) {
// 	if p.cost > cost {
// 		p.cost = cost
// 		p.parents = []state{parent}
// 	} else if p.cost == cost {
// 		p.parents = append(p.parents, parent)
// 	}
// }

// type solver struct {
// 	grid     []string
// 	pq       map[int][]state
// 	cheapest int
// 	highest  int
// 	end      state
// 	visited  map[state]int
// 	prov     map[state]*provenance
// }

// func (s *solver) add(v, prev state, cost int) {
// 	if cost < s.cheapest {
// 		log.Fatalf("Trying to add %v at cost %d but cheapest is %d", v, cost, s.cheapest)
// 	}
// 	p := s.prov[v]
// 	if p == nil {
// 		p = &provenance{cost: cost}
// 		s.prov[v] = p
// 	}
// 	p.maybeAdd(prev, cost)
// 	if c, ok := s.visited[v]; !ok || cost < c {
// 		s.visited[v] = cost
// 		s.pq[cost] = append(s.pq[cost], v)
// 		if cost > s.highest {
// 			s.highest = cost
// 		}
// 	}
// }

// func (s *solver) printgrid() {
// 	g := make([][]byte, len(s.grid))
// 	for r, l := range s.grid {
// 		g[r] = []byte(strings.Clone(l))
// 	}
// 	q := []state{s.end}
// 	var zero state
// 	for len(q) > 0 {
// 		v := q[0]
// 		q = q[1:]
// 		if v != zero {
// 			q = append(q, s.prov[v].parents...)
// 		}
// 		var d byte
// 		switch v.dir {
// 		case east:
// 			d = '>'
// 		case west:
// 			d = '<'
// 		case north:
// 			d = '^'
// 		case south:
// 			d = 'v'
// 		}
// 		g[v.pos.row][v.pos.col] = d
// 	}
// 	for _, l := range g {
// 		fmt.Println(string(l))
// 	}
// }

// func (s *solver) pop(cost int) state {
// 	v := s.pq[cost][0]
// 	s.pq[cost] = s.pq[cost][1:]
// 	return v
// }

// func (s *solver) lookup(p position) byte { return s.grid[p.row][p.col] }

// func (s *solver) isend(p position) bool { return s.lookup(p) == end }

// func (s *solver) isopen(p position) bool { return s.lookup(p) != wall }

// func solve(grid []string, start state) *solver {
// 	s := &solver{grid: grid, pq: map[int][]state{}, visited: map[state]int{}, prov: map[state]*provenance{}}
// 	s.add(start, state{}, 0)
// 	for {
// 		for len(s.pq[s.cheapest]) == 0 {
// 			if s.cheapest > s.highest {
// 				log.Fatalf("Ran out of priority queue: %d > %d", s.cheapest, s.highest)
// 			}
// 			s.cheapest++
// 		}
// 		v := s.pop(s.cheapest)
// 		if s.isend(v.pos) {
// 			s.end = v
// 			return s
// 		}
// 		straight, left, right := v.possible()
// 		if s.isopen(straight.pos) {
// 			s.add(straight, v, s.cheapest+1)
// 		}
// 		if s.isopen(left.pos) {
// 			s.add(left, v, s.cheapest+1000)
// 		}
// 		if s.isopen(right.pos) {
// 			s.add(right, v, s.cheapest+1000)
// 		}
// 	}
// }

// func part1(lines []string) string {
// 	start := state{pos: position{row: len(lines) - 2, col: 1}, dir: east}
// 	if lines[start.pos.row][start.pos.col] != 'S' {
// 		start = state{pos: position{row: 1, col: len(lines[0]) - 2}, dir: south}
// 	}
// 	s := solve(lines, start)
// 	s.printgrid()
// 	return fmt.Sprintf("%d", s.cheapest)
// }

// func part2(lines []string) string {
// 	start := state{pos: position{row: len(lines) - 2, col: 1}, dir: east}
// 	if lines[start.pos.row][start.pos.col] != 'S' {
// 		start = state{pos: position{row: 1, col: len(lines[0]) - 2}, dir: south}
// 	}
// 	s := solve(lines, start)
// 	// s.printgrid()
// 	seen := make(map[position]bool)
// 	q := []state{s.end}
// 	var zero state
// 	for len(q) > 0 {
// 		v := q[0]
// 		q = q[1:]
// 		if v != zero {
// 			seen[v.pos] = true
// 			q = append(q, s.prov[v].parents...)
// 		}
// 	}
// 	return fmt.Sprintf("%d", len(seen))
// }

// func Day16() {
// 	f, err := os.ReadFile("inputs/input16.txt")
// 	utils.HandleErr(err)
// 	lines := strings.Fields(strings.TrimSpace(string(f)))
// 	fmt.Println(part1(lines))
// 	fmt.Println(part2(lines))
// }