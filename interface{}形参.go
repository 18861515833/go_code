package main

import "fmt"

func main() {
	a:=1
	test(&a)
}

func test(p interface{}) {
	data:=p.(*int)
	fmt.Println(*data)
}
