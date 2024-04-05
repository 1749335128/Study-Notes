package main

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}

type Shopping interface {
	Buy(goods *Goods)
}

type AmericanShopping struct {
}

func (a *AmericanShopping) Buy(goods *Goods) {
	if goods == nil {
		fmt.Printf("去美国购物,什么也没买到。\n")
		return
	}
	fmt.Printf("去美国购物,买了%s\n", goods.Kind)
}

type KoreaShopping struct {
}

func (k *KoreaShopping) Buy(goods *Goods) {
	if goods == nil {
		fmt.Printf("去韩国购物,什么也没买到。\n")
		return
	}
	fmt.Printf("去韩国购物,买了%s\n", goods.Kind)
}

//海外代理

type OverSeaProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) *OverSeaProxy {
	return &OverSeaProxy{
		shopping: shopping,
	}
}

func (o *OverSeaProxy) Buy(goods *Goods) {
	if o.differentiate(goods) {
		o.shopping.Buy(goods)
		o.bringBack(goods)
	}
}

func (o *OverSeaProxy) differentiate(goods *Goods) bool {
	if goods.Fact {
		fmt.Printf("对%s物品进行辨别,结果为真。\n", goods.Kind)
	} else {
		fmt.Printf("对%s物品进行辨别,结果为假,取消购买。\n", goods.Kind)
	}
	return goods.Fact
}
func (o *OverSeaProxy) bringBack(goods *Goods) {
	fmt.Printf("将%s物品进行海关安检，并带回国内...\n", goods.Kind)

}

func main() {
	var g1 = Goods{
		Kind: "爱马仕包包",
		Fact: true,
	}
	//proxy := NewProxy(&AmericanShopping{})
	proxy := NewProxy(&KoreaShopping{})
	proxy.Buy(&g1)
}
