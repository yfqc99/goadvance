package mod
import "fmt"
//将属性首字母小写
type person struct{
	Name string
	age int
	sal float64
}
//提供工程模式的函数
func NewPerson(name string) *person {
	return &person{
		Name : name,
		//因为age和sal不是对外公开的所以不能直接访问
		//如果在此处赋值可以通过指针直接查看到对应的值
	}
}
//首字母大写的Set和Get方法，可以加入数据验证的业务逻辑
func (p *person) SetAge(age int)  {
	if age >18 && age <150{
		p.age = age
		return
	}
	fmt.Println("年龄不正确")
}
func (p *person) GetAge() int  {
	return p.age
}
func (p *person) SetSal(sal float64)  {
	if sal >=3000 && sal <= 30000{
		p.sal = sal
		return
	}
	fmt.Println("薪资不正确")
}
func (p *person) GetSal() float64 {
	return p.sal
}