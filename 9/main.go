package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
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
	s := "A"
	r := convert(s, 1)
	fmt.Println(r)

}
