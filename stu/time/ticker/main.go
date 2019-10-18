package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Microsecond * 300)
	fmt.Println("开始：", time.Now())
	for i := 0; i < 10; i++ {
		now, ok := <-ticker.C
		if !ok {
			fmt.Println("获取时间失败 。。。。")
		}
		fmt.Println("时间：", now)
	}
}
