package main

func hammingWeight(num uint32) int {
	count := 0
	mask := uint32(1)
	for i := 0; i < 32; i++ {
		if (num & mask) != 0 {
			count++
		}
		mask = mask << 1
	}
	return count
}

func main() {
}
