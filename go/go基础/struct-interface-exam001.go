//通道、结构体、接口、slice与排序之间的应用为目标
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Iface interface {
	Getval() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Getval() string {
	return p.Name

}

func (p *Person) Setval(name string) {
	p.Name = name
}

//Student继承了Person相关的属性和方法
type Student struct {
	Person
	Score int
}

//定义一个接口用于Stuent这个结构体对Person结构体的功能扩展,实现自己的Fly()方法
type Flyable interface {
	Fly()
}

func (p *Student) Fly() {
	fmt.Printf("通过接口调用,我%v实现了想要飞的功能...\n", p.Name)
}

//声明一个叫做Man的切片，其类型为Person结构体类型,用于后面实现对结体的排序
type Man []Student

//实现sort中的Interface接口Len,Less,Swap方法
func (a Man) Len() int {
	return len(a)
}

func (a Man) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a Man) Swap(i, j int) {
	/*
		temp := a[i]
		a[i] = a[j]
		a[j] = temp
	*/
	a[i], a[j] = a[j], a[i]
}

func main() {
	//map类型的chan,通过make初始化
	ch := make(chan map[string]string, 10)

	//结构构类型的chan，通过make初始化
	ch2 := make(chan Student, 20)

	//初始化一个map,进行赋值，再将数据写入ch通道中
	mp1 := make(map[string]string, 10)
	mp1["city1"] = "hangzhou"
	ch <- mp1

	//实例化一个结构体，并通过Setval方法修改name的值
	var p1 Student
	p1 = Student{
		Person{"tim", 18},
		804,
	}
	fmt.Println("使用结构体方法取数据", p1.Getval())
	fmt.Println("子继续中读取到自己的数据", p1.Score)
	p1.Fly()

	//使用结构体方法设置数据
	p1.Setval("jack")

	//结构体实例p1实现了Iface接接口，所以接口ph可以直接调用Getval()方法
	var ph Iface = &p1
	fmt.Println("使用接口方法取数据:", ph.Getval())

	ch2 <- p1
	var prt = fmt.Println
	fmt.Println("读出通道ch中的数据", <-ch)
	fmt.Println("读出通道ch2中的数据", <-ch2)

	fmt.Println("slice切片是引用类型,sort.Ints()对整型slice进行排序")
	var intslice = []int{10, 20, -3, 78, -9}
	prt("slice排序前", intslice)
	sort.Ints(intslice)
	prt("slice排序后", intslice)

	var mans Man
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 10; i++ {
		var m Student = Student{
			Person{Name: fmt.Sprintf("工人%d", r.Intn(100)), Age: r.Intn(100)}, r.Intn(10)}
		mans = append(mans, m)

	}
	for _, v := range mans {
		prt(v)
	}

	prt("排序后的情况")
	sort.Sort(mans)
	for _, v := range mans {
		prt(v)
	}

}
