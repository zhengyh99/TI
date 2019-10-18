package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//在网络network上连接地址address，并返回一个Conn接口
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("net dial error:", err)
	}
	defer conn.Close()
	//开启子go程，读取客户端标准输入，回写至服务器
	go func() {
		sin := make([]byte, 4096)
		for {
			//读取标准输入
			n, err := os.Stdin.Read(sin)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			//回写至服务器端
			conn.Write(sin[:n])
		}
	}()
	//主go程读取服务器发来的数据
	for {
		sread := make([]byte, 4096)

		n, err := conn.Read(sread)
		//当n为0时表示服务器端已经断开链接
		if n == 0 {
			fmt.Println("服务器断开链接。。。。。")
			return
		}
		if err != nil {
			fmt.Println("n= ", n)
			fmt.Println("conn read error :", err)
		}
		//打印服务器发来的数据
		fmt.Println("服务器发来的数据 ：", string(sread[:n]))
	}

}
