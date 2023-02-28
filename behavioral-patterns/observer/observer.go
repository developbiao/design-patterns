package main

import "fmt"

// -----  Abstract Layout ------

// Abstract Observer
type Listener interface {
	OnTeacherComming()
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// ------- Implementation Layout --------
// Observer student
type StuZhangSan struct {
	Badthing string
}

func (s *StuZhangSan) OnTeacherComming() {
	fmt.Println("ZhangSan stopping do ", s.Badthing)
}

func (s *StuZhangSan) DoBadthing() {
	fmt.Println("ZhangSan doing ", s.Badthing)
}

type StuLiSi struct {
	Badthing string
}

func (s *StuLiSi) OnTeacherComming() {
	fmt.Println("LiSi stopping do ", s.Badthing)
}

func (s *StuLiSi) DoBadthing() {
	fmt.Println("LiSi doing ", s.Badthing)
}

type StuWangWu struct {
	Badthing string
}

func (s *StuWangWu) OnTeacherComming() {
	fmt.Println("WangWu stopping do ", s.Badthing)
}

func (s StuWangWu) DoBadthing() {
	fmt.Println("WangWu doing ", s.Badthing)
}

// Notifier monitor
type ClassMonitor struct {
	listenerList []Listener // Listener collections
}

func (m *ClassMonitor) AddListener(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}

func (m *ClassMonitor) RemoveListener(listener Listener) {
	for index, oneListener := range m.listenerList {
		if listener == oneListener {
			// Link the elments before and after the deleted point
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
			break
		}
	}
}

func (m *ClassMonitor) Notify() {
	for _, listener := range m.listenerList {
		// Each call all observer
		listener.OnTeacherComming()
	}
}

func main() {
	s1 := &StuZhangSan{Badthing: "Read comics"}
	s2 := &StuLiSi{Badthing: "Playing Game"}
	s3 := &StuWangWu{Badthing: "Fall in love"}

	classMonitor := new(ClassMonitor)

	fmt.Println("There is class, but the teacher is not here, and the students are busy with their own game")
	s1.DoBadthing()
	s2.DoBadthing()
	s3.DoBadthing()

	classMonitor.AddListener(s1)
	classMonitor.AddListener(s2)
	classMonitor.AddListener(s3)

	fmt.Println("Teacher is comming, class monitor gave xue he a look ")
	// LiSi continue paying game
	classMonitor.RemoveListener(s2)
	classMonitor.Notify()

	fmt.Println("Main func done!")
}
