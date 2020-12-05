export GOROOT=/usr/local/go
export GOPATH=/usr/local/gopath
PATH=$PATH:/usr/local/go/bin:/root/go/bin
export PATH


дgo����Ŀ��Ŀ¼/root/gowork/
��Ŀ¼commonʹ��go.mod��ʽ
��Ŀ¼toolsʹ��package��ʽ

[root@d1 gowork]# tree /root/gowork/ 
������ common
��?? ������ btest.go
��?? ������ ctest.go
��?? ������ go.mod
������ go.mod
������ main.go
������ runtime001.go
������ tools
    ������ mydb.go

[root@d1 gowork]# cat /root/gowork/main.go
package main

import (
        "fmt"
        "app/tools"
        "common"
)

func begin() {
        fmt.Println("bebgin...")
}

func main(){
    begin()

    tools.Do_this()

    common.Test2()
    common.Test3()

    //check_os in runtime001.go
    check_os()
}
[root@d1 gowork]# cat runtime001.go
package main

import (
        "fmt"
        "runtime"
)

func check_os() {
        slice := make([]int, 0, 100)
        hash := make(map[int]bool, 10)
        fmt.Println(slice)
        fmt.Println(hash)

        ch := make(chan string)

        go func() {
                windos := runtime.GOOS

                ch <- windos
        }()

        msg := <-ch
        fmt.Println(msg)
}

[root@d1 gowork]# cat /root/gowork/go.mod
module app

go 1.14

require (
        common v1.0.1
        github.com/gin-gonic/gin v1.6.3
        github.com/go-sql-driver/mysql v1.5.0
)
replace common => ./common

[root@d1 gowork]# cat /root/gowork/common/go.mod
module common

go 1.14
[root@d1 gowork]# cat /root/gowork/common/btest.go
package common

func Test1() {
        println("test1 function in: common/btest.go file")
}
func Test3() {
        println("Test3 function in: common/btest.go file")
}

[root@d1 gowork]# cat /root/gowork/common/ctest.go
package common
func Test2() (int, string) {
        println("Test2() function in: common/btest.go file")
        const number int = 100
        str_one := "test2"
        return number, str_one

}
[root@d1 gowork]# cat /root/gowork/tools/mydb.go
package tools
import (
        "fmt"
)

func Do_this() {
        fmt.Println("Do_this function in: tools/mydb.go file")
}

�������
go run *.go
��
go build .
./app
