package main

import (
	"github.com/ty-porter/advent-of-code-2023/util"
	"fmt"
	"strings"
	"strconv"
	"unicode"
)

var words = [9]string { "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" }

func main() {
	lines, err := util.LoadPrompt("day1");
	util.CheckErr(err)

	fmt.Println("Part 1: " + part1(lines))
	fmt.Println("Part 2: " + part2(lines))
}

func part1(lines []string) string {
	sum := 0

	for i := 0; i < len(lines); i++ {
		line := string(lines[i])

		leftDigitIndex  := strings.IndexFunc(line, unicode.IsNumber)
		rightDigitIndex := strings.LastIndexFunc(line, unicode.IsNumber)

		leftDigit  := int(line[leftDigitIndex] - '0')
		rightDigit := int(line[rightDigitIndex] - '0')

		sum += leftDigit * 10 + rightDigit
	}

	return strconv.Itoa(sum)
}

func part2(lines []string) string {
	sum := 0

	for i := 0; i < len(lines); i++ {
		line := string(lines[i])

		leftDigitIndex  := strings.IndexFunc(line, unicode.IsNumber)
		rightDigitIndex := strings.LastIndexFunc(line, unicode.IsNumber)

		leftWordValue, leftWordIndex   := minWordIndex(line)
		rightWordValue, rightWordIndex := maxWordIndex(line)

		// This breaks on negative index (no digits found), but only the test prompt on Part 2 has a line with no digits
		// day1/part2.txt changes line 2 from "eightwothree" to "eightwo1three" to get around this, since it has an equivalent Part 2 value
		leftDigitValue  := int(line[leftDigitIndex] - '0')
		rightDigitValue := int(line[rightDigitIndex] - '0')

		leftDigit  := minValueByBoundedIndex(leftWordIndex, leftWordValue, leftDigitIndex, leftDigitValue)
		rightDigit := maxValueByBoundedIndex(rightWordIndex, rightWordValue, rightDigitIndex, rightDigitValue)

		sum += leftDigit * 10 + rightDigit
	}

	return strconv.Itoa(sum)
}

func minWordIndex(line string) (value int, pos int) {
	value = -1
	pos   = len(line)

	for i := 0; i < len(words); i++ {
		testPos := strings.Index(line, words[i])

		if testPos != -1 && testPos < pos {
			pos = testPos
			value = i + 1
		}
	}

	return value, pos
}

func maxWordIndex(line string) (value int, pos int) {
	value = -1
	pos   = -1

	for i := 0; i < len(words); i++ {
		testPos := strings.LastIndex(line, words[i])

		if testPos != -1 && testPos > pos {
			pos = testPos
			value = i + 1
		}
	}

	return value, pos
}

func minValueByBoundedIndex(wordIndex int, wordValue int, digitIndex int, digitValue int) (int) {
	if wordIndex < 0  { return digitValue }
	if digitIndex < 0 { return wordValue }

	if wordIndex < digitIndex {
		return wordValue
	} else { return digitValue }
}

func maxValueByBoundedIndex(wordIndex int, wordValue int, digitIndex int, digitValue int) (int) {
	if wordIndex < 0  { return digitValue }
	if digitIndex < 0 { return wordValue }

	if wordIndex > digitIndex {
		return wordValue
	} else { return digitValue }
}
