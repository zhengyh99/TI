package controler

import (
	"model"
	"net/http"
)

func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("TI/stu/http/webSite/static"))))
	http.HandleFunc("/", model.IndexFunc)
}
