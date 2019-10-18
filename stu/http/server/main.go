package main

import (
	"fmt"
	"net/http"
	"os"
)

func openLocalFile(fileName string, w http.ResponseWriter) {

	f, err := os.Open("D:\\360安全浏览器下载\\" + fileName)
	if err != nil {
		fmt.Println("op.open error: ", err)
	}
	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf)
		w.Write(buf[:n])
	}
}
func victorHtml(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World !!! "))
	openLocalFile(r.URL.String(), w)
	fmt.Println("r.Header", r.Header)
	fmt.Println("r.url", r.URL)
	fmt.Println("r.Host", r.Host)
	fmt.Println("r.Method", r.Method)
	fmt.Println("r.Proto", r.Proto)
	fmt.Println("r.method", r.Method)
	fmt.Println("r.body", r.Body)
}
func main() {
	http.HandleFunc("/", victorHtml)
	http.ListenAndServe("127.0.0.1:8089", nil)
}
