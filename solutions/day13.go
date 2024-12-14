package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func gcd(a, b int) int {
	if a == 0 { return b }
	return gcd(b % a, a)
}

func lcm(a, b int) int {
	return (a/gcd(a,b)) * b
}

func getLeastMultipliers(a, b int) (int, int) {
	cm := lcm(a,b)

	return cm/b, cm/a
}

func day13Pt1(s []string, acc *int) {
	nums := make([]int, 6)
	for i := range 6 {
		nums[i] = utils.StrToI(s[i])
	}
	yM, xM := getLeastMultipliers(nums[2], nums[3])
	btnA := (nums[4] * xM - nums[5] * yM) / (nums[0] * xM -nums[1] * yM)
	isClean := (nums[4] * xM - nums[5] * yM) % (nums[0] * xM -nums[1] * yM)
	if isClean != 0 || btnA > 100 || btnA < 0 { return }
	btnB := (nums[4] - nums[0] * btnA) / nums[2]
	isClean = (nums[4] - nums[0] * btnA) % nums[2]
	if isClean != 0 || btnB > 100 || btnB < 0 { return }
	*acc += btnA * 3 + btnB
}
func day13Pt2(s []string, acc *int) {
	nums := make([]int, 6)
	for i := range 6 {
		nums[i] = utils.StrToI(s[i])
	}
	nums[4] += 10000000000000
	nums[5] += 10000000000000
	yM, xM := getLeastMultipliers(nums[2], nums[3])
	btnA := (nums[4] * xM - nums[5] * yM) / (nums[0] * xM -nums[1] * yM)
	isClean := (nums[4] * xM - nums[5] * yM) % (nums[0] * xM -nums[1] * yM)
	if isClean != 0 || btnA < 0 { return }
	btnB := (nums[4] - nums[0] * btnA) / nums[2]
	isClean = (nums[4] - nums[0] * btnA) % nums[2]
	if isClean != 0 || btnB < 0 { return }
	*acc += btnA * 3 + btnB
}

func Day13() {
	f, err := os.ReadFile("inputs/input13.txt")
	utils.HandleErr(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	re := regexp.MustCompile(`\d+`)
	res1 := 0
	res2 := 0
	for _, machine := range lines {
		s := re.FindAllString(machine, -1)
		day13Pt1(s, &res1)
		day13Pt2(s, &res2)
	}
	fmt.Println("pt 1:", res1)
	fmt.Println("pt 2:", res2)
}