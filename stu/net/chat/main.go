package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

//客户端用户结构体定义
type Client struct {
	Name string
	Addr string
	//客户端消息通道
	C chan string
}

//在线用户列表（map)
var onlineMap map[string]Client

//消息通道
var message = make(chan string)

//消息管理
func Manage() {
	fmt.Println("Manage 初始化开始。。。。。")
	//初始化在线用户列表
	onlineMap = make(map[string]Client)
	//监听消息
	for {
		fmt.Println("等待系统消息。。。。。")
		msg := <-message
		//接收到消息，将消息内容塞入客户端消息通道
		for _, clnt := range onlineMap {
			fmt.Println("用户名：", clnt.Name)
			clnt.C <- msg
		}
		fmt.Println(">>>>得到消息：", msg)
	}
}

//向客户端消息通道写入消息
func WritetoClient(conn net.Conn, clnt Client) {
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

//格式化消息内容
func MakeMSG(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + " : " + msg
	return

}

//客户端逻辑处理
func HandlerConnect(conn net.Conn) {
	//监听客户端是否退出
	quit := make(chan bool)
	//监听客户端是否处于活跃状态
	active := make(chan bool)
	defer conn.Close()
	//获取客户地址信息
	cAddr := conn.RemoteAddr().String()
	fmt.Println(cAddr, "连接服务器.....")
	//初始化客户端实例
	clnt := Client{cAddr, cAddr, make(chan string)}
	//将当客户端写入在线用户列表
	onlineMap[cAddr] = clnt
	//定义go程向所有在线客户端写入消息
	go WritetoClient(conn, clnt)
	//用户登陆状态写入消息通道
	message <- cAddr + "=====>is onlined !!!"

	//定义go程 接收客户端数据
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			//客户端关闭
			if 0 == n {

				quit <- true
				runtime.Goexit()
			}
			if err != nil {
				fmt.Println("conn.read error:", err)
				quit <- true
				runtime.Goexit()
			}
			//客户端请示退出
			if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" || string(buf[:n]) == "exit" {
				fmt.Printf("客户端%v发出退出请示\n", conn.RemoteAddr())
				quit <- true
				runtime.Goexit()
			}
			//将客户端发来的数据写入消息通道
			message <- MakeMSG(clnt, string(buf[:n]))
			//设计用户状态为：活跃
			active <- true
		}
	}()
	//监听客户状态
	for {
		select {
		//客户端已退出
		case <-quit:
			message <- MakeMSG(clnt, " is logout !!!")
			fmt.Println("用户", clnt.Addr, "推出，准备杀死当前go程。。。")
			delete(onlineMap, clnt.Addr)
			close(clnt.C)
			runtime.Goexit()
		//客户端状态活跃
		case <-active:
		//客户端超时监听
		case <-time.After(time.Minute * 2):
			message <- MakeMSG(clnt, " is timeout and logout !!!")
			fmt.Println("用户", clnt.Addr, "超时退出 。。。")
			delete(onlineMap, clnt.Addr)
			close(clnt.C)
			runtime.Goexit()

		}
	}

}

func main() {
	//打开服务器监听
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("net.listen error : ", err)
	}
	fmt.Println("服务器打开监听。。。。。。")
	defer listener.Close()
	//消息管理go程
	go Manage()
	for {
		//接收客户端请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.accept error : ", err)
		}
		//客户端逻辑处理go程
		go HandlerConnect(conn)
	}
}
