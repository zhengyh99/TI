package main

import (
	"fmt"
	"net"
	"strings"
)

func acc(uConn net.UDPConn) {
	buf := make([]byte, 4096)
	for {
		n, cAddr, err := uConn.ReadFromUDP(buf)
		fmt.Printf("开始接收客户端%v发来的数据。。。。\n", cAddr.String())
		if 0 == n {
			fmt.Println("客户端已关闭。。。。。")
			return
		}
		if err != nil {
			fmt.Println("uconn.readfromupd error : ", err)
			return
		}
		uper := strings.ToUpper(string(buf[:n]))
		uConn.WriteToUDP([]byte(uper), cAddr)
	}

}

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.resoveupdaddr error : ", err)
	}
	uConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("net.listenupd error : ", err)
	}
	fmt.Println("服务器开始监听数据。。。。。。。")
	acc(*uConn)
}
