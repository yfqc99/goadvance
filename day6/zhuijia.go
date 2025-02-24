package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	filePath :="G:/goStudy/src/Ajinjie/document/wr.txt"
	//写文件模式，并且在原文件后追加
	file, err :=os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败",err)
		return
	}
	defer file.Close()
	str :="ABCDEFG\r\n"
	writer :=bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}