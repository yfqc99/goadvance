package main
import (
	"fmt"
	"encoding/json"
)
type Monster struct{
	//指定json字符串的key属性名称，利用反射机制
	Name string `json:"monster_name"`
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}
func testStruct()  {
	monster := Monster{
		Name : "张三",
		Age : 18,
		Birthday : "2006-06-06",
		Sal : 6666.66,
		Skill : "挣钱",
	}
	//序列化
	// Marshal的参数是空接口，所以任何类型都可以
	//[...]int是数组类型，...是一个语法糖，让编译器根据初始化时提供的元素个数来确定数组长度
	// 返回byte类型切片
	data,err :=json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误",err)
	}
	fmt.Println("struct序列化：",string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "周六"
	a["age"] = 20
	a["address"] = [2]string{"猫","极道"}
	//make是指针类型
	data,err :=json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误",err)
	}
	fmt.Println("map序列化：",string(data))
}

func testSlice()  {
	var slice []map[string]interface{}
	var m1  map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "李四"
	m1["age"] = 30
	m1["address"] = "猫"
	slice = append(slice,m1)
	var m2  map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "钱五"
	m2["age"] = 35
	m2["address"] = "猫"
	slice = append(slice,m2)
	data,err :=json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误",err)
	}
	fmt.Println("slice序列化：",string(data))
}

func testFloat64()  {
	var num float64 = 234.56
	//也可以传入非指针类型
	data,err :=json.Marshal(num)
	if err != nil {
		fmt.Println("序列化错误",err)
	}
	fmt.Println("float64序列化：",string(data))
}
func main() {
	//序列化后的顺序可能是不一样的
	//map 结构体和切片序列化
	testStruct()
	/* 	{"Name":"张三","Age":18,"Birthday":"2006-06-06"
		,"Sal":6666.66,"Skill":"挣钱"} */
	/* 	struct序列化： {"monster_name":"张三","monster_age":18
		,"Birthday":"2006-06-06","Sal":6666.66,"Skill":"挣钱"} */
	testMap()
	// map序列化： {"address":["猫","极道"],"age":20,"name":"周六"}
	testSlice()
	/* 	slice序列化： [{"address":"猫","age":30,"name":"李四"}
		,{"address":"猫","age":35,"name":"钱五"}] */
	testFloat64()
	//相当于转为了字符串
	// float64序列化： 234.56
}