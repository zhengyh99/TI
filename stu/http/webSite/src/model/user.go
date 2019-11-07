package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"uname"`
	Age      int    `json:"age"`
	Password string
	Hobby    []string
}

func NewUser(uname string, age int, pwd string) *User {
	return &User{Username: uname, Age: age, Password: pwd}
}

func UserList(rw http.ResponseWriter, req *http.Request) {
	userList := make([]User, 0)
	userList = append(userList, User{Username: "xixi", Age: 12})
	userList = append(userList, User{Username: "haha", Age: 24})
	userList = append(userList, User{Username: "mimi", Age: 32})
	ul, err := json.Marshal(userList)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	rw.Header().Set("Content-Type", "application/json;charset=uft-8")
	rw.Write(ul)

}
