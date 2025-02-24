package main
import (
	"fmt"
	"time"
)
func sayHello(){
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}
func test()  {
	// 使用defer+recover
	// 在协程出现恐慌时不影响其他
	defer func ()  {
		if err := recover();err !=nil {
			fmt.Println("test painc：",err)
		}
	}()
	var myMap map[int]string
	// panic: assignment to entry in nil map
	myMap[0]="golang"
}
func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("hello,main")
		time.Sleep(time.Second)
	}
	/* 	hello,main
		test painc： assignment to entry in nil map
		hello,main
		hello,world */
}