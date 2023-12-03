package day3

import (
	"bufio"
	"fmt"
	"strings"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Solve() {
	file, err := os.Open("./data/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	defer file.Close()

	fmt.Println(partOne(sc))
}

func partOne(sc *bufio.Scanner) int {
	matrix := [][]rune{}
	result := 0

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
					if isAdjacentToSymbol(matrix, position[0], position[1]) {
						isNumberValid = true
					}
				}

				if !isNumberValid {
					isNumberValid = isValid(matrix, positions[0][0], positions[0][1]-1) || isValid(matrix, positions[len(positions)-1][0], positions[len(positions)-1][1]+1)
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

func isAdjacentToSymbol(matrix [][]rune, row, col int) bool {
	return (
    isValid(matrix, row+1, col) ||
		isValid(matrix, row-1, col) ||
		isValid(matrix, row+1, col+1) ||
		isValid(matrix, row-1, col-1) ||
		isValid(matrix, row+1, col-1) ||
		isValid(matrix, row-1, col+1))
}

func isValid(matrix [][]rune, row, col int) bool {
	isValid := !isIndexOutOfBounds(matrix, row, col) && strings.ContainsRune("*#$+-=/!&%@_1234567890", matrix[row][col])

	return isValid
}

func isIndexOutOfBounds(matrix [][]rune, row, col int) bool {
	return row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[row])
}
