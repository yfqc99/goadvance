package main
import (
	"fmt"
	"io"
	"os"
	"bufio"
)
func CopyFile(dstFileName string,srcFileName string) (written int64,err error) {
	srcfile, err :=os.Open(srcFileName)
	if err !=nil {
		fmt.Println("文件打开错误",err)
	}
	defer srcfile.Close()
	reader :=bufio.NewReader(srcfile)
	//因为文件可能不存在
	dstfile, err1 :=os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err1 !=nil {
		fmt.Println("文件打开错误",err1)
	}
	defer dstfile.Close()
	writer :=bufio.NewWriter(dstfile)
	return io.Copy(writer,reader)
}
func main() {
	srcfile :="F:/downloads/未命名文件.png"
	/* 	open G:/goStudy/src/Ajinjie/document: is 
		a directory */
	// dstfile :="G:/goStudy/src/Ajinjie/document"
	dstfile :="G:/goStudy/src/Ajinjie/document/image.png"
	_,err :=CopyFile(dstfile,srcfile)
	if err == nil {
		fmt.Println("拷贝完成")
	}else {
		fmt.Println(err)
	}
	// 拷贝完成
}