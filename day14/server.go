package main
import (
	"fmt"
	"net"
)
func process(conn net.Conn)  {
	//循环接收客户端数据
	defer conn.Close()
	for{
		//创建一个新的切片
		buf := make([]byte, 1024)
		//客户端没有发送信息会一直等待，协程会阻塞在这里
		// fmt.Println("服务器在等待客户端发送信息")
		n,err := conn.Read(buf)
		if err != nil{
			//err是一个字符串
			fmt.Println("接收失败",err)
			//可能客户端已经关闭
			return
		}
		 /* 服务器在等待客户端发送信息
		接收失败 read tcp 127.0.0.1:8888->127.0.0.1:64597: wsarecv: 
		An existing connection was forcibly closed by the remote host. */
		//显示客户端发送内容到服务器终端
		//在客户端发送时会将\n读入
		//否则会将剩余的全部输出
		fmt.Print(string(buf[:n]))
		// helloworld,abc!
	}
}
func main() {
	fmt.Println("服务器开始监听")
	//通过tcp协议监听
	//127.0.0.1只支持IPv4
	//0.0.0.0支持IPv4和IPv6
	listen, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败",err)
		return
	}
	//监听一下结束
	fmt.Println("listen=",listen)
	/* 	服务器开始监听
		listen= &{0xc000100a00 {<nil> 0}} */
	//循环等待连接接口
	for{
		fmt.Println("等待客户端连接")
		//返回连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("当前连接失败",err)
		}else{
			fmt.Println("当前连接成功，conn=",conn)
			// conn.RemoteAddr()返回的是Addr接口类型，调用里面的String方法获得IP字符串
			fmt.Println("客户端IP=",conn.RemoteAddr().String())
			// 客户端IP= 127.0.0.1:63486
			//在主线程写接收，可能会发生阻塞
			go process(conn)
		}
		//在终端使用telnet测试是否连接，ctrl+]终止，quit
		//开启一个协程，为当前客户端服务
		/* 	当前连接成功，conn= &{{0xc000090280}}
			等待客户端连接 */
	}
	// defer listen.Close()
}