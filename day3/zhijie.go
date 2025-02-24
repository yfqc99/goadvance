package main
import "fmt"

type Goods struct {
	Name string
	price float64
}
type Brand struct {
	Name string
	Address string
}
type TV struct {
	Goods
	Brand
}
type Latiao struct {
	*Goods
	*Brand
}
type Test struct{
	int
	n int
}
func main() {
	var t Test
	t.int = 66
	t.n = 99
	fmt.Println(t)
	// {66 99}
	tv :=TV{ 
		Goods{
			Name : "电视机",
			price : 9999.9,
		},
		Brand{"美的","大北极"},
	}
	fmt.Println(tv)
	// {{电视机 9999.9} {美的 大北极}}
	latiao := Latiao{
		&Goods{
			Name : "卫龙超级大辣条特惠仅需",
			price : 6.66,
		},
		&Brand{"卫龙","卫龙总部"},
	}
	fmt.Println(latiao)
	// {0xc000008078 0xc000026400}
	fmt.Println(*latiao.Goods," ",*latiao.Brand)
	// {卫龙超级大辣条特惠仅需 6.66}   {卫龙 卫龙总部}
}