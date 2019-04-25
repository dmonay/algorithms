package main

import (
	"fmt"
)

func inversion() {
	input := []int{3, 1, 5, 2, 4, 6}
	fmt.Println(sortCountInv(input))
}

func sortCountInv(arr []int) ([]int, int) {
	n := len(arr)
	if n == 1 {
		return arr, 0
	}

	left, right := split(arr)
	b, x := sortCountInv(left)
	c, y := sortCountInv(right)
	d, z := mergeCountSplitInv(b, c)
	return d, x + y + z
}

func mergeCountSplitInv(left, right []int) ([]int, int) {
	// While merging the two sorted subarrays, keep running count of
	// total number of split inversions. Takes a presorted array.

	// running time of this subroutine: O(n). Thus time complexity of
	// entire algorithm is O(nlog(n)) since main routine calls this one
	// recursively. But it calls itself (sourtCountInv) twice, so why is it not
	// O(nlog(n^2)) ?
	var c []int
	x := 0
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
			numInvInLeft := len(left) - i
			x += numInvInLeft
			c = append(c, right[j])
			j++
			if j == len(right) {
				newLeft := left[i:len(left)]
				c = append(c, newLeft...)
				break
			}
		}
	}
	return c, x
}

func split(arr []int) ([]int, []int) {
	len := len(arr)
	midway := len / 2
	a := []int{}
	b := []int{}

	for index, value := range arr {
		if index < midway {
			a = append(a, value)
		} else {
			b = append(b, value)
		}
	}
	return a, b
}
