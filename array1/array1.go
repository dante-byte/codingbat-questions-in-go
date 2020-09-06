package main

import (
	"fmt"
)




func main() {

	fmt.Println("hello")
	
}


func firstLast6(nums []int) bool {

	if nums[0] == 6 || nums[len(nums)-1] == 6 {

		return true

	}

	return false

}


func sameFirstLast(nums []int) bool  {

	/*

	Given an array of ints, return true if
	the array is length 1 or more, and the
	first element and the last element are
	equal.

	 */

	if len(nums) < 1 {

		return false

	}

	if nums[0] == nums[len(nums)] {

		return true
	}

	return false;
}

func makePi() []int {

	rult := make([]int,3)

	rult[0] = 3
	rult[1] = 1
	rult[2] = 4

	return rult;



}



