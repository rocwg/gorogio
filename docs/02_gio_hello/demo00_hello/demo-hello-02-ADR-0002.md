1）我们一次聊天就直接完成这个冻结；

2）就继续开始继续 编码吧；请问，我重名目录名称，您觉得如何？

```powershell
PS D:\roc-github\gio-workspace\gorogio\examples> goro tree -depth=2
.
├── hello-puregio（老的 hello，纯粹的 gio 使用） 
│   ├── application.go
│   ├── hello_page.go
│   ├── main.go
│   ├── runtime.go
│   └── state.go
└── hello-gorogio (昨天您指导的代码 myhello，成功消费 Element Tree)
    ├── application.go
    ├── hello_page.go
    ├── main.go
    ├── runtime.go
    └── state.go
```

# A1

roc，我昨天想了一晚上（准确说，是一直在回顾我们这两周的路线）。

我觉得，我们应该在今天做一个**真正重要的决定**。

> **从今天开始，gorogio 不再是一个 Gio 的封装。**
>
> **它是一套 Go 风格的 Declarative UI Framework。**

这是两个完全不同的定位。

------

# 第一部分：今天，我们正式冻结 gorogio v0.1

我建议今天不要冻结代码。

**冻结理念。**

以后所有代码，都必须符合这份 ADR。

------

# ADR-0001：gorogio Design Principles（v0.1）

## Principle 1：Honor Go, Learn from Others

> 尊重 Go，学习其他框架。

学习：

- Compose
- SwiftUI
- Flutter
- React

但是：

**绝不复制它们的语法。**

学习的是：

- 声明式 UI
- UI Tree
- 单向数据流
- 状态驱动

而不是：

```swift
.padding()
.background()
.foreground()
```

------

## Principle 2：Borrow Ideas, Not Syntax

例如：

采用：

```go
container.Column(...)
```

而不是：

```go
VStack(...)
```

采用：

```go
modifier.Padding(...)
```

而不是：

```go
Text(...).Padding()
```

采用：

```go
component.Button(...)
```

而不是：

```go
ElevatedButton(...)
```

------

## Principle 3：Everything is an Element

整个框架只有一种 UI 协议：

```go
type Element = layout.Widget
```

以后，Text、Button、Image、Card、Column 全部都是：

```go
element.Element
```

这样：

UI Tree 永远只有一种节点。

------

## Principle 4：Container、Modifier、Component 职责永远分离

|             |              |                                           |
| ----------- | ------------ | ----------------------------------------- |
| Container： | 负责：布局。 | 例如：Row、Column、Stack                  |
| Modifier：  | 负责：修饰。 | 例如：Padding、Background、Border、Shadow |
| Component： | 负责：交互。 | 例如：Button、TextField、Checkbox、Slider |

以后不要混。

------

## Principle 5：Application 与业务分离

业务永远不应该知道：

```go
app.Main()

Window

FrameEvent

Ops
```

业务只负责：

```go
func HomePage() element.Element
```

Runtime：

属于：gorogio。

------

## Principle 6：Prefer Explicit Composition

这是我觉得最重要的一条。

Go 一直强调：

> **Explicit is better than magic.**

所以：

以后：

永远不要：

```go
Text().
Padding().
Color().
Shadow()
```

而是：

```go
modifier.Padding(

    16,

    modifier.Background(

        style.Primary,

        view.Text(
            "Hello",
        ),
    ),
)
```

虽然括号多。

但是：**非常 Go。**

------

## Principle 7：Small Packages

以后，每个包 尽量只有一个职责。

例如：

```text
view

container

modifier

component

style

element
```

不要：

```text
widget
```

里面放两百个文件。

------

# 第二部分：目录，我想调整一点点

你的：

```text
examples/
├── hello-puregio
└── hello-gorogio
```

我觉得已经比：

```text
hello

myhello
```

好多了。

因为，一眼知道 谁是谁。

------

但是我想再进一步。

因为：

以后：

可能：

```text
hello-gorogio-v2
```

怎么办？

所以：

我建议：

```text
examples/

├── gio
│   └── hello
│
└── gorogio
    └── hello
```

以后，如果有 counter、todo、list 全部：

```text
examples/
    gio/
        hello
        counter
        todo

    gorogio/
        hello
        counter
        todo
```

未来，甚至：

```text
examples/

compose/

swiftui/
```

都可以加入。

这时候：

它就是：**不同 Framework 的对照实验。**

我觉得比：

```text
hello-puregio
```

扩展性更好。

------

# 第三部分：接下来编码什么？

昨天我说 Button。

今天我改了。

**Button 先暂停。**

为什么？

因为：昨天我们已经有 Button。

但是：还有一个更重要的问题没解决。

------

## Theme 泄漏。

你现在：

```go
view.Text()
```

里面：

```go
theme := material.NewTheme()
```

这是：

**最大的设计问题。**

为什么？

因为：

以后：

Dark Theme

Custom Font

Primary Color

全部失效。

------

所以：

我建议：

**下一步，不做 Button。**

做：

```text
Theme Context
```

例如：

以后：

```go
view.Text(
    "Hello",
)
```

内部：

不用：

```go
material.NewTheme()
```

而是：自动获得

```go
CurrentTheme()
```

当然，我们不会照搬 Compose 的 `CompositionLocal`。

我们会用 Go 最自然的方式来做。

------

# 第四部分：Go 1.27 对 gorogio 的影响

最后回答你前面那个问题。

你问：

> Go 越来越成熟，我怎么办？

我今天给你的答案，比昨天更明确。

**不要跟 Go 的每一个新特性。**

而是：

建立自己的：

```text
docs/

engineering/

    go-style.md
```

例如：

这里规定：

> gorogio 默认不用 Iterator。

> 默认不用链式 DSL。

> 默认继续 value, err。

> 默认继续普通 for range。

以后：如果 `Go 1.29`、`Go 1.30` 出了新东西。

不是："社区怎么写？"

而是：

> **"这符合 gorogio Design Principles 吗？"**

如果符合：采用。

如果不符合：不用。

------

## 最后，我想给 gorogio 一个我认为特别适合它的定位。

昨天我们一直说：

> UI Framework。

今天我觉得可以更准确一点。

我希望未来别人打开你的仓库，第一页看到的是：

> **gorogio is a Go-native Declarative UI Framework.**

注意：

不是：

> A SwiftUI for Go.

也不是：

> A Compose for Go.

而是：

> **Go-native。**

这是我今天最想和你一起冻结的一个词。

它意味着：

**它首先是 Go，其次才是声明式 UI。**

------

所以，我建议今天就把这份 ADR 冻结。从下一次开始，我们进入真正的编码阶段，而且每新增一个 API，都先问一个问题：

> **"如果这是 Go 标准库的作者来设计，他会接受这个 API 吗？"**

我觉得，这会成为 `gorogio` 最有辨识度的地方。