package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)
func main() {
	filePath :="G:/goStudy/src/Ajinjie/document/wr.txt"
	//读写文件模式，并且在原文件后追加
	file, err :=os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败",err)
		return
	}
	defer file.Close()
	//读取原文件内容，并显示
	reader := bufio.NewReader(file)
	for {
		str, err :=reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
		/* 	你好！chain
			你好！chain
			你好！chain
			你好！chain
			你好！chain
			ABCDEFG
			ABCDEFG
			ABCDEFG
			ABCDEFG
			ABCDEFG */
	}
	str :="上学ing~~~\r\n"
	writer :=bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}