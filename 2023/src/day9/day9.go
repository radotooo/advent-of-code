package day9

import (
	"fmt"
	"github/radotooo/advent-of-code/src/utils"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve() {
	bytesRead, err := os.ReadFile("./data/day9.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(bytesRead)

	fmt.Println(partOne(strings.Split(fileContent, "\n")))
	fmt.Println(partTwo(strings.Split(fileContent, "\n")))
}

func partOne(lines []string) int {
	result := 0

	for _, line := range lines {
		nums := strings.Fields(line)

		if line == "" {
			continue
		}

		ints := make([]int, len(nums))

		for i, s := range nums {
			ints[i], _ = strconv.Atoi(s)
		}

    result += calculate(ints)
	}

	return result
}


func partTwo(lines []string) int {
	result := 0

	for _, line := range lines {
		nums := strings.Fields(line)

		if line == "" {
			continue
		}

		ints := make([]int, len(nums))

		for i, s := range nums {
			ints[i], _ = strconv.Atoi(s)
		}

    slices.Reverse(ints)
    result += calculate(ints)
	}

	return result
}

func calculate(ints []int) int{
     result := 0

		for !utils.Every(ints, func(num int) bool {
			return num == 0
		}) {
			temp := make([]int, len(ints)-1)

			for i := len(ints) - 1; i > 0; i-- {
				if i > 0 {
					if i == len(ints)-1 {
						result += ints[i]
					}

					temp[i-1] = ints[i] - ints[i-1]
				}
			}

			ints = temp
		}

    return result
}
