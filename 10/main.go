package main

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

// 示例 1:

// 输入: 123
// 输出: 321
//  示例 2:

// 输入: -123
// 输出: -321
// 示例 3:

// 输入: 120
// 输出: 21
// 注意:

// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

import (
	"fmt"
)

func reverse(x int) int {
	rs := 0
	intMax := int(^uint32(0) >> 1) //最大数
	intMin := ^intMax              //最小数
	for {
		rs = rs*10 + x%10
		x = x / 10
		if x == 0 {
			break
		}
	}

	if rs < intMin || rs > intMax {
		return 0
	}
	return rs
}

func main() {
	fmt.Println(reverse(123))
}
