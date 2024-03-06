package main

import (
	"daniel/golang-parallel-sort/utils"
	"sync"
)

const size = 100000000
const numWorker = 8

func main() {
	s := utils.Random(size)

	// Divide the slice into smaller sub-slices
	subSliceSize := (size + numWorker - 1) / numWorker

	// Make channels to communicate between goroutines
	sortedChan := make(chan []int, numWorker)

	var wg sync.WaitGroup
	wg.Add(numWorker)

	for i := 0; i < numWorker; i++ {
		go func(start, end int) {
			defer wg.Done()

			// Sort the sub-slice
			if end > len(s) {
				end = len(s)
			}
			subSlice := s[start:end]
			quicksort(subSlice)

			// Send the sorted sub-slice to the channel
			sortedChan <- subSlice
		}(i*subSliceSize, (i+1)*subSliceSize)
	}

	go func() {
		wg.Wait()
		close(sortedChan)
	}()

	// Merge the sorted sub-slices
	var sortedData []int
	for sortedSubSlice := range sortedChan {
		sortedData = merge(sortedData, sortedSubSlice)
	}

	// fmt.Println(sortedData)
}

// Quicksort implementation
func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, pivotIndex := 0, len(arr)-1
	for i := range arr {
		if arr[i] < arr[pivotIndex] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[pivotIndex] = arr[pivotIndex], arr[left]

	quicksort(arr[:left])
	quicksort(arr[left+1:])
}

// Merge two sorted slices into one sorted slice
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}
