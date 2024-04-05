package main

import "fmt"

// facade 装饰模式

type FamilyCinema struct {
	tv        *TV
	lamp      *Lamp
	micro     *MicroPhone
	projector *Projector
	sound     *Sound
}

// TvMode 电影模式
func (f *FamilyCinema) TvMode() {
	fmt.Println("开启电影模式")
	f.tv.ON()
	f.lamp.OFF()
	f.projector.ON()
}

// KtvMode ktv模式
func (f *FamilyCinema) KtvMode() {
	fmt.Println("ktv模式")
	f.tv.ON()
	f.sound.ON()
	f.projector.OFF()
}

// OffAll 断电
func (f *FamilyCinema) OffAll() {
	fmt.Println("开启离家模式")
	f.tv.OFF()
	f.micro.OFF()
	f.sound.ON()
	f.projector.OFF()
	f.lamp.OFF()
}

type Device interface {
	ON()
	OFF()
}

type TV struct{}

func (t *TV) ON() {
	fmt.Println("打开电视")
}
func (t *TV) OFF() {
	fmt.Println("关闭电视")
}

type Lamp struct{}

func (t *Lamp) ON() {
	fmt.Println("打开灯")
}
func (t *Lamp) OFF() {
	fmt.Println("关闭灯")
}

type MicroPhone struct{}

func (t *MicroPhone) ON() {
	fmt.Println("打开麦克风")
}
func (t *MicroPhone) OFF() {
	fmt.Println("关闭麦克风")
}

type Sound struct{}

func (t *Sound) ON() {
	fmt.Println("打开音响")
}
func (t *Sound) OFF() {
	fmt.Println("关闭音响")
}

type Projector struct{}

func (t *Projector) ON() {
	fmt.Println("打开投影仪")
}
func (t *Projector) OFF() {
	fmt.Println("关闭投影仪")
}

func NewFamilyCinema(tv *TV, lamp *Lamp, projector *Projector, sound *Sound, micro *MicroPhone) *FamilyCinema {
	return &FamilyCinema{
		tv:        tv,
		lamp:      lamp,
		projector: projector,
		sound:     sound,
		micro:     micro,
	}
}

func main() {
	c := NewFamilyCinema(new(TV), new(Lamp), new(Projector), new(Sound), new(MicroPhone))
	c.KtvMode()
	fmt.Println("-----")
	c.TvMode()
	fmt.Println("-----")
	c.OffAll()
}
