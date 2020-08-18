package main

import "fmt"

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr, sc := 1, 1
	newColor := 2
	fmt.Println(floodFill(image, sr, sc, newColor))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	sColor := image[sr][sc]
	if sColor == newColor {
		return image
	}
	rows := len(image)
	columns := len(image[0])
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	for i := 0; i < len(queue); i++ {
		if image[queue[i][0]][queue[i][1]] == sColor {
			image[queue[i][0]][queue[i][1]] = newColor
		}
		if queue[i][0]-1 >= 0 && queue[i][0]-1 < rows && image[queue[i][0]-1][queue[i][1]] == sColor {
			queue = append(queue, []int{queue[i][0] - 1, queue[i][1]})
		}
		if queue[i][0]+1 >= 0 && queue[i][0]+1 < rows && image[queue[i][0]+1][queue[i][1]] == sColor {
			queue = append(queue, []int{queue[i][0] + 1, queue[i][1]})
		}
		if queue[i][1]-1 >= 0 && queue[i][1]-1 < columns && image[queue[i][0]][queue[i][1]-1] == sColor {
			queue = append(queue, []int{queue[i][0], queue[i][1] - 1})
		}
		if queue[i][1]+1 >= 0 && queue[i][1]+1 < columns && image[queue[i][0]][queue[i][1]+1] == sColor {
			queue = append(queue, []int{queue[i][0], queue[i][1] + 1})
		}
	}
	return image

}
