package main
import (
	"fmt"
	"reflect"
)
func test(b interface{})  {
	rVal := reflect.ValueOf(b)
	/* panic: reflect: reflect.Value.SetInt
	 using unaddressable value */
	fmt.Println("rVal的kind：",rVal.Kind())
	// rVal的kind： ptr
	// rVal.SetInt(20)
	//Elem返回v持有接口保管的值，或v持有指针指向的值的Value封装
	//类似于指针获得变量地址后，指向该地址后修改变量
	rVal.Elem().SetInt(20)
}
func main() {
	num := 66
	//相当于值拷贝
	// test(num)
	test(&num)
	fmt.Println(num)
	// 20
	
	str := "张三"
	// fs := reflect.ValueOf(str) string类别
	// reflect.Value.SetString using unaddressable value
	// fs.SetString("周六")
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("周六")
	fmt.Println(str)
	// 周六
}