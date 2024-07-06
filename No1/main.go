package main

import "fmt"

func charScale(c byte) int {
	return int(c-'a') + 1
}

func weightStrings(s string, arr []int) []string {
	weightMap := make(map[int]bool)
	var weight int

	for i := 0; i < len(s); i++ {
		weight = 0
		for j := i; j < len(s) && s[j] == s[i]; j++ {
			weight += charScale(s[j])
			weightMap[weight] = true
		}
	}

	result := make([]string, len(arr))
	for i, num := range arr {
		if weightMap[num] {
			result[i] = "Yes"
		} else {
			result[i] = "No"
		}
	}

	return result
}

func main() {
	s := "abbcccd"
	arr := []int{1, 3, 9, 8}
	result := weightStrings(s, arr)
	fmt.Println(result)

	s = "aaabccc"
	arr = []int{1, 2, 3, 6, 8, 10}
	result = weightStrings(s, arr)
	fmt.Println(result)

	s = "dddeefgg"
	arr = []int{4, 5, 6, 10, 12, 14}
	result = weightStrings(s, arr)
	fmt.Println(result)
}
