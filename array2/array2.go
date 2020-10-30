package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println("hello world")

	n := 2

	if n % 10 == 0 {

		fmt.Print("sound evidence")
	}





	//fmt.Print(has12([]int {1,2,3}))









	
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

func sum13(nums []int) int {

	/*
		Return the sum of the numbers in
		the array, returning 0 for an empty
		array. Except the number 13 is very
		unlucky, so it does not count and numbers
		that come immediately after a 13 also do not count.

		sum13([1, 2, 2, 1]) → 6
		sum13([1, 1]) → 2
		sum13([1, 2, 2, 1, 13]) → 6
	*/

	sum := 0

	for i := 0; i < len(nums); i++ {


		if i != len(nums) && nums[i] == 13 {

			i++;
		} else {

			sum += nums[i]
		}
	}

	return sum
}

func sum67(nums []int) int {


	/*
		Return the sum of the numbers in
		the array, except ignore sections
		of numbers starting with a 6 and
		extending to the next 7 (every 6
		will be followed by at least one 7).
		Return 0 for no numbers.

		sum67([1, 2, 2]) → 5
		sum67([1, 2, 2, 6, 99, 99, 7]) → 5
		sum67([1, 1, 6, 7, 2]) → 4
	*/

	sum := 0


	for i := 0; i < len(nums); i++ {

		if nums[i] == 6 {

			for nums[i] == 6 {

				i++
			}
		} else {

			sum += nums[i]
		}
	}

	return sum
}


func has22(nums []int) bool {

	/*
		Given an array of ints, return true
		if the array contains a 2 next to a
		2 somewhere.

		has22([1, 2, 2]) → true
		has22([1, 2, 1, 2]) → false
		has22([2, 1, 2]) → false
	*/

	for i := 1; i < len(nums); i++ {

		if nums[i] == 2 && nums[i-1] == 2 {

			return true
		}
	}

	return false
}

func lucky13(nums []int) bool {


	/*
	Given an array of ints, return
	true if the array contains no 1's and no 3's.

	lucky13([0, 2, 4]) → true
	lucky13([1, 2, 3]) → false
	lucky13([1, 2, 4]) → false

	 */



	i := 0

	for i < len(nums) {

		if nums[i] == 1 || nums[i] == 13 {

			return false

		}

		i++
	}

	return true
}


func sum28(nums []int) bool {

	/*
	Given an array of ints, return true
	if the sum of all the 2's in the
	array is exactly 8.

	sum28([2, 3, 2, 2, 4, 2]) → true
	sum28([2, 3, 2, 2, 4, 2, 2]) → false
	sum28([1, 2, 3, 4]) → false
	 */

	sum := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 2 {

			sum += nums[i]
		}
	}

	return sum == 8
}

func more14(nums []int) bool {

	/*
	Given an array of ints, return true
	if the number of 1's is greater than
	the number of 4's

	more14([1, 4, 1]) → true
	more14([1, 4, 1, 4]) → false
	more14([1, 1]) → true
	 */

	one := 0
	four := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {

			one++;
		}

		if nums[i] == 4 {

			four++;
		}
	}

	return one > four
}

func fizzArray(n int) []string {

	/*
	Given a number n, create and return
	a new int array of length n, containing
	the numbers 0, 1, 2, ... n-1. The given
	n may be 0, in which case just return
	a length 0 array. You do not need a
	separate if-statement for the length-0

	case; the for-loop should naturally
	execute 0 times in that case, so it
	just works. The syntax to make a new
	int array is: new int[desired_length]
	(See also: FizzBuzz Code)

	fizzArray(4) → [0, 1, 2, 3]
	fizzArray(1) → [0]
	fizzArray(10) → [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	 */



	rult :=  []string{}

	for i := 0; i < n; i++ {

		rult[i] = strconv.Itoa(i)
	}

	return rult

}

func only14(nums []int) bool {

	/*
	Given an array of ints, return
	true if every element is a 1 or a 4.

	only14([1, 4, 1, 4]) → true
	only14([1, 4, 2, 4]) → false
	only14([1, 1]) → true
	 */

	i := 0

	for i < len(nums) {

		if nums[i] != 4 && nums[i] != 1 {

			return false

		}

		i++
	}

	return true
}


func no14(nums []int) bool {

	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {

			j := i

			for j < len(nums) {

				if nums[j] == 4 {

					return false

				}
				j++
			}
		}

		if nums[i] == 4 {

			j := i

			for j < len(nums) {

				if nums[j] == 1 {

					return false

				}
				j++
			}

		}
	}

	return true
}


func isEveryWhere(nums []int, val int) bool {

	if len(nums) < 1 {


		return true
	}

	n := 0


	if nums[0] == val {

		n = nums[0]
	} else {

		n = nums[1]
	}

	for i := n; i < len(nums); i+=2 {

		if nums[i] != val {

			return false

		}
	}

	return true
}

func either24(nums []int) bool {

	four := 0
	two := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 4 && nums[i-1] == 4 {
			four++

		}

		if nums[i] == 2 && nums[i-1] == 2 {

			two++

		}
	}

	if four > 0 && two == 0 || two > 0 && four == 0 {

		return true
	}

	return false
}

func matchUp(nums1 []int, nums2 []int) int {

	cnt := 0

	for i := 0; i < len(nums1); i++ {

		if math.Abs(float64(nums1[i]) - float64(nums2[i])) <= 2 && nums2[i] != nums1[i] {

			cnt++

		}
	}

	return cnt
}


func has12(nums []int) bool {

	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {

			j := i

			for j < len(nums) {

				if nums[j] == 2 {

					return true

				}

				j++
			}

		}
	}

	return false
}


func modThree(nums []int) bool {

	if len(nums) < 2 {

		return false
	}


	for i := 2; i < len(nums); i++ {

		if nums[i]  % 2 == 0 && nums[i-1] % 2 == 0 && nums[i-2] % 2 == 0 {

			return true

		}

		if nums[i]  % 2 == 1 && nums[i-1] % 2 == 1 && nums[i-2] % 2 == 1 {

			return true

		}



	}

	return false
}

func haveThree(nums []int) bool {

	if len(nums) < 4 {

		return false

	}

	rult := false
	cnt := 0

	j := 0

	for j < len(nums) {

		if nums[j] == 3  {

			cnt++

		} else {

			if cnt > 3 {

				return false

			}
		}
		j++
	}

	for i := 4; i < len(nums); i++ {

		if nums[i] == 3 && nums[i-2] == 3 && nums[i-4] == 3{

			rult = true

		}
	}

	if cnt == 3 && rult {

		return true
	}

	return false
}

func twoTwo(nums []int) bool {

	/*
	Given an array of ints, return true if every 2
	that appears in the array is next to another 2.

	twoTwo([4, 2, 2, 3]) → true
	twoTwo([2, 2, 4]) → true
	twoTwo([2, 2, 4, 2]) → false
	 */
	
	rult := false
	two := true

	if len(nums) == 1 && nums[0] == 2 {
		return false
	}
	
	for i := 1; i < len(nums); i++ {

		if nums[i] == 2 {
			
			two = false

		}

		if i < len(nums)-1 && nums[i] != 2 && nums[i+1] == 2  {

			if nums[i-1] != 2 || nums[i-1] == 2 {
				return false
				}
		} else {

			if nums[i] == 2 && nums[i-1] == 2 {

				rult = true

			}
		}
	}

	if !two {

		return rult
	}

	return two
}

func sameEnds2(nums []int, v int) bool {

	l := len(nums)

	for i := 0; i<v; i++ {

		if nums[i] != nums[l-v+1] {

			return false
		}
	}

	return true
}

func tripleUp(nums []int) bool {

	for i := 0; i<len(nums)-2; i++ {

		if nums[i] == nums[i+1]-1 && nums[i+1] == nums[i+2]-1 {

			return false
		}
	}

	return true

}

func fizzArray3(start int, end int) []int {

	rult := make([]int,end-start)


	size := len(rult)
	for i := 0; i<size; i++ {

		rult[i] = start
		start++

	}

	return rult
}

func shiftLeft(nums []int) []int {

	if len(nums) <= 1 {

		return nums
	}

	i := 0

	for i < len(nums) {

		cur := nums[i];
		nums[i] = nums[i+1]
		nums[i+1] = cur
		i++
	}

	return nums
}


func tenRun(nums []int) []int {

	if len(nums) < 1 {
		return nums
	}

	cur := nums[0]

	i := 0

	for i < len(nums) {

		if nums[i] % 10 == 0 {

			cur = nums[i]

		}

		if cur != 1 {

			nums[i] = cur
		}

		i++
	}

	return nums
}

func pre4(nums []int) []int {

	n := 0

	for i := 0; i<len(nums); i++ {

		if nums[i] != 4 {

			n++

		} else {

			break
		}
	}

	rult := make([]int,n)

	for i := 0; i<len(rult); i++ {

		rult[i] = nums[i]
	}

	return rult
}

func post4(nums []int) []int {

	none :=  make([]int,0)

	for i := len(nums); i>0; i-- {

		if nums[i] == 4 {

			rult := make([]int,i)

			for j := 0; j<len(rult); j++ {

				rult[j] = nums[i + j + 1]
			}

		}

		return none
	}
}





