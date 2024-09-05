package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{10, 100, 250}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Array Size %v", size), func(sub *testing.B) {
			data := make([]int, size)
			sub.ResetTimer()
			for i := 0; i < b.N; i++ {
				sub.StopTimer()
				for j := 0; j < size; j++ {
					data[j] = rand.Int()
				}
				sub.StartTimer()
				sortAndTotal(data)
			}
		})
	}

}
