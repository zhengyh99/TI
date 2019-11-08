package main

import (
	"TI/stu/http/webSite/src/controler"
	"net/http"
)

func main() {
	// server := http.Server{Addr: ":8866"}
	// server.ListenAndServe()

	http.ListenAndServe(":8866", controler.GetRouter())
}
