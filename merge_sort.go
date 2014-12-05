package main

import (
	"fmt"
)

func main() {
	input := []int{5, 4, 1, 8, 7, 2, 6, 3}
	fmt.Println(mergeSort(input))
	// fmt.Println(merge([]int{3, 7, 8, 11, 29}, []int{1, 4, 6, 12, 14}))
}

func mergeSort(arr []int) []int {
	// assume input has even number of elements
	var c []int

	a, b := split(arr)
	if len(a) == 1 {
		c = merge(a, b)
	} else {
		split(a)
		split(b)
	}

	return c
}

func split(arr []int) ([]int, []int) {
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
	// assume arrays are of equal size since input has even
	// number of elements
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
