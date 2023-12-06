package main

import (
	"github.com/ty-porter/advent-of-code-2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.LoadInput("day6");
	util.CheckErr(err)

	fmt.Println("Part 1: " + part1(lines))
	fmt.Println("Part 2: " + part2(lines))
}

func part1(lines []string) string {
	times   := strings.Fields(lines[0])
	records := strings.Fields(lines[1])

	margin := 1

	for i, time := range times {
		if i == 0 { continue }

		time, _ := strconv.Atoi(time)
		record, _ := strconv.Atoi(records[i])

		margin *= calculateMargin(time, record)
	}

	return strconv.Itoa(margin)
}

func part2(lines []string) string {
	times   := strings.Fields(lines[0])
	records := strings.Fields(lines[1])

	timeStr   := ""
	recordStr := ""

	for i, time := range times {
		if i == 0 { continue }

		timeStr += time
		recordStr += records[i]
	}

	time, _ := strconv.Atoi(timeStr)
	record, _ := strconv.Atoi(recordStr)

	return strconv.Itoa(calculateMargin(time, record))
}

func calculateMargin(D int, L int) int {
	/*
	Distance traveled (L) is given by:

		L <= (D - T) * T

	where D = race duration and T - time held.

	This forms a quadratic equation:

		0 <= -T^2 + DT - L

			or, in terms of x:

		0 <= -x^2 + Dx - L

	Roots of this equation are:

		-D +/- sqrt( D^2 + 4(-L) )
		--------------------------
		    			  -2

	Margin will be:
		Range of [ ⌈ minimum root ⌉, ⌊ maximum root ⌋ ]
	*/
	a := float64(-1)
	b := float64(D)
	c := float64(-L)

	r1 := (-b + math.Sqrt(b*b - 4*a*c)) / (2 * a)
	r2 := (-b - math.Sqrt(b*b - 4*a*c)) / (2 * a)

	// Handle exclusive less than in case roots are whole numbers
	if math.Remainder(r1, 1.0) == 0 { r1 += 1.0 }
	if math.Remainder(r2, 1.0) == 0 { r2 -= 1.0 }

	return int(math.Floor(r2) - math.Ceil(r1)) + 1
}
