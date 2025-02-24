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
func (p phone) Call()  {
	fmt.Println("正在打电话")
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
type Computer struct{
}
func (c Computer) Working(usb Usb)  {
	usb.Start()
	if ph,ok :=usb.(phone); ok{
		ph.Call()
	}
	usb.Stop()
}
func main() {
	var usb [3]Usb
	fmt.Println(usb)
	usb[0] = phone{"小米"}
	usb[1] = Cammer{"ccd"}
	usb[2] = phone{"华为"}
	var com Computer
	for _, v := range usb {
		com.Working(v)
	}
	/* 	手机开始工作
		正在打电话
		手机停止工作
		相机开始工作
		相机停止工作
		手机开始工作
		正在打电话
		手机停止工作 */
}