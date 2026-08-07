### ==A0== 

master 先恭喜一下。🙂

这次我不是单纯说“运行成功”。

我认为你刚刚完成的是 **gorogio 第一个真正意义上的 Component 生命周期闭环**。

而且非常巧：

你自己增加 `Clear` 按钮，刚好触碰到了我们下一步 ADR-0005 的核心。

------

# 一、先评价你的实现

整体评价：

> **设计方向正确，已经达到 gorogio v0.3.1 的目标。**

你的链路：

```
Application
    |
    v
FormScreen
    |
    +----------------+
    |                |
 Input            Button
    |                |
 Editor          Clickable
```

已经跑通。

------

# 二、Input Component 现在已经具备完整三要素

我们之前约定：

每个 Component 思考三个问题：

```
① 拥有什么状态？

② 对外暴露什么行为？

③ 如何生成 Element？
```

我们看你的 Input。

------

## ① 状态 Ownership

你的：

```go
type Input struct {

    editor widget.Editor

    label string
}
```

非常关键。

因为：

状态属于 Input。

也就是说：

```
Input
 |
 +-- Editor State
```

不是：

```
FormScreen
 |
 +-- Name string
 +-- Age string
 +-- City string
```

这是一种合理选择。

------

## ② 行为

你增加：

```go
func (i *Input) Value() string
```

很好。

然后：

```go
func (i *Input) SetValue(value string)
```

这个更重要。

因为它让：

Screen

可以控制：

Component。

现在关系：

```
FormScreen

    调用

Input.SetValue()
```

非常自然。

------

## ③ Element

你的：

```go
func (i *Input) Element(
    th *style.Theme,
)
```

保持：

```
Input
    |
    Element
    |
    Gio Widget
```

正确。

------

# 三、Clear 按钮其实帮我们发现一个重要设计点

你的代码：

```go
screen.Clear =
component.NewButton("清空").
OnClick(
    func() {

        screen.Name.SetValue("")

        screen.Age.SetValue("")

        screen.City.SetValue("")

    },
)
```

现在：

完全正确。

因为：

Clear 是 Screen 的业务行为。

为什么？

因为：

“清空表单”

不是 Input 的职责。

Input 只知道：

```
我要设置自己的值
```

但是：

```
什么时候三个 Input 一起清空？
```

这是 Screen 的业务。

所以：

现在边界：

非常漂亮。

------

# 四、你已经无意中完成了一次状态分层

现在：

## Component State

例如：

Input:

```go
editor
```

Button:

```go
clickable
```

属于：

```
Component Internal State
```

------

## Screen State

以后：

例如：

```go
type FormScreen struct {

    UserID string

    Loading bool

    Error string

}
```

属于：

```
Screen Business State
```

------

## Application State

例如：

```
Current Screen

Theme

Window

Router
```

属于：

```
Application Runtime State
```

你已经开始形成三层。

------

# 五、但是我要指出两个小地方，可以修正

## 问题 1：Update 里面不要留 TODO

现在：

```go
func (i *Input) Update(
    gtx layout.Context,
) {
    //TODO
}
```

虽然没问题。

但是从设计语言看：

容易误导。

建议：

改成：

```go
func (i *Input) Update(
    gtx layout.Context,
) {
    // 当前版本无需额外事件处理。
    // Gio Editor 在 Element 渲染阶段处理输入事件。
}
```

为什么？

因为 TODO 的含义：

“未来必须实现”。

但是这里：

不是。

------

## 问题 2：Input 的 SetValue 可以稍微保护一下

现在：

```go
func (i *Input) SetValue(value string) {

    i.editor.SetText(value)

}
```

没问题。

以后：

如果有：

```
onChange
validation
```

这里会成为入口。

所以保留。

------

# 六、现在最重要：进入 ADR-0005

我认为：

现在时机成熟。

因为你已经遇到了真实问题：

> 状态到底属于 Screen，还是 Component？

我们可以开始设计。

------

# ADR-0005 初步草案

标题：

```
ADR-0005 Component Lifecycle Ownership
```

------

## Context

gorogio 中：

Component 需要保存交互状态。

例如：

```
Button

Input

Counter
```

同时：

Screen 需要组合多个 Component。

问题：

状态应该归谁管理？

------

## Options

## Option A

Component 自己拥有状态

例如：

```go
type Input struct {

    editor widget.Editor

}
```

优点：

简单。

组件独立。

缺点：

业务状态同步困难。

------

## Option B

Screen 拥有全部状态

例如：

```go
type FormScreen struct {

    Name string

}
```

Input:

只是：

View。

优点：

业务清晰。

缺点：

Component 变薄。

------

## Option C

混合模式（推荐）

规则：

```
Component

拥有 UI 交互状态


Screen

拥有业务状态
```

例如：

Input:

拥有：

```
Editor cursor
Focus
Selection
```

Screen:

拥有：

```
Username
Age
City
```

------

# 七、我建议 gorogio 选择 C

原因：

非常符合 Gio + Go。

冻结：

```
Component Ownership

Component owns interaction state.

Screen owns business state.
```

中文：

> Component 负责自己的交互生命周期状态；Screen 负责业务数据和业务流程状态。

------

# 八、但是这里还有一个下一阶段问题

如果采用 C：

未来 Input 会变成：

```go
Input(
    value string,
    onChange func(string),
)
```

类似：

Compose：

```kotlin
TextField(
    value,
    onValueChange
)
```

而不是：

```go
Input.Value()
Input.SetValue()
```

也就是说：

你现在版本：

v0.3.1

是：

```
Component owns state
```

未来：

可能演进：

```
Controlled Component
```

------

但是！

现在不要改。

因为：

你现在正在学习：

- struct
- pointer
- method
- ownership

现在这个版本非常适合 Go 学习。

------

# 九、我建议你下一步动作

不要继续写代码。

先：

1. commit：

```
feat(component): add Input component and form example
```

1. 写：

```
docs/adr/ADR-0005-component-lifecycle-ownership.md
```

冻结当前理解。

1. 然后我们讨论：

**v0.3.2 是否引入 controlled component。**

------

最后，我想特别评价一点：

你这次不是“复制我的代码”。

你自己：

- 增加 Clear
- 发现 SetValue
- 验证状态变化
- 观察控制台输出

这个过程非常重要。

因为真正的软件能力，就是：

> 写一点 → 发现问题 → 调整边界 → 固化设计。

现在 gorogio 已经开始像一个真正的小型工程了。🙂



### ==A1==

master，非常好。👍

我认为现在**正式进入 ADR-0005 是正确时机**。

而且我想强调一下：

这不是“为了写文档而写 ADR”。

你刚刚完成：

```
Input
 ↓
FormField
 ↓
FormScreen
```

之后，自然产生了一个真实工程问题：

> FormField 内部拥有 Input，那么状态到底应该属于谁？

这个问题如果不回答，后面的：

- Checkbox
- Select
- Dialog
- Table
- Tree
- Tabs
- Navigation

都会不断重复遇到。

所以 ADR-0005 正好。

------

# ADR-0005 Component Lifecycle Ownership

我建议标题：

```text
ADR-0005: Component Lifecycle Ownership
```

中文：

```text
组件生命周期与状态所有权设计
```

------

# 1. Context（背景）

当前 gorogio v0.3.x：

结构：

```
Application
    |
    |
 Screen
    |
    |
 Component
    |
    |
 Gio Widget
```

例如：

```
FormScreen

    FormField

        Input

            widget.Editor
```

目前状态：

```
Input owns Editor state
```

也就是：

```go
type Input struct {

    editor widget.Editor

}
```

Input 自己管理：

- 输入内容
- 光标
- selection
- focus

------

# 2. 问题

随着组件增加，会遇到：

## 问题 A

组件内部状态：

例如：

```go
type Checkbox struct {

    checked bool

}
```

谁负责：

```go
checked=true
```

？

------

## 问题 B

Screen 是否应该知道：

```go
checkbox.checked
```

？

------

## 问题 C

是否应该支持：

类似 React：

```text
Controlled Component
```

？

例如：

```go
Input(
    Value:"hello",
    OnChange:func(v string)
)
```

------

# 3. 候选方案

------

# Option A

## Component Owns State

组件拥有状态。

模型：

```
Screen

    Button

    Input
        |
        editor state


    Checkbox
        |
        checked state
```

代码：

```go
type Checkbox struct {

    checked bool

}
```

使用：

```go
checkbox.Toggle()
```

------

优点：

符合 Gio。

原因：

Gio 本身：

```go
widget.Clickable

widget.Editor

widget.List
```

都是：

对象持有状态。

非常 Go。

------

缺点：

外部控制稍弱。

例如：

表单初始化：

需要：

```go
input.SetValue()
```

------

# Option B

## Screen Owns State

状态提升。

类似 React：

```
Screen

    username string


        Input
```

代码：

```go
type Input struct {

    Value string

}
```

使用：

```go
Input{
    Value:screen.Name,
}
```

优点：

数据流清晰。

缺点：

Go/Gio 下会变复杂。

尤其：

每个 widget：

都需要同步。

------

# Option C

## Hybrid

混合模式。

基础组件：

Component owns state。

高级场景：

支持 controlled。

例如：

普通：

```go
input := component.NewInput()
```

内部：

```
editor
```

------

高级：

```go
input.Bind(&username)
```

------

# 我的建议

gorogio 当前阶段：

## 选择：

# Option A

Component Owns State

原因：

非常符合你的目标：

> Learning Go Through Gio

------

因为你现在不是做：

React clone。

你是在学习：

```
Go
 +
 Gio Immediate Mode
```

------

Gio 的哲学：

不是：

```
State
 ↓
Render
```

而是：

```
Object
 |
 owns state
 |
 participates in frame
```

------

# 4. Decision（决定）

建议 ADR 写：

------

## Decision

gorogio v0.3.x 采用：

```
Component Owned State Model
```

原则：

1. Component 拥有内部交互状态。
2. Screen 负责组合 Component。
3. Screen 不直接操作 Component 内部字段。
4. Component 对外暴露语义方法。

例如：

正确：

```go
input.SetValue("")
```

不允许：

```go
input.editor.SetText("")
```

------

Checkbox:

正确：

```go
checkbox.SetChecked(true)
```

不允许：

```go
checkbox.checked=true
```

------

# 5. Consequences（影响）

## 好处

Component:

更加独立。

例如：

```go
component.Input
```

可以迁移到：

其他 Screen。

------

## 代价

复杂表单：

可能需要更多 API。

例如：

```go
form.Reset()
```

需要：

```
Screen
    ↓
Component
    ↓
SetValue
```

------

# 6. Future

不要现在做。

但是记录：

未来 v1.x：

可能增加：

```
Controlled Component
```

例如：

```go
InputOptions{

    Value string

    OnChange func(string)

}
```

------

# 我建议你的 ADR 最终结论：

一句话：

> gorogio follows Gio's state ownership philosophy. Components own their interaction state, while Screens own composition and business flow.

------

# 然后 ADR 完成后，下一步编码路线：

我建议：

不要马上 Checkbox。

顺序：

```
ADR-0005
   |
   |
component Checkbox
   |
   |
settings Screen
   |
   |
发现新的状态问题
   |
   |
ADR-0006
```

因为：

Checkbox 是最小状态组件：

```go
checked bool
```

它会验证 ADR-0005。

------

master，你现在 gorogio 的路线其实已经出现一个很好的节奏：

```
v0.2
Gio封装

v0.3
Component模型

v0.3.4
Component组合

v0.4
状态模型

v0.5
页面导航

...
```

不要急。

今天 FormField 这一小步，已经让 gorogio 从“Gio wrapper”开始变成“小型 UI runtime 设计实验”。这一步很关键。👍
