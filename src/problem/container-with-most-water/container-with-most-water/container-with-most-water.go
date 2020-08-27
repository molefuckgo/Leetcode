package main

func main() {
	
}

func maxArea(height []int) int {
	heightLen := len(height)
	left :=0
	right := heightLen - 1
	result := 0
	for left < right {
		minHeight := min2(height[left], height[right])
		result = max2(result, (right - left) * minHeight)
		if minHeight == height[left] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min2(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max2(x, y int) int {
	if x >= y{
		return x
	}
	return y
}