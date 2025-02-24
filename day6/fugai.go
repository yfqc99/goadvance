package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	filePath :="G:/goStudy/src/Ajinjie/document/wr.txt"
	//写文件模式，并且清空原文件
	file, err :=os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败",err)
		return
	}
	defer file.Close()
	str :="你好！chain\r\n"
	writer :=bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}