package main

import (
        "fmt"
        "net"
        "strings"
)

func Handleclient(conn net.Conn) {
		//通过客户与服务器连接的conn，以便与客户端进行信息交互
        defer conn.Close()
        for {
                buf := make([]byte, 1024)
                //fmt.Printf("server waiting client %s,send message\n",conn.RemoteAddr().String())
                size_rd, err := conn.Read(buf)

                if err != nil {
                        fmt.Printf("client quit err=%v\n", err)
                        return
                }

                //显示客户端发过来的内容
                msg := string(buf[:size_rd])
                fmt.Printf(msg)

                //回复信息给客户端
                msg = strings.Trim(msg,"\n")
/*
                if msg == "hello" {
                        conn.Write([]byte("I just hello server..."))
                }else {
                        conn.Write([]byte("I from server:"+"welcome..."))
                }
*/
                switch msg {
                case "hello":
                        conn.Write([]byte("I just hello server..."))
                case "good":
                        conn.Write([]byte("I just good server..."))
                default:
                        conn.Write([]byte("I just default server..."))
                }
        }
}

func main() {

        listen, err := net.Listen("tcp", "0.0.0.0:81")
        fmt.Println("server is working...")
        if err != nil {
                return
        }
        defer listen.Close()

        for {
				//等待用户来连接
                conn, err := listen.Accept()
                if err != nil {
                        fmt.Println("Accept() err=", err)
                }else {
                        fmt.Printf("Accept() successful conn=%v,客户端ip=%v\n",conn,conn.RemoteAddr())
                }
                go Handleclient(conn)
        }
}
