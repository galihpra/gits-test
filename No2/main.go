package main

import "fmt"

func balancedBracket(bracket string) string {
	var result = "YES"
	var tmp []rune

	mapBracket := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, v := range bracket {
		switch v {
		case '(', '{', '[':
			tmp = append(tmp, v)
		case ')', '}', ']':
			if len(tmp) == 0 || mapBracket[tmp[len(tmp)-1]] != v {
				result = "NO"
			}
			tmp = tmp[:len(tmp)-1]
		}
	}

	return result
}
func main() {
	brackets := "{ [ ( ) ] }"
	fmt.Println(balancedBracket(brackets))

	brackets = "{ [ ( ] ) }"
	fmt.Println(balancedBracket(brackets))

	brackets = "{ ( ( [ ] ) [ ] ) [ ] }"
	fmt.Println(balancedBracket(brackets))

}
