package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(9, 12))
}

func rangeBitwiseAnd2(m int, n int) int {

	for n > m {
		n &= (n - 1)
	}
	return n
}

func rangeBitwiseAnd(m int, n int) int {

	shift := 0
	for m != n {
		m >>= 1
		n >>= 1
		shift += 1
	}
	return m << shift
}
