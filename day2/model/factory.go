package model
type Student struct{
	Name string
	Score float64
}
//相当于私有
type student struct{
	Name string
	// Score float64
	score float64
}
//用工场模式代理生成实例
func NewStudent(name string,score float64) *student{
	return &student{
		Name : name,
		// Score : score,
		score : score,
	}
}

// fmt.Println("姓名：",stu1.Name,"，成绩：",stu1.score)报错
//当结构体中的属性名首字母小写，不可以在其他包直接访问
//解决
func (stu *student) GetScore() float64 {
	return stu.score
}