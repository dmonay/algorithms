package main

import (
	"fmt"
	"math"
)

func testArrayMinDistance() {
	arr1 := []int{3, 5, 4, 2, 6, 5, 6, 6, 5, 4, 8, 3}
	x1 := 3
	y1 := 6
	arrayMinDistance(arr1, x1, y1)

	arr2 := []int{2, 5, 3, 5, 4, 4, 2, 3}
	x2 := 3
	y2 := 2
	arrayMinDistance(arr2, x2, y2)
}

// Given an unsorted array and two numbers x and y, find the minimum distance
// between x and y in the array
func arrayMinDistance(inputArray []int, x, y int) (minDistance int) {
	for i := 0; i < len(inputArray); i++ {

		if inputArray[i] == x {
			var xPos, yPos int
			xPos = i

			for j := 0; j < len(inputArray); j++ {
				if inputArray[j] == y {
					yPos = j
					localMinDistance := int(math.Abs(float64(xPos) - float64(yPos)))
					if minDistance == 0 {
						minDistance = localMinDistance
						break
					}

					if localMinDistance < minDistance {
						minDistance = localMinDistance
						break
					}
				}
			}
		}
	}

	fmt.Printf("Array minimum distance is %d\n", minDistance)

	return
}
