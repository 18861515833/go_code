package main

import "fmt"

func main() {
	//	Go 语言中 range 关键字用于 for 循环中迭代
	// 数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	// 在数组和切片中它返回元素的索引和索引对应的值，
	// 在集合中返回 key-value 对的 key 值。

	//rangeArray()
	//rangeSlice()
	//rangeChannel()
	rangeMap()

}

func rangeArray() {
	//坑点：采用month2拿到的值v是一个副本，无法通过他修改原来数组的值
	var data = [5]int{1, 2, 3, 4, 5}
	fmt.Println("method1:")
	for i := range data {
		fmt.Println("index:", i, ",data:", data[i], ",data addr:", &data[i])
	}
	fmt.Println("method2:")
	for i, v := range data {
		fmt.Println("index:", i, ",data:", v, ",data addr:", &v)
	}
}

func rangeSlice() {
	var data = []int{1, 2, 3, 4, 5}
	fmt.Println("method1:")
	for i := range data {
		fmt.Println("index:", i, ",data:", data[i], ",data addr:", &data[i])
	}
	fmt.Println("method2:")
	for i, v := range data {
		fmt.Println("index:", i, ",data:", v, ",data addr:", &v)
	}
}

func rangeChannel() {
	var data chan int
	data = make(chan int, 10)
	for i := 0; i < 10; i++ {
		data <- i
	}
	//这里如果不关闭通道，遍历的时候可能出错
	close(data)

	for v := range data {
		fmt.Println("data：", v)
	}
}

func rangeMap() {
	var data = map[int]string{1: "123", 2: "456", 3: "789"}
	//data=make(map[int]string)
	fmt.Println("method1:")
	for k := range data {
		fmt.Println("key:", k, ",value:", data[k])
	}
	fmt.Println("method2:")
	for k, v := range data {
		fmt.Println("key:", k, ",value:", v)
	}
}
