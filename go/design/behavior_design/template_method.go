package main

import "fmt"

// 模板方法 模式
// 煮咖啡 1:煮开水 2:冲泡咖啡 3:倒入杯中  4:添加牛奶和糖
// 煮茶  1:煮开水  2:冲泡茶叶 3:倒入杯中 4:加柠檬
// 煮泡面,煮鸡蛋...
// 模板  1:煮开水  2:冲泡    3：倒入杯中 4:添加小料

type CookFood interface {
	BoilWatter()
	Brew()
	PourToCup()
	AddThings()
}

type makeTemplate struct {
	c CookFood
}

func (m *makeTemplate) Cook() {
	if m.c == nil {
		return
	}
	m.c.BoilWatter()
	m.c.Brew()
	m.c.PourToCup()
	m.c.AddThings()
}

type MakeCoffee struct {
	makeTemplate
}

func (m *MakeCoffee) BoilWatter() {
	fmt.Println("开始煮开水")
}
func (m *MakeCoffee) Brew() {
	fmt.Println("开始冲泡咖啡")
}
func (m *MakeCoffee) PourToCup() {
	fmt.Println("将冲泡好的咖啡倒入马克杯中")
}
func (m *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func NewMakeCoffeeTemplate() *MakeCoffee {
	m := new(MakeCoffee)
	m.c = m
	return m
}

func main() {
	m := NewMakeCoffeeTemplate()
	m.Cook()
}
