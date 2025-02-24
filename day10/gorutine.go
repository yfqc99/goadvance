package main
import (
	"fmt"
	"time"
)
//每个1秒输出helloworld
func test()  {
	for i := 1; i <=10; i++ {
		fmt.Println("test hello,world",i)
		//休眠
		time.Sleep(time.Second)
	}
}
func main() {
	//test执行完后再执行main
	// test()
	//test和main穿插执行
	//开启一个协程
	go test()
	for i := 1; i <=10; i++ {
		fmt.Println("main hello,world",i)
		time.Sleep(time.Second)
	}
	/* 	main hello,world 1
		test hello,world 1
		test hello,world 2
		main hello,world 2
		main hello,world 3
		test hello,world 3
		test hello,world 4
		main hello,world 4 */
}