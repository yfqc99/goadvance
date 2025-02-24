package main
import "fmt"
type Point struct {
	x int
	y int
}
func main() {
	var a interface{}
	var point Point = Point{1,2}
	a = point
	var b Point
	/* cannot use a (variable of type interface{}) 
	as Point value in assignment: need type assertion */
	// b = a
	//断言，当a可以转换为Point类型时可以转换成功
	b = a.(Point)
	fmt.Println(b)
	// {1 2}
	c := a.(int)
	// interface {} is main.Point, not int
	fmt.Println(c)

}
