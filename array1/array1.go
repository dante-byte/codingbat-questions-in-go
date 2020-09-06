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

	//rult := make([]int,3)
	//
	//rult[0] = 3
	//rult[1] = 1
	//rult[2] = 4
	//
	//return rult;


	rult := []int {3,1,4}

	return rult



}

func commonEnd(a []int, b []int) bool {

	if len(a) < 1 || len(b) < 1 {

		return false

	}

	if a[0] == b[0] || a[len(a)-1] == b[len(b)-1] {

		return true
	}

	return false



}

func sum3(nums []int) int {

	/*
	Given an array of ints length 3, return the sum of all the elements.

	sum3([1, 2, 3]) → 6
	sum3([5, 11, 2]) → 18
	sum3([7, 0, 0]) → 7
	 */

	sum := 0


	for i := 0; i < len(nums); i++ {

		sum += nums[i];
	}

	return sum
}


func rotateLeft3(nums []int) []int {

	/*
	Given an array of ints length 3, return an array with the elements "rotated left" so {1, 2, 3} yields {2, 3, 1}.

	rotateLeft3([1, 2, 3]) → [2, 3, 1]
	rotateLeft3([5, 11, 9]) → [11, 9, 5]
	rotateLeft3([7, 0, 0]) → [0, 0, 7]
	 */

	rotateTwoHelper(0,1,nums)
	rotateTwoHelper(1,2,nums)

	return nums
}

func rotateTwoHelper(start int, end int, nums []int)  {


	for start <= end-1 {

		cur := nums[start]
		nums[start] = nums[end]
		nums[end] = cur
		start++
		end--


	}
}


func reverse3(nums []int) []int {

	/*
	Given an array of ints length 3, return a new array with the elements in reverse order, so {1, 2, 3} becomes {3, 2, 1}.

	reverse3([1, 2, 3]) → [3, 2, 1]
	reverse3([5, 11, 9]) → [9, 11, 5]
	reverse3([7, 0, 0]) → [0, 0, 7]
	 */

	start := 0
	end := len(nums)-1

	for start <= end {

		cur := nums[start]
		nums[start] = nums[end]
		nums[end] = cur
		start++
		end--
	}

	return nums
}








