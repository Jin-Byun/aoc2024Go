package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

// from leetcode question wordbreak
func day19Pt1(pSet map[string]bool, lP int, designs []string) {
	res := 0
	for _, d := range designs {
		n := len(d)
		dp := make([]bool, n+1)
		dp[0] = true
		for i := 1; i <= n; i++ {
			for j := i-1; j >= 0; j-- {
				if i-j > lP { break }
				if dp[j] && pSet[d[j:i]] {
					dp[i] = true
					break
				}
			}
		}
		if dp[n] { res++ }
	}
	fmt.Println("part 1:", res)
}

func day19Pt2(pSet map[string]bool, designs []string) {
	memo := map[string]int{"": 1}
	// from leetcode question wordbreak II 
	var wordBreak func(string) int	
	wordBreak = func (s string) int {
		if v, ok := memo[s]; ok { return v }
		sum := 0
		for i := range s {
			word := s[:i+1]
			if pSet[word] {
				sum += wordBreak(s[i+1:])
			}
		}
		memo[s] = sum
		return sum
	}
	res := 0
	for _, d := range designs {
		res += wordBreak(d)
	}
	fmt.Println("part 2:", res)
}

func Day19() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input19.txt")
	utils.HandleErr(err)
	input := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	pSet, longestPatternLength := processPattern(strings.Split(input[0], ", "))
	designs := strings.Fields(input[1])
	day19Pt1(pSet, longestPatternLength, designs)
	day19Pt2(pSet, designs)
}

func processPattern(pattern []string) (map[string]bool, int) {
	comp := func(a, b string) int {
			return len(a) - len(b)
		}
	lim := len(slices.MaxFunc(pattern, comp))
	patternSet := map[string]bool{}
	for _, p := range pattern {
		patternSet[p] = true
	}
	return patternSet, lim
}