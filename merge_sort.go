package main

import (
	"fmt"
)

func mergeSortMain() {
	input := []int{5, 4, 1, 8, 7, 12, 6, 3, 9}

	fmt.Println("Unsorted array: ", input)
	fmt.Println("Sorted array: ", mergeSort(input))
}

func mergeSort(arr []int) []int {
	a, b := splitMergeSort(arr)
	if len(arr) <= 1 {
		return arr
	}
	left := mergeSort(a)
	right := mergeSort(b)

	return merge(left, right)
}

func splitMergeSort(arr []int) ([]int, []int) {
	len := len(arr)
	midway := len / 2
	a := []int{}
	b := []int{}

	// 1. split array into two halves
	// NOTE: first argument to append must be a slice
	for index, value := range arr {
		if index < midway {
			a = append(a, value)
		} else {
			b = append(b, value)
		}
	}

	return a, b
}

func merge(left, right []int) []int {
	// takes a presorted array
	var c []int
	max := len(left) * 2
	i := 0
	j := 0
	for k := 0; k < max; k++ {
		if left[i] < right[j] {
			c = append(c, left[i])
			i++
			if i == len(left) {
				newRight := right[j:len(right)]
				c = append(c, newRight...)
				break
			}

		} else {
			c = append(c, right[j])
			j++
			if j == len(right) {
				newLeft := left[i:len(left)]
				c = append(c, newLeft...)
				break
			}

		}
	}

	return c
}
