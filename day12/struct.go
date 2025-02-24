package main
import (
	"fmt"
	"reflect"
)
type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32
	Sex string
}
func (s Monster) Print()  {
	fmt.Println("开始")
	fmt.Println(s)
	fmt.Println("结束")
}
func (s Monster) GetSum(n1,n2 int) int {
	return n1+n2
}
func (s Monster) Set(name string,age int,score float32,sex string)  {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{})  {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	//reflect包下的常量
	if kd != reflect.Struct{
		fmt.Println("不是结构体")
		return
	}
	num := val.NumField()
	fmt.Println("结构体有",num,"个字段")
	for i := 0; i < num; i++ {
		//val.Field获取到变量里的值，但是不可以参与运算，需要转为对应的类型
		fmt.Printf("第%d个字段，值=%v \n",i,val.Field(i))
		//不可以使用val，因为val返回的是变量里的值，type返回的是StructField结构体
		//里面Tag字段是StructTag结构体类型，里面的Get方法可以获取到变量里的值返回为string类型
		tagVal :=typ.Field(i).Tag.Get("json")
		if tagVal != ""{
			fmt.Printf("第%d个字段，标签=%v \n",i,tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Println("结构体有",numOfMethod,"个方法")
	//获取到第二个方法并执行，排序默认是按照函数名排序
	//nil代表没有参数传入
	val.Method(1).Call(nil)
	//call中传入参数是reflect.Value切片
	//调用结构体的第1个方法
	var params []reflect.Value
	//将变量转为reflect.Value类型进行操作
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	//返回结果是[]reflect.Value
	fmt.Println("res=",res[0].Int())
	/* 	结构体有 4 个字段
		第0个字段，值=张三 
		第0个字段，标签=name        
		第1个字段，值=18 
		第1个字段，标签=monster_age 
		第2个字段，值=66.6 
		第3个字段，值=女 
		结构体有 3 个方法
		开始
		{张三 18 66.6 女}
		结束
		res= 50 */
}
func main() {
	a := Monster{"张三",18,66.6,"女"}
	//在使用时，创建一个实例后交给函数，不需要做其他工作
	//在函数内部通过反射调用各个方法
	TestStruct(a)
	
}