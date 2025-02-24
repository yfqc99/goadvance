package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)
func main() {
	file , err :=os.Open("G:/goStudy/src/Ajinjie/document/test.txt")
	if err != nil{
		fmt.Println("open file err",err)
	}
	//当函数退出时，及时关闭file，否则会有内存泄露
	defer file.Close()
	//创建Reader，带缓冲
	const (
		defaultBufsize = 4096 //默认的缓冲区为4096
		//缓冲区可以使读取文件时读一部分处理一部分
		//而不直接将文件读到内存中，使得可以处理比较大的文件
	)
	reader :=bufio.NewReader(file)
	for {
		//循环读取文件内容，读到换行结束一次
		str ,err :=reader.ReadString('\n')
		//io.EOF表示读取到文件末尾
		if err == io.EOF {
			break
		}
		//在读取时会将换行符也读取到
		fmt.Print(str)
		//最后一行的末尾不是\n，是EOF
		/* 	hello world
			我爱北京
			go,hello bad! */
	}
	fmt.Println("文件读取结束")
	// 文件读取结束
}