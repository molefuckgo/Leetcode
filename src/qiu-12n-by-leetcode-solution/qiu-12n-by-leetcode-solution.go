package main

import "fmt"

func main() {
	fmt.Println(sumNums(5))
	//fmt.Println(quickMultiBinary(7, 6))
	//fmt.Println()
}
func sumNums1(n int) int {
	answer := 0
	var sum func(int) bool
	sum = func(n int) bool {
		answer += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return answer
}

func sumNums(n int) int {
	answer, A, B := 0, n, n+1
	AddABecauseBodd := func() bool {
		answer += A
		return true
	}

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && AddABecauseBodd()
	A <<= 1
	B >>= 1
	return answer >> 1
}
func quickMulti(A, B int) int { // 俄罗斯农民乘法
	ans := 0
	for ; B > 1; B /= 2 {
		ans += (B % 2) * A
		A *= 2
	}
	return ans + A
}

func quickMultiBinary(A, B int) int { // 二进制俄罗斯农民乘法
	ans := 0
	for ; ; B >>= 1 {
		ans += A * (B & 1)
		A <<= 1
		if B&1 > 0 {
			break
		}
	}
	return ans + A
}
