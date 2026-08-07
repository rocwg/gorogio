

```powershell
请您指导、评价、和纠正：
我当下浅薄的认知和安排
（1）私人项目：gorogio 项目（go & gio 的实践和学习）
（2）公司项目：wisdom-forestry（当下进行中；以后，公司会有其它spring boot 项目）

（3）今天，想摸鱼一下，继续 gorogio；但是我忘记了我进行到哪里了？

PS D:\roc-github\gio-workspace> goro tree -depth=2
.
├── gio         # 来源gio官方：git clone https://github.com/gioui/gio.git
├── gio-example # 来源gio官方：git clone https://github.com/gioui/gio-example
├── gio-x       # 来源gio官方：git clone https://github.com/gioui/gio-x.git
├── gioview     # 来源社区：git clone https://github.com/oligo/gioview
├── go.work
├── go.work.sum
├── gobyexample # 来源go官方：git clone https://github.com/mmcgrana/gobyexample.git
└── gorogio # 我自己新建的项目，在您的指导下完成
    ├── .gitattributes
    ├── .gitignore
    ├── README.md
    ├── app
    │   └── app.go
    ├── component # 负责交互和复合能力：有行为、有状态、组合出来的 UI 单元。
    │   ├── button.go
    │   └── component.go
    ├── container # 负责空间关系：东西怎么摆
    │   ├── column.go
    │   ├── container.md
    │   ├── expand.go
    │   ├── flex.go
    │   ├── options.go
    │   ├── row.go
    │   └── stack.go
    ├── element   # 所有 UI 节点协议
    │   └── element.go
    ├── examples
    │   ├── go   # 案例：对 go 的学习、实践、验证
    │   ├── gio  # 案例：纯 gio 实现的
    │   └── gorogio # 案例：在 纯 gio 实现的案例，重新使用 gorogio 改造的。
    ├── go.mod
    ├── go.sum
    ├── internal # 当下空，占位。
    ├── modifier # 负责外观修饰：东西怎么装饰？
    │   └── padding.go
    ├── page
    │   └── page.go
    ├── patterns
    ├── primitive # 最基础的、无状态的视觉元素。
    │   ├── space.go
    │   └── text.go
    ├── state
    │   └── state.go
    └── style
        └── theme.go

(4) 今早，我看了几篇文章，地址如下
4.1 https://tonybai.com/2026/06/04/master-new-tech-in-ai-era-counter-intuitive-learning-guide/
4.2 https://tonybai.com/2026/05/16/go-cured-my-over-engineering-addiction-after-java-ts/
4.3 https://tonybai.com/2026/08/03/go-simplicity-philosophy-debate-reddit/

（5）今早，我看到文章：https://tonybai.com/2026/08/04/wails-v3-go-desktop-framework/
5.1 我思考后，还是舍不得 "go & gio": Learning Go Through Gio

Go（真正的学习目标）
        ▲
        │
        │
Gio（Learning Vehicle）学习 Go 的真实工程场景。
        │
        ▼
Prototype Engine（长期产出）

5.2 也舍不得如下路径，怎么办？
      Software Creation Capability
        ┌──────────────────────┐
        │                      │
        ↓                      ↓
Go + Gio                    API Engineering
创造工具能力                 架构设计能力
        ↓                      ↓
Prototype Engine          Backend Platform
        └──────────┬───────────┘
                   ↓
          Personal Software Factory
```

### ==A1== 

我认真看完了，也把我们这几个月的讨论串起来了。

先说结论：

**我认为你现在的整体方向是对的，而且比两三个月前清晰得多。**
真正需要调整的，不是路线，而是**每条路线的定位和节奏**。

------

# 一、先评价你的三个项目安排

你现在实际上已经形成了三个层次。

```
① 公司项目（谋生）
        │
        ▼
Spring Boot
Java
真正业务系统
② 私人长期工程（能力）
API Engineering
Contract First
TypeSpec
Provider
③ 私人实验室（学习）
gorogio
Go
Gio
```

我觉得这是一个很健康的结构。

三者不要互相竞争。

而应该互相服务。

例如：

```
公司项目
      │
      │ 学习业务架构
      ▼

API Engineering
      │
      │ 提炼设计能力
      ▼

gorogio
      │
      │ 学Go语言
      ▼

Prototype Engine
```

它们并不是三个目标。

而是一棵树。

------

# 二、我想纠正一个地方

今天我最想提醒你的，就是这一点。

## gorogio 不是你的最终目标。

而是：

> **Learning Vehicle（学习载体）**

这个定位千万不要丢。

你自己以前其实已经总结出来了：

```
Learning Go Through Gio

Go
↑
真正目标

Gio
↑
Learning Vehicle

Prototype Engine
↑
副产品
```

我仍然坚持这个判断。

原因非常简单。

如果未来 Gio 消失了：

你的 Go 能力不会消失。

如果未来 Gio 不更新：

你的工程能力不会消失。

如果未来你改成：

```
Go + Fyne
```

或者

```
Go + Wails
```

你的 Go 仍然在。

所以真正投资的是：

> Go。

不是 Gio。

------

# 三、但是 Prototype Engine 要不要继续？

我的答案：

**继续。**

但是：**不要为了造 Framework 而造 Framework。**

这是今天 Tony Bai 那几篇文章，其实一直在提醒大家的一件事。

尤其第二篇：

> Go cured my over-engineering addiction...

里面最核心的一句话其实可以概括成：

> 不要把未来假设成今天。

我觉得这句话非常适合你。

因为以前我们讨论的时候，你有一段时间一直想做：

```
Compose
+
Flutter
+
SwiftUI
+
Vue
+
React
+
Qt
+
Godot
+
......
```

最后会变成：

```
超级Framework
```

这是 Go 社区最容易提醒大家避免的方向。

------

# 四、你的 gorogio 已经开始回归正确方向

我看你现在目录：

```
primitive
container
modifier
component
page
style
```

这很好。

因为它没有：

```
Observable

Provider

Compose Runtime

Scheduler

Lifecycle Owner

Remember

Snapshot

Recomposer
```

这些东西。

说明：

你已经开始尊重 Gio。

而不是复制 Compose。

这是很大的进步。

------

# 五、我甚至建议你冻结 v0.2.x

我以前其实就说过。

现在更建议。

例如：

```
v0.2.x

冻结

以后不再折腾API
```

例如：

```
Element

Container

Modifier

Component

Page

Theme
```

就够了。

以后真正增加的，应该是：

```
更多组件
```

例如：

```ini
Card、Input、Dialog、List、Table、Menu、Tree、Toast、Tabs、Navigation
```

而不是：

```
再设计一套新的状态管理
```

或者：

```
再发明一个新的Layout
```

------

# 六、你今天问："我忘了做到哪里了"

我帮你回忆一下。

我们最后一次真正冻结的是：

## v0.2.3

完成：

✅ Page 接口

```
Page
```

完成。

------

完成：

Application

管理 Page 生命周期。

------

Hello Example

改造成：

```
Page
↓
Element()
↓
Render()
```

这一套。

------

然后。

你提出：

```
Container API

是否冻结？
```

我们讨论了：

```
Row

Column

Stack

Expand
```

最终决定：

采用：Column、Row

不要采用：VStack、HStack

原因：更符合 Go 社区。

------

之后，

我们计划进入：

# v0.3

主题叫：

> **Interactive Components（交互组件）**

当时我们还讨论过：

Button 为什么：

```
Update()

Element()
```

分离。

以及：

如何保持：Immediate Mode。

而不要变成：Compose Runtime。

所以，

**我们其实停在了 v0.3 的入口。**

------

# 七、今天看完 Tony Bai 那几篇文章以后，我建议路线微调一下

我会把未来几个月改成这样。

```
gorogio
```

目标：不是 Framework。

而是：**Go 学习项目。**

------

里面每新增一个组件，其实对应学习一种 Go 能力。

| 例如： | 学习：                      |
| ------ | --------------------------- |
| Button | interface、method、callback |
| List   | slice、for、iterator        |
| Tree   | 递归、pointer               |
| Table  | generic、sorting            |
| Editor | buffer、unicode、io         |

这样，每个组件，都变成一次 Go 的工程实践。

而不是一次框架设计。

------

# 八、关于 Wails v3

你今天看到 Wails v3，我反而觉得你的反应很成熟。

以前的你可能会说：

> 我要不要马上换？

今天你的回答是：

> 我舍不得 Go + Gio。

我认为这是一个好现象。

因为你已经不是在追框架。

而是在坚持自己的学习主线。

实际上：

- **Wails** 擅长：用 Web 技术开发桌面应用，生态成熟，适合已有前端栈。
- **Gio** 擅长：纯 Go 原生绘制、跨平台、自定义能力强，更像一个即时渲染 UI 引擎。

它们解决的问题不同，并不是简单的替代关系。

以你目前的目标——**通过真实工程学习 Go，并逐步积累自己的原型开发能力**——继续走 Gio 是完全合理的。

如果未来某个产品需要 Web UI，再引入 Wails 也不晚。

------

# 九、最后，我会把你的长期路线压缩成一句话

这是我觉得经过这几个月讨论后，最准确的一版：

```text
                     公司项目
                (Spring Boot / Java)
                         │
                         ▼
                  学习业务与架构实践
                         │
         ┌───────────────┴───────────────┐
         ▼                               ▼
  API Engineering                 Learning Go Through Gio
 （架构设计能力）                    （语言与工程能力）
         │                               │
         └───────────────┬───────────────┘
                         ▼
              Personal Software Factory
```

如果让我给你一个今天下午最值得做的建议，我不会建议继续设计新的框架抽象，而是回到我们停下来的地方：**启动 gorogio v0.3**，但把它重新定义为 **Interactive Components & Go Engineering Practice**。这样，你写的每一个组件既能丰富原型引擎，也能扎扎实实提升 Go 语言和工程能力。这条路线，与我们前面确定的长期目标是一致的。



### ==A2== 

这个问题我非常愿意认真展开。

因为这四篇文章放在一起看，其实不是四个独立主题。

它们共同回答一个问题：

> **在 AI 时代，一个有多年 Java/企业开发背景的程序员，为什么还值得花几年时间学习 Go，并且如何学习？**

而且非常巧合的是，这四篇文章和你现在的：

- gorogio
- API Engineering
- Personal Software Factory

几乎是一一对应的。

我先给一个总览：

| 文章                  | 表面主题      | 深层主题                             |
| --------------------- | ------------- | ------------------------------------ |
| 4.1 AI时代掌握新技术  | 如何学习 Go   | 如何建立真正能力，而不是 AI 操作能力 |
| 4.2 Go 治好了过度设计 | Go vs Java/TS | 从抽象驱动转向工程价值驱动           |
| 4.3 Go 简洁哲学争论   | Go 为什么保守 | 降低认知负担，长期维护               |
| 4.4 Wails v3          | Go 桌面开发   | Go 生态的应用边界扩展                |

我逐篇分析。

------

# 4.1 《AI时代如何真正掌握一门新技术？》

[AI 时代如何真正掌握一门新技术？这份非主流学习指南建议永久收藏](https://tonybai.com/2026/06/04/master-new-tech-in-ai-era-counter-intuitive-learning-guide/?utm_source=chatgpt.com)

## 核心问题

文章其实不是讲 Go。

它讲：

> AI 可以快速生成答案，但是不能替你建立内部模型。

作者提到一个非常现实的问题：

很多人现在学习：

```
遇到问题

↓

问 AI

↓

复制代码

↓

运行成功
```

但是：

```
代码跑了
≠
能力形成
```

这点我觉得和你非常相关。

------

## 文章提出四个阶段

### 第一阶段：

## Cold Turkey（冷启动）

意思：

刚学习一个东西的时候，不要过度依赖 AI。

例如：

学习 Go：

不要：

```
帮我写一个 Gin CRUD
```

而应该：

自己写：

```go
type User struct {
    ID int
    Name string
}
```

然后遇到：

```
为什么 pointer？
为什么 interface？
为什么 error?
```

再问。

------

### 第二阶段：

## 先建立地图

这是非常重要的一点。

很多人学技术：今天 goroutine，明天 channel，后天 context；但是脑子里面没有结构。

正确：

先知道：

```
Go
|
+-- 基础语法
|
+-- 类型系统
|
+-- interface
|
+-- concurrency
|
+-- package
|
+-- tooling
|
+-- standard library
```

然后填。

([Tony Bai](https://tonybai.com/2026/06/04/master-new-tech-in-ai-era-counter-intuitive-learning-guide/?utm_source=chatgpt.com))

------

### 第三阶段：

## AI 从代码生成器变成导师

这个特别适合我们现在的交流方式。

不是：

> 给我代码

而是：

> 为什么这样设计？
>
> 有什么替代方案？
>
> 这个 abstraction 是否必要？

其实你最近问我的：

> Row / Column 要不要叫 VStack / HStack？

> Component 是否应该存在？

> Element 是否合理？

这些问题就是这种学习方式。

------

### 第四阶段：

## 自己写 → AI Review → 重构

这个我认为是你现在 gorogio 最正确的模式。

比如：

你写：

```go
type Page interface {
    Update(gtx layout.Context)
}
```

然后问：

> 这个 interface 是否过度？

而不是：

> 帮我设计一个 UI Framework。

------

# 对你的启示

这一篇最直接支持：

## gorogio 路线

因为：

Gio 本身不是目的。

它提供：

真实工程环境。

你需要：

- interface
- package
- layout
- state
- event
- rendering

这就是：

> Learning Go Through Gio

------

# 4.2 《Go 语言终于治好了我的过度设计》

[写了10年Java/TS，Go语言终于治好了我的“过度设计”绝症](https://tonybai.com/category/technical-notes/page/13/?utm_source=chatgpt.com)

这一篇，我觉得你应该重点看。

因为它几乎就是：

你的 Java 背景 → Go 转型。

------

## Java/TS 世界容易形成什么习惯？

例如：

一个打印：

Java：

```java
interface Printer {}

class ConsolePrinterImpl 
        implements Printer {}

class PrinterFactory {}

class PrinterConfig {}

@Bean
Printer printer()
```

最后：

打印一句：

```java
System.out.println()
```

变成：

50 个文件。

文章批评的是：

> 抽象成为竞技运动。

([Tony Bai](https://tonybai.com/tag/Go/page/15/?utm_source=chatgpt.com))

------

## Go 的治疗方式是什么？

Go 强迫你：

直接：

```go
fmt.Println("hello")
```

然后：

如果真的出现变化：

再抽象。

也就是：

Java：

```
未来可能变化
        ↓
提前设计
```

Go：

```
真实变化出现
        ↓
设计
```

------

## 这对你的影响非常大

因为你的历史背景：

Spring Boot

DDD

Clean Architecture

六边形

CQRS

Saga

这些都是好东西。

但是风险：

容易：

```
领域模型
↓
Repository
↓
Port
↓
Adapter
↓
Factory
↓
Provider
↓
Interface
```

最后：

业务没几个。

架构很多。

------

但是我要强调：

这不代表：

DDD 错。

而是：

**复杂度必须来自业务，而不是来自架构模板。**

------

# 4.3 《Go 简洁哲学争论》

[Go 对语言演化的保守态度，在未来5年是否仍然正确呢？](https://tonybai.com/2026/07/06/is-go-language-conservative-evolution-still-viable-next-5-years/?utm_source=chatgpt.com)

这一篇和 4.2 是连续的。

核心：

> Go 为什么故意不先进？

------

很多语言：

不断增加：

- syntax sugar
- meta programming
- advanced type system

例如：

Rust：

非常强。

TypeScript：

表达力强。

Scala：

非常强。

但是：

代价：

认知成本。

------

Go 选择：

```
少一点语言能力

换取

团队一致性
长期维护
快速阅读
```

------

举个例子：

Go interface：

```go
type Reader interface {
    Read([]byte)
}
```

没有：

implements。

没有：

annotation。

没有：

继承树。

结果：

新人看代码：

知道。

------

这也是为什么：

很多基础设施：

Kubernetes

Docker

Prometheus

都是 Go。

因为：

维护周期长。

------

# 对你的 gorogio 的影响

这里其实是在提醒你：

不要造：

Compose。

不要造：

Flutter。

你的设计：

现在：

```
Element
Container
Primitive
Modifier
Component
```

很好。

但是不要继续：

增加：

```
Reactive Runtime

Dependency Injection

Lifecycle Manager

Virtual DOM

Compiler
```

否则：

你会重新走 Java 路。

------

# 4.4 《Wails v3》

[使用 Go 构建桌面应用](https://v3.wails.io/zh-cn/?utm_source=chatgpt.com)

这个比较特殊。

因为它触发了你的：

> "是不是应该放弃 Gio?"

这个问题。

------

先说 Wails 是什么。

Wails：

本质：

```
Go Backend

+

Web UI
```

通过系统 WebView。

官方定位：

> 使用 Go 和现代 Web 技术构建桌面应用。([Wails](https://v3.wails.io/zh-cn/?utm_source=chatgpt.com))

类似：

```
Electron
(TS + Chromium)

↓

Wails
(Go + OS WebView)
```

------

## Wails 优势

非常明显：

如果你已经会：Vue、React、Svelte，那么：开发效率极高。

例如：

```
复杂表格

后台系统

管理工具

业务软件
```

非常合适。

------

## Gio 优势

完全不同。

Gio：

```
Go
↓
Immediate Mode UI
↓
GPU rendering
```

你得到：

- Go 深度能力
- UI 系统理解
- 渲染模型理解

它更像：学习计算机图形学 + UI Framework。

------

所以：

不要比较：

```
Wails > Gio
```

这是错误问题。

正确：

```
产品开发效率:
Wails

学习Go和UI原理:
Gio
```

------

# 五、四篇文章合起来，对你的路线有什么指导？

我认为答案非常明确。

你的路线：不要改。

但是调整目标。

## 原来：

可能：

```
我要做 gorogio Framework
```

容易危险。

改：

```
我要通过 gorogio 学 Go

顺便形成 Prototype Engine
```

------

## 原来：

可能：

```
我要设计完美 API Architecture
```

危险。

改：

```
我要形成 API Design Capability
```

------

## 原来：

可能：

```
我要追最新技术
```

危险。

改：

```
我要建立长期复利资产
```

------

# 六、如果把四篇文章浓缩成一句话

我觉得是：

> **AI 时代，不应该竞争写代码速度，而应该培养理解复杂系统、做出取舍、长期维护的能力。**

而你的两个私人项目：

## API Engineering

培养：

```
设计能力
```

## gorogio

培养：

```
创造能力
```

刚好形成：

```
Architecture Capability
+
Software Creation Capability
↓
Personal Software Factory
```

所以我反而觉得：

你今天看到 Wails 后仍然舍不得 Gio，不是执着。

而是说明：

你已经从“追框架”开始转向“培养能力”。

这是一个很好的变化。
