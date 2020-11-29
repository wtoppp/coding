//通过面向对象的的方式例2，重新实现面向过程的实例1全部功能，实现一个账号资金的入账、出账、显示明细管理。
package main

import (
	"fmt"
)

type Jizhang struct {
	//loop控制循环的退出与否
	loop bool
	//key 接收用户选择输入的数字
	key string

	//当前余额
	balance float64
	//进出账金额
	money float64
	//进账原因说明
	note string
	//有收入时用detail进行拼接
	//detail := "收入：\t进账\t余额\t说明\n"
	detail string
	//判断是否有收入或支出
	flag bool
}

func NewJizhang() *Jizhang {
	return &Jizhang{
		key:     "",
		loop:    true,
		balance: 100,
		money:   0.0,
		note:    "",
		flag:    false,
		detail:  "收入：\t进账\t余额\t说明\n",
	}

}

//给结构体实现方法
func (this *Jizhang) inpay() {
	fmt.Println("本次进账")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次入账原因是：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收入\t+%v\t%v\t%v\n", this.money, this.balance, this.note)
	this.flag = true
}

func (this *Jizhang) outpay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//balance_enough用于标记是否有足够的余额支出
	balance_enough := true

	if this.money > this.balance {
		fmt.Println("余额不足")
		balance_enough = false
	}
	if balance_enough {
		this.balance -= this.money
		fmt.Println("本次支出说明：")
		fmt.Scanln(&this.note)
		this.detail += fmt.Sprintf("\n支出\t-%v\t%v\t%v\n", this.money, this.balance, this.note)
		this.flag = true
	}
}

func (this *Jizhang) exit() {
	fmt.Println("你确认要退出吗?y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)

		if choice == "y" || choice == "n" {
			this.loop = false
		}
		fmt.Println("你输入有误,重新输入y/n")
		if choice == "y" {
			break
		}
	}
}

func (this *Jizhang) Showdetail() {
	fmt.Println("---当前收支明细记录---")
	if this.flag {
		fmt.Println(this.detail)
	} else {
		fmt.Println("当前没有入账号出账号记录,试试交易一笔!")
	}
}
func (this *Jizhang) MainMenu() {
	for {
		fmt.Println("------记账项目-----")
		fmt.Println("1.收支明细")
		fmt.Println("2.收入登记")
		fmt.Println("3.支出登记")
		fmt.Println("4.退出")
		fmt.Println("请选择1-4:")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Showdetail()

		case "2":
			this.inpay()

		case "3":
			this.outpay()

		case "4":
			this.exit()

		default:
			fmt.Println("请输入正确的选")

		}
		if !this.loop {
			break
		}
	}

}

func main() {
	fmt.Println("通过面向对象的方式实现记账功能")
	acount := NewJizhang()
	acount.MainMenu()
}
