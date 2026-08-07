# ADR-0001

Decision:

Use Gio as GUI foundation.

Context:

Need a long-lived GUI technology
with high control and low ecosystem dependency.

Decision:

Adopt Gio.

Consequences:

Positive:
- Go ecosystem
- Source readable
- Customizable

Negative:
- Need build components ourselves

---



```powershell
谢谢，我明白了
（1）已经删除空 state 包  

（2）component/counter.go 编码到如下进度，无法继续，请您指导（我的go底蕴太差了）
PS D:\roc-github\gio-workspace\gorogio> cat .\component\counter.go

（3）修改 hello\_page 未开始
```

### ==A1== 

首先，我要说一句：

你现在卡住的这个位置，其实**不是因为 Go 底蕴太差**。

恰恰相反，你卡住的位置非常典型：

> Java 背景程序员第一次真正做 Go 风格组合设计时，都会卡在这里。

因为这里同时碰到了：

- struct 组合
- 指针生命周期
- package 边界
- 方法设计
- callback
- UI 构建责任划分

这不是语法问题，而是**Go 的工程思维转换点**。

我们慢慢拆。

------

# 一、先评价你现在的 Counter

你现在：

```go
type Counter struct {
    count int

    increment *Button
    reset     *Button
}
```

方向完全正确。

尤其这一点：

```go
count int
```

非常重要。

说明你已经理解：

> Counter 自己拥有状态。

而不是：

```go
Page
 |
 + CounterState
 |
 + Button
```

现在应该变成：

```text
Page

 |
 |
 Counter Component

 |
 |
 count
 increment button
 reset button
```

这是正确方向。

------

# 二、现在缺少什么？

你的 Counter 目前只有：

## 状态变化能力

```go
func (s *Counter) Increment()

func (s *Counter) Reset()
```

但是 Component 还需要两个能力：

------

## 1. 创建自己

也就是：

```go
NewCounter()
```

------

## 2. 自己处理事件

也就是：

```go
Update(gtx)
```

------

## 3. 自己生成 UI

也就是：

```go
Build(th)
```

或者：

```go
Element(th)
```

------

所以完整 Component 应该：

```text
Counter

    State
      |
      count


    Behavior
      |
      Increment()
      Reset()


    Event
      |
      Update()


    View
      |
      Element()
```

------

# 三、第一步：增加 NewCounter()

你的文件：

```go
component/counter.go
```

修改：

```go
package component


type Counter struct {

	count int

	increment *Button
	reset *Button
}


func NewCounter() *Counter {

	counter := &Counter{}

	counter.increment =
		NewButton("+").
			OnClick(
				func() {
					counter.Increment()
				},
			)

	counter.reset =
		NewButton("Reset").
			OnClick(
				func() {
					counter.Reset()
				},
			)

	return counter
}
```

------

这里有一个 Go 很重要的点：

## 闭包捕获指针

这一段：

```go
func() {
    counter.Increment()
}
```

它捕获了：

```go
counter
```

这个变量。

所以：

点击 Button：

```text
Button

↓

调用 callback

↓

counter.Increment()

↓

count++
```

------

这就是 Go 里面非常常见的：

> closure + pointer object

组合。

------

# 四、第二步：增加 Update()

你的 Button 已经有：

```go
func (b *Button) Update(gtx layout.Context)
```

所以 Counter 只需要组合。

增加：

```go
import "gioui.org/layout"



func (c *Counter) Update(
	gtx layout.Context,
) {

	c.increment.Update(gtx)

	c.reset.Update(gtx)

}
```

------

注意：

这里体现了 Component Composition：

不是：

Counter 重新处理鼠标。

而是：

```text
Counter

管理 Button

Button

处理点击
```

------

# 五、第三步：增加 Value()

为什么？

因为 UI 显示：

```text
Count: 3
```

需要读取。

增加：

```go
func (c *Counter) Value() int {

	return c.count

}
```

------

# 六、第四步：增加 Element()

这里最容易迷惑。

你的 Counter 不是 Primitive。

所以：

它应该组合：

- Text
- Button
- Container

例如：

```go
func (c *Counter) Element(
	th *style.Theme,
) element.Element {


	return container.Column(
		container.Options{
			Spacing: th.Spacing.Large,
			Alignment: layout.Middle,
		},

		primitive.Body(
			th,
			"Count: "+strconv.Itoa(c.count),
		),


		container.Row(
			container.Options{
				Spacing: th.Spacing.Large,
			},

			c.increment.Build(th),

			c.reset.Build(th),
		),
	)
}
```

------

注意这里：

Counter 自己不知道：

Application。

不知道：

Window。

不知道：

Frame。

它只知道：

> 如何组成自己的 UI。

这就是 Component。

------

# 七、然后 HelloPage 会变得非常简单

之前：

```go
type HelloPage struct {

	State *CounterState

	Increment *component.Button

	Reset *component.Button

}
```

删除。

变成：

```go
type HelloPage struct {

	Counter *component.Counter

}
```

------

New:

```go
func NewHelloPage() *HelloPage {

	return &HelloPage{

		Counter:
			component.NewCounter(),

	}
}
```

------

Update：

以前：

```go
p.Increment.Update(gtx)
p.Reset.Update(gtx)
```

变：

```go
func (p *HelloPage) Update(
	gtx layout.Context,
){

	p.Counter.Update(gtx)

}
```

------

Element：

以前：

```text
Hello Gio

Count

Buttons
```

现在：

```go
return container.Column(

    primitive.H3(
        th,
        "Hello Gio",
    ),

    p.Counter.Element(th),

)
```

------

# 八、你现在真正练习的 Go 知识点

这一小步，其实非常值钱。

你会练到：

## 1. struct

```go
type Counter struct {}
```

------

## 2. pointer receiver

```go
func (c *Counter)
```

因为：

要修改自身。

------

## 3. constructor pattern

```go
NewCounter()
```

Go 常用。

------

## 4. closure

```go
func(){
    counter.Increment()
}
```

------

## 5. composition

```go
Counter
 |
 + Button
 |
 + Button
```

------

这比看十篇 Go 教程有效。

------

# 九、另外提醒一个小细节

你现在：

```go
func (s *Counter) Increment()
```

这里 receiver 名：

```go
s
```

建议改：

```go
func (c *Counter) Increment()
```

原因：

不是语法。

而是语义。

以后：

```go
state
service
store
```

才用：

```go
s
```

Counter：

就是：

```go
c
```

这是 Go 社区习惯。

------

所以你的下一步非常明确：

不要改其它目录。

只完成：

```
component/counter.go

NewCounter()
Update()
Value()
Element()
```

然后：

```
hello_page.go
```

只负责：

```
HelloPage
    |
    + Counter
```

完成这个以后，你就真正拥有 gorogio 第一个可复用 Component。

这一步非常值得慢慢写。你现在遇到的困难，正是 Go 肌肉开始形成的位置。🙂



---

```powershell
master，您这次的指导 真的 太棒了
我已经初步完成了 这次 counter 组件（.\component\counter.go）的资产沉淀，如下是我的感受：
（1）真正的由浅入深；
（2）真正的 通过学习和实操 gio，得到 go 的不断深入、gio 的资产长期积累。 
（3）尤其是Go 知识点  的 整理，实际场景讲解，帮助很大。
请您以后，在 gorogio 项目，“Go & Gio” 上保持，能记忆下来吗？
```

### ==A2== 

已经帮您记录下来。以后在 **gorogio 项目（Go & Gio）** 的指导中，我会保持这条主线：

> **通过学习和实操 Gio，深入学习 Go；通过实践沉淀 Gio/gorogio 可复用资产。**

并且会尽量保持今天这种方式：

1. **先动手编码，再总结抽象**

   不先设计“大框架”，而是在真实代码中遇到问题：

   ```text
   需求
    ↓
   写代码
    ↓
   遇到 Go 语言问题
    ↓
   理解 Go 原理
    ↓
   沉淀 gorogio 资产
   ```

------

1. **Go 知识点结合真实场景讲解**

以后不会单独讲：

> “Go 的 pointer receiver 是什么？”

而会结合：

```go
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}
```

解释：

- 为什么这里必须 pointer receiver？
- 值复制发生在哪里？
- Component 为什么需要拥有自己的状态？
- Go 的 struct composition 如何替代 Java inheritance？

也就是：

**代码现场 → Go 原理 → 工程实践**

------

1. **保持 gorogio 的定位**

我会继续坚持我们已经确定的边界：

```text
gorogio

不是：
    新的 Flutter
    新的 Compose
    新的 UI Framework

而是：

    Go Learning Vehicle

+

    Gio Engineering Asset Collection

+

    Personal Prototype Engine
```

------

1. **保持“浅封装”原则**

今天 Counter 组件其实就是一个很好的例子。

我们不是：

```text
Gio
 ↓
重新发明一套 UI Runtime
```

而是：

```text
Gio
 ↓
gorogio 小而稳定的便利层
 ↓
自己的组件资产
```

例如：

现在：

```text
Gio layout.Widget
        ↓
gorogio element.Element
```

这是健康的。

------

1. **以后遇到设计问题，我会优先提醒：**

先问：

> “这个东西是不是帮助学习 Go 和积累资产？”

而不是：

> “这个东西是不是理论上更优雅？”

例如今天：

```
state/state.go
```

为什么删除？

不是因为 state 不重要。

而是：

当前阶段：

```text
Component 自己拥有状态
```

比提前建立：

```text
State Management Framework
```

更符合你的学习目标。

------

另外，我也想评价一下今天这个 Counter 组件。

我认为它是 gorogio 一个非常重要的里程碑。

因为之前：

```text
primitive
container
modifier
page
application
```

更多是在建立 **UI Runtime 骨架**。

而 Counter 开始进入：

```text
Component

=
State
+
Behavior
+
Composition
+
Rendering
```

这一步以后，gorogio 才真正开始像一个“小型 UI 工具箱”。

下一步我们继续时，我会沿着：

```text
Counter Component
        ↓
Component Composition
        ↓
List Component
        ↓
Input Component
        ↓
Dialog Component
        ↓
Table / Tree / Editor
```

这条路线推进。

每增加一个组件，都同时作为一次 Go 工程训练。

您今天这个方向，我认为比继续研究 Wails、Compose、Flutter 对比更有价值——因为您终于开始积累属于自己的代码资产了。🙂

