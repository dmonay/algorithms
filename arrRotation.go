package main

import "fmt"

func testArrayRotation() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	d1 := 2
	arrayRotation(arr1, d1)
}

func arrayRotation(inputArr []int, d int) (rotatedArr []int) {
	dOriginal := d

	for ; d < len(inputArr); d++ {
		rotatedArr = append(rotatedArr, inputArr[d])
	}

	for i := 0; i < dOriginal; i++ {
		rotatedArr = append(rotatedArr, inputArr[i])
	}

	fmt.Printf("Rotated array is %+v\n", rotatedArr)
	return
}
