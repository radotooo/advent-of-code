package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	file, err := os.Open("./data/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	defer file.Close()

	fmt.Println(partTwo(sc))
}

func partOne(sc *bufio.Scanner) int {
	threshold := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	result := 0

	for sc.Scan() {
		isValid := true
		line := sc.Text()

		dices := map[string]int{}

		r, _ := regexp.Compile("\\d+")
		p, _ := regexp.Compile("[\\d* \\w+]*")
		lineSplited := strings.Split(line, ":")

		firstPart := lineSplited[0]
		secondPart := lineSplited[1]
		gameId, _ := strconv.Atoi(r.FindString(firstPart))
		subsets := p.FindAllString(secondPart, -1)

		for _, subset := range subsets {
			subsetSplited := strings.Split(strings.Trim(subset, " "), " ")

			diceNumber, _ := strconv.Atoi(subsetSplited[0])
			diceColor := subsetSplited[1]
			dices[diceColor] += diceNumber

			if diceNumber > threshold[diceColor] {
				isValid = false
				break
			}

		}

		if isValid {
			result += gameId
		}
	}

	return result
}

func partTwo(sc *bufio.Scanner) int {
	result := 0

	for sc.Scan() {
		line := sc.Text()

		dices := map[string]int{}

		p, _ := regexp.Compile("[\\d* \\w+]*")
		lineSplited := strings.Split(line, ":")

		secondPart := lineSplited[1]
		subsets := p.FindAllString(secondPart, -1)

		for _, subset := range subsets {
			subsetSplited := strings.Split(strings.Trim(subset, " "), " ")

			diceNumber, _ := strconv.Atoi(subsetSplited[0])
			diceColor := subsetSplited[1]

			if diceNumber > dices[diceColor] {
				dices[diceColor] = diceNumber
			}
		}

		sum := 1

		for _, value := range dices {
			sum *= value
		}

		result += sum
	}

	return result
}
