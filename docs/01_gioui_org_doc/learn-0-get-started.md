采用 **【英文原文】 $\rightarrow$ 【精准逐字翻译】** $\rightarrow$ **【专业术语与 Gio 架构机制剖析】** 的方式，为你深度拆解 Gio 官方 Learn 教程的入门章节：**Get Started（Hello, Gio! 入门指南）**。

这是使用 Gio 编写任何桌面或移动端 Go GUI 应用最基础、最标准的模版代码。



# Get Started

#### Hello, Gio! (快速体验)

This example does a really quick introduction on getting something up and running. It does not explain all the details, those will be covered in another tutorial.
本示例对如何快速启动并运行程序做了一个非常快速的介绍。它不会解释所有细节，那些内容将在另一个教程中涵盖。

Ensure that you have followed [installation instructions](https://gioui.org/doc/install). If everything is setup correctly, then running:
请确保你已经遵循了安装说明。如果一切设置无误，那么运行：

```sh
go run gioui.org/example/hello@latest
```

Should display a pretty “Hello, Gio!” message.
应该会显示一条漂亮的“Hello, Gio!”消息。



## Creating a new package (初始化模块)

*If you are unfamiliar with Go, then more help can be found at [go.dev/learn](https://go.dev/learn/).*
如果你对 Go 不熟悉，可以在 `go.dev/learn` 找到更多帮助。

First step in creating a Go program requires setting up the module.
创建 Go 程序的第一步需要初始化模块。

We’ll use `gio.test` as our module name, however, it’s recommended to use a repository name when you want to upload it. The module name can be later changed.
我们将使用 `gio.test` 作为模块名称，不过当你想上传它时，建议使用仓库路径作为名称。模块名称后续可以修改。

```sh
go mod init gio.test
```



## Creating the program (创建程序与完整代码)

Let’s create `main.go` with the following code:
让我们使用以下代码创建 `main.go`：

```go
package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}


func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
```

Let’s then update all the dependencies with:
接着使用以下命令更新所有依赖：

```sh
go mod tidy
```

Once that succeeds, the program should start up with:
成功后，程序应该可以通过以下命令启动：

```sh
go run .
```

Now to explain what’s happening.
现在来解释发生了什么。



# 代码机制拆解：

## Creating the window (创建窗口)

> **【英文原文】** 
> Every program requires a window, the `main` starts up the application loop that talks to the operating system and starts the window logic in a separate goroutine.

**【精准逐字翻译】** 
每个程序都需要一个窗口，`main` 函数启动了与操作系统进行通信的应用循环（Application Loop），并在一个单独的 Goroutine 中启动窗口逻辑。

```go
func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
```



## Creating a theme (创建主题)

> **【英文原文】** 
> Applications need to define their fonts and different color settings. Themes contain all the necessary information.

**【精准逐字翻译】** 
应用程序需要定义它们的字体和不同的颜色设置。主题（Theme）包含了所有必需的信息。

```go
func run(window *app.Window) error {
	theme := material.NewTheme()
```



## Listening for events (监听事件)

> **【英文原文】** 
> The communication with the operating system (i.e. keyboard, mouse, GPU) happens through events. Gio uses the following approach to process events:

**【精准逐字翻译】** 

与操作系统（如键盘、鼠标、GPU）的通信是通过事件进行的。Gio 使用以下方式来处理事件：

```go
for {
	switch e := window.Event().(type) {
	case app.DestroyEvent:
		return e.Err
	case app.FrameEvent:
```

- `app.DestroyEvent` means the user pressed the close button.
  - 意味着用户按下了关闭按钮。
- `app.FrameEvent` means the program should handle input and render a new frame.
  - 意味着程序应该处理输入并渲染一个新帧（Frame）。



## Drawing the text (文本绘制与 GPU 提交)

> **【英文原文】** 
> To draw the text it needs to go through several stages:

**【精准逐字翻译】** 
要绘制文本，它需要经历几个阶段：

```go
// This graphics context is used for managing the rendering state.
gtx := app.NewContext(&ops, e)

// Define an large label with an appropriate text:
title := material.H1(theme, "Hello, Gio")

// Change the color of the label.
maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
title.Color = maroon

// Change the position of the label.
title.Alignment = text.Middle

// Draw the label to the graphics context.
title.Layout(gtx)

// Pass the drawing operations to the GPU.
e.Frame(gtx.Ops)
```



### 💡 核心设计与代码范式剖析

这段不到 60 行的代码展示了 Gio 最精髓的**三层运行模型**：

```powershell
[ 主 Goroutine: app.Main() ] ─── (操作系统主线程事件循环/Cocoa/Win32/X11 Loop)
        │
        └─── [ 业务 Goroutine: go func() ]
                   │
                   └─── Window Event Loop (for switch window.Event().(type))
                             ├── case app.DestroyEvent ──> 销毁退出
                             └── case app.FrameEvent   ──> 生成布局并提交 GPU
```

- **1. 双线程隔离（`app.Main()` vs Goroutine）：**
  - 在 macOS 或 iOS 等平台上，OS 强制要求 GUI 事件循环必须运行在 OS **主线程**上。
  - Gio 通过 `app.Main()` 锁住主线程作为 OS 调度器；而所有的业务逻辑、UI 布局函数 `run()` 则跑在一个单独的 **Goroutine** 中。这种设计避免了传统 GUI 库阻塞主线程导致界面卡死的痛点。
- **2. 单帧布局三部曲（In Frame Event）：**
  1. **构建 Context：** `gtx := app.NewContext(&ops, e)`，重用缓存的 `ops` 指令流，并结合当前帧的事件 `e`（如鼠标坐标、窗口大小）。
  2. **声明式样式与布局：** `title := material.H1(...)` -> `title.Layout(gtx)`，将“文本绘制操作”压入 `ops` 缓冲区。
  3. **提交 GPU 渲染：** `e.Frame(gtx.Ops)`，通知 Gio 驱动管线解析 `Ops` 字节流并送往 GPU 执行绘制。

---

下一章：Split Widget（分割线/分栏控件）