package main
import "fmt"

type A struct{
	Name string
	age int
}
type B struct {
	//不是匿名结构体
	//当要找当前结构体中没有的属性或方法时，会直接报错
	//因为编译器不会向上找属性或方法
	a A
}
func main() {
	var b B
	/* 	b.Name undefined (type B has no field 
		or method Name) */
	// b.Name = "张三"
	//要访问有名结构体字段或方法，必须带上而结构体名
	b.a.Name = "张三"
	fmt.Println(b)
	// {{张三 0}}
}