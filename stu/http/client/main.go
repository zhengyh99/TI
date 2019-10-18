package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8089")
	head := "GET /victor HTTP/1.1\r\nHOST:127.0.0.1:8089\r\n\r\n"
	conn.Write([]byte(head))
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
