master 收到。🙂

这一步非常关键。

因为 **ADR-0005 解决的是“生命周期属于谁”**：

> Component 可以拥有自己的交互生命周期。

而 **ADR-0006 进一步冻结“状态模型”**：

> 什么状态应该留在 Component，什么状态应该提升到 Screen。

我们先冻结设计，再编码。

------

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