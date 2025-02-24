package main
import (
	"fmt"
	"Ajinjie/day3/mod"
)
func main() {
	p := mod.NewPerson("张三")
	p.SetAge(20)
	p.SetSal(6666)
	fmt.Println(p)
	// &{张三 20 6666}
}