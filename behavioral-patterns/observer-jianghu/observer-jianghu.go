package main

import (
	"fmt"
)

const (
	PGaiBang  string = "丐帮"
	PMingJiao string = "明教"
)

// ------- Abstract Layout --------
type Listener interface {
	// 当同派同伴被揍了怎么办?
	OnFriendBeFight(event *Event)
	GetName() string
	GetParty() string
	Title() string
}

type Notifier interface {
	// Add listener
	AddListener(listener Listener)
	// Remove listener
	RemoveListener(listener Listener)
	// Notify
	Notify(event *Event)
}

// Event
type Event struct {
	Noti    Notifier // 被知晓的通知者
	One     Listener // 事件主动发出者
	Another Listener // 被动接收者
	Msg     string   // 具体消息
}

// -------- Implementation Layout --------
type Hero struct {
	Name  string
	Party string
}

func (hero *Hero) Fight(another Listener, baixiao Notifier) {
	msg := fmt.Sprintf("%s 将 %s 揍了...", hero.Title(), another.Title())

	// Genereate event
	event := new(Event)
	event.Noti = baixiao
	event.One = hero
	event.Another = another
	event.Msg = msg

	// 广播事件
	baixiao.Notify(event)
}

func (hero *Hero) Title() string {
	return fmt.Sprintf("[%s]:%s", hero.Party, hero.Name)
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) GetParty() string {
	return hero.Party
}

func (hero *Hero) OnFriendBeFight(event *Event) {
	// 判断是否是当事人
	if hero.Name == event.One.GetName() || hero.Name == event.Another.GetName() {
		return
	}

	// 本帮派同伴将其它门派的揍了，要拍手叫好!
	if hero.Party == event.One.GetParty() {
		fmt.Println(hero.Title(), "得知消息, 拍手叫好!!!")
		return
	}

	// 本帮派被其它门派揍了，要主动报仇反击！
	if hero.Party == event.Another.GetParty() {
		fmt.Println(hero.Title(), " 得知消息，发起报仇反击!!!")
		hero.Fight(event.One, event.Noti)
		return
	}
}

// 百晓生(NOtifier)
type BaiXiao struct {
	heroList []Listener
}

// Add listener
func (b *BaiXiao) AddListener(listener Listener) {
	b.heroList = append(b.heroList, listener)
}

// Remove listener
func (b *BaiXiao) RemoveListener(listener Listener) {
	for index, oneListener := range b.heroList {
		if oneListener == listener {
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
			break
		}
	}
}

func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("【江湖消息】百晓生广播消息: ", event.Msg)
	for _, listener := range b.heroList {
		listener.OnFriendBeFight(event)
	}
}

func main() {
	hero1 := Hero{
		"黄蓉",
		PGaiBang,
	}

	hero2 := Hero{
		"洪七公",
		PGaiBang,
	}

	hero3 := Hero{
		"乔峰",
		PGaiBang,
	}

	hero4 := Hero{
		"张无忌",
		PMingJiao,
	}

	hero5 := Hero{
		"韦一笑",
		PMingJiao,
	}

	hero6 := Hero{
		"金毛狮王",
		PMingJiao,
	}

	baixiao := BaiXiao{}

	baixiao.AddListener(&hero1)
	baixiao.AddListener(&hero2)
	baixiao.AddListener(&hero3)
	baixiao.AddListener(&hero4)
	baixiao.AddListener(&hero5)
	baixiao.AddListener(&hero6)

	fmt.Println("武林江湖一片平静.....")
	hero1.Fight(&hero4, &baixiao)
	fmt.Println("Main func done!")
}
