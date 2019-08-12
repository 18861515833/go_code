//多态测试 利用接口实现
package main

import "fmt"

type animal interface {
	getName() string
	eat()
}
type cat struct {
	name string
}
type dog struct {
	name string
}

func (c *cat) eat() {
	fmt.Println(c.getName(), "吃鱼")
}

func (c *cat) getName() string {
	return c.name
}

func (c *dog) eat() {
	fmt.Println(c.getName(), "吃骨头")
}

func (c *dog) getName() string {
	return c.name
}

func action(a animal) {
	a.eat()
}

func main() {
	action(&cat{"猫"})
	action(&dog{"狗"})
}
