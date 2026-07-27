**第8章：Text（文字）**

**原文地址**：https://gioui.org/doc/architecture/text

---

### 完整中文翻译

# Text

# Text 底层文字管理

## Fonts（字体）

Gio 的文字整形器使用类型 `[]text.FontFace` 来表示可用字体的集合。

包 [gioui.org/font/gofont](https://gioui.org/font/gofont) 中捆绑了一个字体，你可以使用 [gofont.Collection()](https://gioui.org/font/gofont#Collection) 获取一个包含所有 Go 字体变体的 `[]text.FontFace`。

要加载其他字体，可以使用 [gioui.org/font/opentype](https://gioui.org/font/opentype)。使用 [opentype.Parse](https://gioui.org/font/opentype#Parse) 解析字体后，可以把它们追加到 `[]text.FontFace` 中。

## Shapes（形状）

要把字符串转换成裁剪形状，可以使用 [gioui.org/text](https://gioui.org/text) 包。

它包含 [text.Cache](https://gioui.org/text#Cache)，实现了带适当回退的缓存字符串到形状转换。只需把你的字体（`[]text.FontFace`）提供给 `text.NewCache` 即可。

在大多数情况下，你可以使用 [widget.Label](https://gioui.org/widget#Label)，它会处理换行和布局约束。或者当你使用 Material Design 时，使用 [material.LabelStyle](https://gioui.org/widget/material#LabelStyle)。

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