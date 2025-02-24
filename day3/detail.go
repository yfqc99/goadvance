package main
import "fmt"

type A struct {
	Name string
	age int
}
func (a *A) SayOk ()  {
	fmt.Println("A SayOk",a.Name)
}
func (a *A) hello ()  {
	fmt.Println("A hello",a.Name)
}
type B struct{
	A
	Name string
}
func (b *B) SayOk ()  {
	fmt.Println("A SayOk",b.Name)
}
func main() {
	var b B
	//会直接赋值到B结构体的Name变量
	//A中Name不会被赋值
	b.Name = "麻子"
	b.A.Name = "大娃"
	b.age = 18
	b.SayOk()
	b.hello()
	/* 	A SayOk 麻子
		A hello */
	/* 	A SayOk 麻子
		A hello 大娃 */
	
}