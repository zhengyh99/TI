package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//拔号
	client, err := rpc.DialHTTP("tcp", ":8098")
	if err != nil {
		fmt.Println("rpc dialHttp error:", err)
	}
	//请求值与返回值
	var req, resp float32
	req = 1
	//方法一：同步调用
	// err2 := client.Call("MathUtil.CircleArea", req, &resp)
	// if err2 != nil {
	// 	fmt.Println("client call error:", err)
	// }
	// fmt.Println("result:", resp)

	//方法二：异步调用

	synsCall := client.Go("MathUtil.CircleArea", req, &resp, nil)
	replyDone := <-synsCall.Done
	fmt.Println("respDone：", replyDone)
	fmt.Println("result", resp)

}
