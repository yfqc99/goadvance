package main
import "fmt"

type myInt int
type myFloat float32
func (i myInt) print()  {
	fmt.Println("i=",i)
}
func (f *myFloat) change()  {
	/* invalid operation: f + 1 (mismatched 
		types *myFloat and untyped int) */
	// f = f+1
	*f = *f + 1
}
func main() {
	var i myInt = 10
	i.print()
	// myFloat (type) is not an expression
	// var f *myFloat = &myFloat
	/* cannot use 3.2 (untyped float constant) 
	as *myFloat value in assignment */
	var f myFloat = 3.2
	//可以自动转换
	f.change()
	fmt.Println("f=",f)
	// f= 4.2
}