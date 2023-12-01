package main

import (
	"github.com/ty-porter/advent-of-code-2023/util"
	"fmt"
)

func main() {
	lines, err := util.LoadPrompt("day1/prompt.txt");
	util.CheckErr(err)

	fmt.Println("Part 1: " + part1())
	fmt.Println("Part 2: " + part2())
}

func part1() string {
	return "part 1"
}

func part2() string {
	return "part 2"
}
