package main
import "fmt"

type Teacher struct{
	Name string
}
//接收变量的地址值
func (t *Teacher) change()  {
	//两种方式等价
	//t.Name = "周六"
	(*t).Name = "周六"
	/* fmt.Println(t)
	&{周六} */
	fmt.Printf("t的地址：%p",t)
	// t的地址：0xc000042270
}
func main() {
	var t *Teacher=&Teacher{}
	(*t).Name = "王五"
	fmt.Print((*t).Name,"老师被调课为")
	t.change()
	fmt.Println((*t).Name,"老师")
	// 王五老师被调课为周六 老师

	var t1 Teacher
	t1.Name = "张三"
	(&t1).change()
	//编译器底层会做出优化
	//自动识别到需要传入的是指针类型,传入变量的地址值
	t1.change()
	fmt.Println(t1.Name)
	// 周六

	fmt.Printf("t的地址：%p",t)
	// t的地址：0xc000042270
}