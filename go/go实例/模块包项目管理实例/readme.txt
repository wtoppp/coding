go模块包管理应用于zabbix3.X zabbix5.0 api

包管理说明：
包-路径-文件-目录名关系
例：目录结构
bin
---pkg
---src
---pk1
------pk2
---------function1.go
---------function2.go
---index.go

1. function1.go 文件内容：
package pk3
 func Function_test3()  {
    println("function_test3")
}

2. function2.go 文件内容：
package pk3
func Function_test4()  {
    println("function_test4")
}

3. index.go 文件内容
package main
import "pk1/pk2"
 
func main() {
    pk3.Function_test4()
}

4.运行 index.go，输出：
function_test4

得出以下结论：
1、import 导入的参数是路径，而非包名。
2、尽管习惯将包名和目录名保证一致，但这不是强制规定；

3、在代码中引用包成员时，使用包名而非目录名；
4、同一目录下，所有源文件必须使用相同的包名称（因为导入时使用绝对路径，所以在搜索路径下，包必须有唯一路径，但无须是唯一名字）；
5、至于文件名，更没啥限制（扩展名为.go）;
