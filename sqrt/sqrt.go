package sqrt

func mySqrt(x int) int {
	pow := 1

	for x >= pow {
		pow = pow * pow
		pow++
	}

	return pow
}
