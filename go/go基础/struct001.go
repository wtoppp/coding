//1.结构体可以使用嵌套匿名结构体，实现类型继承的功能
//2.匿名结构体字段访问可以简化
package main

import (
	"fmt"
)

type A struct {
	Name string
	Age  int
}

func (a *A) Hello() {
	fmt.Println("A Hello()：", a.Name)
}

func (a *A) Welcome() {
	fmt.Println("A Welcome()：", a.Age)
}

type B struct {
	A
}

func main() {
	var b B
	b.A.Name = "tim"
	b.A.Hello()

	//简化的写法
	b.Age = 19
	b.Welcome()
}
