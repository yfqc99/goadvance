package main
import "fmt"
func TypeJudge(items... interface{})  {
	for _, v := range items {
		switch v.(type) {
		case bool :
			fmt.Println("bool")
		case float32 :
			fmt.Println("float32")
		case float64 :
			fmt.Println("float64")
		case int, int32, int64 :
			fmt.Println("int系列")
		case string :
			fmt.Println("string")
		case Student :
			fmt.Println("Student")
		case *Student :
			fmt.Println("*Student")
		default :
			fmt.Println("不确定")
		}
	}
}
func main() {

}