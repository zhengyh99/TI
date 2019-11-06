package model

import (
	"TI/stu/http/webSite/src/utils"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func IndexFunc(rw http.ResponseWriter, req *http.Request) {
	funcMap := template.FuncMap{"timeFormat": utils.TimeFormat}
	tp := template.New("index.html").Funcs(funcMap)
	tp, err := tp.ParseFiles("TI/stu/http/webSite/src/view/index.html",
		"TI/stu/http/webSite/src/view/header.html", "TI/stu/http/webSite/src/view/footer.html")
	if err != nil {
		fmt.Fprintln(rw, "template parse file error:", err)
	}
	info := make(map[string]interface{})
	user := NewUser("Haha", 12, "password")
	user.Hobby = []string{"足球", "跆拳道", "K-POP"}
	info["user"] = user
	info["balance"] = 20
	info["ctime"] = time.Now()
	info["exTime"] = time.Date(2020, 10, 1, 12, 0, 0, 0, time.Local)

	err = tp.ExecuteTemplate(rw, "index", info)
	if err != nil {
		fmt.Fprintln(rw, "template excute error:", err)
	}
}
