好的，请您帮我完成 ADR-0003：State Management Design

### ==A1==

当然愿意，而且我觉得**这一篇 ADR 将是 gorogio 第一个真正体现设计思想的 ADR**。

不过，在开始写之前，我想先做一个小调整。

------

## 我建议把编号改一下

目前我们已经有：

```text
ADR-0001  Layout API
ADR-0002  Page Lifecycle
```

那么：

```text
ADR-0003  State Management
```

就非常自然。

------

# ADR-0003：State Management Design（初版）

我会尽量保持和你之前 API Engineering 仓库 ADR 的风格一致：**记录设计决策，而不是记录实现细节。**

------

# ADR-0003：State Management Design

## Status

Accepted

------

## Context

gorogio v0.2.3 已完成 UI Runtime 的第一阶段：

```
Application
        │
        ▼
      Page
        │
        ▼
   Element Tree
```

Page 已能够：

- 保存页面状态（State）
- 处理用户事件（Update）
- 构建 UI（Element）

然而，目前 State 仍然只是普通 Go 对象：

```go
type CounterState struct {
    Count int
}
```

虽然能够正常工作，但 Framework 无法感知状态变化，也无法建立统一的状态生命周期。

为了支持未来的响应式 UI（Reactive UI）、页面状态管理以及可观察状态（Observable State），gorogio 需要建立统一的 State 模型。

------

## Problem

目前普通 Go 对象存在以下问题：

- Framework 无法区分普通数据和 UI State。
- 状态修改不会产生统一通知。
- Application 无法管理 State 生命周期。
- Component 无法形成统一的数据流。

因此需要建立 gorogio 自己的 State 抽象。

------

## Decision

gorogio 引入 State 作为 Runtime 的核心组成部分。

Runtime 架构调整为：

```
Application
        │
        ▼
      Page
        │
        ▼
      State
        │
        ▼
   Element Tree
```

其中：

- Application 管理页面生命周期。
- Page 持有 State。
- State 保存页面数据。
- Element Tree 根据 State 构建。

------

## Design Principles

### Principle 1

State 是能够驱动 UI 的数据，而不是普通数据。

只有参与 UI 渲染的数据，才属于 State。

普通计算过程中的局部变量，不属于 State。

------

### Principle 2

State 不负责 UI。

State 不知道：

- Button
- Text
- Dialog
- Layout

State 仅负责保存数据。

UI 根据 State 构建。

依赖方向始终保持：

```
State
    │
    ▼
Page
    │
    ▼
Element Tree
```

禁止反向依赖。

------

### Principle 3

Page 是 State Owner。

一个 Page 负责拥有自己的 State。

例如：

```
LoginPage
    └── LoginState

HomePage
    └── HomeState

SettingsPage
    └── SettingsState
```

State 生命周期与 Page 保持一致。

------

### Principle 4

Component 不拥有业务 State。

Button、Dialog、Navigation 等 Component 可以拥有自身内部运行状态，例如：

- Clickable
- Focus
- Hover
- Animation

但不得保存业务数据。

例如：

```
Counter

属于：

CounterState

而不是：

Button
```

------

### Principle 5

State 修改只能通过公开 API。

禁止外部直接操作内部状态。

未来所有状态更新将通过统一 API 完成。

例如：

```
Increment()

Reset()

Set()

Update()
```

而不是任意修改内部字段。

------

### Principle 6

Framework 不依赖具体业务 State。

Application 只知道：

```
Page
```

Page 自己决定：

```
UserState

OrderState

CounterState
```

Framework 永远不知道业务类型。

------

## Future Evolution

State Management 将分阶段演进。

### v0.3.1

建立统一 State 抽象。

目标：

```
State Object
```

------

### v0.3.2

建立统一 State Interface。

目标：

```
State Interface
```

------

### v0.3.3

引入状态通知机制。

目标：

```
State Changed

↓

Runtime Notification
```

------

### v0.3.4

引入响应式 API。

目标：

```
Reactive State
```

------

## Consequences

完成本 ADR 后：

gorogio Runtime 将演进为：

```
Application
        │
        ▼
      Page
        │
        ▼
      State
        │
        ▼
   Element Tree
```

Runtime 将拥有统一的数据流方向：

```
State

↓

Page

↓

Element Tree
```

未来所有响应式能力、数据绑定、页面刷新、状态同步等功能，都将在此模型基础上逐步实现。

本 ADR 不讨论具体实现，仅冻结 State 在 gorogio Runtime 中的职责与边界。

我还有一个建议，而且这是我看到整个 ADR 后最想补充的一点。

------

## 我建议增加一个 gorogio 自己的设计哲学（Design Philosophy）

以后每一篇 ADR 的最后，都增加一个小节：

```text
## Philosophy
```

例如这篇可以写一句：

> **State is the single source of UI truth.**

但是我们不用完全照搬 React 的表达。

我更喜欢符合 gorogio 的一句话：

```text
UI is a projection of State.
```

翻译成中文就是：

> **UI 是 State 的投影。**

也就是说：

- UI 不保存业务数据。
- UI 不修改业务逻辑。
- UI 只是把当前 State 投影到屏幕上。

这句话其实贯穿了 Compose、SwiftUI、Flutter 和 React 的声明式思想，但用更中性的方式表达出来。

------

最后，我想对 v0.3 的路线做一个小调整。

昨天我说：

> v0.3.1：State Object

今天看完你的 Runtime 和这份 ADR，我认为更自然的路线应该是：

```text
v0.3.1  State<T>（统一状态对象）

↓

v0.3.2  Application 持有 Invalidator（刷新机制）

↓

v0.3.3  State → Invalidate（状态驱动刷新）

↓

v0.3.4  Remember / Local State（页面局部状态）
```

也就是说，**我们不要急着做响应式，而是先建立"State 如何通知 Runtime"这条生命线。**

因为从 Runtime 的角度来看，真正重要的不是 `State<T>` 怎么存值，而是：

> **当 State 改变时，Application 是如何知道"下一帧应该重新构建 UI"的。**

我认为，这才是 gorogio v0.3 最值得设计的核心。

---





### ==A2==

> 请问，Gio 到 gorogio 真的要 做 响应式改造？
> 我感觉 有点 顾虑，好像有违 Go 和 Gio 的思维。



roc，这个问题，我认为是 **gorogio 到目前为止最重要的一个设计问题。**

而且，我想先纠正我昨天说的一件事。

> **昨天我说 "v0.3 做响应式"，今天我觉得这个提法并不准确。**

你的顾虑是有道理的。

------

# 我们先不要谈"响应式"

因为"响应式（Reactive）"这个词，已经被：

- React
- Vue
- Compose
- SwiftUI

赋予了很强的语义。

但是 **Gio 根本不是这么工作的。**

------

## Gio 是怎么工作的？

其实 Gio 非常像游戏引擎。

它每一帧：

```text
Event

↓

Update()

↓

Layout()

↓

Present
```

注意：

**没有 Diff。**

没有：

```text
State

↓

Diff

↓

Patch UI
```

------

例如你的代码：

```go
application.Update(gtx)

application.Draw(gtx)
```

实际上已经是：

```text
Frame

↓

Update()

↓

重新构建 Element Tree

↓

Draw
```

这和 Compose 不一样。

------

# Compose 为什么需要响应式？

因为：

Compose：

```kotlin
Text("$count")
```

并不会每一帧执行。

只有：

```text
State Changed

↓

Recompose
```

才会重新执行。

所以：

Compose 必须知道：

> State 什么时候变了。

------

# Gio 根本不用知道

因为：

下一帧：

```go
func (p *HelloPage) Element(...) {

    primitive.Body(
        strconv.Itoa(
            p.State.Count,
        ),
    )

}
```

每一帧都会重新执行。

所以：

```go
p.State.Count++
```

下一帧：

自然就是：

```text
新的 Count
```

没有任何：

Observer。

没有：

Signal。

没有：

remember。

------

# 所以昨天我忽然意识到

我昨天实际上带了一点 Compose 思维。

但是：

**gorogio 不应该为了像 Compose，而脱离 Gio。**

这一点你提醒得非常及时。

------

# 那 v0.3 应该做什么？

我今天重新思考以后。

我认为：

v0.3 不应该叫：

```text
State Management
```

应该叫：

```text
Application State
```

或者：

```text
State Model
```

重点不是：

响应式。

而是：

> **统一 State 的组织方式。**

例如：

现在：

```go
type CounterState struct {
    Count int
}
```

完全没问题。

我反而觉得：

**这就是 Go。**

------

# Go 的哲学是什么？

Go：

一直都是：

```go
type User struct {

    Name string

}
```

而不是：

```go
Observable<User>

MutableState<User>

Signal<User>
```

------

这也是 Gio 的哲学。

------

# 我越来越倾向一个设计

例如：

```go
type CounterState struct {

    Count int

}
```

保持。

Button：

```go
func() {

    state.Count++

}
```

保持。

Application：

```go
Update()

↓

Element()

↓

Draw()
```

保持。

Framework：

**根本不用知道 Count 改了。**

因为：

下一帧：

自然：

重新：

Element()。

------

# 那 state 包还有意义吗？

有。

但是：

不是：

```text
Reactive State
```

而是：

```text
State Convention
```

例如：

以后：

```text
state/

    lifecycle.go

    persistence.go

    navigation.go

    session.go
```

或者：

```go
type State interface{}
```

仅仅作为：

Framework 概念。

而不是：

Observable。

------

# 你有没有发现一个有趣的地方？

你的 Runtime：

其实已经：

```text
Frame

↓

Update()

↓

Element()

↓

Draw()
```

是不是很像：

游戏：

```text
Frame

↓

Update()

↓

Render()
```

所以：

gorogio：

其实更接近：

**Immediate Mode UI。**

而不是：

React。

------

# 我建议把 ADR-0003 改一下

昨天：

我写：

```text
State

↓

Notification
```

今天：

我会删掉。

改成：

```text
State

↓

Page

↓

Element Tree
```

State：

只是：

数据。

Application：

每一帧：

主动：

重新读取。

所以：

这是：

**Pull。**

不是：

Push。

------

# 我认为 gorogio 应该坚持 Gio 哲学

甚至可以把 Philosophy 改成一句话：

> **State does not notify Runtime. Runtime reads State every frame.**

翻译：

> **State 不通知 Runtime，而是 Runtime 在每一帧主动读取 State。**

这一句话，我认为非常重要。

它意味着：

- 没有观察者（Observer）
- 没有信号（Signal）
- 没有响应式依赖收集
- 没有 remember
- 没有 mutableStateOf

而只有：

```text
Frame

↓

Update

↓

Read State

↓

Build Element Tree

↓

Draw
```

------

## 所以，我现在给你的建议，和昨天相比，有一个比较大的修正。

**我建议 gorogio 不要做 Compose 式响应式。**

原因不是"做不到"，而是**没有必要**。

因为 Gio 已经给了你一个非常优雅、非常符合 Go 思维的 Runtime：

> **每一帧读取状态，而不是让状态驱动 Runtime。**

------

### 我甚至建议，我们把 v0.3 的名字正式改掉。

不要叫：

```text
v0.3
State Management
```

改成：

```text
v0.3
State Model
```

或者更准确一点：

```text
v0.3
Application State Model
```

整个 v0.3 的目标就变成：

- 定义什么是 State。
- 定义 State 应该放在哪里。
- 定义 Page 如何拥有 State。
- 定义 Component 不应该拥有业务 State。

**不引入任何 Observer、Signal、Reactive、remember。**

我认为，这样才是真正尊重 **Go + Gio** 的设计哲学，而不是把 Compose 的机制搬到 Gio 上。



### ==A3==

> 嗯嗯，我也觉得是这样才是真正尊重 Go + Gio 的设计哲学，而不是把 Compose 的机制搬到 Gio 上。
> 请指导我下一步，直接编码吗？

roc，我的答案是：

> **是的，直接编码。**
>
> **但是，不是编码 State，而是编码 Runtime 的下一层抽象。**

这里我想和你分享一个我刚刚重新整理后的路线。

------

# 我们现在的位置

如果把 gorogio 看成一个 UI Framework，它已经有：

```text
Application Runtime        ✅
Page Lifecycle            ✅
Element                   ✅
Container                 ✅
Modifier                  ✅
Primitive                 ✅
Component                 ✅
```

缺什么？

很多人会回答：

```text
State
```

**我现在认为不是。**

因为：

```go
type CounterState struct {
    Count int
}
```

已经很好。

它就是 Go。

------

# 真正缺的是谁？

看看你的 HelloPage：

```go
type HelloPage struct {
    State *CounterState

    Increment *component.Button
    Reset     *component.Button
}
```

然后：

```go
func (p *HelloPage) Update(...)
```

里面：

```go
p.Increment.Update(gtx)
p.Reset.Update(gtx)
```

以后：

```text
Button

TextField

Checkbox

Slider

Dialog

Menu

List
```

会变成：

```go
func (p *HelloPage) Update(...) {

    p.Button1.Update(gtx)
    p.Button2.Update(gtx)
    p.Button3.Update(gtx)
    p.TextField.Update(gtx)
    p.Checkbox.Update(gtx)
    ...
}
```

**这里才是真正的痛点。**

------

# Runtime 还有一层没有抽出来

现在：

```text
Application

↓

Page.Update()

↓

Button.Update()

↓

Button.Update()

↓

Button.Update()
```

其实：

**Page 在做 Component 生命周期管理。**

这意味着：

Page 同时负责：

```text
State

Event

UI

Component Lifecycle
```

职责有点多了。

------

# 我建议 v0.3 真正做这个

不是：

State。

而是：

## Component Lifecycle

例如：

```go
type Component interface {

    Update(
        gtx layout.Context,
    )

    Element(
        th *style.Theme,
    ) element.Element

}
```

是不是很眼熟？

它和：

```go
type Page interface
```

几乎一样。

------

然后 Button：

```go
type Button struct {
    ...
}
```

自然实现：

```go
Component
```

------

以后：

Page：

不用：

```go
p.Button.Update(gtx)
```

而是：

```go
p.Components.Update(gtx)
```

或者：

```go
component.UpdateAll(
    gtx,
    p.Button1,
    p.Button2,
    p.Input,
)
```

------

# 你有没有发现？

现在：

Page：

```text
Update

↓

Button.Update()

↓

Button.Update()

↓

Button.Update()
```

以后：

```text
Page.Update()

↓

Component Runtime

↓

所有 Component
```

Runtime 又干净了一层。

------

# 再往后

Navigation：

Dialog：

Popup：

都可以：

加入：

```text
Component Runtime
```

而不是：

Page 一个个：

```go
Update()
```

------

# 然后才轮到 State

因为：

那时候：

Page：

真正只剩：

```text
State

↓

Build UI
```

生命周期：

全部：

Framework：

负责。

------

# 我现在重新整理了一下 Roadmap

我觉得应该是：

```text
v0.2.3
Application Runtime          ✅

↓

v0.3
Component Lifecycle

↓

v0.4
Navigation

↓

v0.5
Theme System

↓

v0.6
Animation

↓

v0.7
Overlay/Dialog

↓

v0.8
List/Lazy

↓

v0.9
Accessibility

↓

v1.0
Stable
```

你会发现：

**里面没有 Reactive。**

------

# 为什么我现在越来越喜欢这条路线？

因为：

Compose：

最大的难点：

不是：

Button。

而是：

```text
remember

snapshot

recompose

slot table
```

这些都是为了：

**避免每一帧 Build。**

而 Gio：

已经接受：

> 每一帧重新 Layout。

所以：

gorogio：

根本不用解决：

Compose 的问题。

------

# 我的建议（也是 v0.3 的第一步）

我们不要新建 `state.Value[T]`。

而是开始设计：

```go
component/component.go
```

定义：

```go
type Component interface {
    Update(gtx layout.Context)
    Element(th *style.Theme) element.Element
}
```

然后：

让：

```go
Button
```

成为：

```text
第一个真正实现 Component 接口的组件
```

这会让你的 Framework 架构变成：

```text
Application
        │
        ▼
      Page
        │
        ▼
   Component
        │
        ▼
    Element Tree
```

注意，这里的 **Component** 不是目录，而是**生命周期协议**。

------

## 最后，我想给你一个建议

这是我们这几个月合作以来，我最想说的一句话。

你已经发现了一个非常好的倾向：

> **不为了"现代"而设计，而是为了"适合 Go"而设计。**

这也是我希望 gorogio 最终形成的气质：

- 学习 Compose、SwiftUI、Flutter 的**优秀 API**。
- 保留 Gio 的**Immediate Mode Runtime**。
- 保留 Go 的**简单数据模型**。
- 不照搬其他框架的响应式机制。

如果我们一直坚持这个原则，我相信十年以后回头看，你会发现 **gorogio 的价值，不在于它像谁，而在于它终于有了自己的设计语言。**