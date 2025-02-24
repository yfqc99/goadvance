package main
import (
	"fmt"
)
func main() {
	//定义
	var intChan chan int
	intChan = make(chan int,3)
	fmt.Printf("chan本身的地址：%p，chan的值（指向的地址）：%v \n",&intChan,intChan)
	/* chan本身的地址：0xc00000a028，chan的值（指向的地址）：0xc000076100 */
	//写入数据
	intChan<- 10
	num := 211
	intChan<- num
	//长度和容量，管道不能自增长
	fmt.Printf("chan的长度：%v，chan的容量：%v \n",len(intChan),cap(intChan))
	// chan的长度：2，chan的容量：3
	intChan<- 66
	//超出容量报错
	// intChan<- 99
	// fatal error: all goroutines are asleep - deadlock!
	//管道可以边放边取
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)
	fmt.Printf("chan的长度：%v，chan的容量：%v \n",len(intChan),cap(intChan))
	/* 	10
		chan的长度：2，chan的容量：3 */
	//在没有协程的情况下，管道数据全部取出，再次取出报错
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println(num3," ",num4)
	// 211   66
	num5 := <-intChan
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(num5)
}