package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num_list [10]int
	for i := 0; i < 100000; i++ {
		// fmt.Println(rand10())
		num_list[rand10()-1] += 1
	}
	fmt.Println(num_list)
	// var num_list [7]int
	// for i:= 0; i < 10000;i++ {
	// 	fmt.Println(rand7())
	// 	num_list[rand7() - 1] += 1
	// }
	// fmt.Println(num_list)
	// fmt.Println(rand2)
}

func rand10() int {
	row := rand7()
	col := rand7()
	if row > 4 && col < 4 {
		return rand10()
	} else {
		return (row+col)%10 + 1
	}
}

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(7) + 1
	return x
}
