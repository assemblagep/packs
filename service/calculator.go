package service

import (
	"sort"
	"strconv"
)

var packsArray = []int{5000, 2000, 1000, 500, 250} //default

func PackNumber(amount int) map[int]int {
	m := make(map[int]int)
	if amount <= 0 || len(packsArray) == 0 {
		return m
	}

	for _, pack := range packsArray {
		resOfDivision := amount / pack
		if resOfDivision > 0 {
			amount -= pack * resOfDivision
			m[pack] += resOfDivision
		}
	}

	if amount > 0 {
		m[packsArray[len(packsArray)-1]]++
	}

	return m
}

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
