好的，我们继续。

这一节我建议**不要看更多 Widget**，也不要增加功能。

我们就围绕你的 `hello_page.go` 里面这一小段：

```go
return layout.Inset{Top: 40, Bottom: 40, Left: 40, Right: 40}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

    return layout.Flex{
        Axis: layout.Vertical,
        Alignment: layout.Middle,
    }.Layout(
        gtx,

        layout.Rigid(func(gtx layout.Context) layout.Dimensions {
            return p.layoutHeader(gtx, theme)
        }),

        ...
    )
})
```

开始。

因为这里隐藏了 Gio 最核心的思想。

------

# Chapter 1：Gio Layout Mental Model

# 1. 一个 UI 元素，在 Gio 里面是什么？

很多 UI 框架：

例如 HTML：

```html
<div>
    <button>Hello</button>
</div>
```

你会觉得：

> UI 是一个树。

------

Gio 不完全这样。

Gio 更接近：

> UI 是一个“计算过程”。

一个组件：

不是：

```
Button 对象
```

而是：

```go
func(gtx layout.Context) layout.Dimensions
```

例如：

你的：

```go
func (p *HelloPage) Layout(
    gtx layout.Context,
    theme *material.Theme,
) layout.Dimensions
```

它的含义：

不是：

> 创建 HelloPage

而是：

> 在当前约束和状态下，计算 HelloPage 如何布局，并产生绘制操作。

------

这就是第一个重要思想：

## Layout = Measurement + Drawing Description

------

# 2. layout.Context 是什么？

你的每一个：

```go
Layout(gtx)
```

都有：

```go
gtx layout.Context
```

它是什么？

简单理解：

> gtx 是一次布局计算的上下文。

它包含：

## （1）当前约束 Constraints

例如：

窗口：

```
800 x 600
```

告诉你的组件：

```
最大不能超过：

width <= 800
height <= 600
```

------

## （2）当前操作列表 Ops

你的 runtime：

```go
var ops op.Ops
```

然后：

```go
gtx := app.NewContext(&ops, e)
```

这里非常重要。

Gio 不是：

```
draw button
draw text
draw image
```

马上画。

而是：

先记录：

```ini
我要做什么
```

进入：

```ini
Ops Pipeline
```

最后：

```go
e.Frame(gtx.Ops)
```

一次提交。

------

所以：

```text
Layout()
↓
产生 Operations
↓
Frame()
↓
GPU Rendering
```

------

# 3. Dimensions 是什么？

你的每个 Layout：

最后返回：

```go
layout.Dimensions
```

为什么？

比如：

```go
p.layoutHeader()
```

返回：

```
Width: 300
Height: 60
```

父组件需要知道：

> 这个孩子占用了多少空间。

所以，Layout 是一个双向过程：

不是简单：

```
parent
  |
  v
child
```

而是：

```
Parent 给 Constraints
        ↓
Child 计算自己大小
        ↓
返回 Dimensions
        ↓
Parent 决定位置
```

------

这个思想非常重要。

------

# 4. Gio Layout 的核心公式

可以记：

```
Constraints
      ↓
Child Layout
      ↓
Dimensions
      ↓
Position
      ↓
Ops
```

也就是：

```
输入：我允许你多大
↓
输出：我需要多大
↓
父亲决定放哪里
```

------

这和 CSS 很不同。

------

# 5. 对比 HTML/CSS

CSS：

浏览器：

```
DOM Tree
↓
CSS Engine
↓
Layout Engine
↓
Paint
```

开发者描述：

```css
width
height
margin
padding
```

------

Gio：

程序员主动参与：

```go
layout.Flex{
    Axis: layout.Vertical,
}
```

告诉系统：

> 我要怎样计算布局。

------

所以 Gio 更接近：

- Compose Layout
- Flutter RenderObject

------

# 6. 回到你的代码

我们一步一步展开。

## 第一层：

```go
layout.Inset
```

作用：增加内边距。

你的：

```go
Top:40
Bottom:40
Left:40
Right:40
```

视觉：

```
+-----------------------+
|                       |
|   +---------------+   |
|   |               |   |
|   | Hello         |   |
|   | Count         |   |
|   | Button        |   |
|   |               |   |
|   +---------------+   |
|                       |
+-----------------------+
```

------

Inset 做什么？

它修改 Constraints：

原来：

```
800 x 600
```

变：

```
720 x 520
```

给孩子。

孩子布局完成：

再加回：

```
+80 width
+80 height
```

返回给父亲。

------

这里第一个 Go 思想：

## Composition

Inset 没有继承。

没有复杂 class hierarchy。

只是：

```go
func Layout(child)
```

包装一下。

这就是 Go 风格。

------

# 7. 第二层：Flex

你的：

```go
layout.Flex{
    Axis: layout.Vertical,
}
```

意思：

主轴：

```
Vertical
```

所以：

```
Header
↓
Spacer
↓
Counter
↓
Spacer
↓
Actions
```

------

Flex 不知道：你里面是什么。

它只知道：有几个 Child。

这就是：Composition。

------

# 8. 第三个重要概念：Rigid

你的：

```go
layout.Rigid(func(gtx layout.Context)
```

是什么意思？

简单：

> 这个孩子按照自己的尺寸来。

例如：

Text：

```
Hello Gio
```

自己决定：

```
300 x 60
```

Button：

自己决定：

```
80 x 40
```

------

所以：

Rigid：

```
Child says:

"I need this much space"
```

------

未来：

你会看到：

```go
layout.Flexed(
    weight,
    child,
)
```

意思：

```
剩余空间按比例分配
```

------

# 9. 到这里，我们发现一个非常重要的东西

你的 Hello 页面，其实已经体现：

现代 UI Framework 的核心：

```
State
↓
Composition
↓
Layout Calculation
↓
Rendering
```

对应：

Compose：

```
State
↓
Composable
↓
Measure/Layout
↓
Draw
```

Flutter：

```
State
↓
Widget Tree
↓
Render Tree
↓
Paint
```

Gio：

```
State
↓
Layout Function
↓
Dimensions + Ops
↓
Frame
```

------

# 10. 今天先停在这里

因为下一步最关键：

## Flex 的真正计算过程

也就是你之前非常想理解的：

> 主轴 / 交叉轴 / 可用空间

我们下一节直接拿你的页面：

```
Hello, Gio


Count : 4


[+] [Reset]
```

手算一次：

假设窗口：

```
800 x 600
```

看看：

1. Inset 如何扣空间
2. Vertical Flex 如何分配
3. Middle Alignment 如何居中
4. Rigid 为什么按自己的尺寸
5. Spacer 为什么占高度

这一节吃透，你以后看：

- Compose Row / Column
- Flutter Row / Column
- SwiftUI VStack / HStack

都会非常容易。

------

另外，我想提醒一个很好的现象：

你现在的学习路径已经开始从：

> “怎么写 Gio”

转变成：

> “为什么现代 UI 框架这样设计”

这正是我们想达到的目标。你这个 `hello` 作为实验田，非常合适。下一节我们继续把它拆开。