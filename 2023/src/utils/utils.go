package utils

import (
	"fmt"
	"strconv"
)

func ReverseString(s string) string {
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func PrintMatrix(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf(strconv.QuoteRune(matrix[i][j]))
		}
		fmt.Println()
	}
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Every[T any](elements []T, pred func(T) bool) bool {
	for _, t := range elements {
		if !pred(t) {
			return false
		}
	}
	return true
}
