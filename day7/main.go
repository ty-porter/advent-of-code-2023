package main

import (
  "github.com/ty-porter/advent-of-code-2023/util"
  "fmt"
  "math"
  "slices"
  "strconv"
  "strings"
)

type hand struct {
  labels string
  bid    int
  rank   int
}

func main() {
  lines, err := util.LoadInput("day7");
  util.CheckErr(err)

  fmt.Println("Part 1: " + part1(lines))
  fmt.Println("Part 2: " + part2(lines))
}

func part1(lines []string) string {
  hands := make([]hand, 0)

  for _, line := range lines {
    hand := parseLine(line)

    hands = append(hands, hand)
  }

  slices.SortFunc(hands, compareHands)

  total := 0
  for i, hand := range hands {
    total += hand.bid * (i + 1)
  }

  return strconv.Itoa(total)
}

func part2(lines []string) string { return "Implement me!" }

func parseLine(line string) hand {
  fields := strings.Fields(line)

  hand := new(hand)
  hand.labels = fields[0]
  hand.rank   = calculateRank(fields[0])
  hand.bid    = util.ForceInt(fields[1])

  return *hand
}

func calculateRank(line string) int {
  seen := make(map[rune]int)
  counts := make([]int, len(line) + 1)
  i := 1

  for _, c := range line {
    if seen[c] == 0 {
      seen[c] = i
      i += 1
    }
    
    counts[seen[c]] += 1
  }

  slices.Sort(counts)

  rank := 0
  for i := 0; i < len(counts) - 1; i++ {
    v := counts[i + 1]
    rank += v * powInt(10, i)
  }

  return rank
}

func powInt(x, y int) int {
  return int(math.Pow(float64(x), float64(y)))
}

func compareHands(a, b hand) int {
  if a.rank < b.rank { return -1 }
  if a.rank > b.rank { return 1 }

  for i := 0; i < len(a.labels); i++ {
    va := calculateLabelValue(a.labels[i])
    vb := calculateLabelValue(b.labels[i])

    if va < vb { return -1 }
    if va > vb { return 1 }
  }

  return 0
}

func calculateLabelValue(l byte) int {
  v := int(rune(l) - '0')

  if 2 <= v && v <= 9 {
    return v
  }

  switch l {
  case 'A':
    return 14
  case 'K':
    return 13
  case 'Q':
    return 12
  case 'J':
    return 11
  case 'T':
    return 10
  }

  panic(fmt.Sprintf("Invalid card: %s", strconv.QuoteRune(rune(l))))
}
