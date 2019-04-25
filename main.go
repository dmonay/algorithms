package main

func main() {
	// inversion()
	// mergeSortMain()
	// testArrayMinDistance()
	// testArrayRotation()
	bst()

	// arrayLiteral := [...]string{"one", "two", "three"}
	// slice := []string{"four", "five", "six"}

	// for i := 0; i < len(arrayLiteral); i++ {
	// 	fmt.Println(arrayLiteral[i])
	// }

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	// numsToCount := []int{1, 2, 3, 4, 5}
	// res := count(numsToCount...)

	// fmt.Println(res)
}

func count(nums ...int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}
