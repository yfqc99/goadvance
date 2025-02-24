package main
import (
    "fmt"
    "reflect"
)
func test(b interface{})  {
    // 通过反射获得变量的type，kind和值
    //Type反射类型
    rTyp := reflect.TypeOf(b)
    fmt.Printf("rTyp的类型：%T，rTyp的值：%v \n",rTyp,rTyp)
    // rTyp的类型：*reflect.rtype，rTyp的值：int
    rVal := reflect.ValueOf(b)
    fmt.Printf("rVal的类型：%T，rVal的值：%v \n",rVal,rVal)
   /*  rVal的类型：reflect.Value，rVal的值：100 */
   /* cannot convert 2 (untyped int constant) to type struct{typ *reflect.rtype;
     ptr unsafe.Pointer; reflect.flag} */
    //n2 := 2+rVal
    //fmt.Println(n2)
    //返回值的有符号整数
    n2 := 2+rVal.Int()
    fmt.Println(n2)
    // panic: reflect: call of reflect.Value.Float on int Value
    /* n3 := rVal.Float()
    fmt.Println(n3) */
    // 102
    //interface{}
    iV :=rVal.Interface()
    //转成对应类型
    v := iV.(int)
    fmt.Printf("v的类型：%T，v的值：%v \n",v,v)
    // v的类型：int，v的值：100
}
type Student struct {
    Name string
    Age int
}
func test2(b interface{})  {
    rTyp := reflect.TypeOf(b)
    rVal := reflect.ValueOf(b)
    fmt.Println("rTyp=",rTyp," rVal=",rVal)
    //获取变量对应的kind
    fmt.Println("rVal的kind=",rVal.Kind()," rTyp的kind=",rTyp.Kind())
    //kind的范围比type大
    // rVal的kind= struct  rTyp的kind= struct
    iV :=rVal.Interface()
    fmt.Printf("iV的类型：%T，iV的值：%v \n",iV,iV)
    /*  rTyp= main.Student  rVal= {张三 18}
        iV的类型：main.Student，iV的值：{张三 18} */
    //虽然类型和预期一样但是无法取出数据
    // fmt.Println(iV.Name)
    // iV.Name undefined (type any has no field or method Name)
    //因为反射的本质是在运行时的，在运行时可以确定类型和值
    //但是在编译时无法确定所以会报错
    v := iV.(Student)
    fmt.Println(v.Name)
    // 张三
}
func main() {
    var num int = 100
    test(num)
    stu :=Student {"张三",18}
    test2(stu)
}