package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	char := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	encodeChar := strings.Split("zyxwvutsrqponmlkjihgfedcba", "")
	charSplit := strings.Split(string(s.name), "")
	tempChar := []string{}

	for _, split := range charSplit {
		for j, character := range char {
			if split == character {
				tempChar = append(tempChar, encodeChar[j])
			}
		}
	}

	nameEncode = strings.Join(tempChar, "")
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	char := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	decodeChar := strings.Split("zyxwvutsrqponmlkjihgfedcba", "")
	charSplit := strings.Split(string(s.name), "")
	tempChar := []string{}

	for _, split := range charSplit {
		for j, character := range char {
			if split == character {
				tempChar = append(tempChar, decodeChar[j])
			}
		}
	}

	nameDecode = strings.Join(tempChar, "")
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	}
}
