package model

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexFunc(rw http.ResponseWriter, req *http.Request) {

	tp, err := template.ParseFiles("TI/stu/http/webSite/src/view/index.html")
	if err != nil {
		fmt.Fprintln(rw, "template parse file error:", err)
	}
	err = tp.Execute(rw, nil)
	if err != nil {
		fmt.Fprintln(rw, "template excute error:", err)
	}
}
