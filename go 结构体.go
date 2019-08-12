package main

import (
	"fmt"
)

type getId func()int


//注意结构体也有访问权限控制，结构体中的小写字母开头的成员都不能在包外部访问
type student struct {
	name string
	id getId
}

//结构体中的结构体
//注意class和class1的区别
type class struct {
	student
	className string
}

type class1 struct {
	s student
	className string
}


var c =0
//c:=0
func getCount()int {
	id:=c
	c++
	return id
}

func main() {
	s:=student{"www",getCount}
	fmt.Println(s.name,s.id(),s.id())

	myClass:=class{s,"高三11班"}
	fmt.Println(myClass.name)

	myClass1:=class1{s,"高三11班"}
	fmt.Println(myClass1.s.name)
}
