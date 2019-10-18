package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var tmpMax, maxLen, start, end int
	for i := 2; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			left := i - j
			right := i + j
			fmt.Println("left:", left, "right:", right)
			if s[left] != s[right] {
				tmpMax = right - left
				start = left
				end = right
			}

			if tmpMax > maxLen {

				maxLen = tmpMax
			}
			fmt.Println("maxlen", maxLen)
		}

	}
	return s[start:end]

}
func main() {
	s := longestPalindrome("cac")
	fmt.Printf("type: %T,[%v]", s, s)
}

//aaa
