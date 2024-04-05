package main

import "fmt"

//责任链模式

// Handler 接口定义了处理请求的方法和设置后继者的方法
type Handler interface {
	HandleRequest(request int)
	SetSuccessor(successor Handler)
}

// EasyHandler 是具体处理者，处理难度级别低的请求
type EasyHandler struct {
	successor Handler
}

func (e *EasyHandler) HandleRequest(request int) {
	if request <= 10 {
		fmt.Printf("EasyHandler: Handling request %d\n", request)
	} else if e.successor != nil {
		e.successor.HandleRequest(request)
	}
}

func (e *EasyHandler) SetSuccessor(successor Handler) {
	e.successor = successor
}

// MediumHandler 是具体处理者，处理难度级别中等的请求
type MediumHandler struct {
	successor Handler
}

func (m *MediumHandler) HandleRequest(request int) {
	if request > 10 && request <= 50 {
		fmt.Printf("MediumHandler: Handling request %d\n", request)
	} else if m.successor != nil {
		m.successor.HandleRequest(request)
	}
}

func (m *MediumHandler) SetSuccessor(successor Handler) {
	m.successor = successor
}

// HardHandler 是具体处理者，处理难度级别高的请求
type HardHandler struct{}

func (h *HardHandler) HandleRequest(request int) {
	if request > 50 {
		fmt.Printf("HardHandler: Handling request %d\n", request)
	}
}

func (h *HardHandler) SetSuccessor(successor Handler) {
	// 最后一个处理者，不设置后继者
}

func main() {
	// 创建处理者链
	hardHandler := &HardHandler{}
	mediumHandler := &MediumHandler{}
	easyHandler := &EasyHandler{}

	// 设置处理者的后继者
	easyHandler.SetSuccessor(mediumHandler)
	mediumHandler.SetSuccessor(hardHandler)

	// 发送不同难度级别的请求
	requests := []int{5, 15, 60}
	for _, request := range requests {
		easyHandler.HandleRequest(request)
	}
}
