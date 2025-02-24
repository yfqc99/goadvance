package main
import (
	"fmt"
	"Ajinjie/day2/model"
	/* package Ajinjie/day2/model is not in 
	GOROOT (G:\go\src\Ajinjie\day2\model) */
	// 出现上述报错时
	//回到项目目录下执行以下代码
	//或者直接把项目拖到指定的位置（非常不建议）
	/*	初始化模块 go mod init
		下载依赖包 go mod tidy */
)
func main() {
	var stu = model.Student{
		Name : "张三",
		Score : 66.6,
	}
	fmt.Println(stu)
	// {张三 66.6}
	stu1 := model.NewStudent("周六",99.9)
	fmt.Println(*stu1)
	// {周六 99.9}

	/* stu1.score undefined (type *model.student 
		has no field or method score) */
		// fmt.Println("姓名：",stu1.Name,"，成绩：",stu1.score)
	fmt.Println("姓名：",stu1.Name,"，成绩：",stu1.GetScore())
	// 姓名： 周六 ，成绩： 99.9
}