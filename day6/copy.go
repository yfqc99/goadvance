package main
import (
	"fmt"
	"io/ioutil"
)
func main() {
	//将文件内容读取到内存
	//将读取到的内容写入
	filePath :="G:/goStudy/src/Ajinjie/document/wr.txt"
	filePath1 :="G:/goStudy/src/Ajinjie/document/test.txt"
	content,err :=ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("读取错误",err)
		return
	}
	//直接传入切片
	//会清空原文件写入
	err = ioutil.WriteFile(filePath1,content,0666)
	if err != nil{
		fmt.Println("写入错误",err)
	}
}