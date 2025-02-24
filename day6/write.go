package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	//文件名，打开方式（可以组合使用），权限控制（windows下没用）
	//创建新文件
	filePath :="G:/goStudy/src/Ajinjie/document/wr.txt"
	//打开文件
	//写文件模式，该文件不存在就创建该文件
	file, err :=os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败",err)
		return
	}
	//关闭文件句柄
	defer file.Close()
	//写入，因为编辑器对换行的识别不同
	str :="hello,go\r\n"
	//使用带缓存的方式写文件
	writer :=bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为是带有缓存的写入
	//内容一开始在缓存区，所以需要Flush方法写入磁盘
	//否则文件中不存在数据
	writer.Flush()
}