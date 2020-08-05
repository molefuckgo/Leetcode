package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(countAndSay(i))
	}
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		last := countAndSay(n - 1)
		first := last[0]
		num := 1
		var buffer bytes.Buffer
		for i := 1; i < len(last); i++ {
			if last[i] == first {
				//相等,计数就可以
				num++
			} else {
				//不相等,将上个记录的值和数写入buffer,更新新的值和数
				buffer.WriteString(strconv.Itoa(num))
				buffer.WriteString(string(first))
				first = last[i]
				num = 1
			}
		}
		//将最后记录写入
		buffer.WriteString(strconv.Itoa(num))
		buffer.WriteString(string(first))
		return buffer.String()
	}
}
