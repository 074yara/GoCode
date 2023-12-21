package main

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BubbleSort(UnsortedSlice)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = QuickSort(UnsortedSlice)
	}
}
