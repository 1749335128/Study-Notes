package main

import "fmt"

//装饰模式

type Phone interface {
	Show()
}

type Huawei struct{}

func (h *Huawei) Show() {
	fmt.Println("这是华为手机")
}

type XiaoMi struct {
}

func (x *XiaoMi) Show() {
	fmt.Println("这是小米手机")
}

// PhoneDecorator 手机装饰器
type PhoneDecorator struct {
	phone Phone
}

type KeDecorator struct {
	PhoneDecorator
}

func (k *KeDecorator) Show() {
	k.phone.Show()
	//增强方法
	fmt.Println("手机佩戴上了壳")
}

type MoDecorator struct {
	PhoneDecorator
}

func (k *MoDecorator) Show() {
	k.phone.Show()
	//增强方法
	fmt.Println("手机贴上了膜")
}

type OpenDecorator struct {
	PhoneDecorator
}

func (k *OpenDecorator) Show() {
	k.phone.Show()
	//增强方法
	fmt.Println("手机开机了")
}

func NewOpenDecorator(phone Phone) Phone {
	return &OpenDecorator{
		PhoneDecorator{
			phone: phone,
		},
	}
}
func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{
		PhoneDecorator{
			phone: phone,
		},
	}
}
func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{
		PhoneDecorator{
			phone: phone,
		},
	}
}

func main() {
	var phone Phone
	phone = new(XiaoMi)
	//贴膜
	modDecorator := NewMoDecorator(phone)
	modDecorator.Show()
	//壳
	fmt.Println("-----")
	keDecorator := NewKeDecorator(modDecorator)
	keDecorator.Show()
	//开机
	fmt.Println("-----")
	openDecorator := NewOpenDecorator(keDecorator)
	openDecorator.Show()
	//再贴膜
	fmt.Println("-----")
	modDecorator = NewMoDecorator(openDecorator)
	modDecorator.Show()
}
