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
	val := reflect.ValueOf(a)
	kd := val.Kind()
	//获取到的参数类型是指针
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct{
		fmt.Println("不是结构体")
		return
	}
	fmt.Printf("elem()的返回值：%v,elem()的类型：%T\n",val.Elem(),val.Elem())
	val.Elem().Field(0).SetString("周六")
	val.Elem().Method(1).Call(nil)
}
func main() {
	a := Monster{"张三",18,66.6,"女"}
	//传入地址
	TestStruct(&a)
	fmt.Println(a)
	/* 	开始
		{周六 18 66.6 女}
		结束
		{周六 18 66.6 女} */
}