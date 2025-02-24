package main
import (
	"fmt"
	"io/ioutil"
)
func main() {
	//一次性读取
	file := "G:/goStudy/src/Ajinjie/document/test.txt"
	content, err :=ioutil.ReadFile(file)
	// ReadFile内有封装的文件打开和关闭
	if err != nil{
		fmt.Println("read err",err)
	}
	//输出的是[]byte切片
	fmt.Println(content)
	/* 	[104 101 108 108 111 32 119 111 114 108 100 13 10 230 
		136 145 231 136 177 229 140 151 228 186 172 13 10 103 
		111 44 104 101 108 108 111 32 98 97 100 33 13 10] */
	fmt.Println(string(content))
	//在读取的时候EOF行也被读取到
	/* 	hello world
		我爱北京
		go,hello bad!
		*/
}