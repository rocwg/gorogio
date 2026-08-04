package main

type CounterState struct {
	Count int
}

func (s *CounterState) Increment() {
	s.Count++
}

func (s *CounterState) Reset() {
	s.Count = 0
}
