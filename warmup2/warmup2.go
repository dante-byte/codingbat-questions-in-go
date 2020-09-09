package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("warmup2")




	fmt.Print(has271([]int{1,2,7,1}))



}

func countXX(str string) int {

	c := 0;

	for i := 0; i < len(str)-1; i++ {

		if str[i] == 'x' && str[i+1] == 'x' {

			c++
		}
	}

	return c
}

func frontTimes(str string, n int) string {

	var rult string

	len := len(str)

	for i := 0; i < n; i++ {

		if n < len-1 {

			rult += str[0:3];
		} else {

			rult += str;
		}

	}

	return rult
}

func stringTime(str string, n int) string {

	/*
	Given a string and a non-negative int
	n, return a larger string that is n
	copies of the original string.

	stringTimes("Hi", 2) → "HiHi"
	stringTimes("Hi", 3) → "HiHiHi"
	stringTimes("Hi", 1) → "Hi"
	 */

	rult := ""

	for i := 0; i < n; i++ {

		rult += str
	}

	return rult
}




func noTriples(nums []int) bool {

	/*
	Given an array of ints, we'll say
	that a triple is a value appearing
	3 times in a row in the array.
	Return true if the array does not
	contain any triples.

	noTriples([1, 1, 2, 2, 1]) → true
	noTriples([1, 1, 2, 2, 2, 1]) → false
	noTriples([1, 1, 1, 2, 2, 2, 1]) → false
	 */



	n := len(nums)

	if n < 3 {

		return true
	}

	for i := 0; i<n; i++ {

		cur := nums[i]

		if i < n-2 && nums[i] == cur && nums[i+1] == cur && nums[i+2] == cur {
			return false
		}
	}

	return true

}

func has271(nums []int) bool  {

	/*
	Given an array of ints, return
	true if it contains a 2, 7, 1 pattern: a value, followed by the value plus 5, followed by the value minus 1. Additionally the 271 counts even if the "1" differs by 2 or less from the correct value.

	has271([1, 2, 7, 1]) → true
	has271([1, 2, 8, 1]) → false
	has271([2, 7, 1]) → true
	 */


	len := len(nums)

	for i := 0; i < len-2; i++ {

		cur := nums[i]

		if nums[i] == nums[i+1] && math.Abs(float64(nums[i+2] - (cur-1))) <= 2 {

			return true

		}

	}

	return false

}
