package main
import "fmt"
type A interface {
	Say()
}
type B interface {
	Hello()
}
type Monster struct {
}
func (m Monster) Say()  {
	fmt.Println("monster say")
}
func (m *Monster) Hello()  {
	fmt.Println("monster hello")
}
/* func (m Monster) Hello()  {
	fmt.Println("monster hello")
} */
func main() {
	var m Monster
	var a A = m
	/* cannot use m (variable of type Monster) as B value in variable 
	declaration: Monster does not implement B (method Hello has pointer receiver) */
	// var b B = m
	// var b B = &m
	a.Say()
	b.Hello()
	/* 	monster say
		monster hello */
}