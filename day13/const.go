package main
import "fmt"
func main() {
	var num int
	//变量类型可写可不写
	// const tax int
	// missing init expr for tax
	const tax int = 66
	fmt.Println(num," ",tax)
	num = 99
	// tax = 88
	// cannot assign to tax (constant 66 of type int)
	fmt.Println(num," ",tax)
	const b = 9/3
	fmt.Println(b)
	num = 9
	// const c = num/3
	// num / 3 (value of type int) is not constant
	//因为编译器认为变量是有可能改变的
	// fmt.Println(c)
	//同样报错
	// const c = getVal()
	/* 	0   66
		99   66
		3 */
	const(
		a = 1
		d = 2
	)
	fmt.Println(a," ",d)
	//表示给i赋值为0，i1在i的基础上+1，i2在i1的基础上+1
	const(
		i = iota
		i1 
		i2
	)
	fmt.Println(i," ",i1," ",i2)
	/* 	1   2
		0   1   2 */
}