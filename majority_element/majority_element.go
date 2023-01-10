package main

import "fmt"

func majorityElement(nums []int) int {
	maj, counter := 0, 0

	for _, v := range nums {
		if counter == 0 {
			maj = v
			counter++
		} else {
			if v != maj {
				counter--
			} else {
				counter++
			}
		}
	}

	return maj
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
