package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PRUNE = 16777216

func process(secret int, f func(int) int) int {
	return (secret ^ f(secret)) % PRUNE
}

func step1(a int) int {
	return a*64
}
func step2(a int) int {
	return a/32
}
func step3(a int) int {
	return a*2048
}

func day22Pt1(lines []string, steps [3]func(int)int) {
	total := 0
	for _, l := range lines {
		start ,_ := strconv.Atoi(l)
		for range 2000 {
			for _, f := range steps {
				start = process(start, f)
			}
		}
		total += start
	}
	fmt.Println("part 1:", total)
}

func day22Pt2(lines []string, steps [3]func(int)int) {
	bananas := make([][2001]int, len(lines))
	for i, l := range lines {
		start ,_ := strconv.Atoi(l)
		for j := range 2000 {
			bananas[i][j] = start % 10
			for _, f := range steps {
				start = process(start, f)
			}
		}
		bananas[i][2000] = start % 10
	}
	cache := map[[4]int][]int{}
	for _, b := range bananas {
		for k := range cache {
			cache[k][1] = cache[k][0]
		}
		for i := 4; i < 2001; i++ {
			seq := makeSeq(b[i-4], b[i-3], b[i-2], b[i-1], b[i])
			if _, ok := cache[seq]; !ok { cache[seq] = make([]int, 2) }
			if cache[seq][0] != cache[seq][1] { continue }
			cache[seq][0] += b[i]
		}
	}
	maximumBanana := 0
	for _, v := range cache {
		maximumBanana = max(maximumBanana, v[0])
	}
	fmt.Println("part 2:", maximumBanana)
}

func Day22() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input22.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	steps := [3]func(int)int{step1, step2, step3}
	
	day22Pt1(lines, steps)
	day22Pt2(lines, steps)
}

func makeSeq(n1,n2,n3,n4,n5 int) [4]int {
	return [4]int{n2-n1, n3-n2, n4-n3, n5-n4}
}