package rectangle_area

func ComputeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {

	S1 := (ay2 - ay1) * (ax2 - ax1)

	S2 := (by2 - by1) * (bx2 - bx1)

	total := S1 + S2

	if bx1 > ax2 || bx2 < ax1 {
		return total
	}

	if by1 > ay2 || by2 < ay1 {
		return total
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	r := (min(bx2, ax2) - max(ax1, bx1)) * (min(by2, ay2) - max(by1, ay1))

	return total - r
}
