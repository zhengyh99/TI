package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var tmpMax, maxLen, start, end, minLen int
	for i := 0; i < len(s); i++ {
		if i-0 > len(s)-i-1 {
			minLen = len(s) - i - 1
		} else {
			minLen = i - 0
		}
		// for j := i + 1; j < len(s); j++ {
		// 	if s[i] == s[j] {
		// 		t := 0
		// 		if j-0 > len(s)-j-1 {
		// 			minLen2 = len(s) - i - 1
		// 		} else {
		// 			minLen2 = j - 0
		// 		}
		// 		for n := j + 1; n < minLen2; n++ {

		// 			if s[i-n] == s[j+n] {
		// 				t = n
		// 			} else {
		// 				break
		// 			}
		// 		}

		// 		tmpMax = 2*t + j - i
		// 		if tmpMax > maxLen {
		// 			fmt.Println("wwwwwtmpmax:", tmpMax)
		// 			start = i - t
		// 			end = j + t
		// 			fmt.Println("111end:", end, "start:", start)
		// 			maxLen = tmpMax
		// 		}

		// 	} else {

		// 		break
		// 	}
		// }

		// if s[i]==s[i+1] and i+1<len(s){

		// }

		for k := 1; k <= minLen; k++ {

			if s[i+k] == s[i-k] && i >= 2 {
				tmpMax = 2*k - 1
				if tmpMax > maxLen {
					fmt.Println("tmpmax:", tmpMax, "i:", i, "k:", k, "max:", maxLen)
					start = i - k
					end = i + k
					maxLen = tmpMax
					fmt.Println("22222end:", end, "start:", start)
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
	s := longestPalindrome("aabaaaas")
	fmt.Printf("type: %T,[%v]", s, s)
}

//aaa
