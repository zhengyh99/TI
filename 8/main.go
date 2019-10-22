package main

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

// 示例 1：

// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：

// 输入: "cbbd"
// 输出: "bb"

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 || (len(s) == 2 && s[1] == s[0]) {
		return s
	}
	var tmpMax, maxLen, start, end, minLen int
	for i := 0; i < len(s); i++ {
		if i-0 > len(s)-i-1 {
			minLen = len(s) - i - 1
		} else {
			minLen = i - 0
		}
		if i+1 < len(s) && s[i] == s[i+1] { //偶对齐

			if 2 >= maxLen {
				start = i
				end = i + 1
				maxLen = 2
			}
			for n := 1; n <= minLen; n++ {
				if i+1+n < len(s) && s[i-n] == s[i+1+n] {
					tmpMax = 2*n + 2
					if tmpMax >= maxLen {
						start = i - n
						end = i + 1 + n
						maxLen = tmpMax
					}
					continue
				} else {
					break
				}
			}
		}

		for k := 1; k <= minLen; k++ { //奇对齐
			if s[i+k] == s[i-k] && i > 0 {
				tmpMax = 2*k + 1
				if tmpMax >= maxLen {
					start = i - k
					end = i + k
					maxLen = tmpMax
				}
				continue
			} else {

				break
			}

		}

	}

	return s[start : end+1]
}
func main() {
	tString := "aaa"
	fmt.Println(tString)
	s := longestPalindrome(tString)
	fmt.Printf("type: %T,[%v]", s, s)
}

//aaa
