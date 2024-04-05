package main

import "fmt"

// 适配器模式
// (示例:手机需要充电使用5v电,但是家里只有220v的电,所以需要电源适配器来给手机充电,或者其他适配器)

type V5 interface {
	Use5V()
}

type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("使用220V电源")
}

type Phone struct {
	v5 V5
}

// PowerAdapter 电源适配器
type PowerAdapter struct {
	v220 *V220
}

func (a *PowerAdapter) Use5V() {
	a.v220.Use220V()
	fmt.Println("将220v电源进行高频降压至5v")
}
func NewPowerAdapter(v220 *V220) *PowerAdapter {
	return &PowerAdapter{v220: v220}
}

func (p Phone) Charge() {
	p.v5.Use5V()
	fmt.Println("正在给手机充电")
}

func NewPhone(v5 V5) *Phone {
	return &Phone{v5: v5}
}

func main() {
	phone := NewPhone(NewPowerAdapter(new(V220)))
	phone.Charge()
}
