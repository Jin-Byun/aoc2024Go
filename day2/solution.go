package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func handleErr(e error) { if e != nil { panic(e) } }

func sToI(s string) int {
	v, _ := strconv.Atoi(s);
	return v
}

func increasing(prev int, seq []string) int {
	for _, c := range seq {
		curr := sToI(c)
		if curr <= prev || curr - prev > 3 {
			return 0
		}
		prev = curr
	}
	return 1
}

func decreasing(prev int, seq []string) int {
	for _, c := range seq {
		curr := sToI(c)
		if curr >= prev || prev - curr > 3 {
			return 0
		}
		prev = curr
	}
	return 1
}

func pt1(list []string) {
	res := 0
	for _, line := range list {
		strSeq := strings.Fields(line)
		start := sToI(strSeq[0])
		prev := sToI(strSeq[1])
		if prev == start {
			continue
		}
		if prev > start && prev - start <= 3 {
			res += increasing(prev, strSeq[2:])
			continue
		}
		if prev < start && start - prev <= 3 {
			res += decreasing(prev, strSeq[2:])
		}
	}
	fmt.Println("part 1: ", res)
}

func pt2(list[]string) {
	res := 0

	for _, line := range list {
		strSeq := strings.Fields(line)
		for i := range strSeq {
			v := strSeq[i]
			strSeq = slices.Delete(strSeq, i, i+1)
			if processPt2(strSeq) {
				res++
				break
			}
			strSeq = slices.Insert(strSeq, i, v)
		}
	}
	fmt.Println("part 2: ", res)
}

func processPt2(strSeq []string) bool {
	start := sToI(strSeq[0])
	prev := sToI(strSeq[1])
	if prev == start {
		return false
	}
	if prev > start && prev - start <= 3 {
		return increasing(prev, strSeq[2:]) == 1
	}
	if prev < start && start - prev <= 3 {
		return decreasing(prev, strSeq[2:]) == 1
	}
	return false
}

func main() {
	f, err := os.ReadFile("day2/input.txt")
	handleErr(err)
	list := strings.Split(strings.TrimSpace(string(f)), "\n")
	pt1(list)
	pt2(list)
}