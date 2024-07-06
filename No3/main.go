package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func findHighestPalindrome(s string, k int, i, j int) string {
	if k < 0 {
		return "-1"
	}
	if i >= j {
		return s
	}

	if s[i] == s[j] {
		return findHighestPalindrome(s, k, i+1, j-1)
	}

	option1 := findHighestPalindrome(s[:i]+string(s[j])+s[i+1:], k-1, i+1, j-1)
	option2 := findHighestPalindrome(s[:j]+string(s[i])+s[j+1:], k-1, i+1, j-1)

	if option1 == "-1" && option2 == "-1" {
		return "-1"
	} else if option1 == "-1" {
		return option2
	} else if option2 == "-1" {
		return option1
	} else {
		if strings.Compare(option1, option2) > 0 {
			return option1
		} else {
			return option2
		}
	}
}

func highestPalindrome(s string, k int) string {
	if isPalindrome(s) {
		return s
	}

	result := findHighestPalindrome(s, k, 0, len(s)-1)

	if isPalindrome(result) {
		return result
	} else {
		return "-1"
	}
}

func main() {
	s1 := "3943"
	k1 := 1
	fmt.Println(highestPalindrome(s1, k1))

	s2 := "932239"
	k2 := 2
	fmt.Println(highestPalindrome(s2, k2))

	s3 := "532235"
	k3 := 2
	fmt.Println(highestPalindrome(s3, k3))
}
