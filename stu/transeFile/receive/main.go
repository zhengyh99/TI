package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func fileRecv(conn net.Conn, fileName string) {
	//创建本地文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.create error : ", err)
		return
	}
	defer f.Close()
	//解析本地文件的绝对路径
	fpath, _ := filepath.Abs(fileName)
	fmt.Println("本地文件存放的绝对路径:", fpath)
	for {
		//接收发送端数据
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if n == 0 {

			fmt.Println("文件接收完")
			return
		}
		if err != nil {
			fmt.Println("conn.read error : ", err)
			return
		}
		//向文件中写入数据
		f.Write(buf[:n])
	}

}

func main() {
	//打开接收端监听
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	fmt.Println("打开监听。。。。")
	if err != nil {
		fmt.Println("net.listen error : ", err)
	}
	defer listener.Close()
	//准备接收数据
	conn, err := listener.Accept()
	fmt.Println("等待发送端数据。。。。。。")
	if err != nil {
		fmt.Println("listener.accept error : ", err)
	}
	defer conn.Close()

	//接收发送端发来的文件名
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.read error ：", err)
		return
	}
	fileName := string(buf[:n])
	//文件名接收成功
	conn.Write([]byte("ok"))
	//开始接收文件
	fileRecv(conn, fileName)
}
