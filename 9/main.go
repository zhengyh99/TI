package main

import "fmt"

func convert(s string, numRows int) string {
	b := []byte(s)
	tmp := make([][]byte, len(b))
	var n, rowN, col int
	for i := 0; i < len(b); i++ {
		if col%numRows == 0 {
			fmt.Println("bi=", string(b[i]), "i=", i, "col=", col, "n=", n)
			tmp[col] = append(tmp[col], b[i])
			n++
			if n%numRows == 0 {
				col++
			}

		} else {
			rowN = col % numRows
			fmt.Println("rowN", rowN)
			for j := 0; j < numRows; j++ {
				if j == rowN {

					tmp[col] = append(tmp[col], b[i])

				} else {

					tmp[col] = append(tmp[col], byte(32))
				}

				n++

			}
			col++
		}

	}
	fmt.Println(tmp)
	return ""
}

func main() {
	s := "PAYPALIKFG"
	r := convert(s, 3)
	fmt.Println(r)

}
