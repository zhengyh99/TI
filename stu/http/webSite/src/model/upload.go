package model

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Upload(rw http.ResponseWriter, req *http.Request) {

	fileName := req.FormValue("pname")
	file, fileHandle, err := req.FormFile("pimg")
	if err != nil {
		fmt.Println("request form file error:", err)
	}
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ioutil readall error:", err)
	}

	ioutil.WriteFile(fileHandle.Filename, fileContent, 0777)
	tp, err := template.ParseFiles("TI/stu/http/webSite/src/view/success.html")
	if err != nil {
		fmt.Println("template ParseFiles error:", err)
	}
	tp.Execute(rw, fileName)

}
