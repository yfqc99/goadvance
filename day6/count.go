package main
import (
	"fmt"
	"io"
	"os"
	"bufio"
)
type charCount struct{
	chCount int64 //英文个数
	numCount int64 //数字个数
	spaceCount int64 //空格个数
	otherCount int64 //其他个数
}
func main() {
	filePath := "G:/goStudy/src/Ajinjie/document/test.txt"
	file, err :=os.Open(filePath)
	if err != nil {
		fmt.Println("打开失败",err)
		return
	}
	defer file.Close()
	var count charCount
	reader :=bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		// 默认遍历str字符串时按一个字符遍历
		for _, v := range str {
			// v的类型：int32，rune切片
			//byte切片是uint8
			//做if分支函数使用
			switch  {
			case v >= 'a' && v<='z':
				fallthrough
			case v >= 'A' && v<='Z':
				count.chCount++
			case v == ' ' && v == '\t':
				count.spaceCount++
			case v >= '0' && v<='9':
				count.numCount++
			default :
				count.otherCount++
			}
		}
	}
	//换行末尾有\r\n
	fmt.Println(count)
	// {9 7 0 25}
	//[]rune是基于字符串创建的，所以需要先转为字符串
	// strRune := []rune(str)
}