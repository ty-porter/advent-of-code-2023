package main

import (
  "github.com/ty-porter/advent-of-code-2023/util"
  "fmt"
  "strconv"
  "strings"
)

type Point struct {
  id     int  // ID for debugging multiple runners
  x, y   int  // Location on grid
  dx, dy int  // Current velocity Δx, Δy in [-1, 0, 1], in screen coordinates (+ is down, right)
  steps  int  // Step counter
  pipe   rune // Current pipe
}

const target = 'S'

func main() {
  grid, err := util.LoadInput("day10");
  util.CheckErr(err)

  start := findStart(grid)

  fmt.Println("Part 1: " + part1(grid, start))
  fmt.Println("Part 2: " + part2(grid))
}

func part1(grid []string, start Point) string {
  runners := createRunners(start, grid)
  distances := make(map[string]int)

  for {
    done := 0

    for _, position := range runners {
      move(grid, position, distances)

      if position.pipe == target { done++ }
    }

    if done == len(runners) { break }
  }

  maxDist := 0

  for _, v := range distances { 
    if v > maxDist { maxDist = v }
  }

  return strconv.Itoa(maxDist)
}

func part2(grid []string) string { return "Implement me!" }

func findStart(grid []string) Point {
  for y, line := range grid {
    for x, pipe := range line {
      if pipe == target { return Point { x: x, y: y, pipe: pipe } }
    }
  }

  panic("Start position not found!")
}

func createRunners(point Point, grid []string) []*Point {
  invalidPipes := []string { "-7F", "-JL", "|FL", "|7J" }
  directions := [][]int { {0, 1}, {0, -1}, {1, 0}, {-1, 0} }
  points := make([]*Point, 0)

  for i, d := range directions {
    point := Point { id: i, x: point.x, y: point.y, dx: d[0], dy: d[1], pipe: target }
    pipe := rune(grid[point.y + point.dy][point.x + point.dx])

    if strings.ContainsRune(invalidPipes[i], pipe) { continue }
   
    points = append(points, &point)
  }

  return points
}

func move(grid []string, position *Point, distances map[string]int) {
  if position.steps > 0 && position.pipe == target { return }

  xNext := position.x + position.dx
  yNext := position.y + position.dy
  pipe := rune(grid[yNext][xNext])

  switch pipe {
  case '|':
    position.moveRel(0, position.dy)
  case '-':
    position.moveRel(position.dx, 0)
  case 'L':
    if position.dy == 1 {
      position.moveRel(1, 0)
    } else {
      position.moveRel(0, -1)
    }
  case 'J':
    if position.dy == 1 {
      position.moveRel(-1, 0)
    } else {
      position.moveRel(0, -1)
    }
  case '7':
    if position.dx == 1 {
      position.moveRel(0, 1)
    } else {
      position.moveRel(-1, 0)
    }
  case 'F':
    if position.dx == -1 {
      position.moveRel(0, 1)
    } else {
      position.moveRel(1, 0)
    }
  case target:
  default:
    panic(fmt.Sprintf("Cannot move from position %v onto space containing %s at {%d, %d}", position, strconv.QuoteRune(pipe), xNext, yNext))
  }

  position.steps += 1
  position.pipe = pipe

  writeMinDistance(position, distances)
}

func writeMinDistance(point *Point, distances map[string]int) {
  key := fmt.Sprintf("%s,%d", point.x, point.y)

  value, ok := distances[key]

  if !ok || point.steps < value { distances[key] = point.steps }
}

// Relative move by one unit. Adds x, y to point's x, y and sets the movement direction.
func (p *Point) moveRel(x, y int) {
  // Validate -1 >= x,y >= 1 and (x,y) is not (0,0)
  if x > 1 || y > 1 || x < -1 || y < -1 {
    panic(fmt.Sprintf("Only [-1, 0, 1] are valid update values, received x=%d, y=%d", x, y))
  } else if !(x * y== 0 && x + y != 0) {
    panic(fmt.Sprintf("Can only move in one dimension at a time, received x=%d, y=%d", x, y))
  }

  p.x += p.dx
  p.y += p.dy
  p.dx = x
  p.dy = y
}

func (p Point) String() string {
  return fmt.Sprintf("Point< ID: %d x: %d y: %d dx: %d dy: %d, steps: %d pipe: %s>", p.id, p.x, p.y, p.dx, p.dy, p.steps, strconv.QuoteRune(p.pipe))
}
