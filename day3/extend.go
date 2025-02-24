package main
import "fmt"
//利用继承解决代码冗余
type Student struct{
	Name string
	age int
	Score int
}
func (s *Student) showInfo()  {
	fmt.Println("姓名：",s.Name,"年龄：",s.age,"成绩：",s.Score)
}
func (s *Student) SetScore(score int){
	s.Score = score
}
type Pupil struct {
	Student
}
func (p *Pupil) testing1()  {
	fmt.Println("小学生正在考试")
}
type Grudate struct {
	//没有变量名
	Student
}
func (g *Grudate) testing2()  {
	fmt.Println("大学生正在考试")
}
func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "周六"
	pupil.Student.age = 6 
	pupil.testing1()
	pupil.Student.SetScore(66)
	pupil.Student.showInfo()
	/* 	小学生正在考试
		姓名： 周六 年龄： 6 成绩： 66 */
	grudate := &Grudate{}
	//编译器在当前结构体找不到对应的属性和方法时
	//会自动到嵌入的匿名结构体中找
	grudate.Name = "王五"
	grudate.age = 22 
	grudate.testing2()
	grudate.SetScore(222)
	grudate.showInfo()
	/* 	大学生正在考试
		姓名： 王五 年龄： 22 成绩： 222 */
}