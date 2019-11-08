package controler

import (
	"TI/stu/http/webSite/src/model"
	"net/http"

	"github.com/gorilla/mux"
)

// func init() {
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("TI/stu/http/webSite/static"))))
// 	http.HandleFunc("/", model.IndexFunc)
// 	http.HandleFunc("/upload", model.Upload)
// 	http.HandleFunc("/download", model.Download)
// 	http.HandleFunc("/userlist", model.UserList)
// }

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/username/{uname}", model.GetName)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("TI/stu/http/webSite/static"))))
	router.HandleFunc("/", model.IndexFunc)
	router.HandleFunc("/upload", model.Upload)
	router.HandleFunc("/download", model.Download)
	router.HandleFunc("/userlist", model.UserList)
	return router

}
