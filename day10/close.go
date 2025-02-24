package main
import "fmt"
func main() {
	intChan :=make(chan int,3)
	intChan <- 99
	intChan <- 66
	close(intChan)
	// panic: send on closed channel
	// intChan <- 88
	n1 := <-intChan
	fmt.Println("ok",n1)
	// ok 99
}