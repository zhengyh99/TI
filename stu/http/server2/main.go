package main

import (
	"fmt"
	"net/http"
)

func index(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html;charset=utf-8")
	rw.Write([]byte("Hello World 【<font color='blue'><b>Victor</b></font>】 !!!"))
	header := req.Header
	fmt.Fprintln(rw, header)
	fmt.Fprintln(rw, "<br>")
	fmt.Fprintln(rw, len(header["Accept"]))
	fmt.Fprintln(rw, "<br>")
	for _, v := range header["Accept"] {
		fmt.Fprintln(rw, v+"<br>")
	}
	name := req.FormValue("name")
	age := req.FormValue("age")
	fmt.Fprintf(rw, "name = %s <br><br> age = %s<hr>", name, age)

}
func main() {
	server := http.Server{Addr: ":8810"}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
