package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

//定义用户
type User struct {
	UserName  string //本机用户
	OtherUser string //对话用户
	Msg       string //聊天消息
	ServerMsg string //系统消息
}

var (
	//当前用户
	user = new(User)
	//Go程组
	wg sync.WaitGroup
)

func main() {
	//添加Go程组
	wg.Add(1)
	//输入当前和对方用户名
	fmt.Println("输入你的用户名：")
	fmt.Scanln(&user.UserName)
	fmt.Println("输入对方用户名：")
	fmt.Scanln(&user.OtherUser)
	//定义tcp地址
	add, err := net.ResolveTCPAddr("tcp", ":8089")
	if err != nil {
		fmt.Println("net tcp addr error:", err)
	}
	//客户端拨号
	conn, err := net.DialTCP("tcp", nil, add)
	if err != nil {
		fmt.Println("net dial error:", err)
	}
	//发送数据Go程
	go func() {
		fmt.Println("请输入你要输入的信息：")
		for {
			fmt.Scanln(&user.Msg)
			//判断用户是否要退出
			if strings.ToLower(user.Msg) == "exit" {
				fmt.Println("退出聊天系统")
				conn.Close()
				//完成一个Go程任务
				wg.Done()
				os.Exit(0)
			} //发送数据
			conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.UserName, user.OtherUser, user.Msg, user.ServerMsg)))
		}
	}()
	//接收数据Go程
	go func() {
		for {
			//接收数据
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn read error:", err)
				break
			}
			//分析数据
			uInfo := strings.Split(string(buf[:n]), "-")
			if len(uInfo) <= 1 {
				break
			}
			//处理数据
			oUser := new(User)
			oUser.UserName = uInfo[0]
			oUser.OtherUser = uInfo[1]
			oUser.Msg = uInfo[2]
			oUser.ServerMsg = uInfo[3]
			//判断消息类型（用户消息/系统消息）
			if oUser.ServerMsg != "" {
				fmt.Println("\t\t服务器消息：", oUser.ServerMsg)
			} else {
				fmt.Println("\t\t", oUser.UserName, ":", oUser.Msg)
			}
		}
	}()
	//Go组阻塞
	wg.Wait()
}
