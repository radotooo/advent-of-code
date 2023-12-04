package day4

import (
	"bufio"
	"fmt"
	"github/radotooo/advent-of-code/src/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	file, err := os.Open("./data/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	defer file.Close()

	fmt.Println(partTwo(sc))
}

func partOne(sc *bufio.Scanner) int {
	result := 0

	for sc.Scan() {
		line := sc.Text()

    winningPositions, playerPositions := parseLine(line)

		hasWin := false
		currentResult := 0

		for _, position := range playerPositions {
			if utils.Contains(winningPositions, position) {

				if hasWin {
					currentResult *= 2
				} else {
					currentResult++
				}

				hasWin = true
			}
		}

		result += currentResult
	}

	return result
}

func partTwo(sc *bufio.Scanner) int {
	result := 0
	currentCard := 1
  results := make(map[int]int)

	for sc.Scan() {
		line := sc.Text()
    winningPositions, playerPositions := parseLine(line)

		results[currentCard]++

		for i := 0; i < results[currentCard]; i++ {
			card := currentCard

			for _, position := range playerPositions {
				if utils.Contains(winningPositions, position) {

					card++
					results[card]++
				}
			}
		}

		currentCard++
	}

  for _, value := range results {
    result += value
  }

	return result
}

func parseLine(line string) ([]int, []int) {
		lineSplited := strings.Split(line, ":")
		secondPart := lineSplited[1]
		secondPartSplited := strings.Split(secondPart, "|")

		winningPositions := parsePositions(strings.Fields(secondPartSplited[0]))
		playerPositions := parsePositions(strings.Fields(secondPartSplited[1]))

    return winningPositions, playerPositions
}

func parsePositions(data []string) []int {
	positions := []int{}
	for _, field := range data {
		num, _ := strconv.Atoi(field)

		positions = append(positions, num)
	}

	return positions
}
