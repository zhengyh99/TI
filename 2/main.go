package main

import (
	"fmt"
)

//字符串中出现不重复字母的长度
func lenOfNoReapeatSubString(str string) (maxLen int) {
	charOccurredLast := make(map[rune]int)
	start := 0
	for i, ch := range []rune(str) {
		if lastI, ok := charOccurredLast[ch]; ok && lastI >= start {
			start = lastI + i
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		charOccurredLast[ch] = i
	}
	return
}

func main() {
	s := "aab基本都是bcc"
	fmt.Println(lenOfNoReapeatSubString(s))

}
