package main

import (
        "io/ioutil"
        "fmt"
)

func main() {
	//ע��ioutiһ�����ڶ�дС�ļ�
        file1 := "/tmp/log.txt"
        file2 := "/tmp/log2.txt"
        content,err :=ioutil.ReadFile(file1)
        if err !=nil{
                fmt.Println("read file err=%v",err)
                return
        }
        err =ioutil.WriteFile(file2,content,0644)
                if err !=nil{
                fmt.Println("write file err=%v",err)
                return
        }

}

