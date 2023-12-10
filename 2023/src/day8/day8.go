package day8

import (
	"fmt"
	"github/radotooo/advent-of-code/src/utils"
	"log"
	"math"
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

	fmt.Println(partTwo(strings.Split(fileContent, "\n")))
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

func partTwo(lines []string) int {
	data := map[string][]string{}
	r, _ := regexp.Compile("\\w+")

	navigation := lines[0]
	startingPoints := []string{}

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}

		match := r.FindAllString(line, -1)
		data[match[0]] = append(data[match[0]], match[1], match[2])

		if strings.Contains(match[0], "A") {
			startingPoints = append(startingPoints, match[0])
		}
	}

	steps := 0
	pointsCount := make([]int, len(startingPoints))

	for {
		nav := navigation[steps%len(navigation)]
		steps++

		for i, startingPoint := range startingPoints {
			if strings.Contains(startingPoint, "Z") {
				continue
			}

			pointsCount[i]++

			if nav == 'R' {
				startingPoints[i] = data[startingPoint][1]
			} else {
				startingPoints[i] = data[startingPoint][0]
			}
		}

		if utils.Every(startingPoints, func(word string) bool {
			return strings.Contains(word, "Z")
		}) {
			break
		}
	}

	result := pointsCount[0]

	for i := 1; i < len(pointsCount); i++ {
		result = lcm(result, pointsCount[i])
	}

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
