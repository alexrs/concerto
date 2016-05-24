package strop

import (
	"errors"
	"github.com/antzucaro/matchr"
)

func EditDistance(s1, s2 string) int {
	return matchr.Levenshtein(s1, s2)
}

func MinIndex(args []int) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("empty list")
	}
	min := args[0]
	minIndex := 0
	for i, e := range args {
		if e < min {
			min = e
			minIndex = i
		}
	}
	return minIndex, nil
}
