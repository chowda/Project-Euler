/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"strconv"
)

func main() {
	max := 0
	for i := 100; i < 1000; i++ {
		for n := 100; n < 1000; n++ {
			product := i * n
			if product > max && isPalindrome(strconv.Itoa(product)) {
				max = product
			}
		}
	}

	println("The answer is: ", max)
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
