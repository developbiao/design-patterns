package main

import "fmt"

// Television
type TV struct{}

func (t *TV) On() {
	fmt.Println("TV is on")
}

func (t *TV) Off() {
	fmt.Println("TV is power off")
}

// Voice box
type VoiceBox struct{}

func (v *VoiceBox) On() {
	fmt.Println("Voice box is on")
}

func (v *VoiceBox) Off() {
	fmt.Println("Voice box is power off")
}

// Light
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	fmt.Println("Light is power off")
}

// MicroPhone struct
type MicroPhone struct{}

func (m *MicroPhone) On() {
	fmt.Println("MicroPhone is On")
}

func (m *MicroPhone) Off() {
	fmt.Println("MicroPhone is power off")
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is on")
}

func (p *Projector) Off() {
	fmt.Println("Projector is off ")
}

type Xbox struct{}

func (x *Xbox) On() {
	fmt.Println("Xbox is on")
}

func (x *Xbox) Off() {
	fmt.Println("Xbox is power off")
}

// Home Player facade
type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	light Light
	xbox  Xbox
	pro   Projector
	mp    MicroPhone
}

func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("Home Player enter KTV mode")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.vb.On()
	hp.light.Off()
}

func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("Home Player enter game mode")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}

func (hp *HomePlayerFacade) DoMovie() {
	fmt.Println("Home Player enter movie mode")
	hp.tv.Off()
	hp.xbox.On()
	hp.pro.On()
	hp.light.Off()
}

func main() {
	hp := new(HomePlayerFacade)
	hp.DoGame()
	hp.DoKTV()
	hp.DoMovie()
}
