package main
import "fmt"
type Student struct{
	Name string
	Age int
}
//实现String()方法
func (stu *Student) String() string {
	str :=fmt.Sprintf("name=%v,age=%v",stu.Name,stu.Age)
	return str
}
func main() {
	stu :=Student{
		Name : "张三",
		Age : 20,
	}
	fmt.Println(stu)
	//自动调用String方法
	fmt.Println(&stu)
	/* 	{张三 20}
		name=张三,age=20 */
}