**原文地址**：https://gioui.org/doc/architecture/units

---

采用 **【英文原文】 $\rightarrow$ 【精准逐字翻译】** $\rightarrow$ **【专业术语与 Gio 架构机制剖析】** 的方式，为你深度拆解 Gio 官方文档的章节：**Units（单位与双坐标系系统）**。

这是理解 Gio 如何适配不同屏幕 DPI（高清屏/ Retina 屏）以及解决传统 UI 框架“抗锯齿/像素缝隙”问题的核心章节。



# 第7章：Units（单位）

#### Measuring sizes of things (测量事物的大小)



> **【英文原文】** 
>
> Drawing operations use pixel coordinates, ignoring any transformation applied. However, for most use-cases you don’t want to tie your user-interface sizes and positions to screen pixels. People may have screen-scaling enabled and pixel densities vary significantly between devices.

**【逐字精准翻译】** 

绘制操作使用像素坐标，忽略任何已应用的变换。然而，对于大多数使用场景，你并不想把用户界面的尺寸和位置绑定到屏幕像素上。人们可能启用了屏幕缩放，而且不同设备的像素密度差异很大。

- **词汇与概念剖析：**
  - `screen-scaling`：屏幕缩放设置（如 Windows 系统的 125%、150% 缩放）。
  - `pixel densities`：像素密度（DPI/PPI）。



> **【英文原文】** 
>
> In addition to the physical pixel, package [`gioui.org/unit`](https://gioui.org/unit) implements device independent units:
>
> - [`Px`](https://gioui.org/unit#Px) - device dependent pixel. One Px is a pixel on the screen.
> - [`Dp`](https://gioui.org/unit#Dp) - device independent pixel. Takes into account screen-density and the screen-scaling settings.
> - [`Sp`](https://gioui.org/unit#Sp) - device independent pixel for text. An Sp is like a Dp but adjusted for font-scaling.
>
> [`layout.Context`](https://gioui.org/layout#Context) has method [`Px`](https://gioui.org/layout#Context.Px) to convert from [`unit.Value`](https://gioui.org/unit#Value) to pixels
>
> For more information on pixel-density see:
>
> - https://material.io/design/layout/pixel-density.html.
> - https://webplatform.github.io/docs/tutorials/understanding-css-units/

**【逐字精准翻译】** 

除了物理像素之外，`gioui.org/unit` 包实现了设备无关单位：

* **Px（物理像素）**：依赖于设备的像素。一个 Px 就是屏幕上的一个物理像素点。
* **Dp（设备无关像素）**：设备无关像素。综合考量了屏幕像素密度（DPI）和系统屏幕缩放设置。
* **Sp（缩放无关像素）**：专门用于文本的设备无关像素。Sp 类似于 Dp，但会根据系统的字体缩放比例进行调整（如老年人模式调大字体）。

`layout.Context` 提供了 `Px` 方法，用于将 `unit.Value`（如 `unit.Dp(16)` 或 `unit.Sp(14)`）转换为对应当前渲染上下文的像素值（`int`）。

有关像素密度的更多信息，请参见上述链接。

- **代码落地范例（Dp 转 Px）：** 在编写自定义 Widget 或计算布局尺寸时

  ```go
  func customWidget(gtx layout.Context) layout.Dimensions {
      // 将 16 Dp 转换为当前窗口下的实际物理像素 int 值（在 200% 缩放下自动变为 32px）
      padding := gtx.Dp(unit.Dp(16)) 
      // ...
  }
  ```



## Coordinate systems（坐标系）

> **【英文原文】** 
>
> You may have noticed that widget constraints and dimensions sizes are in integer units, while drawing commands such as [`PaintOp`](https://gioui.org/op/paint#PaintOp) use floating point units. That’s because they refer to two distinct coordinate systems, the layout coordinate system and the drawing coordinate system. The distinction is subtle, but important.

**【逐字精准翻译】** 

你可能已经注意到，控件的约束（`Constraints`）和尺寸大小（`Dimensions`）使用的是**整数单位（`int`/`image.Point`）**，而像 `PaintOp` 这样的绘制指令使用的则是**浮点数单位（`float32`/`f32.Point`）**。这是因为它们指向两个截然不同的坐标系统：**布局坐标系（layout coordinate system）** 和 **绘制坐标系（drawing coordinate system）**。这个区别很微妙，但很重要。

- **对比要点：**  
  - **Layout（布局阶段）：** 使用 `int`（整数像素）。  
  - **Drawing/Vector（绘制/渲染阶段）：** 使用 `float32`（浮点数像素）。



> **【英文原文】** 
>
> The layout coordinate system is in integer pixels, because it’s important that widgets never unintentionally overlap in the middle of a physical pixel. In fact, the decision to use integer coordinates was motivated by [conflation issues](https://github.com/flutter/flutter/issues/15035) in other UI libraries caused by allowing fractional layouts.
>
> As a bonus, integer coordinates are perfectly deterministic across all platforms which leads to easier debugging and testing of layouts.

**【逐字精准翻译】** 

布局坐标系采用**整数像素**，是因为确保组件永远不会“无意中”在物理像素的中间发生重叠非常重要。事实上，采用整数坐标的决定，源于其他 UI 库因允许小数/分数布局（fractional layouts）而引发的混合/模糊混淆问题。。

作为一个附加的好处，整数坐标在所有平台上都是**完全确定性的（Deterministic）**，这使得布局的调试和测试变得更加容易。

- **设计哲学解析：** 
  很多 Web 和 Mobile 框架允许子组件拥有 `10.5px` 的宽度，这经常导致亚像素渲染模糊（sub-pixel blurring）以及不同操作系统下舍入策略不一致导致的布局错位。Gio 在布局层面**强制采用整数**，保证了绝对的像素对齐和跨平台一致性。  



> **【英文原文】** 
>
> On the other hand, drawing commands need the generality of floating point coordinates for smooth animation and for expressing inherently fractional shapes such as bézier curves.
>
> It’s possible to draw shapes that overlap at fractional pixel coordinates, but only intentionally: drawing commands directly derived from layout constraints have integer coordinates by construction.

**【逐字精准翻译】** 

另一方面，绘制命令需要浮点坐标的通用性，以实现平滑动画并表达本质上包含小数的形状（例如贝塞尔曲线）。 

在小数像素坐标处绘制重叠的形状是可能的，但这必须是“有意的”：直接源自布局约束的绘制命令，在构造上就已经天然具备整数坐标。

### 布局坐标系 vs. 绘制坐标系对比

| **维度**     | **布局坐标系 (Layout System)**                       | **绘制坐标系 (Drawing System)**                 |
| ------------ | ---------------------------------------------------- | ----------------------------------------------- |
| **数据类型** | 整数（`int` / `image.Point`）                        | 浮点数（`float32` / `f32.Point`）               |
| **核心作用** | 测量大小（Dimensions）、计算边距与布局约束           | 绘制路径（Paths）、位移（Transforms）、平滑动画 |
| **设计考量** | **确定性**：消除不同平台显示卡缝、1px 缝隙与浮点误差 | **高精度**：实现抗锯齿、平滑矢量图与动画过渡    |

### 官方文档全系列终章总结

到此为止，Gio 的核心官方文档全部拆解完毕：

1. **Display（渲染机制）**：基于 GPU Direct Rendering 的即时模式引擎。
2. **Architecture（架构图）**：OS Window $\rightarrow$ Loop $\rightarrow$ Frame Event $\rightarrow$ GPU Ops 的事件闭环。
3. **Ops（绘制指令）**：Canvas/Clip/Paint 基础操作与操作列表流。
4. **Input（事件系统）**：基于 Queue 和 Target 的无状态输入响应。
5. **Widget（控件与上下文）**：`gtx (layout.Context)` 传递机制与生命周期。
6. **Layout（弹性排版）**：Flex/Stack/List 等基于 Constraints 传播的布局算法。
7. **Theme（主题与 Material）**：状态（State）与视觉（Visuals）的彻底解耦。
8. **Units（单位与坐标系）**：Dp/Sp 转换与整数/浮点双坐标系隔离设计。

---

### 深度解读

**1. 为什么需要单位系统？**

现代设备的屏幕差异极大：

- 手机可能是 2x、3x、甚至更高的像素密度
- 桌面显示器有 100%、125%、150%、200% 等系统缩放
- 用户还可能单独调整字体大小（无障碍需求）

如果直接用物理像素写死 `100`，在高密度屏幕上会变得特别小，在低密度屏幕上又会特别大。  
因此几乎所有现代 UI 框架都引入了**密度无关单位**。

**2. 三种单位的本质区别**

| 单位   | 全称                      | 受什么影响                         | 主要用途                         |
| ------ | ------------------------- | ---------------------------------- | -------------------------------- |
| **Px** | Pixel                     | 无（物理像素）                     | 最终绘制、精确到像素的场合       |
| **Dp** | Density-independent Pixel | 屏幕密度 + 系统缩放                | 几乎所有布局尺寸、间距、控件大小 |
| **Sp** | Scale-independent Pixel   | 屏幕密度 + 系统缩放 + **字体缩放** | 文字大小                         |

简单记忆：
- 普通 UI 元素 → 用 **Dp**
- 文字 → 用 **Sp**
- 极少数需要精确到物理像素的场合 → 用 **Px**

**3. 如何在代码中转换？**

`layout.Context` 提供了方便的转换方法：

```go
// 把 16dp 转成当前设备的像素值
px := gtx.Dp(unit.Dp(16))

// 或者更通用的
px := gtx.Metric.Dp(unit.Dp(16))
px = gtx.Metric.Sp(unit.Sp(14))
```

在 material 主题和大多数官方控件里，你几乎看不到硬编码的像素数字，全部使用 Dp/Sp。

**4. 两个坐标系的深刻设计决策**

这是本章最重要、也最容易被忽略的点：

**布局坐标系（Constraints / Dimensions）→ 整数像素**

原因：
- 防止控件在半个物理像素上重叠（导致模糊或闪烁）
- 布局结果在所有平台完全一致，便于测试和调试
- 参考了 Flutter 早期因为允许分数布局而产生的混叠问题

**绘制坐标系（clip、Transform、Paint）→ 浮点坐标**

原因：
- 动画需要亚像素移动才能平滑
- 贝塞尔曲线、圆角、抗锯齿本身就需要浮点精度
- 旋转、缩放等变换会产生分数坐标

**实际影响**：
- 你从布局得到的 `Dimensions.Size` 永远是整数
- 但你可以在绘制时使用 `op.Offset(f32.Pt(0.5, 0))` 做亚像素偏移
- 官方保证：只要你的绘制命令是直接从整数约束派生的，就不会出现“无意的半像素重叠”

**5. 实际开发中的最佳实践**

| 场景                 | 推荐做法                                       |
| -------------------- | ---------------------------------------------- |
| 控件宽高、间距、圆角 | `unit.Dp(xx)`                                  |
| 字体大小             | `unit.Sp(xx)`                                  |
| 边框宽度（想要清晰） | 通常用 `unit.Dp(1)` 或根据密度计算             |
| 动画偏移             | 可以用浮点                                     |
| 自定义布局计算       | 先全部用 Dp 计算，最后转成像素再设 Constraints |
| 调试布局             | 记住 Constraints 是整数，方便打印和对比        |

**6. 与 Material Design 的关系**

Material Design 官方规范就是以 Dp 和 Sp 为基础的。  
Gio 的 material 包完全遵循这套规范，所以你在 Theme 里看到的字号、间距都是 Dp/Sp 值。

**7. 常见踩坑**

1. **在布局里直接写死像素数字**  
   → 不同设备显示效果不一致。

2. **文字用了 Dp 而不是 Sp**  
   → 用户调整系统字体大小时，你的文字不会跟着变。

3. **把浮点尺寸直接塞进 Constraints**  
   → Constraints 只接受整数，需要手动取整。

4. **忘记使用 `gtx.Dp()` / `gtx.Sp()`**  
   → 直接拿 `unit.Dp(16)` 当 int 用会编译错误或结果错误。

**8. 与前后章节的关系**

- 第2章 Drawing：所有绘制最终使用像素（浮点）
- 第4、5章 Widget & Layout：Constraints 和 Dimensions 使用整数像素
- **第7章** 解释了为什么会有这两种坐标系，以及如何正确在它们之间转换
- 第8章 Text 会大量使用 Sp

---

**本章核心记忆点**

> **Dp 用于普通尺寸，Sp 用于文字，Px 是最终物理像素**  
> **布局用整数坐标（确定性 + 避免半像素问题）**  
> **绘制用浮点坐标（支持平滑动画和曲线）**  
> **永远通过 `gtx.Dp()` / `gtx.Sp()` 进行转换**  
> **这是 Gio 为了跨设备一致性和可调试性做出的刻意设计**

---

准备好后，回复“继续第8章”，我会继续给出 Text 章节的完整翻译与深度解读。