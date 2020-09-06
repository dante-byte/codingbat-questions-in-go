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
