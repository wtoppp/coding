//������̵�ʵ��1��ʵ��һ���˺��ʽ�����ˡ����ˡ���ʾ��ϸ����
package main

import (
	"fmt"
)

func main() {
	var loop bool = true
	//�����û�ѡ�����������
	var key = ""

	//��ǰ���
	balance := 100.0
	//�����˽��
	money := 0.0
	//����ԭ��˵��
	note := ""
	//������ʱ��detail����ƴ��
	detail := "���룺\t����\t���\t˵��\n"

	//�ж��Ƿ��������֧��
	var flag bool = false

	for {
		fmt.Println("------������Ŀ-----")
		fmt.Println("1.��֧��ϸ")
		fmt.Println("2.����Ǽ�")
		fmt.Println("3.֧���Ǽ�")
		fmt.Println("4.�˳�")
		fmt.Println("��ѡ��1-4:")

		fmt.Scanln(&key)
		switch key {
		case "1":

			fmt.Println("---��ǰ��֧��ϸ��¼---")
			if flag {
				fmt.Println(detail)
			} else {
				fmt.Println("��ǰû�����˺ų��˺ż�¼,���Խ���һ��!")
			}

		case "2":
			fmt.Println("����Ǽ�")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("��������ԭ���ǣ�")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n����\t%v\t%v\t%v\n", money, balance, note)
			flag = true
		case "3":
			fmt.Println("���˵Ǽ�")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("����")
				break

			}
			balance -= money
			fmt.Println("���γ���ԭ���ǣ�")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n֧��\t%v\t%v\t%v\n", money, balance, note)
			flag = true

		case "4":
			fmt.Println("��ȷ��Ҫ�˳���?y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" {
					break
				}
				fmt.Println("����������,������y/n")
				if choice == "y" || choice == "Y" {
					loop = false
				}
			}

			fmt.Println("�˳�����")
			break

		default:
			fmt.Println("��������ȷ��ѡ")

		}
		if !loop {
			break
		}
	}

}
