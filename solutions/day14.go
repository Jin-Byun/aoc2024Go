package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func day14Pt1(lines []string, re *regexp.Regexp, w, h, t int, disp bool) {
	originX, originY := w >> 1, h >> 1
	n := len(lines)
	pos := make([][2]int, n)
	for i, bot := range lines {
		s := re.FindAllString(bot, 4)
		pos[i] = [2]int{}
		for j := range 2 {
			pos[i][j] = utils.StrToI(s[j]) + utils.StrToI(s[j+2]) * t
		}
		pos[i][0] = ((pos[i][0]%w)+w)%w
		pos[i][1] = ((pos[i][1]%h)+h)%h
	}
	quad := [4]int{}
	for _, p := range pos {
		x, y := p[0], p[1]
		i := 0
		if x == originX || y == originY { continue }
		if x > originX {
			i ^= 1
		}
		if y > originY {
			i ^= 2
		}
		quad[i]++
	}
	res := 1
	for _, v := range quad {
		res *= v
	}
	fmt.Println("part 1:", res)
	if disp {
		displayplot(pos, w, h)
	}
}

func meanXY(nums [][2]int) (float64, float64) {
	y,x := 0, 0
	for _, a := range nums {
			x += a[0]
			y += a[1]
	}
	return float64(x) / float64(len(nums)), float64(y) / float64(len(nums))
}

func varianceXY(nums [][2]int) (float64, float64) {
	var x float64 = 0
	var y float64 = 0
	mx, my := meanXY(nums)
	n := len(nums)

	for _, a := range nums {
			x += (float64(a[0]) - mx) * (float64(a[0]) - mx)
			y += (float64(a[1]) - my) * (float64(a[1]) - my)
	}

	return x / float64(n-1), y / float64(n-1)
}

// Hint taken from the thread of i_have_no_biscuits reddit post in day 14 solution
func day14Pt2(lines []string, re *regexp.Regexp, w, h int) int {
	n := len(lines)
	pos := make([][2]int, n)
	vel := make([][2]int, n)
	for i, bot := range lines {
		s := re.FindAllString(bot, 4)
		pos[i] = [2]int{}
		vel[i] = [2]int{}
		for j := range 2 {
			pos[i][j] = utils.StrToI(s[j])
		}
		for j := 2; j < 4; j++ {
			vel[i][j-2] = utils.StrToI(s[j])
		}
	}
	bx, by := 0, 0
	var bxVar, byVar float64 = 1000, 10000
	for idx := range h {
		// update
		for i, v := range vel {
			pos[i][0], pos[i][1] = (pos[i][0]+v[0]+w)%w, (pos[i][1]+v[1]+h)%h
		}
		xVar, yVar := varianceXY(pos)
		if xVar < bxVar { bx, bxVar = idx+1, xVar }
		if yVar < byVar { by, byVar = idx+1, yVar }
	}
	
	// t ≡ bx (mod w) meaning, t and bx differs by some multiples of w
	// t ≡ by (mod h) meaning, t and by differs by some multiples of h

	// thus, t = bx + k*w and bx + k*w ≡ by (mod h)
	// solve for k -> kw ≡ (by-bx) (mod h)
	// multiply both side by mod inverse of w mod h (inv(w)) because we can't divide in modular arithmetic
	// k 1 (mod h) ≡ (by-bx) inv(w) (modh)
	// solve for t
	// t = bx + ((by-bx)*inv(w) mod h) * w
	
	res := ((modInv(w, h)*(by-bx)) % h)*w + bx
	fmt.Println("part 2:", res)
	return res
}

func modInv(a, b int) int {
	d := gcd(a, b)
	if d != 1 { return -1 }
	for i := 1; i < b; i++ {
		if ((a % b) * (i % b)) % b == 1 { return i } 
	}
	return 1
}

func Day14() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input14.txt")
	utils.HandleErr(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	w, h := 101, 103
	re := regexp.MustCompile(`-?\d+`)
	day14Pt1(lines, re, w, h, 100, false)
	easterEgg := day14Pt2(lines, re, w, h)
	day14Pt1(lines, re, w, h, easterEgg, true)

}

// tried to see the graph
func displayplot(grid [][2]int, w, h int) {
	g := make([][]byte, h)
	for i := range g {
		g[i] = make([]byte, w * 2)
		for j := 0; j < len(g[i]); j += 2 {
			g[i][j] = '.'
			g[i][j+1] = ' '
		}
	}
	for _, pos := range grid {
		x, y := pos[0], pos[1]
		g[y][x*2] = '*'
	}
	for _, line := range g {
		fmt.Println(string(line))
	}
}