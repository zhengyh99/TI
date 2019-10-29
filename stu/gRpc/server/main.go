package main

import (
	"TI/stu/gRpc/message"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

type OrderServer struct {
}

func (os *OrderServer) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (orderInfo *message.OrderInfo, err error) {
	//订单信息存储Map
	orderMap := map[string]message.OrderInfo{
		"20191028001": message.OrderInfo{OrderId: "20191028001", OrderName: "衣服", OrderStatus: "已付款"},
		"20191029001": message.OrderInfo{OrderId: "20191029001", OrderName: "零食", OrderStatus: "已付款"},
		"20191028002": message.OrderInfo{OrderId: "20191028002", OrderName: "食品", OrderStatus: "未付款"},
	}
	//获取当前时间
	current := time.Now().Unix()
	//比较C端和S端的时间
	if request.TimeStamp < current {
		fmt.Println("current:", current)
		fmt.Println("rtime:", request.TimeStamp)
		orderInfo = &message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
		return
	}
	//获取订单信息
	result, ok := orderMap[request.OrderId]
	if ok {
		fmt.Println("result:", result)
		orderInfo = &result
	} else {
		orderInfo = &message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "无此订单"}
	}
	return
}
func main() {

	//创建服务
	server := grpc.NewServer()
	//注册服务和接口实现
	message.RegisterOrderServiceServer(server, new(OrderServer))
	//建立监听
	listen, err := net.Listen("tcp", ":8019")
	if err != nil {
		fmt.Println("net listen error:", err)

	}
	//绑定监听
	server.Serve(listen)
}
