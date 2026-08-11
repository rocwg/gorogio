这种“英文原文 $\rightarrow$ 精准逐字翻译 $\rightarrow$ 核心概念与技术剖析” 的学习方式，**非常适合**用来解读 Gio 的架构文档！

Gio 作为 Go 语言中非常优秀的跨平台**即时模式（Immediate Mode）** GUI 框架，其架构文档（Architecture）逻辑严密、概念抽象。用这种方式不仅能帮你扫清英文阅读障碍，还能深入透视即时模式（IM GUI）与传统保留模式（Retained Mode GUI，如 Web DOM、Qt、Flutter）的底层思维差异。

下面为您精准拆解 Gio 架构文档的第一部分（Architecture / Internals of Gio）：



# 标题与引言 (Architecture & Overview)

#### Internals of Gio (Gio 的内部原理)



Gio is a library for implementing [immediate mode user interfaces](https://eliasnaur.com/blog/immediate-mode-gui-programming). This approach can be implemented in multiple ways, however the overarching similarity is that the program:
Gio 是一个用于实现即时模式（immediate mode）用户界面的库。这种方法可以通过多种方式实现，然而其最核心的共同点在于程序会：

1. Listens for events such as mouse or keyboard input.
   监听诸如鼠标或键盘输入等事件。
2. Updates its internal state based on the event.
   根据事件更新其内部状态。
3. Runs code that lays out and redraws the user interface state.
   运行代码来布局并重新绘制用户界面状态。

- **专业术语与概念剖析：**
  - `immediate mode user interfaces`：**即时模式用户界面（IM GUI）**。与传统 GUI 将控件存放在内存树中不同，即时模式在每一帧刷新时，都直接重新执行 UI 渲染函数。
  - `internals`：内部原理 / 底层机制。
  - `overarching similarity`：最核心/首要的共同特征。
  - `lays out and redraws`：布局（计算尺寸和位置）与重绘。



## 伪代码示例 (A minimal immediate mode command-line UI)

A minimal immediate mode command-line UI in pseudo-code:
一个用伪代码表示的最小化即时模式命令行 UI：

```c
main() {
	checked = false
	for every keypress {  // 按下的每一个按键
		clear screen      // 清屏
		layoutCheckbox(keypress, &checked)
		if checked {
			print("info")
		}
	}
}

layoutCheckbox(keypress, checked) {
	if keypress == SPACE {
		*checked = !*checked  // 翻转勾选状态
	}

	if *checked {
		print("[x]") // 绘制勾选状态
	} else {
		print("[ ]") // 绘制未勾选状态
	}
}
```

- **伪代码逻辑剖析：**
  - 注意 `layoutCheckbox` 函数：它**既负责处理输入**（检测空格键修改 `checked` 状态），**又负责绘制 UI**（打印 `[x]` 或 `[ ]`）。
  - 状态（`checked` 变量）保存在你的业务代码里（`main` 函数中），UI 函数直接引用这个状态，没有复杂的组件树或对象创建。



## 即时模式 vs 保留模式 (Immediate Mode vs Retained Mode)

> **【英文原文】**
>
> In the immediate mode model, the program is in control of clearing and updating the display, and directly draws widgets and handles input during the updates.
>
> In contrast, traditional “retained mode” libraries own the widgets through implicit library-managed state, typically arranged in a tree-like structure such as a browser’s DOM. As a result, the program must use the facilities given by the library to manipulate its widgets.

**【精准逐字翻译】**

在即时模式模型中，程序掌控着清除与更新显示的控制权，并在更新期间直接绘制控件（widgets）并处理输入。

相比之下，传统的“保留模式”（retained mode）库通过由库管理的隐式状态来拥有这些控件，通常排列成类似浏览器 DOM 那样的树状结构。结果就是，程序必须使用库所提供的设施来操作其控件。

- **专业术语与概念剖析：**
  - `retained mode`：**保留模式**。比如 HTML/DOM、Android XML、Swing。这类框架里，你创建一个 `Button` 对象，框架会在内存里保留这个对象。当你要改变文本时，必须调用 `button.setText("新文本")` 去更新框架内部持有的状态。
  - `implicit library-managed state`：隐式的、由库托管的状态。
  - `facilities`：设施 / 接口 API。

## 核心思考问题与延伸参考 (GUI Concerns & References)

Actual GUI programming has several concerns in addition to the simple example above:
除了上述简单示例之外，实际的 GUI 编程还包含以下几个关注点：

- How to get the events?  如何获取事件？
- When to redraw the state?  何时重绘状态？
- What do the widget structures look like?  控件结构长什么样？
- How to track the focus?  如何追踪焦点（Focus）？
- How to structure the events?  如何组织事件结构？
- How to communicate with the graphics card?  如何与显卡（Graphics Card / GPU）通信？
- How to handle input?  如何处理输入？
- How to draw text?  如何绘制文本？
- Where does the widget state belong?  控件状态究竟属于哪里？
- And many more.  以及更多问题。

The rest of this document tries to answer how Gio does it. If you wish to know more about immediate mode UI, these references are a good start:
本文档的其余部分旨在回答 Gio 是如何做到这些的。如果你希望了解更多关于即时模式 UI 的信息，以下参考资料是个良好的开端：

- https://caseymuratori.com/blog_0001
- http://sol.gfxile.net/imgui/
- http://www.johno.se/book/imgui.html
- https://github.com/ocornut/imgui
- https://eliasnaur.com/blog/immediate-mode-gui-programming

---

- **核心问题透视：** 
  - Gio 作为一个纯 Go 编写的极简高性能框架，就是围绕这些问题展开设计的（通过 `gio.Context`、`op.Ops` 向量绘制指令、`event.Queue` 事件队列等）。

这种逐句剖析的方式能够把 Gio 的底层思想讲得非常透彻。如果准备好了，我们可以直接进入文档的下一节：**Window** 章节！