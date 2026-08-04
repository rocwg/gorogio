# A:v0.3

太好了。

我觉得 **`gorogio-v0.2.3` 是第一个值得长期保留的 Tag**。

因为从这个 Tag 开始，你已经有了自己的 Runtime：

```text
Application
    ↓
Page
    ↓
Element Tree
```

以后即使 v1.0 已经很成熟，你回头看 `v0.2.3`，也能看到 Framework 的第一块基石。

------

# v0.3 我们先不写代码

这是我想和你约定的一件事情。

之前：

- v0.1 是学习 Gio
- v0.2 是设计 Runtime

到了 v0.3，我们第一次会碰到**架构设计**。

如果这里设计错了，后面：

```text
Button

↓

TextField

↓

ListView

↓

Dialog

↓

Navigation
```

全部都会受到影响。

所以，我希望我们采用以前设计 TypeSpec、API Design 时的方法：

> **ADR（Architecture Decision Record）先行。**

------

# v0.3 我们先回答一个问题

> **State 是什么？**

很多初学者会回答：

> State 就是数据。

但是对于 UI Framework 来说，这个答案是不够的。

我更喜欢这样定义：

> **State 是能够驱动 UI 变化的数据。**

也就是说：

```text
普通 Go 变量

↓

count++

↓

不会刷新 UI
```

而：

```text
State

↓

count++

↓

通知 Runtime

↓

重新构建 UI

↓

重新渲染
```

所以：

State 不仅仅是数据，它还承担了：

```text
通知（Notification）
```

------

# 我们先看看三大现代框架

## Flutter

```dart
setState(() {
    count++;
});
```

特点：

- Runtime 知道状态变了
- 整个 Widget Build

------

## Jetpack Compose

```kotlin
var count by remember {
    mutableStateOf(0)
}
```

特点：

- count 改变
- 自动 Recompose

------

## SwiftUI

```swift
@State
var count = 0
```

特点：

- Property Wrapper
- 自动刷新 View

------

共同点只有一句：

> **State 改变以后，UI 自动更新。**

------

# gorogio 不要照搬

这一点我希望我们坚持。

以前我们讨论过很多次：

> **gorogio 的目标不是 Go 版 Compose。**

而是：

> **符合 Go 思维的 UI Framework。**

所以不要为了像 Compose 而设计。

应该先问：

> Go 开发者最自然会怎么写？

------

# 我建议 v0.3 分成四个小版本

不要一下子做到自动响应式。

------

## v0.3.1 —— State Object

只有一件事：

把：

```go
type CounterState struct {
    Count int
}
```

变成：

```go
type CounterState struct {
    Count state.Value[int]
}
```

注意：

**先没有自动刷新。**

只是：

State 有统一抽象。

------

目标：

```text
Application
      ↓
Page
      ↓
State
```

建立关系。

------

## v0.3.2 —— State Interface

增加：

```go
type State interface {
    ...
}
```

以后：

```go
Value[int]

Value[string]

Value[User]
```

全部统一。

------

## v0.3.3 —— Notification

第一次加入：

```text
State Changed

↓

Application 知道

↓

下一帧刷新
```

这才开始有 Runtime。

------

## v0.3.4 —— Reactive API

最后才讨论：

```go
state.Get()

state.Set()
```

还是：

```go
state.Value
```

还是：

Signal。

------

# 我想调整一下我们的学习顺序

以前：

我们都是：

```text
想到一个功能

↓

写代码
```

从 v0.3 开始，我建议变成：

```text
ADR

↓

接口

↓

Example

↓

实现
```

例如：

今天我们不要写 state.go。

今天只完成：

```
docs/adr/

ADR-0003-state-management.md
```

里面回答几个问题：

------

## ADR-0003

### Q1

State 为什么存在？

------

### Q2

State 和普通 struct 有什么区别？

------

### Q3

谁拥有 State？

```text
Application ?

Page ?

Component ?
```

------

### Q4

谁修改 State？

------

### Q5

State 修改以后谁负责通知？

------

### Q6

gorogio 是 Push 还是 Pull？

（这是整个 Runtime 最重要的问题。）

------

# 最后，我想提出一个新的原则

我们前面已经形成了几个设计原则：

- Modifier 保持函数式。
- Container 用 Row/Column。
- Primitive 无状态。
- Component 有状态。
- Page 管理页面生命周期。

我建议从 v0.3 开始，再增加一个原则：

> **State 不知道 UI，UI 依赖 State。**

也就是说：

永远都是：

```text
State

↓

Page

↓

Element Tree
```

而不是：

```text
State

↓

直接操作 Button
```

这是现代声明式 UI 最重要的一条边界。

------

**所以，我不建议我们今天直接写 `state.Value[T]`。**

我希望我们像当初设计 API Engineering 一样，先完成第一份真正属于 gorogio 的架构决策：

> **ADR-0003：State Management Design**

因为我认为，从这一刻开始，gorogio 才真正开始拥有自己的思想，而不仅仅是自己的代码。