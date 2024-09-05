package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSum(t *testing.T) {
	testValues := []int{10, 20, 30}
	_, sum := sortAndTotal(testValues)
	expected := 60
	if sum != expected {
		t.Fatalf("Expected %v, Got %v", expected, sum)
	}
}

func TestSort(t *testing.T) {
	slices := [][]int{
		{1, 279, 48, 12, 3},
		{-10, 0, -10},
		{1, 2, 3, 4, 5, 6, 7},
		{1},
	}
	for index, data := range slices {
		t.Run(fmt.Sprintf("Sort #%v", index), func(subT *testing.T) {
			sorted, _ := sortAndTotal(data)
			if !sort.IntsAreSorted(sorted) {
				subT.Fatalf("Unsorted data %v", sorted)
			}
		})
	}
}
