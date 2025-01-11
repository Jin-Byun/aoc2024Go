package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

type Trie struct {
	children [6]*Trie
	freq int
}

func createTrie() *Trie {
	return &Trie{ [6]*Trie{}, 0 }
}

func (t *Trie) Insert(lock []int) {
	for i := range lock {
		if t.children[lock[i]] == nil { t.children[lock[i]] = createTrie() }
		t = t.children[lock[i]]
	}
	t.freq++
}

func (t *Trie) Value(lock []int) int {
	if len(lock) == 0 { return t.freq }
	res := 0
	for i := lock[0]; i < 6; i++ {
		if t.children[i] == nil { continue }
		res += t.children[i].Value(lock[1:])
	}
	return res
}

func createPins(s string, valid byte) []int {
	p := make([]int, 5)
	data := strings.Fields(s)
	for i := range 5 {
		v := 6
		for j := range 6 {
			if data[j][i] != valid { break }
			v--
		}
		p[i] = v
	}
	return p 
}

func Day25() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input25.txt")
	utils.HandleErr(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	
	var lock, key byte = '#', '.'
	tmblr := make(map[byte][][]int, 2)
	tmblr[lock] = [][]int{}
	tmblr[key] = [][]int{}
	for _, l := range lines {
		tmblr[l[0]] = append(tmblr[l[0]], createPins(l, l[0]))
	}
	t := createTrie()
	for _, pins := range tmblr[lock] {
		t.Insert(pins)
	}
	res := 0
	for _, pins := range tmblr[key] {
		res += t.Value(pins)
	}
	fmt.Println("day 25:", res)
}