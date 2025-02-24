package main
import "fmt"

type A interface {
	Say()
}
type B struct{
	Name string
}
func (b B) Say()  {
	fmt.Println("b say")
}
type myInt int
func (m myInt) Say()  {
	fmt.Println("myInt say")
}
func main() {
	var m myInt
	var a1 A = m
	a1.Say()
	// myInt say
	var a A
	// <nil>
	fmt.Println(a)
	var b B
	a = b
	a.Say()
	// b say
}