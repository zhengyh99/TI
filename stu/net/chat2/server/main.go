package main

import (
	"fmt"
	"net"
	"strings"
)

//定义用户
type User struct {
	UserName  string //本机用户
	OtherUser string //对话用户
	Msg       string //聊天消息
	ServerMsg string //系统消息
}

var (
	//在线用户列表存储
	userOnline = make(map[string]net.Conn)
	//当前用户
	user = new(User)
)

func main() {
	//创建tcp地址
	addr, err := net.ResolveTCPAddr("tcp", ":8089")
	if err != nil {
		fmt.Println("net resolve TCPADDR error:", err)
	}
	//打开服务器监听
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("net ListenTCP error:", err)
	}
	for {
		//接收客户端数据
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("net Listen Accept error:", err)
		}
		go func() {
			for {
				//读取客户端数据
				buf := make([]byte, 4096)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("conn read error:", err)
				}
				//分析数据 赋值给user
				uInfo := strings.Split(string(buf[:n]), "-")
				if len(uInfo) <= 1 {
					break
				}
				user.UserName = uInfo[0]
				user.OtherUser = uInfo[1]
				user.Msg = uInfo[2]
				user.ServerMsg = uInfo[3]
				//将当前用户加入到在线用户列表
				userOnline[user.UserName] = conn
				//判断对方用户在线状态
				online, ok := userOnline[user.OtherUser]
				if ok {
					//在线，将消息发送给对方用户
					n, err := online.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.UserName, user.OtherUser, user.Msg, user.ServerMsg)))
					//发送失败或无数写入
					if n <= 0 || err != nil {
						fmt.Println("发送给对方数据失败：", err)
						delete(userOnline, user.OtherUser)
						online.Close()
						conn.Close()
						break
					}
				} else {
					//不在线，返回“对方不在线”
					user.ServerMsg = "对方不在线"
					conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.UserName, user.OtherUser, user.Msg, user.ServerMsg)))
				}
			}
		}()

	}

}
