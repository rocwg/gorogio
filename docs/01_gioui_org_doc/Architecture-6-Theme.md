**第6章：Theme（主题）**

**原文地址**：https://gioui.org/doc/architecture/theme

---

### 完整中文翻译

# Theme

# Theme 让东西看起来一致

同一个抽象控件可以有许多视觉表现，从简单的颜色变化到完全自定义的图形。为了让应用拥有一致的外观，有一个代表特定“主题”的抽象是很有用的。

包 [gioui.org/widget/material](https://gioui.org/widget/material) 实现了一个基于 [Material Design](https://material.io/design) 的主题，[Theme](https://gioui.org/widget/material#Theme) 结构体封装了用于改变颜色、尺寸和字体的参数。

要使用主题，你必须先在应用循环中初始化它：

```go
th := material.NewTheme()
th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

var window app.Window
window.Option(app.Title(title))

var ops op.Ops
for {
	switch e := window.Event().(type) {
	case app.DestroyEvent:
		// 窗口被关闭了。
		return e.Err
	case app.FrameEvent:
		// 为新帧重置 layout.Context。
		gtx := app.NewContext(&ops, e)

		// 根据 e.Queue 中的事件，把状态绘制进 ops。
		draw(gtx, th)

		// 更新显示。
		e.Frame(gtx.Ops)
	}
}
```

然后在你的应用中使用提供的控件：

```go
var isChecked widget.Bool

func themedApplication(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var checkboxLabel string
	isChecked.Update(gtx)
	if isChecked.Value {
		checkboxLabel = "checked"
	} else {
		checkboxLabel = "not-checked"
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(material.H3(th, "Hello, World!").Layout),
		layout.Rigid(material.CheckBox(th, &isChecked, checkboxLabel).Layout),
	)
}
```

[Kitchen 示例](https://git.sr.ht/~eliasnaur/gio-example/tree/main/example/kitchen/kitchen.go) 展示了所有可用的不同控件。

---

### 深度解读

**1. 主题解决的核心问题**

在前面的章节里，我们自己写按钮时，颜色、圆角、字体大小、间距全都是硬编码的。  
如果整个应用有几十个按钮、输入框、列表项，这些数值散落在各处，想统一改成暗色主题或换一套设计语言，就会变成灾难。

Theme 的作用就是把这些“视觉参数”集中管理：

- 主色、强调色、背景色、文字颜色
- 字体家族、字号层级（H1~H6、Body、Caption……）
- 圆角半径、间距、阴影、动画时长等

**2. 状态与视觉彻底分离（Gio 的重要设计）**

这是 Gio 最精妙的地方之一：

| 层级       | 负责什么                | 例子                                                      |
| ---------- | ----------------------- | --------------------------------------------------------- |
| **状态层** | 用户交互产生的数据      | `widget.Bool`、`widget.Clickable`、`widget.Editor`        |
| **视觉层** | 怎么画、用什么颜色/字体 | `material.CheckBox`、`material.Button`、`material.Editor` |

`widget.Bool` 只知道“我现在是 true 还是 false”，完全不知道自己长什么样。  
`material.CheckBox(th, &isChecked, label)` 才决定“用 Material Design 的风格画一个复选框”。

好处非常明显：
- 换主题 = 换一套视觉实现，状态代码一行都不用改
- 可以同时存在多套主题（Material、自定义、暗色、高对比度……）
- 第三方可以发布完全不同的设计系统包

**3. 初始化时必须设置 Shaper**

```go
th := material.NewTheme()
th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
```

`Shaper` 是文字整形器（负责把字符串变成可绘制的字形路径）。  
如果不设置，所有文字相关的 material 控件都会无法正常显示。

官方默认推荐使用 Go 字体（`gofont.Collection()`），它体积小、支持多语言，足够大多数应用使用。需要其他字体时，可以用 `opentype` 包加载。

**4. 实际使用模式**

标准写法是把 `*material.Theme` 作为参数一路传递：

```go
func draw(gtx layout.Context, th *material.Theme) layout.Dimensions {
    return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
        layout.Rigid(material.H3(th, "标题").Layout),
        layout.Rigid(material.Body1(th, "正文内容").Layout),
        layout.Rigid(material.Button(th, &btn, "点击我").Layout),
        // ...
    )
}
```

或者把 Theme 放进自己的应用结构体里：

```go
type App struct {
    Theme *material.Theme
    // 其他状态...
}
```

**5. material 包提供了哪些现成控件？**

官方 Kitchen 示例几乎展示了全部，常见的有：

- 文字：`H1` ~ `H6`、`Body1`、`Body2`、`Caption`、`Label`
- 按钮：`Button`、`IconButton`、`TextButton`
- 输入：`Editor`、`CheckBox`、`RadioButton`、`Switch`
- 反馈：`ProgressBar`、`ProgressCircle`、`Slider`
- 其他：`List`、`Card`、`Divider`、`Loader` 等

这些控件内部都使用了第5章的布局原语（Flex、Inset、Stack 等）和第2、3章的绘制与输入机制。

**6. 如何自定义主题？**

`material.Theme` 是一个普通结构体，里面有很多可导出的字段（颜色、字号等）。你可以：

1. 直接修改字段（最简单）
2. 创建自己的 Theme 结构体，实现相同的方法签名
3. 完全抛弃 material 包，自己写一套视觉系统（很多生产项目会这样做）

**7. 与前后章节的关系**

- 第4章定义了 Widget 的约定（Layout 方法）
- 第5章提供了组合 Widget 的布局工具
- **第6章提供了“开箱即用”的、风格统一的 Widget 实现**
- 第8章 Text 会更深入讲解字体与整形
- 第9章 Color 解释了颜色空间为什么重要

Theme 是让普通开发者能快速做出“看起来像样”的应用的关键一层。没有它，你就得自己从像素级别开始堆所有视觉细节。

**8. 实际建议**

| 场景            | 建议                                             |
| --------------- | ------------------------------------------------ |
| 快速原型 / 学习 | 直接用 `material.NewTheme()`                     |
| 正式产品        | 在 material 基础上调整颜色和字号，或完全自研主题 |
| 多主题切换      | 准备多个 `*material.Theme` 实例，按需切换        |
| 文字显示异常    | 第一时间检查 `th.Shaper` 是否正确设置            |
| 想看所有控件    | 运行官方 Kitchen 示例                            |

---

**本章核心记忆点**

> **Theme = 集中管理视觉参数（颜色、字体、尺寸）**  
> **状态（widget.XXX）与视觉（material.XXX）彻底分离**  
> **必须设置 Shaper，否则文字无法显示**  
> **material 包是官方提供的 Material Design 实现，可替换**  
> **Theme 通常作为参数在整个 UI 树中传递**

---

准备好后，回复“继续第7章”，我会继续给出 Units 章节的完整翻译与深度解读。