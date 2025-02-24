package main
import "fmt"

//声明接口
type Usb interface{
	//声明未实现方法
	Start()
	Stop()
}
type phone struct{
}
func (p phone) Start()  {
	fmt.Println("手机开始工作")
}
func (p phone) Stop()  {
	fmt.Println("手机停止工作")
}
type Cammer struct{
}
func (c Cammer) Start()  {
	fmt.Println("相机开始工作")
}
func (c Cammer) Stop()  {
	fmt.Println("相机停止工作")
}
type Computer struct{
}
//接收一个Usb接口类型变量
//实现接口，就是实现该接口声明的所有方法
func (c Computer) Working(usb Usb)  {
	usb.Start()
	usb.Stop()
}
func main() {
	var c Computer
	var ca Cammer
	var p phone
	c.Working(ca)
	c.Working(p)
	/* 	相机开始工作
		相机停止工作
		手机开始工作
		手机停止工作 */
}