package main

import "fmt"

func main() {
	fmt.Println(spiralMatrixIII(1, 4, 0, 0))
}

//在 R 行 C 列的矩阵上，我们从 (r0, c0) 面朝东面开始
//
// 这里，网格的西北角位于第一行第一列，网格的东南角位于最后一行最后一列。
//
// 现在，我们以顺时针按螺旋状行走，访问此网格中的每个位置。
//
// 每当我们移动到网格的边界之外时，我们会继续在网格之外行走（但稍后可能会返回到网格边界）。
//
// 最终，我们到过网格的所有 R * C 个空间。
//
// 按照访问顺序返回表示网格位置的坐标列表。
//
//
//
// 示例 1：
//
// 输入：R = 1, C = 4, r0 = 0, c0 = 0
//输出：[[0,0],[0,1],[0,2],[0,3]]
//
//
//
//
//
//
// 示例 2：
//
// 输入：R = 5, C = 6, r0 = 1, c0 = 4
//输出：[[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3
//,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0
//],[3,0],[2,0],[1,0],[0,0]]
//
//
//
//
//
//
// 提示：
//
//
// 1 <= R <= 100
// 1 <= C <= 100
// 0 <= r0 < R
// 0 <= c0 < C
//
// Related Topics 数学
// 👍 37 👎 0

// 方法1
type DirectionType struct {
	right int
	down  int
	left  int
	up    int
}

var directionType = DirectionType{right: 0, down: 1, left: 2, up: 3}

//leetcode submit region begin(Prohibit modification and deletion)
func spiralMatrixIII2(R int, C int, r0 int, c0 int) [][]int {
	positionList := make([][]int, 0)
	horizontalEnd := 1 // 记录横向转竖向前需要走的步数
	verticalEnd := 1   // 记录竖向转横向前需要走的步数

	nowDirection := directionType.right //初始化方向
	nowPositionRow := r0
	nowPositionColumn := c0
	startPositionRow := r0
	startPositionColumn := c0

	for len(positionList) < R*C { // 结束标志
		if (nowPositionRow >= 0 && nowPositionRow <= R-1) && (nowPositionColumn >= 0 && nowPositionColumn <= C-1) {
			var position = []int{nowPositionRow, nowPositionColumn}
			positionList = append(positionList, position)
		}

		nowPositionRow, nowPositionColumn = nextTodo(nowDirection, nowPositionRow, nowPositionColumn)

		if nowDirection == directionType.left || nowDirection == directionType.right {
			if nowPositionColumn-startPositionColumn == horizontalEnd || nowPositionColumn-startPositionColumn == -(horizontalEnd) {
				startPositionRow = nowPositionRow
				startPositionColumn = nowPositionColumn
				horizontalEnd++
				nowDirection = nextDirection(nowDirection)
			}
		} else {
			if nowPositionRow-startPositionRow == verticalEnd || nowPositionRow-startPositionRow == -(verticalEnd) {
				startPositionRow = nowPositionRow
				startPositionColumn = nowPositionColumn
				verticalEnd++
				nowDirection = nextDirection(nowDirection)
			}
		}

	}

	return positionList
}

func nextDirection(nowDirection int) int { // 获取下次走的方向
	if nowDirection == directionType.up {
		return directionType.right
	} else {
		return nowDirection + 1
	}
}

func nextTodo(nowDirection int, nowPositionRow int, nowPositionColumn int) (int, int) {
	if nowDirection == directionType.right { // 向右
		nowPositionColumn++
	} else if nowDirection == directionType.down { // 向下
		nowPositionRow++
	} else if nowDirection == directionType.left {
		nowPositionColumn--
	} else if nowDirection == directionType.up {
		nowPositionRow--
	}

	return nowPositionRow, nowPositionColumn
}

//方法2
func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	l := c0
	r := c0
	t := r0
	b := r0

	positionList := make([][]int, 0)
	for len(positionList) < R*C {

		if check(r0, c0, R, C) {
			var position = []int{r0, c0}
			positionList = append(positionList, position)
		}
		c0++

		for c0 <= r {
			if check(r0, c0, R, C) {
				var position = []int{r0, c0}
				positionList = append(positionList, position)
			}
			c0++
		}
		r++
		for r0 <= b {
			if check(r0, c0, R, C) {
				var position = []int{r0, c0}
				positionList = append(positionList, position)
			}
			r0++
		}
		b++

		for c0 >= l {
			if check(r0, c0, R, C) {
				var position = []int{r0, c0}
				positionList = append(positionList, position)
			}
			c0--
		}
		l--
		for r0 >= t {
			if check(r0, c0, R, C) {
				var position = []int{r0, c0}
				positionList = append(positionList, position)
			}
			r0--
		}
		t--
	}
	return positionList
}

func check(x, y int, R, C int) bool {
	if x >= 0 && x < R && y >= 0 && y < C {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
