package main

import (
	"image/color"
	"strconv"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// TimerScreenUI 保存 Gio 控件跨帧需要保留的交互状态。
//
// 它刻意不存放计时业务数据；业务数据属于 TimerState。这样的区分
// 对应 Compose 中“组件暂态”和“上提后的业务状态”各司其职的思路。
type TimerScreenUI struct {
	StartButton   widget.Clickable
	ResetButton   widget.Clickable
	DurationInput widget.Editor
}

// newTimerScreenUI 创建并初始化界面控件。控件只创建一次，不能每帧重建。
func newTimerScreenUI() TimerScreenUI {
	ui := TimerScreenUI{}
	ui.DurationInput.SingleLine = true
	ui.DurationInput.Alignment = text.Middle
	ui.DurationInput.SetText("300")
	return ui
}

// run 管理 Gio 的帧事件：先处理输入和状态，再根据最新状态绘制界面。
func run(w *app.Window) error {
	var (
		ops   op.Ops
		ui    = newTimerScreenUI()
		timer TimerState
	)

	th := material.NewTheme()
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			now := gtx.Now

			updateTimer(gtx, now, &ui, &timer)
			viewState := timer.View(now)
			layoutTimer(gtx, th, now, &ui, viewState)
			e.Frame(gtx.Ops)

		case app.DestroyEvent:
			return e.Err
		}
	}
}

// updateTimer 集中处理按钮事件。这让布局代码只负责呈现状态。
func updateTimer(gtx layout.Context, now time.Time, ui *TimerScreenUI, timer *TimerState) {
	// Reset 是独立的业务意图：它不改变输入框，只清空当前计时进度。
	for ui.ResetButton.Clicked(gtx) {
		timer.Reset()
	}

	// Clicked 可能积累多个未处理事件，因此使用 for 将它们全部消费。
	for ui.StartButton.Clicked(gtx) {
		switch timer.Phase(now) {
		case TimerIdle, TimerFinished:
			if duration, ok := parseSeconds(ui.DurationInput.Text()); ok {
				timer.Start(now, duration)
			}
		case TimerRunning:
			timer.Pause(now)
		case TimerPaused:
			timer.Resume(now)
		}
	}
}

// parseSeconds 位于视图层，因为 string 是 Editor 提供的 UI 输入。
// 它把合法的秒数转换为 TimerState 所需的 time.Duration。
func parseSeconds(input string) (time.Duration, bool) {
	seconds, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil || seconds <= 0 {
		return 0, false
	}
	return time.Duration(seconds * float64(time.Second)), true
}

// layoutTimer 只读取本帧的 TimerViewState，不直接访问可变的 TimerState。
func layoutTimer(gtx layout.Context, th *material.Theme, now time.Time, ui *TimerScreenUI, state TimerViewState) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return Egg(gtx, EggProps{Progress: state.Progress})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layoutDurationInput(gtx, th, &ui.DurationInput)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layoutProgress(gtx, th, now, state)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layoutTimerActions(gtx, th, ui, state)
		}),
	)
}

// layoutDurationInput 绘制输入总秒数的文本框。
func layoutDurationInput(gtx layout.Context, th *material.Theme, input *widget.Editor) layout.Dimensions {
	inset := layout.Inset{Right: unit.Dp(170), Bottom: unit.Dp(40), Left: unit.Dp(170)}
	border := widget.Border{Color: color.NRGBA{R: 204, G: 204, B: 204, A: 255}, CornerRadius: unit.Dp(3), Width: unit.Dp(2)}
	editor := material.Editor(th, input, "seconds")
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return border.Layout(gtx, editor.Layout)
	})
}

// layoutProgress 绘制进度条，并在计时期间请求下一帧以实现动画。
func layoutProgress(gtx layout.Context, th *material.Theme, now time.Time, state TimerViewState) layout.Dimensions {
	if state.Phase == TimerRunning {
		gtx.Execute(op.InvalidateCmd{At: now.Add(time.Second / 25)})
	}
	return material.ProgressBar(th, state.Progress).Layout(gtx)
}

// layoutTimerActions 将开始/暂停和重置操作并排放置。
func layoutTimerActions(gtx layout.Context, th *material.Theme, ui *TimerScreenUI, state TimerViewState) layout.Dimensions {
	buttonLabel := ""
	switch state.Phase {
	case TimerIdle:
		buttonLabel = "Start"
	case TimerRunning:
		buttonLabel = "Pause"
	case TimerPaused:
		buttonLabel = "Resume"
	case TimerFinished:
		buttonLabel = "Start again"
	}

	inset := layout.Inset{Top: unit.Dp(25), Right: unit.Dp(35), Bottom: unit.Dp(25), Left: unit.Dp(35)}
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return Button(gtx, th, &ui.StartButton, ButtonProps{Text: buttonLabel, Variant: ButtonPrimary})
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: unit.Dp(12)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return Button(gtx, th, &ui.ResetButton, ButtonProps{Text: "Reset", Variant: ButtonSecondary})
				})
			}),
		)
	})
}
