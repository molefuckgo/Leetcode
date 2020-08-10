package main

import "fmt"

func main() {
	words := []string{"ab", "ba", "aaab", "abab", "baa"}
	obj := Constructor(words)
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('b'))
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('b'))
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('b'))
	fmt.Println(obj.Query('b'))
	fmt.Println(obj.Query('b'))
	fmt.Println(obj.Query('a'))
}

type StreamChecker struct {
	words         []string
	lastLetterMap map[byte]bool
}

func Constructor(words []string) StreamChecker {
	return StreamChecker{
		words:         words,
		lastLetterMap: make(map[byte]bool),
	}
}

func (this *StreamChecker) Query(letter byte) bool {

	for i := 0; i < len(this.words); i++ {
		if len(this.words[i]) == 0 {
			continue
		} else if len(this.words[i]) == 1 {
			if this.words[i][0] == letter {
				this.words[i] = ""
				this.lastLetterMap[letter] = true
			} else {
				continue
			}
		} else {
			if this.words[i][0] == letter {
				this.words[i] = this.words[i][1:]
			}
			continue
		}
	}

	return this.lastLetterMap[letter]
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
