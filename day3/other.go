package main
import "fmt"

type A struct {
	Name string
	age int
}
type B struct {
	Name string
	score float64
}
type C struct {
	A
	B
}
func main() {
	var c C
	// ambiguous selector c.Name
	//当C结构体中也有Name则不会报错
	//因为编译器根据就近原则直接找到要赋值的变量
	c.A.Name ="张三"
	fmt.Println(c)
	// {{张三 0} { 0}}
}