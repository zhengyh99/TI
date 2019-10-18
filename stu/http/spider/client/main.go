package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.get error：", err)
	}
	fmt.Println("resp.proto　：", resp.Proto)
	fmt.Println("resp.Request:", resp.Request)
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	var str string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err == io.EOF {
			fmt.Println("resp.body.read error : ", err)
			break
		}
		str += string(buf[:n])
	}
	fmt.Println(str)
}
