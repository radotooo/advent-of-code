package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	file, err := os.Open("./data/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	defer file.Close()

	fmt.Println(partTwo(sc))
}

func partOne(sc *bufio.Scanner) int {
	ranges := map[int][][]int{}
	rangeCount := 0

	seeds := []int{}

	for sc.Scan() {
		line := sc.Text()

		if strings.Contains(line, "seeds:") {
			lineSplited := strings.Split(line, ":")
			secondPart := lineSplited[1]

			data, _ := parseData(strings.Fields(secondPart))
			seeds = data

		} else if line != "" {
			data, err := parseData(strings.Fields(line))

			if err != nil {
				rangeCount++
			} else {
				ranges[rangeCount] = append(ranges[rangeCount], data)
			}
		}
	}

	keys := []int{}

	for key := range ranges {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for i := range seeds {
		for _, key := range keys {
			for _, data := range ranges[key] {
				destinationStart, rangeStart, length := data[0], data[1], data[2]
				rangeEnd := rangeStart + length

				if seeds[i] >= rangeStart && seeds[i] <= rangeEnd {
					seeds[i] = (seeds[i] - rangeStart) + destinationStart
					break
				}
			}
		}
	}

	return slices.Min(seeds)
}

func partTwo(sc *bufio.Scanner) int {
	ranges := map[int][][]int{}
	rangeCount := 0

	seeds := []int{}

	for sc.Scan() {
		line := sc.Text()

		if strings.Contains(line, "seeds:") {
			lineSplited := strings.Split(line, ":")
			secondPart := lineSplited[1]

			data, _ := parseData(strings.Fields(secondPart))
			seeds = data

		} else if line != "" {
			data, err := parseData(strings.Fields(line))

			if err != nil {
				rangeCount++
			} else {
				ranges[rangeCount] = append(ranges[rangeCount], data)
			}
		}
	}

	keys := []int{}

	for key := range ranges {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	lowestValue := math.MaxUint32

	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			value := j

			for _, key := range keys {
				for _, data := range ranges[key] {
					destinationStart, rangeStart, length := data[0], data[1], data[2]
					rangeEnd := rangeStart + length

					if value >= rangeStart && value <= rangeEnd {
						value = (value - rangeStart) + destinationStart
						break
					}
				}
			}

			if value < lowestValue {
				lowestValue = value
			}
		}
	}

	return lowestValue
}

func parseData(data []string) ([]int, error) {
	positions := []int{}

	for _, field := range data {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}

		positions = append(positions, num)
	}

	return positions, nil
}
