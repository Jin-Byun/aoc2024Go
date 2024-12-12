package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func day11Pt1(stones []string) {
	defer utils.FuncTimer("pt1")()
	for range 25 {
		n := len(stones)
		for i := range n {
			switch {
			case stones[i] == "0":
				stones[i] = "1"
			case len(stones[i]) & 1 == 0:
				mid := len(stones[i]) / 2
				stones = append(stones, stones[i][:mid])
				for mid < len(stones[i]) - 1 && stones[i][mid] == '0' {
					mid++
				}
				stones[i] = stones[i][mid:]
			default:
				stones[i] = strconv.Itoa(utils.StrToI(stones[i]) * 2024)
			}
		}
	}
	fmt.Println("part 1: ", len(stones))
}

func countForEvenDigits(x int) (int, bool) {
	res := 1
	x /= 10
	for x > 0 {
		x /= 10
		res++
	}
	return res, res & 1 == 0
}
func getMultiplier(x int) int {
	res := 1
	for range x {
		res *= 10
	}
	return res
}
var funcCache map[[2]int]int = map[[2]int]int{}
func day11Pt2(stone, blinkCount int) int {
	if v, exists := funcCache[[2]int{stone, blinkCount}]; exists {
		return v
	}
	if blinkCount == 0 { return 1 }
	res := 0
		if stone == 0 {
			res = day11Pt2(1, blinkCount-1)
		} else if digits, isEven := countForEvenDigits(stone); isEven {
			mul := getMultiplier(digits / 2)
			res = day11Pt2(stone / mul, blinkCount-1) + day11Pt2(stone % mul, blinkCount-1)
		} else {
			res = day11Pt2(stone * 2024, blinkCount-1)
		}
	funcCache[[2]int{stone, blinkCount}] = res
	return res
}

func Day11() {
	defer utils.FuncTimer("main")()
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input11.txt")
	utils.HandleErr(err)
	stones := strings.Fields(strings.TrimSpace(string(f)))
	nums := make([]int, len(stones))
	for i := range stones {
		nums[i] = utils.StrToI(stones[i])
	}
	day11Pt1(stones)
	pt2Blink := 75
	res2 := 0
	for _, n := range nums {
		res2 += day11Pt2(n, pt2Blink)
	}
	fmt.Println("part 2:", res2)
}