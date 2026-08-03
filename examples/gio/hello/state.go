package main

//职责：只保存数据。
//不允许出现 Gio import。

type CounterState struct {
	Count int
}

func (s *CounterState) Increment() {
	s.Count++
}

func (s *CounterState) Reset() {
	s.Count = 0
}
