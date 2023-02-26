package main

import "fmt"

type BeautyWoman interface {
	MakeEyesWithMan()
	HappyWithMan()
}

// Subject
type PanJinLian struct{}

func (p *PanJinLian) MakeEyesWithMan() {
	fmt.Println("Pan jin lian make eyes for me")
}

func (p *PanJinLian) HappyWithMan() {
	fmt.Println("Pan jin lian make romantic date tonight")
}

// Proxy
type WangPo struct {
	woman BeautyWoman
}

func NewProxy(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman}
}

func (p *WangPo) MakeEyesWithMan() {
	p.woman.MakeEyesWithMan()
}

func (p *WangPo) HappyWithMan() {
	p.woman.HappyWithMan()
}

func main() {
	// 西门庆想找点乐子
	wangPo := NewProxy(new(PanJinLian))
	// 王波命令潘金莲做点事
	wangPo.MakeEyesWithMan()
	wangPo.HappyWithMan()
}
