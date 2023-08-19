package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(text string) bool {
	text = strings.ToLower(text)

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("input text : ")
	scanner.Scan()
	text := scanner.Text()

	if isPalindrome(text) {
		fmt.Printf("captured : %s \nPalindrome", text)
	} else {
		fmt.Printf("captured : %s \nBukan Palindrome", text)
	}

}
