package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else if x == 0 {
		return true
	}
	x_resrve := 0
	for ; x > x_resrve; x /= 10 {
		x_resrve = x_resrve*10 + x%10

	}
	return x_resrve == x || x_resrve/10 == x

}
