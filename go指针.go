package main

import "fmt"

type student struct {
	name string
	age int
}

//注意这两种方式的区别
/*
func (t *student)addAge(){
	fmt.Println(&t.age)
	t.age++
}
*/
func (t student)addAge(){
	fmt.Println(&t.age)
	t.age++
}

var data0=map[int]string {1:"123",2:"456",3:"789"}
type mymap map[int]string
//引用类型方法实现测试

func (t mymap)reload(){
	t=data0
}

/*
func (t *mymap)reload(){
	*t=data0
}
*/




func main() {
	/*
	s := &student{"www",21}
	fmt.Println(&s.age)
	s.addAge()
	*/

	var data =mymap{1:"1",2:"2",3:"3",4:"4",5:"5"}
	data.reload()
	fmt.Println(data)
}
