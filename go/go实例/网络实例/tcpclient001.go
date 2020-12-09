package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func main() {
        server :="serverip:81"
        conn, err := net.Dial("tcp", server)
        if err != nil {
                fmt.Println("Dial failed:", err)
                os.Exit(1)
        }

        for {
                reader := bufio.NewReader(os.Stdin)
                line, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Println("ReadString failed:", err)
                }
                //控制用户的退出
                line = strings.Trim(line, "\r\n")
                if line == "exit" {
                        fmt.Println("客户端退出...")
                        break
                }

                //发信息给服务器
                size_wr, err := conn.Write([]byte(line + "\n"))
                if err != nil {
                        fmt.Println("Write failed:", err)
                }
                fmt.Printf("发信息给服务器size_wr大小为:%v\n", size_wr)

                //显示收到服务器回复到信息
                buf := make([]byte, 1024)
                rd_size, err := conn.Read(buf)
                data := string(buf[:rd_size])
                if err != nil {
                        fmt.Println("read data from server err=%v\n", err)
                }
                fmt.Printf("收到服务器的回复信息大小:%v,数据为:%v\n", rd_size, data)
        }
}
