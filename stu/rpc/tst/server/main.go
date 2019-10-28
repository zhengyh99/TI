package main

import (
	"TI/stu/rpc/utils"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	//初始化
	mathUtil := new(utils.MathUtil)
	//rpc注册
	err := rpc.Register(mathUtil)
	if err != nil {
		fmt.Println("rpc register error:", err)
	}
	//rpc绑定http
	rpc.HandleHTTP()

	//打开http 监听
	listen, err := net.Listen("tcp", ":8098")
	if err != nil {
		fmt.Println("net listen error:", err)
	}
	http.Serve(listen, nil)
}
