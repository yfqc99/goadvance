package main
import "fmt"
func main() {
	intChan :=make(chan int,10)
	for i := 0; i <10; i++ {
	intChan <- i
	}
	stringChan :=make(chan string,5)
	for i := 0; i <5; i++ {
		stringChan <- "hello"+fmt.Sprintf("%d",i)
	}
	//传统遍历需要关闭管道，但关闭时机不好确定
	label:
	for{
		select{
			//管道一直没有关闭，不会一直阻塞
			//会自动到下一个case匹配
			case v := <- intChan :
				fmt.Println("intChan读取到数据：",v)
			case v := <- stringChan :
				fmt.Println("stringChan读取到数据：",v)
			default :
				fmt.Println("都取不到")
				//退出函数
				// return
				break label
		}
	}
	/* 	intChan读取到数据： 0
		stringChan读取到数据： hello0
		stringChan读取到数据： hello1
		intChan读取到数据： 1
		intChan读取到数据： 2
		stringChan读取到数据： hello2
		intChan读取到数据： 3
		intChan读取到数据： 4
		stringChan读取到数据： hello3
		intChan读取到数据： 5
		stringChan读取到数据： hello4
		intChan读取到数据： 6
		intChan读取到数据： 7
		intChan读取到数据： 8
		intChan读取到数据： 9
		都取不到 */
}