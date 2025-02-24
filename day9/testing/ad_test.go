package test
import (
	_"fmt"
	//引入go的testing框架包
	"testing"
)
//编写测试用例
//Test后的第一个首字母必须大写
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55{
		// fmt.Println("期望值：",55,"实值：",res)
		//输出日志的同时停止程序
		/* 	(*testing.common).Fatalf call has 
			arguments but no formatting directives */
		t.Fatalf("期望值：%v，实值：%v",55,res)
	}
	//正确输出日志
	t.Logf("执行正确")
	// 执行正确
}