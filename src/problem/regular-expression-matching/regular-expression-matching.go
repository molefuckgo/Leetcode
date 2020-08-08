package main

import "fmt"

func main() {
	fmt.Println(isMatch("bcc", "a*bc*"))
}

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	if pLen == 0 {
		return sLen == 0
	}
	firstMatch := sLen != 0 && (p[0] == '.' || p[0] == s[0])
	if pLen >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}

}
