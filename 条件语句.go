package main

import "fmt"

func main() {
	//if 后面跟两条语句  第一条赋值正常执行 第二条作为条件判断
	if a := 1; a > 0 {
		fmt.Println(a)
	}
}
