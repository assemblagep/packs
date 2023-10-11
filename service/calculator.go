package service

import (
	"sort"
	"strconv"
)

// packsArray is default sorted list of available packs
var packsArray = []int{5000, 2000, 1000, 500, 250}

// PackNumber return result map of packs and amounts
func PackNumber(amount int) map[int]int {
	m := make(map[int]int)
	if amount <= 0 || len(packsArray) == 0 {
		return m
	}

	// go through packsArray from biggest pack
	for _, pack := range packsArray {
		resOfDivision := amount / pack
		if resOfDivision > 0 {
			amount -= pack * resOfDivision
			m[pack] += resOfDivision
		}
	}

	// add 1 more pack if there is some amount less than min pack
	if amount > 0 {
		m[packsArray[len(packsArray)-1]]++
	}

	return m
}

// Packs accept string array of available packs, sort them in descending order and redefine packsArray
func Packs(in []string) error {
	packsArray = make([]int, len(in))

	for i, s := range in {
		num, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		packsArray[i] = num
	}
	sort.Sort(sort.Reverse(sort.IntSlice(packsArray)))

	return nil
}
