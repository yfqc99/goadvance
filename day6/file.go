package main
import (
	"fmt"
	"os"
)
func main() {
	//打开文件
	file , err :=os.Open("G:/goStudy/src/Ajinjie/document/test.txt")
	if err != nil{
		fmt.Println("open file err",err)
	}
	//输出文件
	fmt.Println(file)
	// &{0xc000100780}
	/* 	open file err open document/test.txt: 
	The system cannot find the path specified.
		<nil> */
	//关闭文件
	err = file.Close()
	if err != nil{
		fmt.Println("close file err",err)
	}
}