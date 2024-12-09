package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)
func makeSlices(x, n int) []int {
	res := make([]int, n)
	for i := range n {
		res[i] = x
	}
	return res
}

func day9Pt1(line string) {
	blocks := []int{}
	gaps := []int{}
	id := 0
	for i, r := range line {
		n := int(r ^ '0')
		switch i & 1 {
			case 0:
				blocks = slices.Concat(blocks, makeSlices(id / 2, n))
			case 1:
				gaps = slices.Concat(gaps, makeSlices(id / 2, n))
		}
		id++
	}
	l, g, r := 0, 0, len(blocks) - 1
	id = 0
	pos := 0
	var sum uint64
	for l <= r {
		for l <= r && blocks[l] == id {
			sum += uint64(id * pos)
			l++
			pos++
		}
		for g < len(gaps) && l <= r && gaps[g] == id {
			sum += uint64(blocks[r] * pos)
			r--
			pos++
			g++
		}
		id++
	}
	fmt.Println("part 1:", sum)
}

// {startIdx, length, id}
type Block [3]int

func day9Pt2(line string) {
	blocks := []Block{}
	gaps := []Block{}
	id, prev := 0, 0
	for i, r := range line {
		n := int(r ^ '0')
		if r != '0' {
			switch i & 1 {
				case 0:
					blocks = append(blocks, Block{prev, n, id / 2})
				case 1:
					gaps = append(gaps, Block{prev, n, id / 2})
			}
			prev += n
		}
		id++
	}
	for i := len(blocks) - 1; i > 0; i-- {
		b := blocks[i]
		for j := 0; j < len(gaps) && gaps[j][2] < b[2]; j++ {
			if gaps[j][1] >= b[1] {
				blocks[i][0] = gaps[j][0]
				gaps[j][1] -= blocks[i][1]
				gaps[j][0] += blocks[i][1]
				break
			}
		}
	}
	sum := 0
	for _, b := range blocks {
		for i := range b[1] {
			sum += (b[0] + i) * b[2]
		}
	}
	fmt.Println("part 2:", sum)
}



func Day9() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input9.txt")
	utils.HandleErr(err)
	line := strings.TrimSpace(string(f))
	
	day9Pt1(line)
	day9Pt2(line)
}