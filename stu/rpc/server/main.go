package main

import (
	"TI/stu/rpc/message"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type OrderService struct {
}

func (os *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	//订单信息存储Map
	orderMap := map[string]message.OrderInfo{
		"20191028001": message.OrderInfo{OrderId: "20191028001", OrderName: "衣服", OrderStatus: "已付款"},
		"20191029001": message.OrderInfo{OrderId: "20191029001", OrderName: "零食", OrderStatus: "已付款"},
		"20191028002": message.OrderInfo{OrderId: "20191028002", OrderName: "食品", OrderStatus: "未付款"},
	}
	//获取当前时间
	current := time.Now().Unix()
	//比较C端和S端的时间
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
		return nil
	}
	//获取订单信息
	result, ok := orderMap[request.OrderId]
	if ok {
		*response = result
	} else {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "无此订单"}
	}
	return nil
}

func main() {
	//初始化
	orderService := new(OrderService)
	//rpc注册
	err := rpc.Register(orderService)
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
