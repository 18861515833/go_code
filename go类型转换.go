package main

import "fmt"

func main() {
	var a int = -1
	var b uint = 1
	if uint(a) > b {
		fmt.Println("a>b")
	}
}
