package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
  "github/radotooo/advent-of-code/src/utils"
)

func Solve() {
	file, err := os.Open("./data/day1.txt")
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
		r, _ := regexp.Compile("\\d+")

		match := strings.Join(r.FindAllString(sc.Text(), -1), "")

		firstChar := match[0]
		lastChar := match[len(match)-1]
		digit, _ := strconv.Atoi(string(firstChar) + string(lastChar))

		result += digit
	}

	return result
}

func partTwo(sc *bufio.Scanner) int {
	result := 0

	for sc.Scan() {
		text := sc.Text()

		firstDigit := getDigit(text, false)
		secondDigit := getDigit(utils.ReverseString(text), true)

		digin, _ := strconv.Atoi(firstDigit + secondDigit)
		result += digin
	}

	return result
}

func getDigit(line string, isReversed bool) string {
	textToNumber := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	isFound := false
	letters := ""
	digit := ""

	for _, char := range line {
		letters += string(char)

		if unicode.IsDigit(char) {
			isFound = true
			digit = string(char)
		}

		for key, value := range textToNumber {
			keyValue := key

			if isReversed {
				keyValue = utils.ReverseString(key)
			}

			if strings.Contains(letters, keyValue) {
				isFound = true
				digit = value
			}
		}

		if isFound {
			break
		}
	}

	return digit
}

