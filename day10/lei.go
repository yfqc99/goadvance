package main
import "fmt"

type Cat struct {
	Name string
	Age int
}
func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string,10)
	m1 := make(map[string]string,20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"
	m2 := make(map[string]string,20)
	m2["hero1"] = "蜘蛛侠"
	m2["hero2"] = "钢铁侠"
	mapChan <- m1
	mapChan <- m2
	/* 	invalid operation: cannot indirect mapChan 
		(variable of type chan map[string]string) */
	// fmt.Println("mapChan:",*mapChan)

	var catChan chan Cat
	catChan = make(chan Cat,10)
	cat1 := Cat{"tom",18}
	cat2 := Cat{"jery",20}
	catChan <- cat1
	catChan <- cat2

	var catChan1 chan *Cat
	catChan1 = make(chan *Cat,10)
	catChan1 <- &cat1
	catChan1 <- &cat2

	var allChan chan interface{}
	allChan = make(chan interface{},10)
	allChan <- cat1
	allChan <- 10
	allChan <- "张三"
	allChan <- m1
	newCat := <-allChan
	//在运行的时候可以确定是Cat类型不会报错
	fmt.Printf("newCat的类型：%T，newCat的值：%v \n",newCat,newCat)
	//在编译时此处仍认为是空接口类型，所以会报错
	/* 	newCat.Name undefined (type interface{} 
		has no field or method Name) */
	a := newCat.(Cat)
	fmt.Println("名字：",a.Name)
	/* 	newCat的类型：main.Cat，newCat的值：{tom 18} 
		名字： tom */
}