package main

import (
	"TI/stu/gRpc/message"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//grpc拔号
	conn, err := grpc.Dial("localhost:8019", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc dial error:", err)
	}
	//延迟关闭资源
	defer conn.Close()
	//建立客户端
	client := message.NewOrderServiceClient(conn)
	//获取数据
	orderInfo, err := client.GetOrderInfo(context.Background(), &message.OrderRequest{OrderId: "20191029001", TimeStamp: (time.Now().Unix() - 5000000)})
	if err != nil {
		fmt.Println("client GetOrderInof error:", err)
	}
	fmt.Println("=======orderInfo===============")
	fmt.Printf("orderid:%s\n", orderInfo.GetOrderId())
	fmt.Printf("ordername:%s\n", orderInfo.GetOrderName())
	fmt.Printf("orderstatus:%s", orderInfo.GetOrderStatus())

}
