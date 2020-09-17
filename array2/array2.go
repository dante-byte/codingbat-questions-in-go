package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("hello world")



	
}


func countEvens(nums []int) int {


	i := 0
	cnt := 0

	for i < len(nums) {

		if nums[i] % 2 == 0 {

			cnt++
		}

		i++
	}

	return cnt
}

func bigDiff(nums []int) int {

	/*
	Given an array length 1 or more
	of ints, return the difference
	between the largest and smallest
	values in the array. Note: the
	built-in Math.min(v1, v2) and
	Math.max(v1, v2) methods return the
	smaller or larger of two values.

	bigDiff([10, 3, 5, 6]) → 7
	bigDiff([7, 2, 10, 9]) → 8
	bigDiff([2, 10, 7, 2]) → 8
	 */



	min := nums[0]
	max := nums[0]

	i := 0

	for i < len(nums) {

		min = int(math.Min(float64(min),float64(nums[i])))

		max = int(math.Max(float64(min),float64(nums[i])))

	}

	rult := int(math.Abs(float64(min) - float64(max)))

	return rult
}

func centeredAverage(nums []int) int {

	/*
	Return the "centered" average of an array
	of ints, which we'll say is the mean
	average
	of the values, except ignoring the largest
	and smallest values in the array. If there
	are multiple copies of the smallest value,
	ignore just one copy, and likewise for the
	largest value. Use int division to produce
	the final average. You may assume that the
	array is length 3 or more.

	centeredAverage([1, 2, 3, 4, 100]) → 3
	centeredAverage([1, 1, 5, 5, 10, 8, 7]) → 5
	centeredAverage([-10, -4, -2, -4, -2, 0]) → -3
	 */

	max := nums[0]
	min := nums[0]
	avg := 0

	for i := 0; i < len(nums); i++ {

		min = int(math.Min(float64(min), float64(nums[i])))

		max = int(math.Max(float64(max), float64(nums[i])))

		avg += nums[i]
	}

	avg = avg - (max - min)

	n := len(nums)-2

	avg = avg / n

	return avg
}
