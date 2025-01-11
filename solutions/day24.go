package solutions

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day24Pt1(data []string) {
	wires := map[string]uint8{}
	for _, l := range strings.Split(data[0], "\n") {
		wires[l[:3]] = l[5] - '0'
	}
	q := [][3]string{} // wires, gate type, res Wire
	zLen := 0
	for _, l := range strings.Split(data[1], "\n") {
		gateData := strings.Fields(l)
		if gateData[4][0] == 'z' { zLen++ }
		q = append(q, [3]string{gateData[0]+gateData[2], gateData[1], gateData[4]})		
	}
	res := make([]int, zLen)
	for len(q) > 0 && zLen > 0 {
		gate := q[0]
		w, g, r := gate[0], gate[1], gate[2]
		q = q[1:]
		if _, ok := wires[w[:3]]; !ok {
			q = append(q, gate)
			continue
		}
		if _, ok := wires[w[3:]]; !ok {
			q = append(q, gate)
			continue
		}
		var v uint8
		switch g {
		case "XOR":
			v = wires[w[:3]] ^ wires[w[3:]]
		case "OR":
			v = wires[w[:3]] | wires[w[3:]]
		default:
			v = wires[w[:3]] & wires[w[3:]]
		}
		if r[0] == 'z' {
			idx, _ := strconv.Atoi(r[1:])
			res[idx] = int(v)
			zLen--
		} else {
			wires[r] = v
		}
	}
	var sum int
	for i, v := range res {
		sum += v << i
	}
	fmt.Println("part 1:", sum)
}

func day24Pt2(res int) {
	fmt.Println("part 2:", res)
}

type Gate struct {
	g, w1, w2, w3, paste, forward string
}

func createGate(s string) *Gate {
	g := strings.Fields(s)
	return &Gate{g[1], g[0], g[2], g[4], "", ""}
}

func (g *Gate) updatePaste(s string) {
	g.paste = s
}
func (g *Gate) updateForward(inMap map[string][]*Gate) {
	wire := g.w3
	if wire[0] == 'z' {
		g.forward = "FINAL"
		return
	}
	res := make([]string, len(inMap[wire]))
	for i := range res {
		res[i] = inMap[wire][i].g
	}
	slices.Sort(res)
	g.forward = strings.Join(res, ",")
}
func (g *Gate) Pattern() string {
	return g.paste + "." + g.g + "." + g.forward
}

func Day24() {
	// f, err := os.ReadFile("test.txt")
	f, err := os.ReadFile("inputs/input24.txt")
	utils.HandleErr(err)
	data := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	day24Pt1(data)
	outputGates := map[string]*Gate{}
	inputGates := map[string][]*Gate{}
	for _, l := range strings.Split(data[1], "\n") {
		g := createGate(l)
		outputGates[g.w3] = g
		if inputGates[g.w1] == nil { inputGates[g.w1] = []*Gate{} }
		if inputGates[g.w2] == nil { inputGates[g.w2] = []*Gate{} }
		inputGates[g.w1] = append(inputGates[g.w1], g)
		inputGates[g.w2] = append(inputGates[g.w2], g)
	}
	gates := initGates(outputGates, inputGates)
	badWires := []string{}
	for {
		patterns := map[string][]string{}
		for _, g := range gates {
			p := g.Pattern()
			if patternSpecial(g, p) { continue }
			if patterns[p] == nil { patterns[p] = []string{} }
			patterns[p] = append(patterns[p], g.w3)
		}
		for k, list := range patterns {
			if len(list) > 8 { continue }
			fmt.Println(k, list)
			badWires = append(badWires, list...)
		}
		if len(badWires) == 8 { break }
		nextGates := []*Gate{}
		for _, g := range gates {
			wire := g.w3
			if wire[0] == 'z' || slices.Contains(badWires, wire) { continue }
			for _, next := range inputGates[wire] {
				next.updatePaste(g.paste+" "+g.g)
				next.updateForward(inputGates) 
				nextGates = append(nextGates, next)
			}
		}
		gates = nextGates
	}
	slices.Sort(badWires)
	fmt.Println(strings.Join(badWires, ","))
	day24Pt2(len(badWires))
}

func patternSpecial(g *Gate, p string) bool {
	switch p {
	case "xy.XOR.FINAL", "xy.AND.AND,XOR":
		return g.w1[1:] == "00"
	case "xy AND.OR.FINAL":
		return g.w3 == "z45"
	}
	return false
}

func initGates(oGates map[string]*Gate, iGates map[string][]*Gate) []*Gate {
	starter := make([]*Gate, 90)
	typeToInt := map[string]int{"AND": 0, "XOR": 1}
	for _, g := range oGates {
		if g.w1[0] != 'x' && g.w1[0] != 'y' { continue }
		i, _ := strconv.Atoi(g.w1[1:])
		i = i * 2 + typeToInt[g.g]
		starter[i] = g
		g.updatePaste("xy")
		g.updateForward(iGates)
	}
	return starter
}
