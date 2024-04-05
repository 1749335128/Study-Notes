package main

import "fmt"

//抽象工厂设计模式
//产品族苹果,香蕉,葡萄(产品族最好固定,后期也不添加)。
//可以有不同产品等级(中国,日本,马来西亚等)
//符合开闭原则

type FruitFactory interface {
	BuyApple() AbstractApple
	BuyBanana() AbstractBanana
	BuyGrape() AbstractGrape
}
type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}

type AbstractGrape interface {
	ShowGrape()
}

// 实现层

type ChinaApple struct{}

func (c *ChinaApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct{}

func (c *ChinaBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinaGrape struct{}

func (c *ChinaGrape) ShowGrape() {
	fmt.Println("中国葡萄")
}

type JapanApple struct{}

func (c *JapanApple) ShowApple() {
	fmt.Println("日本苹果")
}

type JapanBanana struct{}

func (c *JapanBanana) ShowBanana() {
	fmt.Println("日本香蕉")
}

type JapanGrape struct{}

func (c *JapanGrape) ShowGrape() {
	fmt.Println("日本葡萄")
}

// 中国工厂

type ChinaFactory struct{}

func (c ChinaFactory) BuyApple() AbstractApple {
	var abstractApple AbstractApple
	abstractApple = new(ChinaApple)
	return abstractApple
}
func (c ChinaFactory) BuyGrape() AbstractGrape {
	return &ChinaGrape{}
}
func (c ChinaFactory) BuyBanana() AbstractBanana {
	return &ChinaBanana{}
}

// 日本工厂

type JapanFactory struct{}

func (c JapanFactory) BuyApple() AbstractApple {
	var abstractApple AbstractApple
	abstractApple = new(JapanApple)
	return abstractApple
}
func (c JapanFactory) BuyGrape() AbstractGrape {
	return &JapanGrape{}
}
func (c JapanFactory) BuyBanana() AbstractBanana {
	return &JapanBanana{}
}

func main() {
	var factory FruitFactory
	factory = new(ChinaFactory)
	factory.BuyGrape().ShowGrape()
	factory.BuyBanana().ShowBanana()
	factory.BuyGrape().ShowGrape()
	factory = &JapanFactory{}
	factory.BuyGrape().ShowGrape()
	factory.BuyBanana().ShowBanana()
	factory.BuyGrape().ShowGrape()
}
