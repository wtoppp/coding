//面向过程的实例1，实现一个账号资金的入账、出账、显示明细管理。
package main

import (
	"fmt"
)

func main() {
	var loop bool = true
	//接收用户选择输入的数字
	var key = ""

	//当前余额
	balance := 100.0
	//进出账金额
	money := 0.0
	//进账原因说明
	note := ""
	//有收入时用detail进行拼接
	detail := "收入：\t进账\t余额\t说明\n"

	//判断是否有收入或支出
	var flag bool = false

	for {
		fmt.Println("------记账项目-----")
		fmt.Println("1.收支明细")
		fmt.Println("2.收入登记")
		fmt.Println("3.支出登记")
		fmt.Println("4.退出")
		fmt.Println("请选择1-4:")

		fmt.Scanln(&key)
		switch key {
		case "1":

			fmt.Println("---当前收支明细记录---")
			if flag {
				fmt.Println(detail)
			} else {
				fmt.Println("当前没有入账号出账号记录,试试交易一笔!")
			}

		case "2":
			fmt.Println("收入登记")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次入账原因是：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v\n", money, balance, note)
			flag = true
		case "3":
			fmt.Println("出账登记")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break

			}
			balance -= money
			fmt.Println("本次出账原因是：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%v\t%v\t%v\n", money, balance, note)
			flag = true

		case "4":
			fmt.Println("你确认要退出吗?y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" {
					break
				}
				fmt.Println("你输入有误,请输入y/n")
				if choice == "y" || choice == "Y" {
					loop = false
				}
			}

			fmt.Println("退出程序")
			break

		default:
			fmt.Println("请输入正确的选")

		}
		if !loop {
			break
		}
	}

}
