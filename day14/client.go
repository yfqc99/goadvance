package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)
func main() {
	//连接协议，服务器IP+端口
	conn ,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil{
		fmt.Println("客户端连接失败",err)
		return
	}
	fmt.Println("连接成功，conn=",conn)
	// 连接成功，conn= &{{0xc000100a00}}
	//从控制台输入发送给服务器，在服务器接收显示
	//接收键盘输入，可以一行行输入
	// os.Stdin标准输入，一般代表键盘输入
	//终端输入先写到一个文件中，在通过该文件读取到程序里
	reader := bufio.NewReader(os.Stdin)
	//从终端读取一行用户输入，并准备发送到服务器
	//返回字符串
	for{
		line, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("读取失败",err)
		}
		//去除\r\n和空格
		line = strings.Trim(line," \r\n")
		if line == "exit"{
			fmt.Println("客户端退出")
			break
		}
		//发送
		//参数是一个切片
		// n,err := conn.Write([]byte(line))
		_,err1 := conn.Write([]byte(line + "\n"))
		if err1 != nil{
			fmt.Println("当前发送失败",err1)
		}
	}
	// fmt.Println("客户端发送",n,"字节的数据")
	/* 	helloworld,abc!
		客户端发送 17 字节的数据 */
}