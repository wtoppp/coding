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
                //�����û����˳�
                line = strings.Trim(line, "\r\n")
                if line == "exit" {
                        fmt.Println("�ͻ����˳�...")
                        break
                }

                //����Ϣ��������
                size_wr, err := conn.Write([]byte(line + "\n"))
                if err != nil {
                        fmt.Println("Write failed:", err)
                }
                fmt.Printf("����Ϣ��������size_wr��СΪ:%v\n", size_wr)

                //��ʾ�յ��������ظ�����Ϣ
                buf := make([]byte, 1024)
                rd_size, err := conn.Read(buf)
                data := string(buf[:rd_size])
                if err != nil {
                        fmt.Println("read data from server err=%v\n", err)
                }
                fmt.Printf("�յ��������Ļظ���Ϣ��С:%v,����Ϊ:%v\n", rd_size, data)
        }
}
