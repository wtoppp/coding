//1.类型断言,因为接口不知道具体类型，如果要转换成具体类型，则需要使用类型断言。

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var a interface{}

	var p Person = Person{Name: "jack", Age: 17}
	a = p
	fmt.Println(p)
	fmt.Println(a)

	var b Person
	//b = a 不能这样写，而应该通过类型断言来实现
	//b = a.(Person)类型断言,即判断接口变量a是否是指向Person类型的变量（如果不是则会报错），如果是就将a转换成Person类型并赋给b变量。
	b = a.(Person)
	fmt.Println(b)

	//例2类型断言时带上检测机制。
	var x float64 = 3.14
	a = x
	//var y float64
	//y = a 不能这样写，而应该通过类型断言来实现

	/*
		y, ok := a.(float64)
		if ok {
			fmt.Printf("convert successful:y类型是%T,值%v", y, y)
		} else {
			fmt.Printf("convert Fail:y类型是%T,值%v", y, y)
		}
	*/

	//以下简化写法与上面/*...*/注释的部分功能相同
	if y, ok := a.(float64); ok {
		fmt.Printf("convert successful:y类型是%T,值%v", y, y)
	} else {
		fmt.Printf("convert Fail:y类型是%T,值%v", y, y)
	}

}
