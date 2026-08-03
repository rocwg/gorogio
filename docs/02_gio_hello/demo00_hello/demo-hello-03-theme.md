谢谢，冻结的好；我们进入真正的编码阶段，而且每新增一个 API，都先问一个问题：

"如果这是 Go 标准库的作者来设计，他会接受这个 API 吗？"

第一件事 是不是 自动获得 CurrentTheme()，且用 Go 最自然的方式来做。

# A2

roc，我认为**是，但不是 `CurrentTheme()`。**

这是今天我们第一次要遵守刚刚冻结的 ADR。

> **如果这是 Go 标准库作者设计，他会接受这个 API 吗？**

我的答案是：

> **`CurrentTheme()` 我不会接受。**

------

# 为什么？

因为它有一点像：

```java
SecurityContextHolder.getContext()

LocaleContextHolder.getLocale()

RequestContextHolder...
```

或者：

Compose：

```kotlin
MaterialTheme.current
```

React：

```jsx
useContext()
```

这些都是：

> **全局 Context。**

Go 标准库其实一直不喜欢这种设计。

Go 更喜欢：

> **依赖显式传递。**

例如：

```go
http.Request.Context()

context.WithValue()

sql.DB

template.Execute(...)
```

几乎没有：

```go
CurrentXXX()
```

------

# 我建议我们先问一个更根本的问题

Theme 属于什么？

我觉得：

> Theme 不是全局变量。

Theme 是：

> **Render Environment（渲染环境）**

也就是说：

今天我们其实缺的不是：

```go
CurrentTheme()
```

而是：

```go
UI Context
```

------

# 我建议重新设计 Element

现在：

```go
type Element = layout.Widget
```

其实就是：

```go
func(layout.Context) layout.Dimensions
```

这里只有：

```text
Layout Context
```

但是没有：

```text
Theme
```

------

# 我更喜欢这样（未来方向）

例如：

```go
type Context struct {

    Layout layout.Context

    Theme *style.Theme

}
```

然后：

```go
type Element func(
    Context,
) layout.Dimensions
```

以后：

Text：

```go
func Text(
    value string,
) element.Element
```

内部：

```go
func(gtx element.Context) layout.Dimensions {

    return material.Body1(

        gtx.Theme.Material,

        value,

    ).Layout(

        gtx.Layout,

    )

}
```

注意。

**没有 CurrentTheme。**

------

# 但是……

这里有一个问题。

这样意味着：

我们不能再：

```go
type Element = layout.Widget
```

而要变成：

```go
type Element func(
    Context,
) layout.Dimensions
```

这意味着：

今天：

所有：

```go
layout.Rigid(child)
```

都会失效。

因为：

Gio：

只认识：

```go
layout.Widget
```

------

# 所以，我反而建议今天不要改 Element。

这是我今天最大的建议。

> **v0.1 保持 Element = layout.Widget。**

为什么？

因为：

我们昨天刚刚完成：

> 消费 Element Tree。

如果今天：

改：

Element。

我们几乎整个框架重写。

我觉得：

太早了。

------

# 那 Theme 怎么办？

这里我想到了一个我认为非常 Go 的方案。

它也是 Gio 自己很多地方采用的思想。

就是：

**Theme Injection（主题注入）**

例如：

不要：

```go
view.Text(
    "Hello",
)
```

而是：

```go
view.Text(

    th,

    "Hello",

)
```

是不是有点丑？

是。

但是：

非常 Go。

------

例如：

Button：

今天：

```go
component.NewButton(

    th.Material,

    "OK",

)
```

其实就是：

Theme Injection。

------

Text：

改：

```go
func Text(

    th *style.Theme,

    value string,

) element.Element
```

里面：

直接：

```go
material.Body1(

    th.Material,

    value,

)
```

结束。

------

# 但是我觉得还能更进一步

昨天我就在想：

为什么：

Compose：

```kotlin
Text("Hello")
```

Go：

却要：

```go
view.Text(

    theme,

    "Hello",

)
```

因为：

Compose：

有：

CompositionLocal。

Go：

没有。

那么：

Go 有没有：

更自然的方法？

我觉得：

有。

------

# 就是 Page。

例如：

以后：

```go
type HelloPage struct {

    Theme *style.Theme

}
```

然后：

```go
func (p *HelloPage) Text(

    value string,

) element.Element {

    return view.Text(

        p.Theme,

        value,

    )

}
```

于是：

页面：

```go
return container.Column(

    p.Text("Hello"),

    p.Text("World"),

)
```

业务：

永远：

不用：

Theme。

------

# 我为什么喜欢这个？

因为：

Theme 本来就是：

**Page Environment。**

不是：

全局 Environment。

------

# 所以，我建议今天冻结一个新的原则

> **Theme 不主动获取（CurrentTheme），而是由上层注入（Inject）。**

这是 Go 的风格。

------

# 我建议我们的路线改成这样

今天不要做：

```text
CurrentTheme()
```

而是：

第一步：

✅ 修复：

```go
view.Text()

component.Button()
```

全部：

改成：

Theme Injection。

例如：

```go
view.Text(

    theme,

    "Hello",

)
```

------

第二步：

Page：

封装：

```go
p.Text()

p.Button()
```

业务：

再也看不到：

Theme。

------

第三步（以后）

等：

gorogio v0.5

我们真的觉得：

Theme 到处传太麻烦。

再考虑：

是不是需要：

UI Context。

------

## 所以，我今天想否决 `CurrentTheme()`，不是因为它做不到，而是因为它**太早了**。

我们刚刚冻结了一条原则：

> **"如果这是 Go 标准库作者设计，他会接受这个 API 吗？"**

我觉得，Go 标准库作者更可能接受：

```go
view.Text(theme, "Hello")
```

而不是：

```go
view.Text("Hello") // 内部偷偷 CurrentTheme()
```

因为前者的依赖关系是**显式**的，后者则引入了隐藏状态。

------

**所以，我建议我们今天真正开始编码的第一件事，不是实现 `CurrentTheme()`，而是完成整个 `Theme Injection` 链路。**

这是一个小改动，却会让整个 `gorogio` 的设计更稳，也更符合我们刚刚冻结的 **Go-native Declarative UI Framework** 的方向。



# A3

您是如下意思吗？
（1）当前阶段，使用 Theme Injection（主题注入），并接受显式依赖 view.Text(theme, "Hello")；
（2）且，始终保持 由上层注入（Inject）是吗？
gio-x 是如何设计的？我应该怎么选择？



roc，这次我想非常严谨地回答，因为**这是 gorogio 第一个真正的架构决策（Architecture Decision）**。

先回答你的两个问题。

> **（1）当前阶段，使用 Theme Injection（主题注入），并接受显式依赖 `view.Text(theme, "Hello")`；**
>
> **（2）且，始终保持由上层注入（Inject）是吗？**

**我的答案是：第一句对，第二句我想修正一下。**

不是：

> **始终保持 Theme Injection。**

而是：

> **始终保持依赖显式（Explicit Dependency）。**

这是两个不同层次。

------

# 第一层：原则（十年不变）

我希望 gorogio 冻结的是：

> **依赖必须是显式的（Explicit Dependency）。**

例如：

```go
view.Text(...)
```

它需要 Theme。

Theme 从哪里来？

一定能够顺着调用链找到。

不能：

```go
CurrentTheme()
```

不能：

```go
GlobalTheme
```

不能：

```go
Singleton
```

这条原则，我希望十年都不要改。

------

# 第二层：实现（可以变化）

**Theme 如何传递？**

这是实现。

它可以演进。

例如：

------

## v0.1

直接传：

```go
view.Text(
    theme,
    "Hello",
)
```

这是最简单、最 Go。

------

## v0.2

Page 帮你包装：

```go
p.Text("Hello")
```

里面：

```go
return view.Text(
    p.Theme,
    value,
)
```

Theme 还是显式注入。

只是业务不用关心。

------

## v0.5

如果以后真的有：

```text
Window
    ↓
Page
        ↓
Dialog
            ↓
Popup
                ↓
Button
```

Theme 已经要传五层。

这时候再考虑：

```go
UIContext
```

例如：

```go
type UIContext struct {

    Layout layout.Context

    Theme *Theme

}
```

依赖依然是显式。

只是：

Theme 放进 Context。

------

所以：

**我们冻结的是原则。**

不是：

Theme Injection 这个具体实现。

------

# 这就是 Go 标准库一直干的事情。

例如：

最开始：

```go
func Copy(
    dst Writer,
    src Reader,
)
```

后来：

增加：

```go
CopyBuffer(...)
```

原则没变。

实现不断升级。

------

# 下面回答 gio-x

这个问题昨天我也去重新想了一遍。

其实：

gio-x 的设计目标，

和你现在做的 gorogio，

**完全不同。**

------

## gio

定位：

Low-level UI Toolkit。

提供：

```text
layout

widget

material

op

paint
```

------

## gio-x

定位：

Extra Widgets。

例如：

```text
calendar

colorpicker

richtext

markdown

flow
```

它没有重新设计 Gio。

只是：

> **补组件。**

所以：

gio-x 很少去做：

```text
Application

Runtime

Theme System

Pattern Language
```

------

# 而你现在做的是什么？

我觉得已经不是：

> gio-x。

而是：

更像：

Compose Runtime。

或者：

SwiftUI Framework。

例如：

你已经有：

```text
element

container

modifier

view

component

style

app
```

这些都不是：

Extra Widget。

而是：

**Framework Layer。**

------

# 所以我反而建议：

以后：

**不要参考 gio-x 的架构。**

为什么？

因为：

目标不同。

------

我更建议：

分别学习三样东西。

------

## 第一学习对象

Gio。

学习：

```text
Layout

Constraint

Ops

Widget
```

这是底层。

------

## 第二学习对象

Compose。

学习：

```text
UI Tree

State

Composition

Remember

Navigation
```

这是框架思想。

------

## 第三学习对象

Go 标准库。

学习：

API Taste。

例如：

为什么：

```go
io.Copy()
```

长这样？

为什么：

```go
http.Server
```

长这样？

为什么：

```go
template.Execute()
```

长这样？

**这才是 gorogio 最应该学习的对象。**

------

# 我昨天晚上突然想到一句话。

我觉得以后我们所有 API 都应该过三关。

以后每新增一个 API。

都问：

------

## 第一关

如果这是：

Go 标准库作者。

会不会接受？

------

## 第二关

如果这是：

Compose 作者。

会不会觉得：

Declarative。

------

## 第三关

如果这是：

一个 Go 新手。

第一次看到。

能不能猜出来？

例如：

```go
container.Column()
```

不用文档。

知道：

就是：

竖着排。

------

如果：

三个都是：

Yes。

这个 API：

基本就是好 API。

------

# 最后，我想把今天的第一个编码任务再收敛一点。

我建议：

**今天不要急着设计 UIContext。**

也不要实现 CurrentTheme。

我们先完成一个非常小、但非常重要的重构。

## 当前

```go
view.Text() {

    material.NewTheme()

}
```

↓

## 第一步（v0.1）

```go
view.Text(

    theme,

    "Hello",

)
```

↓

## 第二步

```go
page.Text(

    "Hello",

)
```

↓

以后再考虑：

UIContext。

------

## 我最后再补一句，也是我今天最大的建议。

昨天我们冻结了一句话：

> **Honor Go, Learn from Others.**

今天，我想补充它的下一句。

> **Freeze Principles, Evolve Implementations.**

中文就是：

> **冻结原则，演进实现。**

对于 `gorogio` 来说：

- **依赖显式**：冻结。
- **Theme 怎么传**：可以演进。
- **Element 是 UI Tree 的唯一节点**：冻结。
- **内部如何优化 Runtime**：可以演进。

我认为，这种分层思维，会让 `gorogio` 在未来几年里既保持稳定的 API，又有足够的演进空间。

