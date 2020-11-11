package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("hello")



	
}

func maxSpan(nums []int) int {

	/*

	Consider the leftmost and righmost
	appearances of some value in an
	array. We'll say that the "span"
	is the number of elements between
	the two inclusive. A single value
	has a span of 1. Returns the
	largest span found in the given
	array. (Efficiency is not a priority.)


	maxSpan([1, 2, 1, 1, 3]) → 4
	maxSpan([1, 4, 2, 1, 4, 1, 4]) → 6
	maxSpan([1, 4, 2, 1, 4, 4, 4]) → 6
	 */

	max := 0

	if len(nums) == 1 {
		return 1
	}

	for i := 0; i < len(nums); i++ {

		cur := nums[i]

		for j := len(nums)-1; j > 0; j-- {

			if nums[j] == cur {

				max = int(math.Max(float64(max), float64(j-i+1)))

			}
		}
	}

	return max
}
