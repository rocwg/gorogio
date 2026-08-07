master 收到。🙂

这一步非常关键。

因为 **ADR-0005 解决的是“生命周期属于谁”**：

> Component 可以拥有自己的交互生命周期。

而 **ADR-0006 进一步冻结“状态模型”**：

> 什么状态应该留在 Component，什么状态应该提升到 Screen。

我们先冻结设计，再编码。

------



### ==A0== 

## 一、ADR-0006 Component State Model

建议文件：

```
docs/adr/ADR-0006-component-state-model.md
```

内容：

# ADR-0006 Component State Model

## Status

Accepted

## Context

gorogio 中 Component 开始承担：

- UI 渲染
- 用户交互
- 内部生命周期管理

例如：

- Button
- Counter
- Input

随着 Component 数量增加，需要明确：

状态应该属于 Component，还是属于 Screen。

如果没有明确规则，未来容易出现：

- Component 过度智能
- Screen 管理大量 UI 细节状态
- 状态来源混乱

## Decision

gorogio 采用：

# Hybrid Component State Model

混合状态模型。

规则：

## 1. Component 拥有 Interaction State

Component 负责管理：

- 鼠标状态
- 键盘状态
- Focus
- Selection
- 输入过程状态
- 动画状态

例如：

Input:

```go
editor widget.Editor
```

Button:

```go
widget.Clickable
```

这些状态属于 Component。

------

## 2. Screen 拥有 Business State

Screen 负责：

- 业务数据
- 页面流程
- 业务规则

例如：

FormScreen:

```go
type FormState struct {
    Name string
    Age int
}
```

Component 不应该知道业务含义。

------

## 3. Controlled API 作为高级能力存在

默认：

Component 内部管理交互状态。

但是对于复杂场景：

允许 Screen 注入：

- Value
- OnChange
- Callback

形成受控组件。

例如：

```go
Input(
    value,
    onChange,
)
```

Controlled 模式不是默认模式。

------

## Consequences

优点：

1. 保持 Gio Immediate Mode 思想。
2. 保持 Go 简单的数据模型。
3. Component 可以独立复用。
4. Screen 保持业务清晰。

缺点：

1. 简单 Component 和复杂 Component 可能存在两种 API。
2. 需要在实践中不断平衡。

## Example

Input:

Component:

```go
type Input struct {
    editor widget.Editor
}
```

负责：

- 输入事件
- 光标
- Focus

Screen:

```go
type FormScreen struct {

    username string

}
```

负责：

- 用户数据
- 提交逻辑

## Future

未来版本继续验证：

- Checkbox
- Select
- Toggle
- Form

是否符合该模型。

------

# 二、开始编码：component/checkbox.go

这次不要做复杂。

目标：

学习：

- bool 状态
- 点击事件
- Component 生命周期

新增：

```
component/checkbox.go
```

代码：

```go
package component

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)


// Checkbox
//
// 一个简单复选框组件。
//
// 自己拥有：
// checked 状态
type Checkbox struct {

	click widget.Bool

	label string
}


// NewCheckbox
func NewCheckbox(
	label string,
) *Checkbox {

	return &Checkbox{
		label: label,
	}
}


// Update
//
// 每一帧处理点击事件。
func (c *Checkbox) Update(
	gtx layout.Context,
) {


	c.click.Update(gtx)

}


// Checked
//
// 获取当前状态。
func (c *Checkbox) Checked() bool {

	return c.click.Value

}


// Toggle
//
// 修改状态。
func (c *Checkbox) Toggle() {

	c.click.Value = !c.click.Value

}


// Element
func (c *Checkbox) Element(
	th *style.Theme,
) element.Element {


	return func(
		gtx layout.Context,
	) layout.Dimensions {


		return material.CheckBox(
			th.Material,
			&c.click,
			c.label,
		).Layout(gtx)

	}

}
```

------

# 三、新增 settings 示例

目录：

```
examples/
└── gorogio/
    └── settings/
        ├── main.go
        ├── runtime.go
        └── settings_screen.go
```

先复制：

```
hello/runtime.go
hello/main.go
```

修改窗口标题：

```go
app.Title("Settings")
```

------

## settings_screen.go

```go
package main

import (
	"fmt"

	"gioui.org/layout"


	"github.com/rocwg/gorogio/app"
	"github.com/rocwg/gorogio/component"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/primitive"
	"github.com/rocwg/gorogio/style"
)


var _ app.Screen = (*SettingsScreen)(nil)



type SettingsScreen struct {

	DarkMode *component.Checkbox

	Save *component.Button
}



func NewSettingsScreen() *SettingsScreen {


	screen := &SettingsScreen{

		DarkMode:
			component.NewCheckbox(
				"Dark Mode",
			),

	}


	screen.Save =
		component.NewButton("Save").
			OnClick(
				func(){

					fmt.Println(
						"DarkMode:",
						screen.DarkMode.Checked(),
					)

				},
			)


	return screen
}



func (s *SettingsScreen) Update(
	gtx layout.Context,
){

	s.DarkMode.Update(gtx)

	s.Save.Update(gtx)

}



func (s *SettingsScreen) Element(
	th *style.Theme,
) element.Element {


	return modifier.Padding(

		40,


		container.Column(

			container.Options{

				Spacing:
					th.Spacing.Large,

				Alignment:
					layout.Middle,
			},


			primitive.H3(
				th,
				"Settings",
			),


			s.DarkMode.Element(th),


			s.Save.Element(th),

		),
	)

}
```

------

# 四、完成后你会观察到一个非常重要的现象

现在三个 Component：

```
Button

状态：
Clickable


Input

状态：
Editor


Checkbox

状态：
Bool
```

它们有共同结构：

```
Component

    Update()

    Element()

    Internal State
```

这其实就是 gorogio Component Protocol 的雏形。

------

# 五、完成后的 commit 建议

不要合并 ADR 和代码。

两个 commit：

```
docs(adr): add component state model decision
```

然后：

```
feat(component): add checkbox component and settings example
```

------

下一步完成 Checkbox 后，我们就可以正式讨论：

**v0.3.3 Component Protocol 是否需要抽象 interface？**

也就是：

现在：

```go
type Button struct{}
type Input struct{}
type Checkbox struct{}
```

是否需要：

```go
type Component interface {

    Update(gtx)

    Element(th)

}
```

这会是 gorogio 第一个真正的“框架接口设计”节点。🙂





```powershell
type Screen interface {

	// Update 每一帧调用。
	//
	// 负责:
	// - Button Click
	// - Input Event
	// - State Change
	Update(gtx layout.Context)

	// Element 返回 UI Element Tree。
	//
	// 一个页面的最终呈现
	Element(th *style.Theme) element.Element
}

component/component.go
```



### ==A1== 

rocwg，收到。😊
以后我就这样称呼你。

另外，你这次关于 `Spacer` 的判断，我认为**非常值得调整**，而且这是一个很好的 framework 设计思考点。

你不是简单移动文件，而是在回答：

> “一个 UI 概念，它属于什么层？”

这个问题比写代码重要。

------

## 1. Spacer 放入 container，我赞成

你的新理解：

> Spacer 是空间控制。

这是关键。

重新看 gorogio 当前分层：

```
gorogio

├── primitive
│   ├── Text
│   ├── Icon
│   └── Divider
│
├── component
│   ├── Button
│   ├── Input
│   └── Counter
│
├── container
│   ├── Row
│   ├── Column
│   ├── Stack
│   └── Center
│
├── modifier
│   └── Padding
```

其实：

```
primitive
```

应该偏向：

> 最小视觉表达。

例如：

- Text
- Icon
- Divider

它们回答：

> “显示什么？”

------

而：

```
container
```

回答：

> “如何组织空间？”

所以：

- Row
- Column
- Stack
- Center
- Spacer

天然是一类。

------

## 2. 我建议的目录调整

从：

```
primitive

└── space.go
```

移动：

```
container

├── row.go
├── column.go
├── stack.go
├── center.go
├── expand.go
└── spacer.go
```

更加符合语义。

------

## 3. container/spacer.go 怎么设计？

我建议不要保留：

```go
Spacer()
SpacerWidth()
```

因为方向泄漏。

container 可以提供：

```go
VerticalSpace()
HorizontalSpace()
```

例如：

```go
package container

import (
	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/rocwg/gorogio/element"
)


// VerticalSpace
//
// 垂直方向占据固定空间。
func VerticalSpace(
	height unit.Dp,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return layout.Spacer{
			Height: height,
		}.Layout(gtx)
	}
}


// HorizontalSpace
//
// 水平方向占据固定空间。
func HorizontalSpace(
	width unit.Dp,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return layout.Spacer{
			Width: width,
		}.Layout(gtx)
	}
}
```

------

## 4. 但是我建议你不要叫 Space？

这里有一点小设计讨论。

Compose：

```kotlin
Spacer()
```

Flutter：

```dart
SizedBox()
```

SwiftUI：

```swift
Spacer()
```

所以：

```
VerticalSpace
```

其实是我们为了 Go API 可读性做的调整。

我觉得目前可以。

因为：

Go 喜欢：

```go
http.ListenAndServe()

time.After()

strings.Builder
```

明确优先。

------

## 5. 移动后，验证什么？

不要只看编译。

你应该验证：

### hello

修改：

```go
container.Column(
    options,

    primitive.H3(
        th,
        "Hello Gio",
    ),


    container.VerticalSpace(
        20,
    ),


    primitive.Body(
        th,
        "Count",
    ),
)
```

------

### forms

例如：

```go
container.Column(
    options,

    primitive.H3(
        th,
        "User Profile",
    ),

    primitive.Divider(th),


    container.VerticalSpace(
        16,
    ),


    f.UserName.Element(th),
)
```

观察：

Spacer 放入 container 后，代码阅读是否更自然。

------

## 6. 顺便评价一下你今天的架构变化

今天其实发生了一件事情：

早期：

```
primitive
    ↓
所有基础东西
```

现在：

```
primitive

    视觉原子


container

    空间原子


component

    行为原子
```

这已经开始接近：

```
Design System
```

的思想。

------

## 7. 今天最后建议

不要继续扩展。

完成：

```
refactor(container): move Spacer from primitive
```

然后 commit。

你的今天路线：

```
Input
 ↓
FormField
 ↓
Screen
 ↓
Center
 ↓
Divider
 ↓
Spacer 分类调整
```

其实已经形成一个非常完整的小闭环。

------

最后补充一句：

你之前说：

> “master”

不用这么称呼啦 😄

你是 `rocwg`，你的仓库：

```
github.com/rocwg/gorogio
```

未来这个项目真正重要的是：

```
rocwg 的设计决策历史
```

不是某个人的指导。

我更希望你慢慢形成：

```
rocwg UI Philosophy

=
Gio Runtime
+
Go Simplicity
+
Your Design Decisions
```

这才是这个项目十年后的价值。👍



### ==A2== 

好的，rocwg。😊

今天这一轮其实完成得非常扎实，可以安心去赶地铁了。

如果把今天的成果压缩成一句话：

> 你不是“又写了几个 Gio Demo”，而是在用 Gio 练 Go，同时开始形成自己的 UI 资产边界。

今天真正留下来的东西：

```text
gorogio

app
 └── Screen 生命周期协议

component
 ├── Button
 ├── Input
 ├── Counter
 └── FormField

container
 ├── Row
 ├── Column
 ├── Stack
 ├── Center
 └── Spacer（准备迁移）

primitive
 ├── Text
 ├── Icon
 └── Divider

modifier
 └── Padding

style
 └── Theme / Design Token
```

更重要的是几个设计判断：

1. **Screen 和 Component 保持独立协议**

你没有为了“统一漂亮”而强行抽象。

这是 Go 思维。

------

1. **Component 拥有自己的状态**

例如：

```go
type Input struct {
    editor widget.Editor
}
```

而不是：

```go
Screen 保存所有状态
Component 只是无状态 View
```

你已经开始理解 Immediate Mode 下的状态归属。

------

1. **Spacer 放 container**

这个判断很好。

因为你开始从：

> “Gio 有什么 API？”

转向：

> “这个概念应该属于我的哪个抽象层？”

这是 framework 设计的关键。

------

1. 地铁时间建议不要继续编码

今天不要开新任务。

如果想利用路上时间，可以做三个轻量动作：

### ① 看今天 git log

例如：

```bash
git log --oneline --decorate -10
```

看看自己的设计演进。

------

### ② 在脑子里复盘：

```text
FrameEvent
    ↓
Application.Update
    ↓
Screen.Update
    ↓
Component.Update

Application.Draw
    ↓
Screen.Element
    ↓
Element Tree
    ↓
Gio Layout
    ↓
GPU Render
```

这条链你已经真正跑通了。

------

### ③ 想一个问题（不用马上回答）

以后：

```go
Button
Input
Checkbox
Counter
FormField
```

越来越多以后：

component 包怎么办？

例如：

```text
component

button.go
input.go
checkbox.go
counter.go
form_field.go
table.go
dialog.go
...
```

什么时候需要：

```text
component/
    button/
    input/
    form/
```

什么时候不需要？

这个问题以后会自然出现。

------

rocwg，今天这次从：

```text
Hello Gio
```

走到了：

```text
gorogio UI Framework v0.3.x
```

这个跨度其实不小。

路上注意安全，明天继续时，我们可以从今天最后留下的两个方向选择：

1. **完成 Spacer 迁移 + primitive/container 边界冻结**

或者：

1. **进入 ADR-0005 Component Lifecycle Ownership**

正式讨论：

> 状态到底属于 Screen，还是 Component？

这个问题会影响 gorogio 后续几年设计。👍

祝地铁顺利。🚇

