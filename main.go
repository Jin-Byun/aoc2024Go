package main

import (
	"aoc2024/solutions"
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Set up: [S], Execute: [X]: ")
	cmd, _ := reader.ReadBytes('\n')
	switch cmd[0] {
		case 'S', 's':
			utils.SetUp(reader)
		case 'X', 'x', 13:
			day := utils.GetDay(reader)
			switch day {
				case "2":
					solutions.Day2()
				case "3":
					solutions.Day3()
				case "4":
					solutions.Day4()
				case "5":
					solutions.Day5()
				case "6":
					solutions.Day6()
				case "7":
					solutions.Day7()
				case "8":
					solutions.Day8()
				case "9":
					solutions.Day9()
				case "10":
					solutions.Day10()
				case "11":
					solutions.Day11()
				case "12":
					solutions.Day12()
				case "13":
					solutions.Day13()
				case "14":
					solutions.Day14()
				case "15":
					solutions.Day15()
				case "16":
					solutions.Day16()
				case "17":
					solutions.Day17()
				case "18":
					solutions.Day18()
				case "19":
					solutions.Day19()
				case "20":
					solutions.Day20()
				case "21":
					solutions.Day21()
				case "22":
					solutions.Day22()
				case "23":
					solutions.Day23()
				case "24":
					solutions.Day24()
				case "25":
					solutions.Day25()
				}
	}
}