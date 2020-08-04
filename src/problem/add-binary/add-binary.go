package main

func main() {

}

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	var maxLen int
	if aLen >= bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}

	return a
}
