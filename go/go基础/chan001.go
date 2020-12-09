//goroutine和channel的协同工作,实现以下两个功能：
//1. 开一个协程A，向管道中写入20个整数
//2. 再开一个协程B,从管道中读取数据
//3. 主线程需要等待A和B都完成工作后才可以退出（通过一个临时的bool型管道Exitchan与主线程交互，解决主线程提前退出的问题）
package main

import (
	"fmt"
	"time"
)

func WriteData(ch1 chan int) {
	for i := 0; i < 20; i++ {
		ch1 <- i
		fmt.Println("writedata:", i)
		time.Sleep(time.Nanosecond * 500)
	}
	close(ch1)
}

func ReadData(Rch chan int, Exitchan chan bool) {
	for {
		v, ok := <-Rch
		if !ok {
			fmt.Println("<<<< readdata breaked")
			break
		}
		fmt.Println("readdata:", v)
	}
	//读完了管道中的数据后，写个true到Exitchan管道中，读主线程去读
	Exitchan <- true
	close(Exitchan)
}

func main() {
	ch1 := make(chan int, 20)
	Exitchan := make(chan bool, 1)
    //开启两个协程
	go WriteData(ch1)
	go ReadData(ch1, Exitchan)

	for {
		//循环读取Exitchan管道中的数据，直到读完。
		rd, ok := <-Exitchan
		if !ok {
			fmt.Println("<<<< master breaked")
			break
		}
		fmt.Println("got value:", rd)
	}
}
