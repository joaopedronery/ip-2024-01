package main

import (
	"fmt"
)

func main() {
	var phrases []string
	var input string

	for true {
		fmt.Scanln(&input)
		if input == "#" {
			break
		}
		phrases = append(phrases, input)
	}
	for _, phrase := range phrases {
		fmt.Println(reverseString(phrase))
	}
}

func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}