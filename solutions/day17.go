package solutions

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day17Pt1(insts []int, combo [7]int) {
	out := []int{}
	for i := 0; i < len(insts); i+=2 {
		opcode, operand := insts[i], insts[i+1]
		switch opcode {
		case 0:
			denom := 1 << combo[operand]
			combo[4] /= denom
		case 1:
			combo[5] ^= operand
		case 2:
			combo[5] = combo[operand] % 8
		case 3:
			if combo[4] != 0 {
				i = operand - 2
			}
		case 4:
			combo[5] ^= combo[6]
		case 5:
			v := combo[operand] % 8
			out = append(out, v)
		case 6, 7:
			denom := 1 << combo[operand]
			combo[opcode-1] = combo[4] / denom
		}
	}
	fmt.Println()
	for _, v := range out {
		fmt.Printf(",%d", v)
	}
	fmt.Println()
}

func day17Pt2(res int) {
	fmt.Println("part 2:", res)
}

func Day17() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input17.txt")
	utils.HandleErr(err)
	splitter := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	re := regexp.MustCompile(`\d+`)
	combo := [7]int{}
	for i := range 4 {
		combo[i] = i
	}
	reg := re.FindAllString(splitter[0], 3)
	for i, v := range reg {
		combo[4+i] = utils.StrToI(v)
	}
	tmp := re.FindAllString(splitter[1], -1)
	insts := make([]int, len(tmp))
	for i := range tmp {
		insts[i] = utils.StrToI(tmp[i])
	}
	day17Pt1(insts, combo)
	res := 0
	comp := make([]int, len(insts))
	for idx := len(insts)-1; idx >= 0; idx-- {
		for j := range math.MaxInt {
			curr := res + (1 << (idx*3)) * j
			combo[4], combo[5], combo[6] = curr, 0, 0
			v := 0
			for i := 0; i < len(insts); i+=2 {
				opcode, operand := insts[i], insts[i+1]
				switch opcode {
				case 0:
					denom := 1 << combo[operand]
					combo[4] /= denom
				case 1:
					combo[5] ^= operand
				case 2:
					combo[5] = combo[operand] % 8
				case 3:
					if combo[4] != 0 {
						i = operand - 2
					}
				case 4:
					combo[5] ^= combo[6]
				case 5:
					comp[v] = combo[operand] % 8
					v++
				case 6, 7:
					if combo[operand] < 0 {
						fmt.Println(combo, operand, idx, j, res)
					}
					denom := 1 << combo[operand]
					combo[opcode-1] = combo[4] / denom
				}
			}
			if v == len(insts) && checkEqualProgram(comp[idx:], insts[idx:]) {
				fmt.Println(combo, comp, insts)
				fmt.Println(strconv.FormatInt(int64(curr), 8))
				res = curr
				break
			} 
		}
	}
	fmt.Println(combo)
	day17Pt2(res)
}

func checkEqualProgram(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] { return false }
	}
	return true
}