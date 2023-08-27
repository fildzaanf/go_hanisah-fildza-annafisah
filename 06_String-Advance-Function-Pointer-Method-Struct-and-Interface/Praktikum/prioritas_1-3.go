package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {

	result := ""

	for i := range b {

		subStr := b[:len(b)-i]

		if strings.Contains(a, subStr) {
			result = subStr
			break
		}
	}

	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}
