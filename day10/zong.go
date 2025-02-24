package main
import "fmt"
//传入管道后，使该管道只能写入
func writeData (intChan chan<- int){
    for i := 1; i <= 50; i++ {
        intChan <- i
        fmt.Println("write=",i)
    }
    //用close做结束标记
    close(intChan)
}
func readData(intChan chan int,exitChan chan bool )  {
    for{
        //ok取值是否成功，当读到关闭时读取失败
        v,ok := <- intChan
        if !ok {
            break
        }
        fmt.Println("v=",v)
    }
    //读取完成后
    exitChan <- true
    close(exitChan)
}
func main() {
    intChan := make(chan int)
    exitChan := make(chan bool)
    go writeData(intChan)
    // fatal error: all goroutines are asleep - deadlock!
    // go readData(intChan,exitChan)
    //没有做任何工作主线程会很快停止工作
    for{
        _,ok := <- exitChan
        if !ok{
            break
        }
    }
    
   /*   write= 1
        v= 1
        v= 2
        write= 2
        write= 3
        v= 3
        v= 4
        write= 4
        write= 5
        v= 5
        v= 6
        write= 6
        write= 7
        v= 7
        v= 8
        write= 8
        write= 9
        v= 9
        v= 10
        write= 10
        write= 11
        v= 11
        v= 12
        write= 12
        write= 13
        v= 13
        v= 14
        write= 14
        write= 15
        v= 15
        v= 16
        write= 16
        write= 17
        v= 17
        v= 18
        write= 18
        write= 19
        v= 19
        v= 20
        write= 20
        write= 21
        v= 21
        v= 22
        write= 22
        write= 23
        v= 23
        v= 24
        write= 24
        write= 25
        v= 25
        v= 26
        write= 26
        write= 27
        v= 27
        v= 28
        write= 28
        write= 29
        v= 29
        v= 30
        write= 30
        write= 31
        v= 31
        v= 32
        write= 32
        write= 33
        v= 33
        v= 34
        write= 34
        write= 35
        v= 35
        v= 36
        write= 36
        write= 37
        v= 37
        v= 38
        write= 38
        write= 39
        v= 39
        v= 40
        write= 40
        write= 41
        v= 41
        v= 42
        write= 42
        write= 43
        v= 43
        v= 44
        write= 44
        write= 45
        v= 45
        v= 46
        write= 46
        write= 47
        v= 47
        v= 48
        write= 48
        write= 49
        v= 49
        v= 50
        write= 50 */
}