package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexFunc(rw http.ResponseWriter, req *http.Request) {

	tp, err := template.ParseFiles("TI/stu/http/webSite/view/index.html")
	if err != nil {
		fmt.Fprintln(rw, "template parse file error:", err)
	}
	err = tp.Execute(rw, nil)
	if err != nil {
		fmt.Fprintln(rw, "template excute error:", err)
	}
}

func main() {
	server := http.Server{Addr: ":8866"}
	http.HandleFunc("/", indexFunc)
	server.ListenAndServe()
}
