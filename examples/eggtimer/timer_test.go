package main

import (
	"testing"
	"time"
)

func TestTimerReset(t *testing.T) {
	now := time.Date(2026, time.July, 21, 12, 0, 0, 0, time.UTC)
	var timer TimerState
	timer.Start(now, 5*time.Minute)
	timer.Pause(now.Add(time.Minute))

	timer.Reset()

	if timer.running || timer.duration != 0 || timer.elapsed != 0 || !timer.startedAt.IsZero() {
		t.Fatalf("Reset() 后状态未清空: %+v", timer)
	}
	if progress := timer.Progress(now); progress != 0 {
		t.Fatalf("Reset() 后进度 = %v，期望 0", progress)
	}
}

func TestTimerPhase(t *testing.T) {
	now := time.Date(2026, time.July, 21, 12, 0, 0, 0, time.UTC)
	var timer TimerState

	assertPhase(t, timer.Phase(now), TimerIdle)

	timer.Start(now, time.Minute)
	assertPhase(t, timer.Phase(now), TimerRunning)

	timer.Pause(now.Add(10 * time.Second))
	assertPhase(t, timer.Phase(now.Add(10*time.Second)), TimerPaused)

	timer.Resume(now.Add(10 * time.Second))
	assertPhase(t, timer.Phase(now.Add(11*time.Second)), TimerRunning)
	assertPhase(t, timer.Phase(now.Add(2*time.Minute)), TimerFinished)
}

func TestTimerView(t *testing.T) {
	now := time.Date(2026, time.July, 21, 12, 0, 0, 0, time.UTC)
	var timer TimerState
	timer.Start(now, time.Minute)

	view := timer.View(now.Add(30 * time.Second))
	assertPhase(t, view.Phase, TimerRunning)
	if view.Progress != 0.5 {
		t.Fatalf("进度 = %v，期望 0.5", view.Progress)
	}
}

func assertPhase(t *testing.T, actual, want TimerPhase) {
	t.Helper()
	if actual != want {
		t.Fatalf("阶段 = %v，期望 %v", actual, want)
	}
}
