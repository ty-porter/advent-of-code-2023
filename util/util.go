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

func LoadInput(path string) ([]string, error) {
	file, err := os.Open(path + "/" + inputName() + ".txt")
	CheckErr(err)

	var lines []string
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, scanner.Err()
}

func inputName() string {
	if (len(os.Args) > 1) {
		return os.Args[1]
	}

	return "input"
}