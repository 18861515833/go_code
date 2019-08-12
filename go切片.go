package main

import "fmt"

func main() {
	var data = []int{1, 2, 3, 4, 5}

	//append

	//copy
	fmt.Println("改变副本")
	cdata := make([]int, 5)
	copy(cdata, data)
	cdata[1] = 100
	fmt.Println(data)
	//len
	fmt.Println(len(data))
	//cap
	fmt.Println(cap(data))

	//slice 的增删改查
	//增
	fmt.Println("增加前", data)
	data = append(data, 6)
	fmt.Println("增加后", data)
	//删
	fmt.Println("删除前", data)
	data = append(data[:2], data[3:]...)
	fmt.Println("删除后", data)
	//改
	fmt.Println("改变前", data)
	data[0] = 100
	fmt.Println("改变后", data)
	//查
	fmt.Println("查找5", data)
	find := 5
	for i, v := range data {
		if v == find {
			fmt.Println("index:", i, "value:", v)
		}
	}

	//slice 作为函数参数传递
	//slice 的底层实现是一个地址指向内存空间，两个int形变量，一个length，一个capacity
	//go的函数参数传递都是通过值传递的，所以为了实现引用传递，需要传递指针

	//典型示例，通过函数来进行make
	var arr []int
	fmt.Println("before getMemory", len(arr))
	getMemory(arr)
	fmt.Println("after getMemory", len(arr))
}

func getMemory(arr []int) {
	arr = make([]int, 10) //通过函数改变了len
}
