package main

import (
	"time"
)

// TimerState 是煮蛋计时器的业务状态。
//
// 它只在 Gio 的窗口事件循环中读写；进度通过当前时间推导，
// 因而不需要额外的 goroutine 或锁。
type TimerState struct {
	duration  time.Duration // 本次计时的总时长。
	elapsed   time.Duration // 暂停前已经走过的时长。
	startedAt time.Time     // 最近一次开始或继续计时的时刻。
	running   bool          // 当前是否正在计时。
}

// TimerPhase 是界面需要知道的计时器阶段。
// TimerState 内部如何用 elapsed、startedAt 和 running 表示这些阶段，
// 不再需要泄漏给界面代码。
type TimerPhase uint8

const (
	TimerIdle     TimerPhase = iota // 尚未开始，或已被重置。
	TimerRunning                    // 正在计时。
	TimerPaused                     // 已暂停，可继续计时。
	TimerFinished                   // 已到达设定时长。
)

// TimerViewState 是某一帧交给 UI 渲染的只读快照。
//
// 它只包含界面当前需要的事实，而不暴露 TimerState 的内部计时细节。
// 这对应 Compose 中由状态层提供给界面的 UiState。
type TimerViewState struct {
	Phase    TimerPhase
	Progress float32
}

// View 从业务状态推导当前帧的 UI 快照。
func (t *TimerState) View(now time.Time) TimerViewState {
	return TimerViewState{
		Phase:    t.Phase(now),
		Progress: t.Progress(now),
	}
}

// Phase 根据当前时刻推导计时器阶段。
func (t *TimerState) Phase(now time.Time) TimerPhase {
	switch {
	case t.Finished(now):
		return TimerFinished
	case t.running:
		return TimerRunning
	case t.elapsed > 0:
		return TimerPaused
	default:
		return TimerIdle
	}
}

// Elapsed 返回 now 时刻实际已经走过的时长，并将其限制为总时长。
func (t *TimerState) Elapsed(now time.Time) time.Duration {
	elapsed := t.elapsed
	if t.running {
		elapsed += now.Sub(t.startedAt)
	}
	if elapsed > t.duration {
		return t.duration
	}
	return elapsed
}

// Progress 把经过时长换算为 0 到 1 的进度值，供进度条和鸡蛋颜色使用。
func (t *TimerState) Progress(now time.Time) float32 {
	if t.duration <= 0 {
		return 0
	}
	return float32(float64(t.Elapsed(now)) / float64(t.duration))
}

// Finished 表示计时器是否已经到达设定的总时长。
func (t *TimerState) Finished(now time.Time) bool {
	return t.duration > 0 && t.Elapsed(now) >= t.duration
}

// Start 按给定时长开始一轮全新的计时。
func (t *TimerState) Start(now time.Time, duration time.Duration) {
	if duration <= 0 {
		return
	}
	t.duration = duration
	t.elapsed = 0
	t.startedAt = now
	t.running = true
}

// Resume 从暂停的位置继续计时。
func (t *TimerState) Resume(now time.Time) {
	if !t.running && !t.Finished(now) && t.elapsed > 0 {
		t.startedAt = now
		t.running = true
	}
}

// Pause 记录当前已走过的时间，并停止计时。
func (t *TimerState) Pause(now time.Time) {
	if !t.running {
		return
	}
	t.elapsed = t.Elapsed(now)
	t.running = false
}

// Reset 清空当前计时进度。输入框的内容属于 TimerScreenUI，因此不会被重置。
func (t *TimerState) Reset() {
	t.duration = 0
	t.elapsed = 0
	t.startedAt = time.Time{}
	t.running = false
}
