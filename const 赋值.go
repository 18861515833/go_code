package main

import "fmt"

type level int8

//const 初始化可以沿用上面的赋值表达式
//iota没出现一个const变量就加1
const (
	level1 level = iota - 1
	level2
	level3
	level4
)

func main() {
	fmt.Println(level1)
	fmt.Println(level2)
	fmt.Println(level3)
	fmt.Println(level4)
}
