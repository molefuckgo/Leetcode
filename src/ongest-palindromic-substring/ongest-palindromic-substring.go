package main

import (
	"fmt"
)

func main() {
	s := "012345"
	fmt.Println(string(s[0]) + string(s[1]))
}

// public static String Manacher(String s) {
//    if (s.length() < 2) {
//        return s;
//    }
//    // 第一步：预处理，将原字符串转换为新字符串
//    String t = "$";
//    for (int i=0; i<s.length(); i++) {
//        t += "#" + s.charAt(i);
//    }
//    // 尾部再加上字符@，变为奇数长度字符串
//    t += "#@";
//    // 第二步：计算数组p、起始索引、最长回文半径
//    int n = t.length();

//    // p数组
//    int[] p = new int[n];
//    int id = 0, mx = 0;
//    // 最长回文子串的长度
//    int maxLength = -1;
//    // 最长回文子串的中心位置索引
//    int index = 0;

//    for (int j=1; j<n-1; j++) {
//        // 参看前文第五部分
//        p[j] = mx > j ? Math.min(p[2*id-j], mx-j) : 1;
//        // 向左右两边延伸，扩展右边界
//        while (t.charAt(j+p[j]) == t.charAt(j-p[j])) {
//            p[j]++;
//        }
//        // 如果回文子串的右边界超过了mx，则需要更新mx和id的值
//        if (mx < p[j] + j) {
//            mx = p[j] + j;
//            id = j;
//        }

//        // 如果回文子串的长度大于maxLength，则更新maxLength和index的值
//        if (maxLength < p[j] - 1) {
//            // 参看前文第三部分
//            maxLength = p[j] - 1;
//            index = j;
//        }
//    }

//    // 第三步：截取字符串，输出结果
//    // 起始索引的计算参看前文第四部分
//    int start = (index-maxLength)/2;
//    return s.substring(start, start + maxLength);
// }

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// // 第一步：预处理，将原字符串转换为新的字符串
	// var t  = "$"
	// for i := 0; i < len(s); i++ {
	// 	t += "#" + string(s[i])
	// }
	//
	// // 尾部再加上字符@，变为奇数长度字符串
	// t += "#@"
	// // 第二步：计算数组p、起始索引、最长回文半径
	// n := len(t)
	// // p数组
	// var p [...]int
	// id, mx := 0, 0
	// maxLength := -1
	// index := 0
	//
	// for j := 1; j < n - 1; j++ {
	// 	p[j] =
	// }
	// return ""
	return s
}

func min(m int, n int) int {
	if m < n {
		return m
	}
	return n
}
