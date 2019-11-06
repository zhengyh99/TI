package model

type User struct {
	Username string
	Age      int
	Password string
	Hobby    []string
}

func NewUser(uname string, age int, pwd string) *User {
	return &User{Username: uname, Age: age, Password: pwd}
}
