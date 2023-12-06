package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println(partTwo())
}

func partOne() int {
	timesString, distancesString := parseInput()
	result := 1
  times := parsePositions(timesString)
  distances := parsePositions(distancesString)

	for i, value := range times {
		speed := 1
		wins := 0

		for j := 0; j < value; j++ {
			if (times[i]-speed)*speed > distances[i] {
				wins++
			}
			speed += 1
		}

		result *= wins
	}

	return result
}

func partTwo() int {
	timesString, distancesString := parseInput()
	result := 1

	time, _ := strconv.Atoi(strings.Join(timesString, ""))
	distance, _ := strconv.Atoi(strings.Join(distancesString, ""))
	speed := 1
	wins := 0

	for j := 0; j < time; j++ {
		if (time-speed)*speed > distance {
			wins++
		}
		speed += 1
	}

	result *= wins

	return wins
}

func parseInput() ([]string, []string) {
	file, err := os.Open("./data/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	defer file.Close()

	times := []string{}
	distances := []string{}

	for sc.Scan() {
		line := sc.Text()

		parsedLine := strings.Split(line, ":")

		firstPart := parsedLine[0]
		secondPart := parsedLine[1]

		switch firstPart {
		case "Time":
			times = strings.Fields(secondPart)
		case "Distance":
			distances = strings.Fields(secondPart)
		}
	}

	return times, distances
}

func parsePositions(data []string) []int {
	positions := []int{}
	for _, field := range data {
		num, _ := strconv.Atoi(field)

		positions = append(positions, num)
	}

	return positions
}
