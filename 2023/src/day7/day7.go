package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	file, err := os.Open("./data/day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	defer file.Close()

	fmt.Println(partTwo(sc))
}

type Hand struct {
	cards    string
	bet      int
	bestHand int
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

func (hand *Hand) GetBestHand() {
	types := map[rune]int{}

	for _, val := range hand.cards {
		types[val] += 1
	}

	bestHand := ONE_PAIR

	if len(types) == 1 {
		bestHand = FIVE_OF_A_KIND
	} else if len(types) == 2 {
		if count(types, 4) == 1 {
			bestHand = FOUR_OF_A_KIND
		} else {
			bestHand = FULL_HOUSE
		}
	} else if len(types) == 3 {
		if count(types, 3) == 1 {
			bestHand = THREE_OF_A_KIND
		} else {
			bestHand = TWO_PAIR
		}
	} else if len(types) == 5 {
		bestHand = HIGH_CARD
	}

	hand.bestHand = bestHand
}

func GetBestHandPartTwo(hand *Hand, hasJoker bool) {
	types := map[rune]int{}

	for _, val := range hand.cards {
		types[val] += 1
	}

	hand.GetBestHand()
	if !hasJoker {
		return
	} 

	bestHand := ONE_PAIR

	if types['J'] >= 4 {
		bestHand = FIVE_OF_A_KIND
	} else if types['J'] == 3 {
		if hand.bestHand == FULL_HOUSE {
			bestHand = FIVE_OF_A_KIND
		} else {
			bestHand = FOUR_OF_A_KIND
		}
	} else if types['J'] == 2 {
		if hand.bestHand == TWO_PAIR {
			bestHand = FOUR_OF_A_KIND
		} else if hand.bestHand == ONE_PAIR {
			bestHand = THREE_OF_A_KIND
		} else if hand.bestHand == FULL_HOUSE {
			bestHand = FIVE_OF_A_KIND
		}
	} else if types['J'] == 1 {
		if hand.bestHand == THREE_OF_A_KIND {
			bestHand = FOUR_OF_A_KIND
		} else if hand.bestHand == TWO_PAIR {
			bestHand = FULL_HOUSE
		} else if hand.bestHand == ONE_PAIR {
			bestHand = THREE_OF_A_KIND
		} else if hand.bestHand == FOUR_OF_A_KIND {
			bestHand = FIVE_OF_A_KIND
		}
	}

	hand.bestHand = bestHand
}

var letterValues = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func partOne(sc *bufio.Scanner) int {
	hands := []Hand{}
	result := 0

	for sc.Scan() {
		line := sc.Text()

		splited := strings.Split(line, " ")
		bet, _ := strconv.Atoi(splited[1])

		hand := Hand{cards: splited[0], bet: bet, bestHand: 0}
		hand.GetBestHand()
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].bestHand == hands[j].bestHand {
			for x := range hands[i].cards {
				if hands[i].cards[x] != hands[j].cards[x] {
					return letterValues[hands[i].cards[x]] < letterValues[hands[j].cards[x]]
				}
			}
		}
		return hands[i].bestHand < hands[j].bestHand
	})

	for i, hand := range hands {
		result += (hand.bet * (i + 1))
	}

	return result
}

func partTwo(sc *bufio.Scanner) int {
	hands := []Hand{}
	result := 0
  letterValues['J'] = 1

	for sc.Scan() {
		line := sc.Text()

		splited := strings.Split(line, " ")
		bet, _ := strconv.Atoi(splited[1])

		hand := Hand{cards: splited[0], bet: bet, bestHand: 0}
		GetBestHandPartTwo(&hand, strings.Contains(hand.cards, "J"))

		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].bestHand == hands[j].bestHand {
			for x := range hands[i].cards {
				if hands[i].cards[x] != hands[j].cards[x] {
					return letterValues[hands[i].cards[x]] < letterValues[hands[j].cards[x]]
				}
			}
		}
		return hands[i].bestHand < hands[j].bestHand
	})

	for i, hand := range hands {
		result += (hand.bet * (i + 1))
	}

	return result
}

func count(input map[rune]int, target int) int {
	count := 0
	for _, value := range input {
		if value == target {
			count++
		}
	}

	return count
}
