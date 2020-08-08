package main

import "fmt"

func main() {
	啊, 逼 := 1, 2
	啊 ^= 逼
	逼 ^= 啊
	啊 ^= 逼
	fmt.Println(啊)
	fmt.Println(逼)
}
