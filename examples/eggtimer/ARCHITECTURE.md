# Egg Timer：Gio 的状态驱动结构

这个案例不追求复刻 Jetpack Compose 的 API，而是借用它的三条原则：

1. UI 由状态推导；
2. 用户事件集中修改状态；
3. 组件只承担明确、稳定的视觉职责。

## 每帧链路

```text
app.FrameEvent
  -> updateTimer：消费按钮事件，修改 TimerState
  -> timer.View(now)：从业务状态生成只读 TimerViewState
  -> layoutTimer：仅根据 TimerViewState 组合并绘制 UI
  -> e.Frame：提交本帧的绘制操作
```

当计时器处于 `TimerRunning`，`layoutProgress` 会提交
`op.InvalidateCmd`，请求约 1/25 秒后的下一帧。因此时间流逝时，UI
仍然遵循同一条链路，而不需要后台 goroutine 直接修改进度。

## 状态分层

| 类型 | 保存内容 | 作用 |
| --- | --- | --- |
| `TimerState` | 时长、已过时间、开始时间、是否运行 | 可变的业务模型 |
| `TimerPhase` | Idle / Running / Paused / Finished | 对 UI 有意义的业务阶段 |
| `TimerViewState` | `Phase`、`Progress` | 当前帧交给 UI 的只读快照 |
| `TimerScreenUI` | `Clickable`、`Editor` | Gio 控件必须跨帧保留的交互状态 |

`TimerState` 不知道按钮、输入框和颜色；`TimerViewState` 不暴露
`elapsed`、`startedAt` 等内部细节；布局函数不直接修改 `TimerState`。

## 文件职责

| 文件 | 职责 |
| --- | --- |
| `main.go` | 创建窗口并启动 `run` |
| `timer.go` | 纯计时业务模型、阶段和 UI 快照 |
| `view.go` | 帧循环、事件处理、输入解析和页面布局 |
| `button.go` | 案例内的薄按钮组件封装 |
| `egg.go` | 根据 `EggProps` 绘制鸡蛋 |
| `timer_test.go` | 计时状态、阶段和 UI 快照的测试 |

当前不要继续拆子目录。这个案例只有一个屏幕，文件已经按职责分开；
此时新增 `ui/`、`domain/` 等目录只会增加跳转成本。只有出现第二个屏幕，
或多个案例确实共享同一组件时，才考虑提取 package 或 `internal/` 目录。

## 维护规则

- 新的业务规则优先放进 `TimerState`；
- 新的点击事件在 `updateTimer` 中消费；
- 新的显示数据加入 `TimerViewState`，再由布局读取；
- 一个视觉元素在至少两个案例中稳定复用后，才考虑移出本案例；
- 不为了模仿 Compose 而包装 `layout.Flex`、`layout.Inset` 等 Gio 基础能力。

## 验证

```powershell
go test -race ./examples/eggtimer
go vet ./examples/eggtimer
go run ./examples/eggtimer
```
