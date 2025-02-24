package main
import "fmt"
func main() {
	//普通for不要使用，因为管道长度会变化
	intChan :=make(chan int,50)
	for i := 0; i <50; i++ {
		intChan <- i*2
	}
	//因为管道里不会返回下标，因为它只能顺序遍历
	//如果没有关闭，则会一直遍历，因为没有结束的标志
	//一直等待就会报错
	// fatal error: all goroutines are asleep - deadlock! 
	close(intChan)
	for v := range intChan {
		fmt.Println("v=",v)
	}
}