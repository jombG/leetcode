package main

func maxDepth(s string) int {
	max := 0
	currentDepth := 0
	// stack := Stack[rune]{}
	for _, e := range s {
		if e == '(' {
			//stack.Push(e)
			currentDepth++
			if currentDepth > max {
				max = currentDepth
			}
		} else if e == ')' {
			currentDepth--
		}
	}

	return max
}
