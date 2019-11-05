package main

import (
	_ "controler"
	"net/http"
)

func main() {
	server := http.Server{Addr: ":8866"}
	server.ListenAndServe()
}
