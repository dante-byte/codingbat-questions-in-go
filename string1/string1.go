package main

import (
	"fmt"
)

func main() {

	fmt.Println("string1")

	//var nums1 []int
	//var nums2 []int

	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}

	m := 3
	n := 3

	head := m-1
	tail := n-1

	//for i := len(nums1); i >= 0; i-- {
	//
	//
	//
	//	if nums1[head] > nums2[head] {
	//
	//		nums1[i] = nums1[head]
    //
	//	} else {
	//
	//		nums1[i] = nums2[tail]
	//	}
	//}

	i := (m + n)-1

	for i >= 0 {

		if nums1[head] > nums2[head] {

			fmt.Println("test.....")

			if n == 1 {
				break
			}

			nums1[i] = nums1[head]
			head--

		} else {

			nums1[i] = nums2[tail]
			tail--
		}





		i--
	}

	fmt.Println(nums1)




	
}

func helloName(name string) string {

	return "Hello " + name + "!"
}

func makeAbba(a string, b string) string {

	return a + b + b + a
}

func makeTags(tag string, word string) string {

	return "<" + tag + ">" + word + "</" + tag + ">"
}

func makeOutWord(out string, word string) string {

	/*
	Given an "out" string length 4, such as "<<>>",
	and a word, return a new string where the word
	is in the middle of the out string, e.g. "<<word>>".
	Note: use str.substring(i, j) to extract the String
	starting at index i and going up to but not including index j.

	makeOutWord("<<>>", "Yay") → "<<Yay>>"
	makeOutWord("<<>>", "WooHoo") → "<<WooHoo>>"
	makeOutWord("[[]]", "word") → "[[word]]"

	 */

	return out[0:2] + word + out[len(out)-2:len(out)]
}




