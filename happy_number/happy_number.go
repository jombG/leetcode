package main

func isHappy(n int) bool {
	m := make(map[int]struct{})
	currentNumber := n
	for {
		currentNumber = getSumNumbers(currentNumber)
		if currentNumber == 1 {
			return true
		}

		if _, ok := m[currentNumber]; !ok {
			m[currentNumber] = struct{}{}
		} else {
			return false
		}
	}

}

func getSumNumbers(n int) int {
	tmp := 0
	for i := n; i > 0; i = i / 10 {
		j := i % 10
		tmp += j * j
	}
	return tmp
}

func main() {
	isHappy(1111111)
}
