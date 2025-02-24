package main
import (
	"fmt"
	"flag"
)
func main() {
	//定义变量接收命令行参数值
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user,"u","","获取用户名，默认为空")
	flag.StringVar(&pwd,"pwd","","获取密码，默认为空")
	flag.StringVar(&host,"h","localhost","获取主机名，默认为localhost")
	flag.IntVar(&port,"port",3306,"获取端口号，默认为3306")
	//从切片中解析注册的flag，在所有flag注册好后，并没有访问前执行
	flag.Parse()
	fmt.Println("user=",user," ","pwd=",pwd," ","host=",host," ","port=",port)
	/* 	.\flag.exe -u root -pwd root -h 127.0.0.5 -port 8080
		user= root   pwd= root   host= 127.0.0.5   port= 8080 */
}