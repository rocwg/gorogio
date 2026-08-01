package main

//职责：只保存数据。
//不允许出现 Gio import。
//这是非常重要的边界。

type CounterState struct {
	Count int
}

func (s *CounterState) Increment() {

	s.Count++
}
