package main
import "fmt"

type Teacher struct{
	Name string
}
//接收者为值类型时，不能将指针类型的数据直接传递
//相反一样
func test(t Teacher)  {
	fmt.Println("test：",t.Name)
	// test： 张三
}
func test2(t *Teacher)  {
	fmt.Println("test2：",t.Name)
	// test2： 张三
}

//接收者为值类型时，可以直接用指针类型的变量调用方法
//相反一样
func (t Teacher)test1()  {
	t.Name = "王五"
	fmt.Println("test1：",t.Name)
	// test1： 王五
}
func (t *Teacher)test3()  {
	t.Name = "麻子"
	fmt.Println("test3：",t.Name)
	// test3： 麻子
}
func main() {
	t := Teacher{
		Name :"张三",
	}
	test(t)
	// test： 张三
	/* cannot use &t (value of type *Teacher) 
	as Teacher value in argument to test */
	// test(&t)
	/* cannot use t (variable of type Teacher) 
	as *Teacher value in argument to test2 */
	// test2(t)
	t2 := &Teacher{
		Name :"张三",
	}
	test2(t2)
	// test2： 张三

	t1 := Teacher{"周六"}
	t1.test1()
	// test1： 王五
	//底层优化，识别到是值类型后会自动进行值拷贝
	(&t1).test1()
	fmt.Println("main的test1：",t1.Name)
	// main的test1： 周六
	t1.test3()
	fmt.Println("main的test3：",t1.Name)
	// main的test3： 麻子

}