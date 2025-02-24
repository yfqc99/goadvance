package main
import "fmt"
func main() {
	//只写
	var intChan chan<- int
	//无缓冲管道，在通信时要求读取和写入同时准备好
	// intChan = make(chan int)
	//有缓冲管道，可以在没有读取操作的情况下缓存一定数量的数据
	intChan = make(chan int,2)
	intChan <- 66
	/* invalid operation: cannot receive from 
	send-only channel intChan (variable of type chan<- int) */
	// num := <- intChan
	// fmt.Println(num)
	//只读
	var intChan1 <-chan int
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send]:
	// intChan1 = make(<-chan int,2)
	intChan1 = make(chan int,2)
	/* invalid operation: cannot send to 
	receive-only channel intChan1 (variable of type <-chan int) */
	// intChan1 <- 66
	num := <- intChan1
	fmt.Println(num)
}