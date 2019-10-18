package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//发送文件
func sendFile(conn net.Conn, filePath string) {
	//打开本地文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open error : ", err)
	}
	defer f.Close()

	for {
		buf := make([]byte, 4096)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil {
			//文件发送完成
			if err == io.EOF {
				fmt.Println("本地文件读取完成")

			} else {
				fmt.Println("f.read other error : ", err)
			}
			return
		}
		//向服务端发送文件数据
		_, err2 := conn.Write(buf[:n])
		if err2 != nil {
			fmt.Println("conn.write error : ", err)
		}

	}
}

func main() {
	//接收命令行参数，第一个是程序名
	cmdList := os.Args
	if len(cmdList) != 2 {
		fmt.Println("格式：go run ***.go 传送文件的绝对路径")
	}
	//接收第二个参数，文件绝对路径
	filePath := cmdList[1]
	//解决文件状态，返回文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat error : ", err)
	}
	//开始连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	fmt.Println("开始连接服务器。。。。")
	if err != nil {
		fmt.Println("net.Dial error : ", err)
	}
	defer conn.Close()
	//向服务器传送文件名
	conn.Write([]byte(fileInfo.Name()))
	fmt.Println("向服务器发送文件名。。。。")

	buf := make([]byte, 16)
	//接收文件名传输状态 “ok”表服务接收成功
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.read error : ", err)
	}
	if "ok" == string(buf[:n]) {
		fmt.Println("服务器成功接收到文件名称，开始向服务器端发送数据 。。。。")
		//开始发送文件数据
		sendFile(conn, filePath)

	} else {
		fmt.Println("文件传送失败。。。。")
	}

}
