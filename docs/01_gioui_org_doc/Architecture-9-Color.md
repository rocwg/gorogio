**原文地址**：https://gioui.org/doc/architecture/color

---

采用 **【英文原文】 $\rightarrow$ 【精准逐字翻译】** $\rightarrow$ **【专业术语与 Gio 色彩混合架构机制剖析】** 的方式，为你深度拆解 Gio 官方文档的章节：**Color（颜色与色彩空间混合）**。

这是理解 Gio 为什么在渲染渐变、透明度叠加和图像缩放时比许多传统 UI 库更清晰、无“发灰/发脏（Muddy）”现象的核心章节。



# 第9章：Color（颜色）

#### Understanding color and blending (理解颜色与混合)



> **【英文原文】** 
>
> Color handling is something that we don’t usually don’t think about. However, a framework can make many tradeoffs while handling color.
>
> The short explanation is that Gio uses sRGB colors for input but uses linear color space for blending. This results in the color blending being correct without manually converting usual color values to linear color space.
>
> If the short explanation wasn’t sufficient, then there’s a longer one below.
>
> *Note: the following will oversimplify things to make them more understandable. For all the gritty details, read the linked articles.*

**【逐字精准翻译】** 

颜色处理通常是我们不会主动去思考的事情。然而，UI 框架在处理颜色时会做出许多权衡取舍。

简短的解释就是：**Gio 使用 sRGB 颜色空间进行输入，但使用线性颜色空间（Linear Color Space）进行混合（Blending）。**这样既保证了颜色混合的正确性，又无需开发者手动将日常使用的颜色值转换为线性颜色空间。

如果简单的解释还不够充分，下面提供了更详细的说明。

*注意：以下内容会简化一些内容，使它们更容易理解。*
*要了解所有细节，请阅读链接的文章。*



## Color primer（颜色入门）

> **【英文原文】** 
>
> Most programs represent colors with red, green, and blue lightness values. The simplest approach would be to represent the exact lightness value with the value you use in your RGB color. However, eyes are more sensitive to darker colors than lighter colors. With just 8 bits available per color channel, a linear mapping of lightness wastes bits to represent lighter values that people cannot differentiate.
>
> One approach is [gamma correction](https://en.wikipedia.org/wiki/Gamma_correction) that encodes lightness values with a function that stretches the darker color range at the cost of compressing the lighter color range.
>
> Usually the gamma transformations look like:

**【逐字精准翻译】** 

大多数程序使用红（R）、绿（G）、蓝（B）的明度值来表示颜色。最简单的方法是用 RGB 颜色中的数值直接表示精确的明度值。然而，人眼对暗色比对亮色更加敏感。在每个颜色通道只有 8 位（Bit）可用（即 0-255）的情况下，明度的线性映射会浪费宝贵的位数去表示人眼根本区分不开的高光细节。

一种解决方法是**伽马校正（Gamma Correction）**，它使用一个函数对亮度值进行编码，这个函数会拉伸较暗的颜色范围，代价是压缩较亮的颜色范围。

通常，Gamma 变换如下所示：

```go
// 将线性颜色转换为 Gamma 压缩颜色
gamma_color  := math.Pow(linear_color, gamma)
// 将 Gamma 压缩颜色转换为线性颜色
linear_color := math.Pow(gamma_color, 1/gamma)

// 其中
linear_color = [0..1]
gamma_color  = [0..1]
gamma        = 通常是 2.2 或 2.4
```

- 伽马变换公式：
  - $$\text{gamma\_color} = \text{linear\_color}^{\gamma}$$ 
  - $$\text{linear\_color} = \text{gamma\_color}^{1/\gamma}$$ 
  - *(其中 $\gamma$ 通常为 $2.2$ 或 $2.4$)* 



> **【英文原文】** 
>
> One of the problems with this function is that the [rate of color change is near infinite](https://en.wikipedia.org/wiki/SRGB#Transfer_function_("gamma")). To avoid this boundary condition there is a lightness value transformation called [sRGB color space](https://en.wikipedia.org/wiki/SRGB). sRGB conversion looks like:

**【逐字精准翻译】** 

这种函数的问题之一在于其零点附近的***颜色变化率（斜率）***趋近于无穷大。为了避免这种边界条件，出现了一种名为 **sRGB 颜色空间** 的明度变换。sRGB 转换看起来像这样：

```go
// 把线性颜色转换成 sRGB 颜色
if linear_color <= 0.0031308 {
	srgb_color = 12.92 * linear_color
} else { // linear_color > 0.0031308
	srgb_color = 1.055 * math.Pow(linear_color, 1/2.4) - 0.055
}

// 把 sRGB 颜色转换成线性颜色
if srgb_color <= 0.04045 {
	linear_color = srgb_color / 12.92
} else { // srgb_color > 0.04045
	linear_color = math.Pow((srgb_color + 0.055) / 1.055, 2.4)
}
```

sRGB 的转换公式如下：

- **线性转 sRGB：**
  - 当 $\text{linear\_color} \le 0.0031308$ 时，$\text{srgb\_color} = 12.92 \times \text{linear\_color}$ 
  - 当 $\text{linear\_color} > 0.0031308$ 时，$\text{srgb\_color} = 1.055 \times \text{linear\_color}^{1/2.4} - 0.055$ 
- **sRGB 转线性：**
  - 当 $\text{srgb\_color} \le 0.04045$ 时，$\text{linear\_color} = \text{srgb\_color} / 12.92$ 
  - 当 $\text{srgb\_color} > 0.04045$ 时，$\text{linear\_color} = \left(\frac{\text{srgb\_color} + 0.055}{1.055}\right)^{2.4}$ 



> **【英文原文】** 
>
> The details of the sRGB vs gamma corrected colors aren’t that important for the discussion, so we’ll keep using the gamma transformation, because it’s shorter to write than the sRGB conversions.

**【逐字精准翻译】** 

sRGB 与 Gamma 校正颜色的细节对于本次讨论并不那么重要，因此我们将继续使用 Gamma 变换，因为它比 sRGB 转换写起来更短。



## Problems with sRGB（sRGB 的问题）

> **【英文原文】** 
>
> One of the problems that sRGB and gamma-corrected colors have is that when you directly compute the sum of them, you don’t get the correct color mixing.
>
> Let’s take an example of mixing `linear_color_alpha` and `linear_color_beta`:

**【逐字精准翻译】** 

sRGB 和 Gamma 校正颜色存在的问题之一是：**当你直接计算它们的加权和（线性叠加）时，无法得到正确的色彩混合效果。**

我们以混合 `linear_color_alpha` 和 `linear_color_beta` 为例：

```go
// 使用线性颜色空间混合颜色
linear_color = 0.5*linear_color_alpha + 0.5*linear_color_beta

// 使用 sRGB 颜色空间混合颜色
linear_color = math.Pow(
	0.5 * math.Pow(linear_color_alpha, gamma) +
	0.5 * math.Pow(linear_color_beta, gamma),
	1/gamma)
```

- 在线性颜色空间中直接混合（正确方式）：

  $$\text{linear\_color} = 0.5 \times \text{linear\_color\_alpha} + 0.5 \times \text{linear\_color\_beta}$$ 

- 在 sRGB 颜色空间中直接混合（传统错误方式）：

  $$\text{linear\_color} = \left( 0.5 \times \text{linear\_color\_alpha}^{\gamma} + 0.5 \times \text{linear\_color\_beta}^{\gamma} \right)^{1/\gamma}$$ 



> **【英文原文】** 
>
> When you experiment with this example, you should notice that blending in sRGB results often in a darker or grayer color, which ends up causing muddied colors in blending.
>
> The blending issues have been discussed in more detail in:

**【逐字精准翻译】** 

当你动手实验这个例子时会发现，**在 sRGB 空间中直接混合往往会导致颜色偏暗或发灰，最终导致混合处的色彩变脏/变浊（Muddied）。**

混合问题在以下文章中有更详细的讨论：

- [How software gets color wrong](https://bottosson.github.io/posts/colorwrong/) (软件是如何把颜色搞错的)
- [Linear Gamma vs Higher Gamma](https://ninedegreesbelow.com/photography/linear-gamma-blur-normal-blend.html) (线性 Gamma 与高 Gamma 对比)
- [Gamma error in picture scaling](http://www.ericbrasseur.org/gamma.html) (图片缩放中的 Gamma 误差)



## Choice for frameworks（框架的选择）

> **【英文原文】** 
>
> Overall, frameworks need to choose a color space to work with. Historically, the most common choice was sRGB because of the darker color benefit. Similarly, as an accident or for performance reasons, people ended up using sRGB blending. This also leads to bugs related to [resizing images](http://www.ericbrasseur.org/gamma.html).
>
> So, due to the historical importance of sRGB, there are a few choices for a UI framework:
>
> 1. Use sRGB for input and blending: this causes incorrect blending and muddy colors. However, this behavior is similar to all other programs.
> 2. Use linear colors for input and blending: this has correct blending. However, people cannot use their usual “color pickers” (because they work in sRGB) and must manually convert images from sRGB to linear.
> 3. Use sRGB colors when providing input; however, blend using linear colors: this is compatible with programs for color selection. Mixing colors is going to be different from sRGB blending.
>
> Gio has chosen approach **3**, because it’s a pragmatic choice that has correct blending and does not have the annoyances of color conversion.
>
> *Sidenote: of course, there are more choices, such as using higher bit-depth or wide-gamut color spaces, but for usual UI applications, there isn’t a significant benefit from them.*

**【逐字精准翻译】** 

总的来说，框架需要选择一个颜色空间来工作。历史上，最常见的选择是 sRGB，因为它对较暗颜色有好处。同样，作为意外或出于性能原因，人们最终使用了 sRGB 混合。这也导致了与[图像缩放相关的 bug](http://www.ericbrasseur.org/gamma.html)。

因此，由于 sRGB 的历史重要性，UI 框架有几种选择：

1.  **输入和混合均使用 sRGB**：这会导致不正确的色彩混合和发脏的颜色。不过，这种行为和大多数传统软件一致。

2.  **输入和混合均使用线性颜色**：这拥有正确的色彩混合。但开发者无法直接使用日常的“取色器”（因为取色器输出 sRGB），且必须手动将图片从 sRGB 转换为线性空间。

3.  **输入时使用 sRGB 颜色；但在渲染混合时使用线性颜色**：这既兼容了常见的取色工具，又能获得物理上正确的色彩混合效果。

Gio 选择了**第 3 种方法**，因为它是一个务实的选择，既有正确的混合，又没有颜色转换的烦恼。

*旁注：当然，还有更多选择，例如使用更高的位深度（bit-depth）或广色域（wide-gamut）色彩空间，但对于常规的 UI 应用程序来说，这些并不会带来显著的好处。*

### Gio 颜色处理逻辑总结

| 阶段               | 使用的色彩空间                 | 开发者视角 / 架构行为                                        |
| ------------------ | ------------------------------ | ------------------------------------------------------------ |
| **API 输入阶段**   | **sRGB** (`image/color.NRGBA`) | 开发者直接传入标准十六进制或 `color.NRGBA{R:0xff, G:0x00, B:0x00, A:0xff}`。 |
| **GPU 着色器渲染** | **Linear RGB**                 | Gio 在提交 Op 到 GPU 或 Shader 计算时，自动将 sRGB 转化为 Linear，进行 Alpha 混合与像素采样，避免边缘暗环。 |



---

### 深度解读

**1. 为什么颜色章节放在最后？**

因为前面所有绘制、控件、主题最终都会落到颜色上。颜色处理看似简单，却是很多 UI 框架埋下视觉问题的地方。

**2. 人眼感知与 8 位颜色的矛盾**

人眼对暗部更敏感。如果直接用线性 0~255 表示亮度：
- 暗部变化被压缩，细节丢失
- 亮部浪费了很多人眼分辨不出的级数

因此诞生了**伽马校正**和后来的 **sRGB**（一种标准化的近似伽马曲线）。  
我们平时在取色器、CSS、设计软件里看到的 `#FF0000`、`rgb(255,0,0)` 几乎都是 **sRGB 值**。

**3. 混合为什么会出问题？**

假设你要做 50% 红色 + 50% 绿色的半透明混合。

- **在 sRGB 空间直接平均**：结果会偏暗、发灰（“muddy colors”）
- **在线性空间平均后再转回 sRGB**：结果符合物理光照，看起来正确

很多老框架（包括早期浏览器、部分游戏引擎）直接在 sRGB 上混合，导致半透明、抗锯齿、渐变经常出现脏色。

**4. Gio 的务实选择（方法 3）**

Gio 做了目前公认较好的折中：

| 环节                             | 使用的空间  | 好处                           |
| -------------------------------- | ----------- | ------------------------------ |
| **你输入的颜色**                 | sRGB        | 可以直接用取色器、设计稿的数值 |
| **内部混合、抗锯齿、半透明叠加** | 线性空间    | 混合结果正确，不会发脏         |
| **最终输出到屏幕**               | 再转回 sRGB | 符合显示器特性                 |

对你这个开发者来说：**几乎无感**。你继续写：

```go
color.NRGBA{R: 0xFF, G: 0x80, B: 0x00, A: 0xFF}  // 直接用 sRGB 值
```

Gio 内部会自动处理转换。

**5. 实际影响与注意事项**

- **半透明叠加、阴影、抗锯齿**会比很多框架更干净。
- **图片**：`paint.ImageOp` 也会按同样规则处理（假设输入是 sRGB）。
- **不要自己手动转线性**：Gio 已经帮你做了，再转一次反而错误。
- **与其他程序对比**：如果你把 Gio 的半透明效果和直接在 sRGB 混合的程序放一起，颜色会略有不同（Gio 更正确）。

**6. 与前后章节的关系**

- 第2章 Drawing：`paint.ColorOp` 接收的就是 sRGB 的 `color.NRGBA`
- 第6章 Theme：主题里定义的所有颜色都是 sRGB 值
- 第8章 Text：文字颜色同样走这套流程
- **第9章**解释了为什么你可以放心直接用普通颜色值，却能获得正确的混合结果

---

**本章核心记忆点**

> **Gio 输入用 sRGB，混合用线性空间**  
> **你只需要直接写普通的 sRGB 颜色值即可**  
> **内部自动转换，保证半透明和抗锯齿颜色正确**  
> **这是目前 UI 框架中较为先进且务实的选择**

---

### 九章整体架构回顾

到这里，Gio Architecture 的 9 大核心章节全部解读完毕。我们可以用一张逻辑图把它们串起来：

```
第1章 Window
    ↓ 提供 FrameEvent 循环 + op.Ops
第2章 Drawing
    ↓ 提供 clip + paint 绘制原语
第3章 Input
    ↓ 提供基于 clip 的事件注册与路由
第4章 Widget
    ↓ 把 Drawing + Input + 尺寸计算打包成可复用单元
第5章 Layout
    ↓ 提供组合多个 Widget 的约束布局系统
第6章 Theme
    ↓ 提供统一的视觉风格（Material Design）
第7章 Units
    ↓ 解决跨设备尺寸一致性（Dp / Sp）
第8章 Text
    ↓ 解决字体与文字渲染
第9章 Color
    ↓ 解决正确的颜色混合
```

**Gio 的核心哲学总结**：

1. **即时模式**：每帧完整描述 UI，无隐式控件树
2. **零分配**：操作是值类型 + `Add` 方法
3. **操作列表**：绘制与输入的统一中间表示
4. **约束布局**：父给约束，子返回尺寸
5. **状态与视觉分离**：`widget` 管状态，`material` 管长相
6. **务实的正确性**：颜色、单位、坐标系都做了深思熟虑的选择

如果你后续想深入某个具体主题（例如自定义控件完整示例、List 性能优化、多窗口、动画最佳实践、或者把这九章整理成一份完整 Markdown/PDF），随时告诉我即可。