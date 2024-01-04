package main

import (
	"fmt"

	"adibhanna/palindrome/pkg/palindrome"
)



func main() {
    input := "babad"
    result := palindrome.LongestPalindrome(input)
    fmt.Println("The longest palindromic substring is:", result)
}
