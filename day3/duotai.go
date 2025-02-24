package main
import "fmt"

type Usb interface{
	Start()
	Stop()
}
type phone struct{
	Name string
}
func (p phone) Start()  {
	fmt.Println("手机开始工作")
}
func (p phone) Stop()  {
	fmt.Println("手机停止工作")
}
type Cammer struct{
	Name string
}
func (c Cammer) Start()  {
	fmt.Println("相机开始工作")
}
func (c Cammer) Stop()  {
	fmt.Println("相机停止工作")
}
func main() {
	//普通数组只能存放一种指定类型的数据
	//可以通过接口，实现存放多种类型的数据
	var usb [3]Usb
	fmt.Println(usb)
	usb[0] = phone{"小米"}
	usb[1] = Cammer{"ccd"}
	usb[2] = phone{"华为"}
	fmt.Println(usb)
	/* 	[<nil> <nil> <nil>]
		[{小米} {ccd} {华为}] */
}