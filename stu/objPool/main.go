package main

import (
	"TI/stu/objPool/oPool"
	"fmt"
)

func main() {
	pool := oPool.GetPool(10)
	for i := 0; i < 10; i++ {
		if o, err := pool.GetObject(); err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Printf("%T\n", o)
		}
	}
}
