package main

import (
	"fmt"
	"math"
)




func main() {

	fmt.Println("fmt")




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

		sum += nums[i]
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

func maxEnd3(nums []int) []int {

	/*
	Given an array of ints length 3, figure out which is larger, the first or last element in the array, and set all the other elements to be that value. Return the changed array.

	maxEnd3([1, 2, 3]) → [3, 3, 3]
	maxEnd3([11, 5, 9]) → [11, 11, 11]
	maxEnd3([2, 11, 3]) → [3, 3, 3]
	 */


	n := math.Max(float64(nums[0]), float64(nums[len(nums)-1]))

	for i := 0; i < len(nums); i++ {

		nums[i] = int(n)

	}

	return nums

}

func sum2(nums []int) int  {

	/*
	Given an array of ints, return the sum of the first 2 elements in the array. If the array length is less than 2, just sum up the elements that exist, returning 0 if the array is length 0.

	sum2([1, 2, 3]) → 3
	sum2([1, 1]) → 2
	sum2([1, 1, 1, 1]) → 2
	 */

	n := 0

	rult := 0

	if len(nums) == 0 {

		return 0
	} else if len(nums) == 1 {

		return nums[0]

	}

	for n < 2 {

		rult += nums[n]
		n++

	}


	return rult

}

func middleWay(a []int, b []int) []int  {

	/*
	Given 2 int arrays, a and b, each length 3,
	return a new array length 2 containing their
	middle elements.

	middleWay([1, 2, 3], [4, 5, 6]) → [2, 5]
	middleWay([7, 7, 7], [3, 8, 0]) → [7, 8]
	middleWay([5, 2, 9], [1, 4, 5]) → [2, 4]
	 */

	n := len(a) / 2



	rult := []int{a[n],b[n]}

	return rult

}

func makeEnds(nums []int) []int {


	/*
	Given an array of ints, return a new
	array length 2 containing the first
	and last elements from the original
	array. The original array will be length
	1 or more.

	makeEnds([1, 2, 3]) → [1, 3]
	makeEnds([1, 2, 3, 4]) → [1, 4]
	makeEnds([7, 4, 6, 2]) → [7, 2]
	 */

	rult := []int{nums[0],nums[len(nums)-1]}

	return rult

}

func has23(nums []int) bool {

	/*
	Given an int array length 2,
	return true if it contains a
	2 or a 3.

	has23([2, 5]) → true
	has23([4, 3]) → true
	has23([4, 5]) → false
	 */

	if nums == nil {

		return false
	}

	i := 0

	for i < len(nums) {

		if nums[i] == 2 || nums[i] == 3 {

			return true

		}
		i++
	}

	return false

}


func no23(nums []int) bool {

	/*
	Given an int array length 2, return
	true if it does not contain a 2 or 3.

	no23([4, 5]) → true
	no23([4, 2]) → false
	no23([3, 5]) → false
	 */

	if nums == nil {

		return false

	}

	var i int = 0

	for i < len(nums) {

		if nums[i] == 2 || nums[i] == 3 {

			return false
		}
		i++
	}

	return true
}


func makeLast(nums []int) []int {

	/*
	Given an int array, return a new
	array with double the length where
	its last element is the same as the
	original array, and all the other
	elements are 0. The original array will
	be length 1 or more. Note: by default,
	a new int array contains all 0's.

	makeLast([4, 5, 6]) → [0, 0, 0, 0, 0, 6]
	makeLast([1, 2]) → [0, 0, 0, 2]
	makeLast([3]) → [0, 3]
	 */



	rult := make([]int,len(nums) * 2)

	rult[len(rult)-1] = nums[len(nums)-1]

	return rult
}

func double23(nums []int) bool {

	/*
	Given an int array, return true if
	the array contains 2 twice, or 3 twice.
	The array will be length 0, 1, or 2.

	double23([2, 2]) → true
	double23([3, 3]) → true
	double23([2, 3]) → false
	 */

	var two int = 0
	var three int = 0

	i := 0

	for i < len(nums) {

		if nums[i] == 2 {

			two++

		}

		if nums[i] == 3 {

			three++

		}
		i++
	}

	if two >= 2 || three >= 3 {

		return true

	}

	return false



}

func fix23(nums []int) []int {


	i := len(nums)

	for i > 0 {

		if nums[i] == 3 && nums[i-1] == 2 {

			nums[i] = 0

		}

		i--
	}

	return nums

}


func start1(a []int, b []int) int {

	/*
	Start with 2 int arrays, a and
	b, of any length. Return how
	many of the arrays have 1 as
	their first element.

	start1([1, 2, 3], [1, 3]) → 2
	start1([7, 2, 3], [1]) → 1
	start1([1, 2], []) → 1
	 */

	count := 0

	if len(a) != 0 && a[0] == 1 {

		count++
	}

	if len(b) != 0 && a[0] == 1 {
		count++
	}

	return count
}

func biggerTwo(a []int, b []int) []int {

	suma := 0
	sumb := 0

	for i := 0; i < len(a); i++ {

		suma += a[i];
		sumb += b[i];
	}

	if suma > sumb || suma == sumb {

		return a
	}

	return b
}


func makeMiddle(nums []int) []int {


	/*
	Given an array of ints of even
	length, return a new array length
	2 containing the middle two elements
	from the original array. The original
	array will be length 2 or more.

	makeMiddle([1, 2, 3, 4]) → [2, 3]
	makeMiddle([7, 1, 2, 3, 4, 9]) → [2, 3]
	makeMiddle([1, 2]) → [1, 2]
	 */

	len := len(nums)/2


	rult := []int{nums[len-1],nums[len]}

	return rult
}


func plusTwo(a []int, b []int) []int {

	/*
	Given 2 int arrays, each length 2,
	return a new array length 4
	containing all their elements.

	plusTwo([1, 2], [3, 4]) → [1, 2, 3, 4]
	plusTwo([4, 4], [2, 2]) → [4, 4, 2, 2]
	plusTwo([9, 2], [3, 4]) → [9, 2, 3, 4]
	 */

	rult := make([]int, len(a) + len(b))

	alen := 0
	blen := 0

	i := 0

	for i < len(rult) {

		if i < 2 {

			rult[i] = a[alen]
			alen++
		} else {

			rult[i] = b[alen]
			blen++
		}

		i++
	}

	return rult



}

func swapEnds(nums []int) []int {
	/*
	Given an array of ints, swap
	the first and last elements
	in the array. Return the modified
	array. The array length will be
	at least 1.

	swapEnds([1, 2, 3, 4]) → [4, 2, 3, 1]
	swapEnds([1, 2, 3]) → [3, 2, 1]
	swapEnds([8, 6, 7, 9, 5]) → [5, 6, 7, 9, 8]
	 */

	first := nums[0]
	nums[0] = nums[len(nums)-1]
	nums[len(nums)-1] = first
	return nums
}


func midThree(nums []int) []int {

	/*
	Given an array of ints of odd
	length, return a new array
	length 3 containing the elements
	from the middle of the array.
	The array length will be at least 3.

	midThree([1, 2, 3, 4, 5]) → [2, 3, 4]
	midThree([8, 6, 7, 5, 3, 0, 9]) → [7, 5, 3]
	midThree([1, 2, 3]) → [1, 2, 3]
	 */

	len := len(nums)/2

	rult := make([]int,3)

	rult = []int{nums[len-1],nums[len],nums[len+1]}

	return rult
}


func maxTriple(nums []int) int {

	/*
	Given an array of ints of odd length,
	look at the first, last, and middle
	values in the array and return the
	largest. The array length will be a
	least 1.

	maxTriple([1, 2, 3]) → 3
	maxTriple([1, 5, 3]) → 5
	maxTriple([5, 2, 3]) → 5
	 */

	max := nums[0]
	i := 0

	for i < len(nums) {


		max = int(math.Max(float64(max), float64(nums[i])))
	}

	return max
}









