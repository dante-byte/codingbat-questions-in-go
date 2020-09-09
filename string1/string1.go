package main

import (
	"fmt"
)

func main() {

	fmt.Println("string1")
	
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




