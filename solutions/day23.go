package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func day23Pt1(connection map[string][]string) {
	validSet := map[[3]string]struct{}{}
	for k, v := range connection {
		n := len(v)
		if n < 2 { continue }
		if chCheck(k) {
			for _, node := range v {
				next := connection[node]
				for _, node2 := range next {
					if node2 == k { continue }
					if slices.Contains(v, node2) {
						validSet[sortedStrings(k, node, node2)] = struct{}{}
					}
				}
			}
		}
	}
	fmt.Println("part 1:", len(validSet))
}

func chCheck(a string) bool {
	return strings.HasPrefix(a, "t")
}

func sortedStrings(a,b,c string) [3]string {
	tmp := []string{a,b,c}
	slices.Sort(tmp)
	return [3]string{tmp[0],tmp[1],tmp[2]}
}

func day23Pt2(connection map[string][]string) {
	resMap := updatedMap(connection, connection)
	for len(resMap) > 1 {
		resMap = updatedMap(resMap, connection)
	}
	var res string
	for k := range resMap {
		res = k
	}
	fmt.Println("part 2:", res)
}

func updatedMap(prev, original map[string][]string) map[string][]string {
	resMap := map[string][]string{}
	for k, v := range prev {
		for _, node := range v {
			tmp1 := append(strings.Split(k, ","), node)
			slices.Sort(tmp1)
			tmp2 := []string{}
			for _, n1 := range v {
				if slices.Contains(original[node], n1) {
					tmp2 = append(tmp2, n1)
				}
			}
			resMap[strings.Join(tmp1, ",")] = tmp2
		}
	}
	return resMap
}

func Day23() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input23.txt")
	utils.HandleErr(err)
	lines := strings.Fields(strings.TrimSpace(string(f)))
	connection := map[string][]string{}
	for _, l := range lines {
		pc := strings.Split(l, "-")
		if _, exists := connection[pc[0]]; !exists {
			connection[pc[0]] = []string{}
		}
		if _, exists := connection[pc[1]]; !exists {
			connection[pc[1]] = []string{}
		}
		connection[pc[0]] = append(connection[pc[0]], pc[1])
		connection[pc[1]] = append(connection[pc[1]], pc[0])
	}
	day23Pt1(connection)
	day23Pt2(connection)
}