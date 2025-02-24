package main
import "fmt"

type A interface{
	test()
	B
	C
}
type B interface{
	test1()
}
type C interface{
	test2()
}
type Student struct{
}
func (stu Student) test()  {
	fmt.Println("test")
}
func (stu Student) test1()  {
	fmt.Println("test1")
}
/* cannot use stu (variable of type Student) as A value in variable 
declaration: Student does not implement A (missing method test2) */
func (stu Student) test2()  {
	fmt.Println("test2")
}
func main() {
	var stu Student
	var a A = stu
	a.test()
	a.test1()
	a.test2()
	/* 	test
		test1
		test2 */
	var t interface{} = stu
	fmt.Println(t)
	// {}
	var num int =66
	t = num
	fmt.Println(t)
	// 66
}