package main

import (
	"fmt"
	"log"
	"sort"
)

func sortAndTotal(vals []int) (sorted []int, total int) {
	var logger = log.New(log.Writer(), "sortAndTotal: ", log.Ldate|log.Lmsgprefix)
	logger.Printf("Invoked with %v values", len(vals))
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	logger.Printf("Sorted data: %v", sorted)
	for _, val := range sorted {
		total += val
		// total++
	}
	logger.Printf("Total: %v", total)
	return
}
func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := sortAndTotal(nums)
	fmt.Println("Sorted Data:", sorted)
	fmt.Println("Total:", total)
}
