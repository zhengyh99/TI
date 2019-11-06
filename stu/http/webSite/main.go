package main

import (
	_ "TI/stu/http/webSite/src/controler"
	"net/http"
)

func main() {
	server := http.Server{Addr: ":8866"}
	server.ListenAndServe()
}
