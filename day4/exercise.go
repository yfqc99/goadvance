package main
import (
	"fmt"
	"sort"
	"math/rand"
)
type Hero struct{
	Name string
	Age int
}
type HeroSlice []Hero
//实现interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
//Less排序方式
func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Age < hs[j].Age
}
//交换
func (hs HeroSlice) Swap(i,j int) {
	/* 	temp := hs[i]
		hs[i] = hs[j]
		hs[j] = temp */
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	//定义数组或者切片
	var slice = []int{0,-1,10,7,90}
	//排序
	//因为slice是引用类型，不需要再加&
	sort.Ints(slice)
	fmt.Println(slice)
	// [-1 0 7 10 90]
	//对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			// rand随机数
			Name : fmt.Sprintf("英雄 %d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heroes = append(heroes,hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	/* 	{英雄 83 69}
		{英雄 84 82}
		{英雄 65 55}
		{英雄 99 8}
		{英雄 41 57}
		{英雄 41 91}
		{英雄 28 37}
		{英雄 80 47}
		{英雄 5 56}
		{英雄 73 13} */
	sort.Sort(heroes)
	fmt.Println("排序后")
	for _, v := range heroes {
		fmt.Println(v)
	}	
/* 	排序后
	{英雄 99 8}
	{英雄 73 13}
	{英雄 28 37}
	{英雄 80 47}
	{英雄 65 55}
	{英雄 5 56}
	{英雄 41 57}
	{英雄 83 69}
	{英雄 84 82}
	{英雄 41 91} */
}