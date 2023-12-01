package util

import (
	"bufio"
	"os"
)

func CheckErr(e error) {
	if e != nil {
			panic(e)
	}
}

func LoadPrompt(path string) ([]string, error) {
	file, err := os.Open(path)
	CheckErr(err)

	var lines []string
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, scanner.Err()
}