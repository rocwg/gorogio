收到了！你直接跳转到了规范中非常关键的一章：**语句（Statements）** 及其基础语法定义，并引出了**终止语句（Terminating statements）** 的概念。

我们继续采用【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】 的方式，为你逐字对齐解读：

## 章节：语句 (Statements)

### 段落 1

> **【英文原文】**
>
> Statements¶
>
> Statements control execution.

**【逐字精准翻译】**

语句¶

语句用于控制（程序的）执行。

- **词汇与句式剖析：**
  - `Statement`：语句（构成程序执行逻辑的基本单元，与产生值的“表达式 Expression”不同，语句通常用于改变状态或控制流程）。
  - `control execution`：控制执行（流）。

### EBNF 产生式定义

> **【英文原文】**
>
> EBNF
>
> ```
> Statement  = Declaration | LabeledStmt | SimpleStmt |
>              GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
>              FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
>              DeferStmt .
> 
> SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
> ```

**【逐字精准翻译】**

EBNF

```
语句       = 声明 | 标签语句 | 简单语句 |
             Go语句 | Return语句 | Break语句 | Continue语句 | Goto语句 |
             Fallthrough语句 | 块 | If语句 | Switch语句 | Select语句 | For语句 |
             Defer语句 .

简单语句   = 空语句 | 表达式语句 | 发送语句 | 自增自减语句 | 赋值语句 | 短变量声明 .
```

- **词汇剖析与分支拆解：**
  - `Declaration`：声明（如 `var`, `const`, `type`）。
  - `LabeledStmt`：带标签的语句（如 `MyLoop: for ...`）。
  - `GoStmt` / `DeferStmt`：`go` 协程启动语句 / `defer` 延迟调用语句。
  - `Block`：代码块（用花括号 `{}` 包裹的一组语句）。
  - `SendStmt`：发送语句（针对 Channel 的发送操作，如 `ch <- value`）。
  - `IncDecStmt`：自增自减语句（如 `i++` 或 `i--`；注意 Go 中 `i++` 是**语句**而非表达式！）。
  - `ShortVarDecl`：短变量声明（使用 `:=` 的声明，如 `x := 1`）。

> **规范深度概念——为什么单独定义 `SimpleStmt`（简单语句）？**
>
> 在 Go 语法中，很多复合语句（如 `if`, `switch`, `for`）的头部都允许附带一个初始化语句，例如：
>
> ```
> if err := doSomething(); err != nil
> ```
>
> 这里的 `err := doSomething()` 就是一个 `SimpleStmt`。Go 规范通过将 `SimpleStmt` 提取为一个独立的范式，使得这些控制流语句的语法定义极度统一和紧凑。

### 标题引出

> **【英文原文】**
>
> Terminating statements¶

**【逐字精准翻译】**

终止语句¶

- **词汇剖析：**
  - `Terminating`：终止的 / 结束的。
  - **预告概念：** “终止语句”在 Go 语言规范中有着非常严格的控制流含义（例如：一个带有返回值类型的函数，其控制流的末尾必须是一个终止语句，否则编译器会报错 `missing return`）。

你引用的文本正好停在 **Terminating statements** 这一小节的开头。

如果你手头有这一小节接下来的原文，请随时发出来，我们继续逐字翻译并深入剖析它的具体判定规则！

非常关键的一个章节！**终止语句（Terminating statements）** 是 Go 编译器进行**控制流分析（Control Flow Analysis）** 的核心依据——例如编译器如何判断一个有返回值的函数是否在所有路径上都有 `return`，全靠本节定义的这套法则。

我们继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，逐字逐句进行深度解析：

## 章节：终止语句 (Terminating statements)

### 段落 1：定义与总纲

> **【英文原文】**
>
> Terminating statements¶
>
> A terminating statement interrupts the regular flow of control in a block. The following statements are terminating:

**【逐字精准翻译】**

终止语句¶

终止语句会中断块中常规的控制流。以下语句属于终止语句：

- **词汇与句式剖析：**
  - `interrupts`：中断 / 打断。
  - `regular flow of control`：常规控制流（指代码自上而下顺序执行的流程）。
  - `block`：代码块（即 `{}` 包裹的区域）。

### 条目 1：Return / Goto 语句

> **【英文原文】**
>
> A "return" or "goto" statement.

**【逐字精准翻译】**

"return"（返回）或 "goto"（跳转）语句。

- **逻辑拆解：** `return` 会直接退出当前函数，`goto` 会直接跳转到其他位置，因此它们必然会中断当前代码块的顺序执行。

### 条目 2：Panic 调用

> **【英文原文】**
>
> A call to the built-in function panic.

**【逐字精准翻译】**

对内置函数 `panic` 的调用。

- **词汇剖析：**
  - `built-in function`：内置函数（如 `panic`, `recover`, `len` 等）。
- **逻辑拆解：** 调用 `panic` 会引发恐慌并启动栈展开（Stack unwinding），除非被 `recover` 捕获，否则程序不会接着往下顺序执行。

### 条目 3：以终止语句结尾的代码块

> **【英文原文】**
>
> A block in which the statement list ends in a terminating statement.

**【逐字精准翻译】**

其语句列表以终止语句结尾的代码块。

- **逻辑拆解：** 只要一个花括号 `{ ... }` 包裹的代码块，最后一条有效语句是终止语句，整个代码块就被视为一个终止语句。

### 条目 4：If 语句

> **【英文原文】**
>
> An "if" statement in which:
>
> - the "else" branch is present, and
> - both branches are terminating statements.

**【逐字精准翻译】**

满足以下条件的 "if" 语句：

- 存在 "else" 分支，且
- 两个分支都是终止语句。

- **词汇剖析：**

  - `branch`：分支。
  - `is present`：存在 / 提供了。

- **工程示例：**

  ```go
  if cond {
      return 1 // 分支 1：终止
  } else {
      return 2 // 分支 2：终止
  } // 整个 if-else 组合被视作一个终止语句！
  ```

### 条目 5：For 循环语句

> **【英文原文】**
>
> A "for" statement in which:
>
> - there are no "break" statements referring to the "for" statement, and
> - the loop condition is absent, and
> - the "for" statement does not use a range clause.

**【逐字精准翻译】**

满足以下条件的 "for" 语句：

- 不存在引用该 "for" 语句的 "break" 语句，且
- 循环条件不存在（即省略了条件表达式），且
- 该 "for" 语句没有使用 range 子句。

- **词汇与句式剖析：**
  - `referring to`：引用 / 指向（指针对当前 `for` 的 `break` 或 `break Label`）。
  - `is absent`：不存在 / 省略了（例如 `for { ... }` 死循环）。
  - `range clause`：range 子句（如 `for k, v := range slice`）。
- **逻辑拆解：** 只有**无条件死循环**（`for { ... }`）且**内部没有任何 `break` 能跳出**时，这个 `for` 语句才被判定为终止语句。

### 条目 6：Switch 语句

> **【英文原文】**
>
> A "switch" statement in which:
>
> - there are no "break" statements referring to the "switch" statement,
> - there is a default case, and
> - the statement lists in each case, including the default, end in a terminating statement, or a possibly labeled "fallthrough" statement.

**【逐字精准翻译】**

满足以下条件的 "switch" 语句：

- 不存在引用该 "switch" 语句的 "break" 语句，
- 存在一个 default（默认）分支，且
- 每个 case（包括 default 分支）中的语句列表均以终止语句结尾，或者以一个（可能带标签的）"fallthrough" 语句结尾。

- **词汇与句式剖析：**
  - `default case`：默认分支。
  - `possibly labeled`：可能带有标签的。
  - `fallthrough`：贯穿语句（Go 中用于让 `switch` 强制执行下一个 case 块的关键字）。

### 条目 7：Select 语句

> **【英文原文】**
>
> A "select" statement in which:
>
> - there are no "break" statements referring to the "select" statement, and
> - the statement lists in each case, including the default if present, end in a terminating statement.

**【逐字精准翻译】**

满足以下条件的 "select" 语句：

- 不存在引用该 "select" 语句的 "break" 语句，且
- 每个 case（包括存在时的 default 分支）中的语句列表均以终止语句结尾。

- **逻辑拆解：** `select` 必须穷尽所有可能的通道分支（若有 `default` 也包含在内），且每个分支都必须以终止语句结尾，且不能中途 `break`。

### 条目 8：带标签的终止语句

> **【英文原文】**
>
> A labeled statement labeling a terminating statement.

**【逐字精准翻译】**

标记了一个终止语句的带标签语句。

- **示例：** `MyLabel: return`（由于 `return` 是终止语句，因此给它打上标签后的整个语句依然是终止语句）。

### 补充规则：非终止声明与列表结尾定义

> **【英文原文】**
>
> All other statements are not terminating.
>
> A statement list ends in a terminating statement if the list is not empty and its final non-empty statement is terminating.

**【逐字精准翻译】**

所有其他语句均不是终止语句。

如果一个语句列表非空，且其最后一个非空语句是终止语句，则该语句列表以终止语句结尾。

- **词汇与句式剖析：**
  - `statement list`：语句列表（即多条连续书写的语句）。
  - `final non-empty statement`：最后一个非空语句（排除掉了空分号 `;` 等空语句）。

### 预告：下一小节

> **【英文原文】**
>
> Empty statements¶

**【逐字精准翻译】**

空语句¶

你引用的这一段完全展现了 Go 编译器在语法检查时逻辑的极其严密性！

如果准备好继续，我们可以直接推进 **空语句 (Empty statements)** 及其后续的 **表达式语句 (Expression statements)** 等内容！

继续保持精准的【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】 标准，为你剖析 **空语句**、**标签语句** 和 **表达式语句** 的底层规范：

## 章节：空语句 (Empty statements)

> **【英文原文】**
>
> Empty statements¶
>
> The empty statement does nothing.
>
> ```
> EmptyStmt = .
> ```

**【逐字精准翻译】**

空语句¶

空语句不执行任何操作。

```
空语句 = .
```

- **词汇与句式剖析：**
  - `does nothing`：不执行任何操作 / 无动作。
  - **EBNF 产生式解读：** `EmptyStmt = .` 中等号后面直接紧跟句点 `.`，代表该产生式右侧没有任何词法单元（即纯粹的空白或单个独立的半角分号 `;`）。在 Go 中，多个连续的分号（如 `;;;`）会被解析为多个空语句。

## 章节：标签语句 (Labeled statements)

> **【英文原文】**
>
> Labeled statements¶
>
> A labeled statement may be the target of a goto, break or continue statement.
>
> ```
> LabeledStmt = Label ":" Statement .
> Label       = identifier .
> Error: log.Panic("error encountered")
> ```

**【逐字精准翻译】**

标签语句¶

标签语句可以是 goto、break 或 continue 语句的跳转目标。

```
标签语句 = 标签 ":" 语句 .
标签     = 标识符 .
Error: log.Panic("error encountered")
```

- **词汇与句式剖析：**
  - `target`：目标 / 跳转对象。
  - `identifier`：标识符（即合法的变量/函数命名格式）。
  - **代码示例拆解：** `Error: log.Panic("error encountered")` 中，`Error` 是标签（`Label`），`log.Panic(...)` 是被标记的语句（`Statement`）。

## 章节：表达式语句 (Expression statements)

### 段落 1：基本法则

> **【英文原文】**
>
> Expression statements¶
>
> With the exception of specific built-in functions, function and method calls and receive operations can appear in statement context. Such statements may be parenthesized.
>
> ```
> ExpressionStmt = Expression .
> ```

**【逐字精准翻译】**

表达式语句¶

除特定的内置函数外，函数调用、方法调用以及（通道）接收操作可以出现在语句上下文中。此类语句可以使用圆括号括起来。

```
表达式语句 = 表达式 .
```

- **词汇与句式剖析：**
  - `With the exception of`：除……以外 / ……除外。
  - `statement context`：语句上下文（指不需要将表达式的返回值赋值给任何变量，直接作为独立的一行代码执行）。
  - `parenthesized`：被加了圆括号的。

### 段落 2：被禁止独立成句的内置函数列表

> **【英文原文】**
>
> The following built-in functions are not permitted in statement context:
>
> ```
> append cap complex imag len make max min new real
> unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice unsafe.SliceData unsafe.String unsafe.StringData
> ```

**【逐字精准翻译】**

以下内置函数不允许出现在语句上下文中：

```
append cap complex imag len make max min new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice unsafe.SliceData unsafe.String unsafe.StringData
```

- **词汇与句式剖析：**
  - `not permitted`：不允许的 / 被禁止的。
- **规范底层原理：**
  - 这些内置函数都是**纯纯用于产生新值或读取属性**的（无副作用/Side-effects）。
  - 编译器禁止你写单独一行的 `len(s)` 或 `make([]int, 10)`，因为丢弃它们的返回值没有任何工程意义，必定属于程序员的逻辑笔误。相反，`copy` 和 `delete` 等有副作用的内置函数则允许独立成句。

### 段落 3：代码示例分析

> **【英文原文】**
>
> ```go
>h(x+y)
> f.Close()
> <-ch
> (<-ch)
> len("foo")  // illegal if len is the built-in function
> ```

**【逐字精准翻译】**

```go
h(x+y)      // 合法：函数调用
f.Close()   // 合法：方法调用
<-ch        // 合法：通道接收操作（丢弃接收到的值，仅用于等待信号）
(<-ch)      // 合法：带圆括号的通道接收操作
len("foo")  // 非法：如果 len 是内置函数的话
```

- **词汇与句式剖析：**
  - `illegal`：非法的 / 编译不通过的。
  - **注释深入解答：** 为什么注释里要强调 "if len is the built-in function"？因为在前几章讲过，`len` 是预声明标识符（Predeclared identifier）而非关键字。如果你在局部作用域自己定义了一个叫 `len` 的自定义函数，那么 `len("foo")` 就可以独立成句；但如果是内置的 `len`，就会触发编译器报错！

你引用的文本正好停在 **Send statements (发送语句)** 的开头。

准备好后，我们随时继续推进 `Send statements`（通道发送）、`IncDecStmt`（自增自减语句）以及 `Assignment`（赋值语句）！

我们继续保持 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的严密标准，逐字剖析 **发送语句（Send statements）** 与 **自增自减语句（IncDec statements）**：

## 章节：发送语句 (Send statements)

### 段落 1：类型约束与语法结构

> **【英文原文】**
>
> Send statements¶
>
> A send statement sends a value on a channel. The channel expression must be of channel type, the channel direction must permit send operations, and the type of the value to be sent must be assignable to the channel's element type.
>
> ```
> SendStmt = Channel "<-" Expression .
> Channel  = Expression .
> ```

**【逐字精准翻译】**

发送语句¶

发送语句在通道上发送一个值。通道表达式必须是通道类型，通道方向必须允许发送操作，并且要发送的值的类型必须可以赋值给该通道的元素类型。

```
发送语句 = 通道 "<-" 表达式 .
通道     = 表达式 .
```

- **词汇与句式剖析：**
  - `sends a value on a channel`：在通道上发送一个值。
  - `channel direction`：通道方向（指通道是双向 `chan T`，还是只写 `chan<- T`）。
  - `permit send operations`：允许发送操作（即不能对只读通道 `<-chan T` 执行发送）。
  - `assignable to`：可赋值给……（要求类型兼容或满足隐式转换/赋权规则）。
  - `element type`：元素类型（通道里承载的数据类型 `T`）。

### 段落 2：运行时行为（阻塞、Panic 与死锁）

> **【英文原文】**
>
> Both the channel and the value expression are evaluated before communication begins. Communication blocks until the send can proceed. A send on an unbuffered channel can proceed if a receiver is ready. A send on a buffered channel can proceed if there is room in the buffer. A send on a closed channel proceeds by causing a run-time panic. A send on a nil channel blocks forever.

**【逐字精准翻译】**

通道表达式和值表达式都在通信开始之前进行求值。通信会一直阻塞，直到发送操作可以继续进行。如果接收者已准备就绪，在无缓冲通道上的发送操作就可以继续进行。如果缓冲区中有空间，在有缓冲通道上的发送操作就可以继续进行。在已关闭的通道上进行发送操作，会通过引发运行时恐慌（run-time panic）来继续（终止）。在 nil（空）通道上的发送操作会永远阻塞。

- **词汇与句式剖析：**
  - `evaluated`：被求值 / 被计算出来。
  - `Communication blocks`：通信阻塞。
  - `can proceed`：可以继续进行 / 可以执行。
  - `unbuffered channel / buffered channel`：无缓冲通道 / 有缓冲通道。
  - `causing a run-time panic`：引发运行时恐慌（`panic: send on closed channel`）。
  - `blocks forever`：永远阻塞（会导致 Goroutine 泄漏或死锁）。

### 段落 3：代码示例

> **【英文原文】**
>
> ```
> ch <- 3  // send value 3 to channel ch
> ```

**【逐字精准翻译】**

```
ch <- 3  // 将值 3 发送到通道 ch
```

### 段落 4：泛型类型参数下的约束 (Go 1.18+)

> **【英文原文】**
>
> If the type of the channel expression is a type parameter, all types in its type set must be channel types that permit send operations, they must all have the same element type, and the type of the value to be sent must be assignable to that element type.

**【逐字精准翻译】**

如果通道表达式的类型是一个类型参数（泛型），则其类型集中的所有类型都必须是允许发送操作的通道类型，它们必须都具有相同的元素类型，并且要发送的值的类型必须可以赋值给该元素类型。

- **词汇与句式剖析：**
  - `type parameter`：类型参数（即泛型约束中的类型变量）。
  - `type set`：类型集（泛型接口约束所定义的具体类型集合）。

## 章节：自增自减语句 (IncDec statements)

### 段落 1：基本法则与可寻址性要求

> **【英文原文】**
>
> IncDec statements¶
>
> The "++" and "--" statements increment or decrement their operands by the untyped constant 1. As with an assignment, the operand must be addressable or a map index expression.
>
> ```
> IncDecStmt = Expression ( "++" | "--" ) .
> ```

**【逐字精准翻译】**

自增自减语句¶

"++" 和 "--" 语句使其操作数增加或减少无类型常量 1。与赋值语句一样，该操作数必须是可寻址的，或者是 map 的索引表达式。

```
自增自减语句 = 表达式 ( "++" | "--" ) .
```

- **词汇与句式剖析：**
  - `increment / decrement`：增加 / 减少。
  - `operand`：操作数。
  - `untyped constant 1`：无类型常量 1（因此不仅适用于 `int`，同样适用于 `float64` 等数值类型）。
  - `addressable`：可寻址的（即必须是一个可以获取内存地址的左值变量，不能对字面量或函数返回值如 `5++` 或 `foo()++` 执行自增）。
  - `map index expression`：map 索引表达式（例如 `m["key"]++`；虽然 map 的元素不可寻址，但规范特许其用于自增自减）。

### 段落 2：等价赋值语义对比

> **【英文原文】**
>
> The following assignment statements are semantically equivalent:
>
> Plaintext
>
> ```
> IncDec statement    Assignment
> x++                 x += 1
> x--                 x -= 1
> ```

**【逐字精准翻译】**

以下赋值语句在语义上是等价的：

Plaintext

```
自增自减语句         赋值语句
x++                 x += 1
x--                 x -= 1
```

- **词汇剖析：**
  - `semantically equivalent`：语义上等价的。
- **规范底层强调：**
  1. 在 Go 语言中，`x++` 是一个**语句（Statement）**，而不是**表达式（Expression）**！所以你**绝对不能**写 `y = x++` 或 `func(x++)`。
  2. Go 语言**没有前缀自增** `++x` 或 `--x`。

你引用的文本正好停在 **Assignment statements (赋值语句)** 的标题处。

如果准备好了，请发送赋值语句及后续的原文，我们接着逐字对齐解读！

继续保持 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，为您深度剖析 **赋值语句（Assignment statements）** 的底层规范：

## 章节：赋值语句 (Assignment statements)

### 段落 1：基本定义

> **【英文原文】**
>
> Assignment statements¶
>
> An assignment replaces the current value stored in a variable with a new value specified by an expression. An assignment statement may assign a single value to a single variable, or multiple values to a matching number of variables.
>
> ```
> Assignment = ExpressionList assign_op ExpressionList .
> assign_op  = [ add_op | mul_op ] "=" .
> ```

**【逐字精准翻译】**

赋值语句¶

赋值用由表达式指定的崭新值替换存储在变量中的当前值。赋值语句可以将单个值赋给单个变量，或者将多个值赋给匹配数量的变量。

```
赋值语句 = 表达式列表 赋值运算符 表达式列表 .
赋值运算符 = [ 加法运算符 | 乘法运算符 ] "=" .
```

- **词汇与句式剖析：**
  - `replaces ... with ...`：用……替换……。
  - `matching number of`：匹配/相同数量的。
  - `assign_op`：赋值运算符（包括普通赋值 `=` 以及复合赋值如 `+=`, `*=`, `&=` 等）。

### 段落 2：左值约束（Left-hand side operands）

> **【英文原文】**
>
> Each left-hand side operand must be addressable, a map index expression, or (for = assignments only) the blank identifier. Operands may be parenthesized.
>
> ```go
>x = 1
> *p = f()
> a[i] = 23
> (k) = <-ch  // same as: k = <-ch
> ```

**【逐字精准翻译】**

左侧的每个操作数必须是可寻址的、map 的索引表达式，或者（仅限 `=` 赋值）空标识符。操作数可以用圆括号括起来。

```go
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // 等同于：k = <-ch
```

- **词汇与句式剖析：**
  - `left-hand side operand`：左侧操作数（简称左值）。
  - `blank identifier`：空标识符（即下划线 `_`；注意：复合赋值运算符如 `_ += 1` 是绝对禁止的！）。

### 段落 3：复合赋值运算符（Compound assignments）

> **【英文原文】**
>
> An assignment operation `x op= y` where `op` is a binary arithmetic operator is equivalent to `x = x op (y)` but evaluates `x` only once. The `op=` construct is a single token. In assignment operations, both the left- and right-hand expression lists must contain exactly one single-valued expression, and the left-hand expression must not be the blank identifier.
>
> ```go
>a[i] <<= 2
> i &^= 1<<n
> ```

**【逐字精准翻译】**

其中 `op` 为二元算术运算符的赋值操作 `x op= y` 等价于 `x = x op (y)`，但对 `x` 仅求值一次。`op=` 结构是一个独立的词法单元（token）。在（复合）赋值操作中，左右两侧的表达式列表都必须恰好包含一个单值表达式，并且左侧表达式绝不能是空标识符。

```go
a[i] <<= 2
i &^= 1<<n
```

- **词汇与句式剖析：**
  - `binary arithmetic operator`：二元算术运算符。
  - `evaluates x only once`：对 x 仅求值一次（如 `getArray()[i] += 1` 中，`getArray()` 函数只会被执行一次，防范了副作用）。
  - `&^=`：位清空（Bit Clear）赋值运算符。

### 段落 4：元组赋值（Tuple assignment）

> **【英文原文】**
>
> A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. There are two forms. In the first, the right hand operand is a single multi-valued expression such as a function call, a channel or map operation, or a type assertion. The number of operands on the left hand side must match the number of values. For instance, if f is a function returning two values,
>
> `x, y = f()` assigns the first value to x and the second to y. In the second form, the number of operands on the left must equal the number of expressions on the right, each of which must be single-valued, and the nth expression on the right is assigned to the nth operand on the left:
>
> ```
> one, two, three = '一', '二', '三'
> ```

**【逐字精准翻译】**

元组赋值将多值操作的各个独立元素赋给变量列表。形式有两种。在第一种形式中，右侧操作数是单个多值表达式，例如函数调用、通道或 map 操作、或者类型断言。左侧操作数的数量必须与值的数量匹配。例如，如果 f 是一个返回两个值的函数，

```
x, y = f()
```

将第一个值赋给 x，将第二个值赋给 y。在第二种形式中，左侧操作数的数量必须等于右侧表达式的数量，右侧每个表达式都必须是单值的，且右侧第 n 个表达式被赋给左侧第 n 个操作数：

```
one, two, three = '一', '二', '三'
```

- **词汇与句式剖析：**
  - `tuple assignment`：元组赋值（实现多变量同时赋值或交换的技术）。
  - `multi-valued operation`：多值操作（如 `v, ok := m[k]` 或 `v, ok := <-ch`）。

### 段落 5：空标识符在赋值中的作用

> **【英文原文】**
>
> The blank identifier provides a way to ignore right-hand side values in an assignment:
>
> ```go
>_ = x       // evaluate x but ignore it
> x, _ = f()  // evaluate f() but ignore second result value
> ```

**【逐字精准翻译】**

空标识符提供了一种在赋值中忽略右侧值的方法：

```go
_ = x       // 对 x 求值，但忽略其结果
x, _ = f()  // 对 f() 求值，但忽略第二个返回值
```

### 段落 6：赋值的两个阶段与求值顺序（关键重点！）

> **【英文原文】**
>
> The assignment proceeds in two phases. First, the operands of index expressions and pointer indirections (including implicit pointer indirections in selectors) on the left and the expressions on the right are all evaluated in the usual order. Second, the assignments are carried out in left-to-right order.

**【逐字精准翻译】**

赋值分两个阶段进行。首先，左侧的索引表达式和指针间接引用（包括选择器中的隐式指针间接引用）的操作数，以及右侧的表达式，均按通常顺序进行求值。其次，赋值按照从左到右的顺序执行。

- **词汇与句式剖析：**
  - `proceeds in two phases`：分两个阶段进行。
  - `pointer indirections`：指针解引用 / 指针间接引用（如 `*p`）。
  - `implicit pointer indirections in selectors`：选择器中的隐式指针间接引用（例如 `p.x` 中 `p` 为指针时，自动解引用为 `(*p).x`）。

### 赋值阶段例证详解（代码段逐行分析）

> **【英文原文与代码例证】**
>
> ```go
>a, b = b, a  // exchange a and b
> 
> x := []int{1, 2, 3}
> i := 0
> i, x[i] = 1, 2  // set i = 1, x[0] = 2
> 
> i = 0
> x[i], i = 2, 1  // set x[0] = 2, i = 1
> 
> x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)
> 
> x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.
> 
> type Point struct { x, y int }
> var p *Point
> x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7
> 
> i = 2
> x = []int{3, 5, 7}
> for i, x[i] = range x {  // set i, x[2] = 0, x[0]
> 	break
> }
> // after this loop, i == 0 and x is []int{3, 5, 3}
> ```

**【逐字精准翻译与逻辑剖析】**

1. `a, b = b, a`
   - **解析：** 交换 a 和 b。第一阶段先算出右侧的 `b` 和 `a` 的值，第二阶段写入左侧。
2. `i, x[i] = 1, 2` $\rightarrow$ **将 i 设为 1，x[0] 设为 2**
   - **原理：** 第一阶段先求值左侧的索引 `x[i]`（此时 `i` 为 0，锁定目标为 `x[0]`）；第二阶段按从左往右赋值，依次赋予 `i = 1`，`x[0] = 2`。
3. `x[i], i = 2, 1` $\rightarrow$ **将 x[0] 设为 2，i 设为 1**
   - **原理：** 第一阶段锁定 `x[i]` 同样为 `x[0]`；第二阶段赋值，`x[0] = 2`，`i = 1`。
4. `x[0], x[0] = 1, 2` $\rightarrow$ **先设 x[0] = 1，接着设 x[0] = 2（最终 x[0] == 2）**
   - **原理：** 第二阶段从左到右覆盖。
5. `x[1], x[3] = 4, 5` $\rightarrow$ **将 x[1] 设为 4，接着在设置 x[3] 时引发 panic**
   - **原理：** 越界发生在第二阶段赋值过程中，此时 `x[1]` 已经被成功写入为 4。
6. `x[2], p.x = 6, 7` $\rightarrow$ **将 x[2] 设为 6，接着在设置 p.x 时引发 panic**
   - **原理：** `p` 为 nil，第二阶段设置 `p.x` 时解引用空指针触发 panic，但 `x[2]` 已被写入为 6。
7. `for i, x[i] = range x` 循环例子：
   - **原理：** 第一阶段迭代 `range x` 得到首个元素索引 0 和值 `x[0]`（即 3）。此时左侧 `x[i]` 的 `i` 仍为 2（锁定 `x[2]`）。第二阶段赋值 `i = 0`, `x[2] = 3`。因此循环一次后 `i` 为 0，切片变为 `[3, 5, 3]`。

### 段落 7：可赋值性与特例（Assignability special cases）

> **【英文原文】**
>
> In assignments, each value must be assignable to the type of the operand to which it is assigned, with the following special cases:
>
> - Any typed value may be assigned to the blank identifier.
> - If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first implicitly converted to its default type.
> - If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first implicitly converted to type bool.

**【逐字精准翻译】**

在赋值中，每个值都必须可以赋值给其所赋给的操作数的类型，但有以下特例：

- 任何带类型的值都可以赋给空标识符。
- 如果将无类型常量赋给接口类型变量或空标识符，则该常量首先被隐式转换为其默认类型。
- 如果将无类型布尔值赋给接口类型变量或空标识符，则其首先被隐式转换为 bool 类型。

- **词汇与句式剖析：**
  - `untyped constant`：无类型常量（如孤立的数字字面量 `42` 或 `3.14`）。
  - `implicitly converted to its default type`：隐式转换为其默认类型（例如整数默认为 `int`，浮点数默认为 `float64`）。

### 段落 8：数据替换与引用复制机制

> **【英文原文】**
>
> When a value is assigned to a variable, only the data that is stored in the variable is replaced. If the value contains a reference, the assignment copies the reference but does not make a copy of the referenced data (such as the underlying array of a slice).

**【逐字精准翻译】**

当将一个值赋给变量时，仅替换存储在该变量中的数据。如果该值包含引用，则赋值会复制该引用，但不会对引用的数据（例如切片的底层数组）制作副本。

- **词汇与句式剖析：**
  - `referenced data`：被引用的数据。
  - `underlying array`：底层数组。
  - `slice descriptor`：切片描述符（包含指针、长度 len、容量 cap 的结构体）。

### 段落 9：引用类型赋值代码例证

> **【英文原文与代码】**
>
> ```go
>var s1 = []int{1, 2, 3}
> var s2 = s1                    // s2 stores the slice descriptor of s1
> s1 = s1[:1]                    // s1's length is 1 but it still shares its underlying array with s2
> s2[0] = 42                     // setting s2[0] changes s1[0] as well
> fmt.Println(s1, s2)            // prints [42] [42 2 3]
> 
> var m1 = make(map[string]int)
> var m2 = m1                    // m2 stores the map descriptor of m1
> m1["foo"] = 42                 // setting m1["foo"] changes m2["foo"] as well
> fmt.Println(m2["foo"])         // prints 42
> ```

**【逐字精准翻译】**

```go
var s1 = []int{1, 2, 3}
var s2 = s1                    // s2 存储了 s1 的切片描述符副本
s1 = s1[:1]                    // s1 的长度变为 1，但它依然与 s2 共享底层数组
s2[0] = 42                     // 修改 s2[0] 也会同时改变 s1[0]
fmt.Println(s1, s2)            // 打印 [42] [42 2 3]

var m1 = make(map[string]int)
var m2 = m1                    // m2 存储了 m1 的 map 描述符副本
m1["foo"] = 42                 // 设置 m1["foo"] 也会同时改变 m2["foo"]
fmt.Println(m2["foo"])         // 打印 42
```

您引用的文本正好停在 **If statements (If 语句)** 的开头。

准备好后，请随时发送 `If statements` 及接下来的原文！
