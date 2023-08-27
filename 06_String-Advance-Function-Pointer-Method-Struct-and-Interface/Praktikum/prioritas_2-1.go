package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	result := ""
	input = strings.ToLower(input)

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			result += string('a' + (char-'a'+int32(offset))%26)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
