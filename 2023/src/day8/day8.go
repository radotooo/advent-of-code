package day8

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func Solve() {
	bytesRead, err := os.ReadFile("./data/day8.txt")

	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytesRead)

	fmt.Println(partOne(strings.Split(fileContent, "\n")))
}

func partOne(lines []string) int {
	data := map[string][]string{}
	r, _ := regexp.Compile("\\w+")

	navigation := lines[0]

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}

		match := r.FindAllString(line, -1)
		data[match[0]] = append(data[match[0]], match[1], match[2])
	}

	targetValue := "AAA"
	steps := 0

	for {
		nav := navigation[steps%len(navigation)]
		steps++

		if nav == 'R' {
			targetValue = data[targetValue][1]
		} else {
			targetValue = data[targetValue][0]
		}

		if targetValue == "ZZZ" {
			break
		}
	}

	return steps
}
