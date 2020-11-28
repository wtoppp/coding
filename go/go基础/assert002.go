//1.类型断言,因为接口不知道具体类型，如果要转换成具体类型，则需要使用类型断言。
//动态传入多个值参，判断其类型。
package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func Checktype(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v参数是%T类型,值是%v\n", index, x, x)
		case float32:
			fmt.Printf("第%v参数是%T类型,值是%v\n", index, x, x)
		case float64:
			fmt.Printf("第%v参数是%T类型,值是%v\n", index, x, x)
		case string:
			fmt.Printf("第%v参数是%T类型,值是%v\n", index, x, x)

		case Student:
			fmt.Printf("第%v参数是%T类型,值是%v\n", index, x, x)
		case *Student:
			fmt.Printf("第%v参数是%T类型,值是%v\n", index, x, x)
		default:
			fmt.Printf("第%v参数是类型不确定,值是%v\n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.3
	var n2 float64 = 2.45
	var name = "tim"
	var t = true

	var st1 Student = Student{"jack", 18}
	var st2 *Student = &Student{Name: "tim", Age: 16}

	Checktype(n1, n2, name, t, st1, st2)
}
