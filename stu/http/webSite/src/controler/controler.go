package controler

import (
	"TI/stu/http/webSite/src/model"
	"net/http"
)

func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("TI/stu/http/webSite/static"))))
	http.HandleFunc("/", model.IndexFunc)
	http.HandleFunc("/upload", model.Upload)
	http.HandleFunc("/download", model.Download)
	http.HandleFunc("/userlist",model.UserList)
}
