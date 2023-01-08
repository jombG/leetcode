package fizz_buzz

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i := 0; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}

	return res
}
