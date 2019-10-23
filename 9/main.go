package main

// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

// L   C   I   R
// E T O E S I I G
// E   D   H   N
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

// 请你实现这个将字符串进行指定行数变换的函数：

// string convert(string s, int numRows);
// 示例 1:

// 输入: s = "LEETCODEISHIRING", numRows = 3
// 输出: "LCIRETOESIIGEDHN"
// 示例 2:

// 输入: s = "LEETCODEISHIRING", numRows = 4
// 输出: "LDREOEIIECIHNTSG"
// 解释:

// L     D     R
// E   O E   I I
// E C   I H   N
// T     S     G
import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	b := []byte(s)
	tmp := make([][]byte, numRows)
	var n, rowN, col int
	for i := 0; i < len(b); i++ {
		if col%(numRows-1) == 0 {

			tmp[n%numRows] = append(tmp[n%numRows], b[i])
			n++
			if n%numRows == 0 {
				col++
			}

		} else {
			rowN = col % (numRows - 1)
			for j := 0; j < numRows; j++ {
				if j == numRows-rowN-1 {

					tmp[j] = append(tmp[j], b[i])

				}
				// else {

				// 	tmp[j] = append(tmp[j], byte(32))
				// }

				n++

			}
			col++
		}

	}
	str := make([]string, numRows)
	for key, val := range tmp {

		str[key] = string(val)

	}
	return strings.Join(str, "")
}

func main() {
	s := "AB"
	r := convert(s, 1)
	fmt.Println(r)

}
