package main

import "fmt"

func main() {

	inputs := []int{3, 4, 5, 2}

	searchOrInsert := searchInsert(inputs, 2)

	fmt.Println(searchOrInsert)
}

func searchInsert(nums []int, target int) int {

	var targetIndex int

	for i, v := range nums {
		if v == target {
			targetIndex = i
			break
		} else {
			// TODO:implement insert missin value and return index
		}
	}

	return targetIndex
}
