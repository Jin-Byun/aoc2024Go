package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func isValidOrder(order []string, rule map[string]bool) bool {
	for _, v := range order {
		if rule[v] { return false }
	}
	return true
}

func Day5() {
	f, err := os.ReadFile("inputs/input5.txt")
	utils.HandleErr(err)
	parts := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	rawRules, rawOrders := strings.Fields(parts[0]), strings.Fields(parts[1])
	rules := make(map[string]map[string]bool, 100)
	for _, l := range rawRules {
		before, after := l[:2], l[3:]
		if rules[before] == nil { rules[before] = map[string]bool{} }
		rules[before][after] = true
	}

	res1 := 0
	res2 := 0
	
	for _, l := range rawOrders {
		order := strings.Split(l, ",")
		n, i := len(order), 0
		for i < n {
			if !isValidOrder(order[:i], rules[order[i]]) { break }
			i++
		}
		if i == n {
			res1 += utils.StrToI(order[n >> 1])
		} else {
			slices.SortFunc(order, func(a, b string) int {
				if rules[a][b] { return 1 }
				if rules[b][a] { return -1 }
				return 0
			})
			res2 += utils.StrToI(order[n >> 1])
		}
	}

	fmt.Println("part 1: ", res1)	
	fmt.Println("part 2: ", res2)	
}