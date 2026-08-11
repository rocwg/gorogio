**原文地址**：https://gioui.org/doc/architecture/text

---

在 GPU 硬件加速的 UI 框架中，文本处理本质上是将字符编码（Unicode）通过字体文件（TTF/OTF）进行 **塑形（Shaping）**，再转换为矢量 Path 或位图，最后由 GPU 渲染。Gio 提供了底层的 `text` 包和高层的 `material.Label` 来完成这一全流程。

章节：**Text（文本渲染与文本塑造）**。



# 第8章：Text（文本）

#### Low-level text management (底层文本管理)



## Fonts（字体集合与加载）

> **【英文原文】** 
>
> Gio’s text shaper uses the type `[]text.FontFace` to represent the collection of available fonts.
>
> There is one font bundled in package [`gioui.org/font/gofont`](https://gioui.org/font/gofont), you can use [`gofont.Collection()`](https://gioui.org/font/gofont#Collection) to get a `[]text.FontFace` containing all of the variants of the Go fonts.
>
> For loading other fonts there is [`gioui.org/font/opentype`](https://gioui.org/font/opentype). After parsing the font(s) using [`opentype.Parse`](https://gioui.org/font/opentype#Parse), you can append them to a `[]text.FontFace`.

**【逐字精准翻译】** 

Gio 的文本整形器（text shaper）使用 `[]text.FontFace` 类型来表示可用字体的集合。

在包 `gioui.org/font/gofont` 中内置绑定了一种字体，你可以使用 `gofont.Collection()` 来获取包含 Go 字体所有变体（如 Bold, Italic, Regular 等）的 `[]text.FontFace` 集合。

对于加载其他自定义字体，可以使用 `gioui.org/font/opentype`。在使用 `opentype.Parse` 解析字体后，你可以将它们追加（append）到 `[]text.FontFace` 中。

- **代码落地范例（自定义 OpenType / TTF 字体加载）：** 

  ```go
  import (
      "gioui.org/font/opentype"
      "gioui.org/text"
  )
  
  func loadCustomFont(fontBytes []byte) ([]text.FontFace, error) {
      // 1. 解析字体字节数据
      face, err := opentype.Parse(fontBytes)
      if err != nil {
          return nil, err
      }
      // 2. 将解析好的字体追加到 FontFace 集合中
      var collection []text.FontFace
      collection = append(collection, face)
      return collection, nil
  }
  ```

- **技术剖析：** 
  - `text.FontFace`：封装了字体的底层二进制 Data 以及 Style（粗细、斜体等元数据）。
  - `opentype.Parse`：能够解析 `.ttf` / `.otf` 标准字体文件，将其转换为 Gio 底层矢量引擎可绘制的 Face 对象。



## Shapes（字形转化与缓存）

> **【英文原文】** 
>
> For converting strings to clip shapes there is the [`gioui.org/text`](https://gioui.org/text) package.
>
> It contains [`text.Cache`](https://gioui.org/text#Cache) that implements cached string to shape conversion, with appropriate fallbacks. Simply provide your fonts (`[]text.FontFace`) to `text.NewCache`.

**【逐字精准翻译】** 

为了将字符串转换为裁切形状（Clip Shapes），Gio 提供了 `gioui.org/text` 包。

它包含了 `text.Cache`（在最新版 Gio 中演进为由 `text.Shaper` 托管），实现了从字符串到图形形状的缓存转换，并附带恰当的后备字体（Fallback）机制。只需将你的字体集合（`[]text.FontFace`）传递给 `text.NewCache`（或 `text.NewShaper`）即可。

- **底层机制说明：** 
  Gio 是纯 GPU 矢量渲染框架。字符串并不能直接绘制，必须先经过 **Text Shaping（文本整形/排版）** 转换为字符轮廓路径（`clip.Path`）。为了避免每一帧重复计算矢量路径，`text.Cache` 会在内存中缓存计算好的字形路径，从而保持高性能。



> **【英文原文】** 
>
> In most cases you can use [`widget.Label`](https://gioui.org/widget#Label) which handles wrapping and layout constraints. Or when you are using material design [`material.LabelStyle`](https://gioui.org/widget/material#LabelStyle).

**【逐字精准翻译】** 

在大多数情况下，你可以直接使用 `widget.Label`，它会自动处理换行（wrapping）和布局约束（layout constraints）。或者当你在使用 Material Design 时，直接使用 `material.LabelStyle`。

- **应用层与底层对比：** 
  - **高层（日常使用）：** 直接调用 `material.Body1(th, "Hello").Layout(gtx)`，框架会在内部自动调用 Shaper 与 Cache 完成测量和排版。
  - **底层（自定义组件）：** 若要实现类似图文混排、富文本编辑器或特殊的弯曲文本，才需要直接操作 `text.Shaper` / `text.Cache`。



### 4. 衍生实践代码：自定义加载 TTF 字体

为了帮助你将这段文档转化为落地代码，以下展示如何在 Gio 中加载自定义的 `.ttf` 字体（如思源黑体或自定义英文字体）：

```go
import (
	"gioui.org/font/opentype"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func loadCustomFont(ttfData []byte) (*material.Theme, error) {
	// 1. 解析 TTF/OTF 二进制数据
	face, err := opentype.Parse(ttfData)
	if err != nil {
		return nil, err
	}

	// 2. 构建 FontFace 集合
	fontFaces := []text.FontFace{face}

	// 3. 创建 Shaper 并注入主题
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(fontFaces))

	return th, nil
}
```



### Gio 文本渲染架构剖析

```powershell
  [ String / UTF-8 ] 
          │
          ▼
   [ text.Shaper ]  <───  ([]text.FontFace / GoFont / Custom TTF)
          │
  (Shaping & Metric Cache)
          │
          ▼
 [ Vector Clip Path / GPU Draw Op ]  ───> [ Rendered Text on Screen ]
```

- **架构要点解析：**
  - **Text Shaping（文本塑形）：** 文本并非简单的把字符贴在屏幕上。塑形引擎需要根据字体表计算字形（Glyph）索引、字距调整（Kerning）、基线（Baseline）以及从左到右/从右到左的排版。
  - **高层与底层分工：**
    - **底层（`text.Shaper`）：** 负责纯字形转换与测量，通常作为 `th.Shaper` 挂载在 `material.Theme` 上。
    - **高层（`material.Label` / `widget.Label`）：** 自动响应 `gtx.Constraints` 约束进行断行（Line Wrapping）、省略号截断（Ellipsis），并将其转化为对应的 `PaintOp` 输出到画布中。

---

### 深度解读

**1. 为什么文字章节这么短？**

官方文档第8章确实非常精简。这是因为：

- 大多数应用**不需要**直接操作底层文字系统
- 高层 API（`widget.Label` 和 `material.Label`）已经封装得很好
- 真正复杂的文字处理（双向文字、连字、emoji、复杂脚本）被藏在 HarfBuzz 和字体引擎里

但理解底层机制，对自定义控件、性能优化和排查文字问题非常重要。

**2. 文字渲染的完整流水线**

Gio 中一个字符串最终变成屏幕上的像素，大致经过以下步骤：

```
字符串
  ↓
Shaper（整形器）—— 根据字体、字号、方向等，把字符变成字形（Glyph）
  ↓
布局（换行、对齐、约束）
  ↓
生成 clip.Path / 字形路径
  ↓
paint 填充（颜色）
  ↓
屏幕像素
```

`text.Cache` 和 `text.Shaper` 就是负责前半段的核心。

**3. 字体系统详解**

**默认字体：Go Font**
```go
faces := gofont.Collection()
```
- 体积小
- 支持基本拉丁、西里尔、希腊等字符
- 官方推荐用于大多数应用

**加载自定义字体（OpenType / TrueType）**
```go
import "gioui.org/font/opentype"

fontData, _ := os.ReadFile("MyFont.ttf")
face, _ := opentype.Parse(fontData)
faces = append(faces, text.FontFace{Font: font.Font{Typeface: "MyFont"}, Face: face})
```

然后把 `faces` 交给 Shaper 或 Cache。

**多字体回退**  
`[]text.FontFace` 是有序的。整形器会按顺序查找能显示某个字符的字体（类似 CSS 的 font-family 回退机制）。这对中文、日文、emoji 混排非常重要。

**4. Shaper vs Cache**

| 组件          | 作用             | 使用场景               |
| ------------- | ---------------- | ---------------------- |
| `text.Shaper` | 核心整形引擎     | 被 material.Theme 使用 |
| `text.Cache`  | 带缓存的整形包装 | 需要手动控制缓存时     |

在第6章 Theme 里我们已经见过：
```go
th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
```

`material` 包内部会用这个 Shaper 来处理所有文字。

**5. 高层推荐用法（90% 的情况）**

**方式一：使用 material 包（最推荐）**
```go
material.H5(th, "标题").Layout(gtx)
material.Body1(th, "正文内容，支持自动换行").Layout(gtx)
material.Label(th, unit.Sp(14), "自定义大小").Layout(gtx)
```

**方式二：使用 widget.Label（更底层一点）**
```go
var label widget.Label
label.Layout(gtx, th.Shaper, font.Font{}, unit.Sp(16), "文字内容")
```

这两种方式都会自动处理：
- 换行
- 布局约束（根据父布局给的最大宽度）
- 基线对齐（Dimensions.Baseline）

**6. 实际开发中的关键点与踩坑**

| 问题                | 原因                | 解决方法                              |
| ------------------- | ------------------- | ------------------------------------- |
| 文字不显示          | Shaper 未设置       | 检查 `th.Shaper`                      |
| 中文/日文显示为方框 | 字体缺少对应字形    | 加载支持 CJK 的字体并加入 Collection  |
| 文字模糊            | 位置在半像素上      | 布局使用整数坐标（第7章）             |
| 性能差（大量文字）  | 每帧重新整形        | 使用 `text.Cache`，或减少不必要的重绘 |
| 想精确控制位置      | 高层 Label 不够灵活 | 手动使用 Shaper 获取字形信息后自己画  |

**7. 进阶：什么时候需要直接用底层 API？**

- 自定义文字效果（渐变字、描边字、路径文字）
- 实现富文本编辑器
- 精确测量文字宽度（做复杂布局）
- 特殊排版需求（竖排、环形文字等）

这时你需要：
1. 用 Shaper 把字符串变成 `[]text.Glyph`
2. 根据 Glyph 的位置和路径自己构建 `clip.Path`
3. 用 `paint` 填充

**8. 与前后章节的关系**

- 第2章 Drawing：文字最终变成 clip 形状 + paint 填充
- 第6章 Theme：Theme 持有 Shaper，所有 material 文字控件都依赖它
- 第7章 Units：文字大小必须用 **Sp**（会跟随系统字体缩放）
- 第9章 Color：文字颜色也是普通的 `color.NRGBA`

文字系统是 Gio 中相对独立但又与布局、绘制紧密耦合的一层。

---

**本章核心记忆点**

> **大多数情况直接用 `material.Label` 或 `widget.Label` 即可**  
> **必须给 Theme 设置 Shaper，否则文字无法显示**  
> **字体集合支持回退（多字体按顺序查找）**  
> **文字大小用 Sp，而不是 Dp**  
> **底层文字最终也会变成 clip 路径被 paint 填充**  
> **需要高性能或特殊效果时，才深入 Shaper / Cache**

---

准备好后，回复“继续第9章”，我会给出最后一章 Color 的完整翻译与深度解读，并在结束后做一个整体架构回顾。