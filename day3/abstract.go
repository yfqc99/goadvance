package main
import "fmt"

type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}
func (a *Account) Deposite(money float64,pwd string)  {
	if pwd != a.Pwd{
		fmt.Println("密码不正确")
		return
	}
	if money <= 0{
		fmt.Println("金额不正确")
		return
	}
	a.Balance += money
	fmt.Println("存款成功")
}
func (a *Account) WithDraw(money float64,pwd string)  {
	if pwd != a.Pwd{
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 || money > a.Balance{
		fmt.Println("金额不正确")
		return
	}
	a.Balance -= money
	fmt.Println("取款成功")
}
func (a *Account) Query(pwd string)  {
	if pwd != a.Pwd{
		fmt.Println("密码不正确")
		return
	}
	fmt.Println("余额：",a.Balance)
}
func main() {
	a := Account{
		AccountNo : "666666",
		Pwd : "123456",
		Balance : 99.9,
	}
	a.Deposite(900,a.Pwd)
	a.WithDraw(333.3 ,a.Pwd)
	a.Query(a.Pwd)
	/* 	存款成功
		取款成功
		余额： 666.5999999999999 */
}