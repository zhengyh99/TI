package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
)

//接收处理客户端发来的数据
func acc(conn net.Conn) {

	defer conn.Close()
	//定义接收数据的字节切片
	b := make([]byte, 4096)
	for {

		//读取客户端发来的数据
		n, err := conn.Read(b)
		//当返回值n为0时，表示客户端已关闭连接
		if n == 0 {
			fmt.Printf("客户端%v关闭连接\n", conn.RemoteAddr())
			//退出当前协程
			runtime.Goexit()
		}
		if err != nil {
			fmt.Println("conn read error:", err)
		}
		//当客户端发来的数据为“exit\n”(除windows外的终端请求）、
		//“exit\r\n"（windows cmd平台）、"exit"（客户端代码请求）时表示客户端请求断开连接
		if string(b[:n]) == "exit\n" || string(b[:n]) == "exit\r\n" || string(b[:n]) == "exit" {
			fmt.Printf("客户端%v发出退出请示\n", conn.RemoteAddr())
			runtime.Goexit()
		}
		//字符串转化成大写
		uper := strings.ToUpper(string(b[:n]))
		//将处理结果回写至客户端
		conn.Write([]byte(uper))
	}

}

func main() {
	//服务器起动监听
	listener, err := net.Listen("tcp", "127.0.0.1:8088")

	if err != nil {
		fmt.Println("net listen error :", err)
	}
	fmt.Println("服务器打开监听。。。。。。")
	defer listener.Close()

	for {
		//等待下一个呼叫，并返回一个该呼叫的Conn接口
		conn, err := listener.Accept()
		//返回客户端ip和端口
		addr := conn.RemoteAddr()
		if err != nil {
			fmt.Println("listener Accept error:", err)
		}
		fmt.Println("服务器开始接收", addr, "的数据 。。。。。")
		//开启go程 ，接收、处理客户端发来的数据
		go acc(conn)
	}

}
