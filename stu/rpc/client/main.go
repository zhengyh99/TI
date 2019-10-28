package main

import (
	"TI/stu/rpc/message"
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	//拔号
	client, err := rpc.DialHTTP("tcp", ":8098")
	if err != nil {
		fmt.Println("rpc dialHttp error:", err)
	}
	timeNow := time.Now().Unix()
	//请求值与返回值
	var resp message.OrderInfo
	req := message.OrderRequest{OrderId: "20191028001", TimeStamp: timeNow}
	//方法一：同步调用
	err2 := client.Call("OrderService.GetOrderInfo", req, &resp)
	if err2 != nil {
		fmt.Println("client call error:", err)
	}
	fmt.Println("result:", resp)

}
