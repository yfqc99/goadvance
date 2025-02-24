package main
import "fmt"

type Teacher struct{
	Name string
}
type Student struct{
	Name string
}
//将该方法和Teacher类型进行绑定
func (t Teacher) teach(){
	t.Name = "周六"
	//因为是值传递，在方法内对数据进行修改
	//不会影响到方法外数据
	fmt.Println(t.Name,"老师正在教课")
	// 周六 老师正在教课
}

func (te Teacher) add(num int) int{
	res := 0
	for i := 1; i <= num; i++ {
		res += i
	}
	return res
}

func main() {
	var t Teacher
	t.Name = "王五"
	fmt.Print(t.Name,"老师被调课为")
	fmt.Println(t.Name,"进行计算为",t.add(99))
	// 王五 进行计算为 4950
	//t在调用方法时会传到方法
	t.teach()
	fmt.Println(t.Name,"老师正在教课")
	// 王五 老师正在教课

	//该方法只能通过Teacher类型进行调用
	/* s.teach undefined (type Student 
		has no field or method teach) */
	/* var s Student
	s.Name = "麻子"
	s.teach() */
	// undefined: teach
	// teach()

	
}