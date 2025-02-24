package main
import "fmt"
func main() {
	var x interface{}
	var b float32 = 1.1
	x = b
	if y,ok := x.(float64); ok {
		fmt.Printf("y的类型：%T，值：%v \n",y,y)
	}else{
		fmt.Println("转换失败")
	}
	fmt.Println("继续执行")
	/* 	转换失败
		继续执行 */
}