package main
import (
	"fmt"
	"os"
)
func main() {
	//获取到命令行里的动态参数
	// .\test.exe com G:\goStudy\src\Ajinjie\document\test.txt 999
	fmt.Println("命令行参数：",len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
	/* 	命令行参数： 4
		G:\goStudy\src\Ajinjie\day7\test.exe
		com
		G:\goStudy\src\Ajinjie\document\test.txt
		999 */
}