package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	countChar := make(map[rune]int)
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	fmt.Print("input : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	segments := []string{}
	numSegments := 4
	segmentLength := len(text) / numSegments
	for i := 0; i < len(text); i += segmentLength {
		last := i + segmentLength
		if last > len(text) {
			last = len(text)
		}
		segments = append(segments, text[i:last])
	}

	for _, segment := range segments {
		wg.Add(1)
		go func(segment string) {
			defer wg.Done()
			processSegment(segment, &mutex, countChar)
		}(segment)
	}

	wg.Wait()

	for char, count := range countChar {
		fmt.Printf("%c : %d\n", char, count)
	}
}

func processSegment(segment string, mutex *sync.Mutex, countChar map[rune]int) {
	for _, char := range segment {
		lowerChar := rune(strings.ToLower(string(char))[0])
		if lowerChar >= 'a' && lowerChar <= 'z' {
			mutex.Lock()
			countChar[char]++
			mutex.Unlock()
		}
	}
}
