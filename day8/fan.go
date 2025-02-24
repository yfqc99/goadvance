package main
import (
	"fmt"
	"encoding/json"
)
type Monster struct{
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}
func unmarshalStruct()  {
	str :="{\"Name\":\"张三\",\"Age\":18,\"Birthday\":\"2006-06-06\",\"Sal\":6666.66,\"Skill\":\"挣钱\"}"
	var monster Monster
	err :=json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Println("反序列化错误：",err)
	}
	fmt.Println("struct反序列化：",monster)
}

func unmarshalMap()  {
	str :="{\"address\":[\"猫\",\"极道\"],\"age\":20,\"name\":\"周六\"}"
	var a map[string]interface{}
	//在反序列化map时，底层自动make，封装到Unmarshal函数
	err :=json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Println("反序列化错误：",err)
	}
	fmt.Println("map反序列化：",a)
}

func unmarshalSlice()  {
	str :="[{\"address\":\"猫\",\"age\":30,\"name\":\"李四\"},"+
	"{\"address\":\"猫\",\"age\":35,\"name\":\"钱五\"}]"
	var slice []map[string]interface{}
	err :=json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Println("反序列化错误：",err)
	}
	fmt.Println("slice反序列化：",slice)
}
func main() {
	unmarshalStruct()
	// struct反序列化： {张三 18 2006-06-06 6666.66 挣钱}
	unmarshalMap()
	// map反序列化： map[address:[猫 极道] age:20 name:周六]
	unmarshalSlice()
	/* 	slice反序列化： [map[address:猫 age:30 name:李四] 
		map[address:猫 age:35 name:钱五]] */
}