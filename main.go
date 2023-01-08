package main

import "fmt"

func main() {
	for _, ch := range "words" {
		fmt.Println(ch - 'a')
	}
}
