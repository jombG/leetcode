package main

import "fmt"

// "Hello World"
//    fly me   to   the moon

func lengthOfLastWord(s string) int {
	lastLen := 0

	counter := 0
	for _, v := range s {
		if v == ' ' {
			if counter != 0 {
				lastLen = counter
			}
			counter = 0
		} else {
			counter++
		}
	}

	if counter != 0 {
		return counter
	}

	return lastLen
}

func main() {
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}
