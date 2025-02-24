package main
import (
	"fmt"
	"os"
)
func PathExists (path string) (bool,error)  {
	_,err := os.Stat(path)
	//返回的错误为nil，说明文件或文件夹存在
	if err == nil {
		return true,nil
	}
	//返回的错误类型使用IsNotExist判断为true，说明文件或文件夹不存在
	if os.IsNotExist(err) {
		return false,nil
	}
	//返回的错误为其他类型，不确定是否存在
	return false,err
}
func main() {
}