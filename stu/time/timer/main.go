package main

import (
	"fmt"
	"time"
)

func runIt(t time.Timer, e chan<- bool) {
	for {
		//规定时间后，chan 返回系统时间
		now, ok := <-t.C
		//当ok返回false 表示 chan无返回
		if !ok {
			//退出函数
			return
		}
		fmt.Println("当前时间：", now)
		// 返回当时时间后，向e chan中写出数据，表示发出退出程序请示
		e <- true

	}

}
func main() {
	//创建一个timer,定时3秒钟
	t := time.NewTimer(time.Second * 3)
	//定义退出chanel
	var eChan = make(chan bool)
	//显示当前时间
	fmt.Println("当前时间：", time.Now())
	//启动go程 ，定时退出程序
	go runIt(*t, eChan)

	for {
		//检测eChan是否写入数据
		r, ok := <-eChan
		//当ok为true时，表示eChan中已被写入数据退出main（）
		if ok {
			fmt.Println("退出请示 r = ", r, "程序退出。")
			return
		}
	}

}
