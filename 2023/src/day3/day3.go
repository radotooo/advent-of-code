package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solve() {
	file, err := os.Open("./data/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	defer file.Close()

	fmt.Println(partTwo(sc))
}

func partOne(sc *bufio.Scanner) int {
	matrix := [][]rune{}
	result := 0
	winningPositions := []int{}

	for sc.Scan() {
		line := sc.Text()
		lineArray := []rune{}

		for _, char := range line {
			lineArray = append(lineArray, char)
		}

		matrix = append(matrix, lineArray)
	}

	for i := 0; i < len(matrix); i++ {
		positions := [][2]int{}
		num := ""
		isNumberValid := false

		for j := 0; j < len(matrix[i]); j++ {
			isDigit := unicode.IsDigit(matrix[i][j])

			if isDigit {
				num += string(matrix[i][j])
				positions = append(positions, [2]int{i, j})
			}

			if !isDigit && num != "" || isDigit && j == len(matrix[i])-1 {
				value, _ := strconv.Atoi(num)

				for _, position := range positions {
					if isAdjacentToSymbol(matrix, position[0], position[1], &winningPositions) {
						isNumberValid = true
					}
				}

				if !isNumberValid {
					isNumberValid = checkLastLetters(matrix, positions, winningPositions)
				}

				if isNumberValid {
					result += value
					isNumberValid = false
				}

				positions = [][2]int{}
				num = ""
			}
		}
	}

	return result
}

func checkLastLetters(matrix [][]rune, positions [][2]int, winningPositions []int) bool {
	return isValid(matrix, positions[0][0], positions[0][1]-1, &winningPositions) ||
		isValid(matrix, positions[len(positions)-1][0], positions[len(positions)-1][1]+1, &winningPositions)

}

func partTwo(sc *bufio.Scanner) int {
	matrix := [][]rune{}
	result := 0
	stars := map[string][]int{}

	for sc.Scan() {
		line := sc.Text()
		lineArray := []rune{}

		for _, char := range line {
			lineArray = append(lineArray, char)
		}

		matrix = append(matrix, lineArray)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if string(matrix[i][j]) == "*" {
				stars[strconv.Itoa(i)+strconv.Itoa(j)] = []int{}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		positions := [][2]int{}
		num := ""
		isNumberValid := false
		winningPositions := []int{}

		for j := 0; j < len(matrix[i]); j++ {
			isDigit := unicode.IsDigit(matrix[i][j])

			if isDigit {
				num += string(matrix[i][j])
				positions = append(positions, [2]int{i, j})
			}

			if !isDigit && num != "" || isDigit && j == len(matrix[i])-1 {
				value, _ := strconv.Atoi(num)

				for _, position := range positions {
					if isAdjacentToSymbol(matrix, position[0], position[1], &winningPositions) {
						isNumberValid = true
					}
				}

				if !isNumberValid {
					isNumberValid = isValid(matrix, positions[0][0], positions[0][1]-1, &winningPositions) ||
						isValid(matrix, positions[len(positions)-1][0], positions[len(positions)-1][1]+1, &winningPositions)
				}

				if isNumberValid {
					isNumberValid = false

					key := strconv.Itoa(winningPositions[0]) + strconv.Itoa(winningPositions[1])
					val, _ := stars[key]

					arr := append(val, value)
					stars[key] = arr

				}
				winningPositions = []int{}
				positions = [][2]int{}
				num = ""
			}
		}

	}

	for _, value := range stars {
		if len(value) >= 2 {
			for i := 0; i < len(value); i += 2 {
				result += (value[i] * value[i+1])
			}
		}
	}

	return result
}

func isAdjacentToSymbol(matrix [][]rune, row, col int, winningPositions *[]int) bool {
	return (isValid(matrix, row+1, col, winningPositions) ||
		isValid(matrix, row-1, col, winningPositions) ||
		isValid(matrix, row+1, col+1, winningPositions) ||
		isValid(matrix, row-1, col-1, winningPositions) ||
		isValid(matrix, row+1, col-1, winningPositions) ||
		isValid(matrix, row-1, col+1, winningPositions))
}

func isValid(matrix [][]rune, row, col int, winningPositions *[]int) bool {
	isValid := !isIndexOutOfBounds(matrix, row, col) && strings.ContainsRune("*", matrix[row][col])

	if isValid {
		*winningPositions = append(*winningPositions, row, col)
	}

	return isValid
}

func isIndexOutOfBounds(matrix [][]rune, row, col int) bool {
	return row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[row])
}
