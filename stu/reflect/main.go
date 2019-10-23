package main

import (
	"fmt"
	"reflect"
)

/*
通过反射对结构体中的数据进行修改
*/
type Student struct {
	Name string
	Age  int
}

func main() {
	stu := &Student{
		Name: "haha",
		Age:  12,
	}
	fmt.Println("stu", stu.Name)
	//ValueOf返回一个初始化为i接口保管的具体值的Value，ValueOf(nil)返回Value零值。
	pStu := reflect.ValueOf(stu)
	//判断是否为反射指针
	if pStu.Kind() == reflect.Ptr {
		eStu := pStu.Elem()
		if eStu.CanSet() { //判断是否可进行修改
			//返回该类型名为name的字段（的Value封装）（会查找匿名字段及其子字段），
			//如果v的Kind不是Struct会panic；如果未找到会返回Value零值。
			sName := eStu.FieldByName("Name")
			//为字段赋值
			sName.SetString("xixi")
			sAge := eStu.FieldByName("Age")
			fmt.Println("sAge's type:", sAge.Kind())
			sAge.SetInt(66)
		}
	}
	//修改后的结构体内容
	fmt.Println("After change:", stu)

}
