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
	file, err := os.Open(path + "/" + promptName() + ".txt")
	CheckErr(err)

	var lines []string
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, scanner.Err()
}

func promptName() string {
	if (len(os.Args) > 1) {
		return os.Args[1]
	}

	return "prompt"
}