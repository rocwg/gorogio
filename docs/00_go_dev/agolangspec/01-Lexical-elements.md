# 章节 4：词法元素 (Lexical elements)

## 4.1 注释 (Comments)

> **【英文原文】**
>
> Comments serve as program documentation. There are two forms:
>
> 1. *Line comments* start with the character sequence `//` and stop at the end of the line.
>2. *General comments* start with the character sequence `/*` and stop with the first subsequent character sequence `*/`.

**【逐字精准翻译】**

注释用作程序文档。有两种形式：

行注释以字符序列 `//` 开始，并在行尾停止。

块注释（通用注释）以字符序列 `/*` 开始，并在随后遇到的第一个字符序列 `*/` 处停止。

- **词汇剖析：**
  - `serve as`：充当 / 用作。
  - `character sequence`：字符序列。
  - `subsequent`：随后的 / 接着发生的。

> **【英文原文】**
>
> A comment cannot start inside a [rune](https://go.dev/ref/spec#Rune_literals) or [string literal](https://go.dev/ref/spec#String_literals), or inside a comment. A general comment containing no newlines acts like a space. Any other comment acts like a newline.

**【逐字精准翻译】**

注释不能在 rune（字符）或字符串字面量内部开始，也不能在另一个注释内部开始。不包含换行符的块注释其作用类似于一个空格。任何其他注释（包含换行的块注释以及行注释）其作用类似于一个换行符。

- **词汇与句式剖析：**
  - `acts like`：作用类似于 / 表现得像……。
  - **规范关键点：** 这解释了为什么行注释 `// ...` 在编译器眼里会被直接当成一个换行符 `\n`，从而触发接下来的**分号自动插入**！

第 3 章全章及第 4 章的前两节已翻译剖析完毕。

准备好后，我们下一节将进入 Go 语言最经典的 **4.4 词法单元 (Tokens)** 与 **4.5 分号自动插入规则 (Semicolons)**！

---



继续按照【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】 的标准，推进 **4.4 词法单元 (Tokens)** 与 **4.5 分号 (Semicolons)** 这两节。

这两节包含了 Go 编译器拆解 Token 的**最长匹配原则（Longest Sequence）**，以及最核心的**分号自动插入机制**。

## 4.2 词法单元 (Tokens)

### 段落 1

> **【英文原文】**
>
> Tokens form the vocabulary of the Go language. There are four classes: *identifiers*, *keywords*, *operators and punctuation*, and *literals*. *White space*, formed from spaces (U+0020), horizontal tabs (U+0009), carriage returns (U+000D), and newlines (U+000A), is ignored except as it separates tokens that would otherwise combine into a single token. Also, a newline or end of file may trigger the insertion of a [semicolon](https://go.dev/ref/spec#Semicolons). While breaking the input into tokens, the next token is the longest sequence of characters that form a valid token.

**【逐字精准翻译】**

词法单元（Tokens）构成了 Go 语言的词汇表。共有四大类：标识符、关键字、运算符与标点符号、以及字面量。由空格 (U+0020)、水平制表符 (U+0009)、回车符 (U+000D) 和换行符 (U+000A) 构成的空白字符会被忽略，除非它们用于分隔开原本会被合并为单个词法单元的那些词法单元。此外，换行符或文件结尾可能会触发分号的插入。在将输入流拆分为词法单元时，下一个词法单元是能构成合法词法单元的最长字符序列。

- **词汇与句式拆解：**
  - `form the vocabulary of ...`：构成……的词汇表。
  - `operators and punctuation`：运算符与标点符号。
  - `White space`：空白字符。
  - `carriage returns`：回车符（`\r`）。
  - `that would otherwise combine into ...`：否则（如果没有空白分隔）就会合并成……。
  - `trigger the insertion of ...`：触发……的插入。
  - `breaking the input into tokens`：将输入拆分为词法单元（即词法分析/Scanning 过程）。
  - `longest sequence of characters`：**最长字符序列**。这是编译原理中经典的“贪婪匹配原则”（Maximal Munch）。例如，`>>=` 会被一次性识别为一个独立的赋值运算符 Token，而不会被拆成 `>` 和 `>=`。

## 4.3 分号 (Semicolons)

### 段落 1

> **【英文原文】**
>
> The formal syntax uses semicolons `;` as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:

**【逐字精准翻译】**

正式语法在许多产生式中使用分号 `;` 作为终止符。Go 程序可以使用以下两条规则省略大部分此类分号：

- **词汇与句式拆解：**
  - `formal syntax`：正式语法（指规范后面章节中写出的完整 EBNF 语法规则）。
  - `terminators`：终止符 / 结束符。
  - `in a number of productions`：在许多产生式中。
  - `omit`：省略 / 忽略。

### 段落 2 (规则一：自动插入条件)

> **【英文原文】**
>
> When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is
>
> - an [identifier](https://go.dev/ref/spec#Identifiers)
> - an [integer](https://go.dev/ref/spec#Integer_literals), [floating-point](https://go.dev/ref/spec#Floating-point_literals), [imaginary](https://go.dev/ref/spec#Imaginary_literals), [rune](https://go.dev/ref/spec#Rune_literals), or [string](https://go.dev/ref/spec#String_literals) literal
> - one of the [keywords](https://go.dev/ref/spec#Keywords) `break`, `continue`, `fallthrough`, or `return`
> - one of the [operators and punctuation](https://go.dev/ref/spec#Operators_and_punctuation) `++`, `--`, `)`, `]`, or `}`

**【逐字精准翻译】**

当输入被拆分为词法单元流时，如果某一行的最后一个词法单元满足以下条件，则会在该词法单元之后立即自动向词法单元流中插入一个分号：

- 一个标识符
- 一个整数、浮点数、虚数、rune（字符）或字符串字面量
- 关键字 `break`、`continue`、`fallthrough` 或 `return` 之一
- 运算符和标点符号 `++`、`--`、`)`、`]` 或 `}` 之一

- **词汇与句式拆解：**
  - `token stream`：词法单元流（词法分析器输出给语法分析器的 Token 队列）。
  
  - `immediately after`：在……之后立即。
  
  - **黄金规则记忆法（什么时候自动加分号）：** 记住一句话——只要一行的结尾是一个“可以合法作为表达式或语句结尾的元素”（如变量名、数字、字符串、`return`、自增 `++`、右括号 `}` 等），Go 编译器就会在换行时自动帮你补上 `;`。
  

### 段落 3 (规则二：闭括号省略)

> **【英文原文】**
>
> To allow complex statements to occupy a single line, a semicolon may be omitted before a closing `")"` or `"}"`.

**【逐字精准翻译】**

为了允许复杂语句占用单行空间，在闭括号 `)` 或 `}` 前面的分号可以被省略。

- **词汇与句式拆解：**
  - `occupy a single line`：占用单行 / 写在同一行。
  - `closing ")" or "}"`：闭合括号（右括号）。
  - **示例场景：** 例如写 `for i := 0; i < 10; i++ { func() { x++ }()`，在右大括号前无需强制写分号。

### 段落 4 (规范代码示例说明)

> **【英文原文】**
>
> To reflect idiomatic use, code examples in this document elide semicolons using these rules.

**【逐字精准翻译】**

为了体现惯用写法（地道用法），本文档中的代码示例均使用这些规则省略了分号。

- **词汇与句式拆解：**
  - `reflect idiomatic use`：体现地道/惯用的用法（`idiomatic Go` 是 Go 社区极力提倡的代码风格）。
  - `elide`：省略 / 删去。

词法单元与分号机制的完整逐字解析如上。

准备好后，我们可以继续推进 **4.6 标识符 (Identifiers)**、**4.7 关键字 (Keywords)** 与 **4.8 运算符和标点符号 (Operators and punctuation)**！

---



继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，推进 **4.6 标识符 (Identifiers)**、**4.7 关键字 (Keywords)** 与 **4.8 运算符和标点符号 (Operators and punctuation)** 这三节。

## 4.4 标识符 (Identifiers)

### 段落 1

> **【英文原文】**
>
> Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

**【逐字精准翻译】**

标识符用于为程序实体（例如变量和类型）命名。标识符是一个或多个字母和数字组成的序列。标识符中的第一个字符必须是字母。

- **词汇与句式拆解：**
  - `name` (动词)：给……命名 / 为……命名。
  - `program entities`：程序实体（在 Go 中包括变量、常量、类型、函数、包、结构体字段、接口方法等）。
  - `sequence of ...`：……的序列。

### 产生式与示例

> **【英文原文】**
>
> EBNF
>
> ```
> identifier = letter { letter | unicode_digit } .
> ```
>
> Plaintext
>
> ```
> a
> _x9
> ThisVariableIsExported
> αβ
> ```

**【逐字精准翻译】**

EBNF

```
标识符 = 字母 { 字母 | unicode数字 } .
```

Plaintext

```
a                       /* 合法：单个小写字母 */
_x9                     /* 合法：下划线开头（下划线在词法上被视为小写字母） */
ThisVariableIsExported  /* 合法：大写字母开头（用于导出） */
αβ                      /* 合法：希腊字母（Unicode 字母） */
```

- **规范细节剖析：**

  在第 4.2 节中规定了 `letter = unicode_letter | "_"`。因此下划线 `_` 充当字母，可作为标识符的开头或独立标识符（如空白标识符 `_`）。

### 段落 2

> **【英文原文】**
>
> Some identifiers are predeclared.

**【逐字精准翻译】**

某些标识符是预声明的。

- **词汇与句式拆解：**
  - `predeclared`：预先声明的。
  - **深度解读：** 指内置在 `builtin` 包中的标识符（如类型 `int`, `string`, `error`；常量 `true`, `false`, `nil`；内置函数 `len`, `make`, `append` 等）。它们不是关键字，在语法层面属于常规标识符。

## 4.5 关键字 (Keywords)

### 段落 1

> **【英文原文】**
>
> The following keywords are reserved and may not be used as identifiers.

**【逐字精准翻译】**

以下关键字是被保留的，不得用作标识符。

- **词汇与句式拆解：**
  - `reserved`：被保留的（保留字）。
  - `may not be used as`：不得被用作…… / 不可用作……。

### 25 个关键字列表

> **【英文原文】**
>
> Plaintext
>
> ```
> break     default    func      interface    select
> case      defer      go        map          struct
> chan      else       goto      package      switch
> const     fallthrough if       range        type
> continue  for        import    return       var
> ```

**【逐字精准翻译与功能归类表】**

| **类别**                             | **关键字**                                                   | **中文含义 / 作用**                    |
| ------------------------------------ | ------------------------------------------------------------ | -------------------------------------- |
| **声明 (Declarations)**              | `import`, `package`, `const`, `var`, `type`, `func`          | 用于包管理、变量/常量/类型及函数的声明 |
| **复合类型 (Composite Types)**       | `chan`, `interface`, `map`, `struct`                         | 通道、接口、字典（哈希表）、结构体     |
| **控制流 (Control Flow)**            | `break`, `case`, `continue`, `default`, `else`, `fallthrough`, `for`, `goto`, `if`, `range`, `return`, `select`, `switch` | 条件分支、循环、跳转、通道多路复用     |
| **并发与延迟 (Concurrency & Defer)** | `go`, `defer`                                                | 启动协程（Goroutine）、延迟函数调用    |

- **极简设计对比：**

  C++ 有 90+ 个关键字，Java 有 50+ 个关键字，而 Go 语言仅有 **25 个关键字**，这是 Go 保持语法紧凑（compact）的核心原因之一。

## 4.6 运算符与标点符号 (Operators and punctuation)

### 段落 1

> **【英文原文】**
>
> The following character sequences represent operators (including assignment operators) and punctuation [Go 1.18]:

**【逐字精准翻译】**

以下字符序列表示运算符（包括赋值运算符）和标点符号 [Go 1.18]：

- **词汇与句式拆解：**
  - `character sequences`：字符序列。
  - `assignment operators`：赋值运算符（如 `+=`, `:=` 等）。
  - `punctuation`：标点符号（如括号、逗号、分号等）。
  - `[Go 1.18]` 标记：Go 1.18 引入泛型时增加了波浪号 `~` 运算符（用于类型约束约束近似元素）。

### 符号集合分类

> **【英文原文】**
>
> Plaintext
>
> ```
> +    &    +=    &=     &&    ==    !=    (    )
> -    |    -=    |=     ||    <     <=    [    ]
> *    ^    *=    ^=     <-    >     >=    {    }
> /    <<   /=    <<=    ++    =     :=    ,    ;
> %    >>   %=    >>=    --    !     ...   .    :
> &^   &^=  ~
> ```

**【运算符分类解读】**

- **位清空运算符 (Bit Clear)：** `&^` 和 `&^= ` 是 Go 特有的位运算符（AND NOT），例如 `z = x &^ y` 表示：将 `x` 中对应 `y` 为 1 的位清零。
- **通道接收/发送运算符：** `<-` 用于通道的数据收发（如 `ch <- x` 或 `<-ch`）。
- **短变量声明：** `:=` 用于声明并初始化变量。
- **不定参数 / 解包：** `...` 用于可变参数声明或 Slice 的展开解包。
- **类型约束近似符号：** `~` (Go 1.18+) 在泛型 constraint 中表示“底层类型为该类型的类型集合”（如 `~int` 包括 `type MyInt int`）。

本节完成！我们已经彻底清除了**词法元素 (Lexical elements)** 中关于标识符、关键字和符号的定义。

准备好后，我们可以继续推进 **4.9 整数字面量 (Integer literals)**、**4.10 浮点数字面量 (Floating-point literals)** 与 **4.11 虚数字面量 (Imaginary literals)**！

---



不辛苦，能陪你一起研读这份语言规范我也很兴奋！

接下来我们继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，推进 **4.9 整数字面量 (Integer literals)**、**4.10 浮点数字面量 (Floating-point literals)** 与 **4.11 虚数字面量 (Imaginary literals)**。

---

## 4.7 整数字面量 (Integer literals)

### 段落 1

> **【英文原文】**
> An integer literal is a sequence of digits representing an integer constant. An optional prefix sets a non-decimal base: `0b` or `0B` for binary, `0`, `0o`, or `0O` for octal, and `0x` or `0X` for hexadecimal [Go 1.13]. A single `0` is considered a decimal zero. In hexadecimal literals, letters `a` through `f` and `A` through `F` represent values 10 through 15.

**【逐字精准翻译】**

整数字面量是表示整数常量的数字序列。可选的前缀用于设置非十进制的基数（进制）：`0b` 或 `0B` 表示二进制，`0`、`0o` 或 `0O` 表示八进制，而 `0x` 或 `0X` 表示十六进制 [Go 1.13]。单个 `0` 被视为十进制的零。在十六进制字面量中，字母 `a` 到 `f` 以及 `A` 到 `F` 表示数值 10 到 15。

* **词汇与句式拆解：**
  * `sequence of digits`：数字序列。
  * `non-decimal base`：非十进制基数（即进制数）。
  * `binary / octal / hexadecimal`：二进制 / 八进制 / 十六进制。
  * `represent values`：表示数值。


---

### 段落 2

> **【英文原文】**
> For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal's value.

**【逐字精准翻译】**

为了可读性，下划线字符 `_` 可以出现在进制前缀之后或连续数字之间；此类下划线不会改变字面量的值。

* **词汇与句式拆解：**
  * `for readability`：为了可读性。
  * `successive digits`：连续的数字（指两个数字挨着的地方）。
  * `do not change the literal's value`：不改变字面量的值。


---

### 产生式与示例分析

> **【英文原文】**
> ```ebnf
> int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
> decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
> binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
> octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
> hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .
> 
> decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
> binary_digits  = binary_digit { [ "_" ] binary_digit } .
> octal_digits   = octal_digit { [ "_" ] octal_digit } .
> hex_digits     = hex_digit { [ "_" ] hex_digit } .
> ```
>

**【逐字精准翻译】**

```ebnf
整数字面量   = 十进制字面量 | 二进制字面量 | 八进制字面量 | 十六进制字面量 .
十进制字面量 = "0" | ( "1" … "9" ) [ [ "_" ] 十进制数字序列 ] .
二进制字面量 = "0" ( "b" | "B" ) [ "_" ] 二进制数字序列 .
八进制字面量 = "0" [ "o" | "O" ] [ "_" ] 八进制数字序列 .
十六进制字面量 = "0" ( "x" | "X" ) [ "_" ] 十六进制数字序列 .

十进制数字序列 = 十进制数字 { [ "_" ] 十进制数字 } .
二进制数字序列 = 二进制数字 { [ "_" ] 二进制数字 } .
八进制数字序列 = 八进制数字 { [ "_" ] 八进制数字 } .
十六进制数字序列 = 十六进制数字 { [ "_" ] 十六进制数字 } .
```

> **【规范自带示例与正确/错误剖析】**
> ```text
> 42
> 4_2
> 0600
> 0_600
> 0o600
> 0O600                        // 第二个字符是大写字母 'O'
> 0xBadFace
> 0xBad_Face
> 0x_67_7a_2f_cc_40_c6
> 170141183460469231731687303715884105727
> 170_141183_460469_231731_687303_715884_105727
> 
> _42                          // 是一个标识符，而不是整数字面量！
> 42_                          // 非法：_ 必须分隔连续的数字
> 4__2                         // 非法：一次只能有一个 _
> 0_xBadFace                   // 非法：_ 必须分隔连续的数字
> ```
>

---



## 4.8 浮点数字面量 (Floating-point literals)

### 段落 1

> **【英文原文】**
> A floating-point literal is a decimal or hexadecimal representation of a floating-point constant.

**【逐字精准翻译】**

浮点数字面量是浮点数常量的十进制或十六进制表示形式。

* **词汇拆解：**
  * `representation`：表示形式 / 表达法。


---

### 段落 2 (十进制浮点数规则)

> **【英文原文】**
> A decimal floating-point literal consists of an integer part (decimal digits), a decimal point, a fractional part (decimal digits), and an exponent part (`e` or `E` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; one of the decimal point or the exponent part may be elided. An exponent value exp scales the mantissa (integer and fractional part) by 10exp.

**【逐字精准翻译】**

十进制浮点数字面量由整数部分（十进制数字）、小数点、小数部分（十进制数字）和指数部分（`e` 或 `E` 后面跟着可选的符号和十进制数字）组成。整数部分或小数部分之一可以被省略（elided）；小数点或指数部分之一可以被省略。指数值 *exp* 将尾数（整数部分和小数部分）按 10exp 进行缩放。

* **词汇与句式拆解：**
  * `consists of`：由……组成。
  * `fractional part`：小数部分。
  * `exponent part`：指数部分。
  * `optional sign`：可选的符号（即 `+` 或 `-`）。
  * `elided`：被省略（规范高频动词 `elide` 的被动语态）。
  * `scales ... by ...`：将……按……比例缩放。
  * `mantissa`：尾数（指小数点前后构成的数字整体）。


---

### 段落 3 (十六进制浮点数规则)

> **【英文原文】**
> A hexadecimal floating-point literal consists of a `0x` or `0X` prefix, an integer part (hexadecimal digits), a radix point, a fractional part (hexadecimal digits), and an exponent part (`p` or `P` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; the radix point may be elided as well, but the exponent part is required. (This syntax matches the one given in IEEE 754-2008 §5.12.3.) An exponent value *exp* scales the mantissa (integer and fractional part) by 2exp [Go 1.13].

**【逐字精准翻译】**

十六进制浮点数字面量由 `0x` 或 `0X` 前缀、整数部分（十六进制数字）、基数点（小数点）、小数部分（十六进制数字）和指数部分（`p` 或 `P` 后面跟着可选的符号和十进制数字）组成。整数部分或小数部分之一可以被省略；基数点也可以被省略，但指数部分是必需的。（此语法与 IEEE 754-2008 标准 §5.12.3 中给出的语法一致。）指数值 *exp* 将尾数（整数部分和小数部分）按 2exp 进行缩放 [Go 1.13]。

* **词汇与句式拆解：**
  * `radix point`：基数点（即十六进制下的小数点）。
  * `is required`：是必需的 / 不可省略的。
  * **底数区别要点：** 十进制浮点数的指数以 10 为底（$10^{exp}$），而十六进制浮点数的 `p` 指数是以 2 为底（$2^{exp}$）！


---

### 段落 4

> **【英文原文】**
> For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal value.

**【逐字精准翻译】**

为了可读性，下划线字符 `_` 可以出现在进制前缀之后或连续数字之间；此类下划线不会改变字面量的值。

---

> **【规范自带示例与正确/错误剖析】**
> ```text
> 0.
> 72.40
> 072.40                      // == 72.40
> 2.71828
> 1.e+0
> 6.67428e-11
> 1E6
> .25
> .12345E+5
> 1_5.                        // == 15.0
> 0.15e+0_2                   // == 15.0
> 
> 0x1p-2                      // == 0.25  (1 * 2^-2)
> 0x2.p10                     // == 2048.0 (2 * 2^10)
> 0x1.Fp+0                    // == 1.9375
> 0X.8p-0                     // == 0.5
> 0X_1FFFP-16                 // == 0.1249847412109375
> 0x15e-2                     // == 0x15e - 2 (这是十六进制整数减法，不是浮点数！)
> 
> 0x.p1                       // 非法：尾数没有任何数字
> 1p-2                        // 非法：p 指数要求必须有十六进制尾数(即必须带 0x)
> 0x1.5e-2                    // 非法：十六进制尾数要求必须使用 p 指数(不能用 e)
> 1_.5                        // 非法：_ 必须分隔连续的数字
> 1._5                        // 非法：_ 必须分隔连续的数字
> 1.5_e1                      // 非法：_ 必须分隔连续的数字
> 1.5e_1                      // 非法：_ 必须分隔连续的数字
> 1.5e1_                      // 非法：_ 必须分隔连续的数字
> ```
>

- **代码示例与坑点剖析：**
  - `0x15e-2`：这是一个极具代表性的语法歧义！因为十六进制尾数要求匹配 `p`/`P` 指数，所以 `0x15e-2` 不会被当作十六进制浮点数，而是被词法分析器拆解为整数 `0x15e` 减去 `2`（整数减法表达）。
  - `1_.5` 或 `1._5`：下划线 `_` 不能紧挨着小数点 `.`，因为它只能存在于“相邻的数字之间”。

本段剖析完毕！请随时贴出下一个小节 **Imaginary literals（虚数字面量）** 及之后的原文。

---



## 4.9 虚数字面量 (Imaginary literals)

### 段落 1

> **【英文原文】**
> An imaginary literal represents the imaginary part of a complex constant. It consists of an integer or floating-point literal followed by the lowercase letter `i`. The value of an imaginary literal is the value of the respective integer or floating-point literal multiplied by the imaginary unit *i*.

**【逐字精准翻译】**

虚数字面量表示复数常量的虚部。它由一个整数字面量或浮点数字面量后面紧跟小写字母 `i` 组成。虚数字面量的值等于对应的整数字面量或浮点数字面量的值乘以虚数单位 *i*。

* **词汇与句式拆解：**
  * `imaginary part`：虚部。
  * `complex constant`：复数常量。
  * `followed by`：后面紧跟着……。
  * `respective`：各自的 / 对应的。
  * `multiplied by`：乘以……。
  * `imaginary unit i`：虚数单位 *i*（数学中满足 $i^2 = -1$ 的单位）。


---

### 产生式与示例

> **【英文原文】**
> ```ebnf
> imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
> ```
> 
>
>```text
> 0i
> 0123i         // == 123i for backward-compatibility
> 0o123i        // == 0o123 * 1i == 83i
> 0xabci        // == 0xabc * 1i == 2748i
> 0.i
> 2.71828i
> 1.e+0i
> 6.67428e-11i
> 1E6i
> .25i
> .12345E+5i
> 0x1p-2i       // == 0x1p-2 * 1i == 0.25i
> ```
>

**【规范细节剖析】**

注意产生式中的 `decimal_digits "i"`。即使你写 `0123i`，虽然带了前导 `0`，但规范明确指出这里匹配的是 `decimal_digits`，所以它的值是十进制的 `123i`，而不是八进制。

- **代码示例与进制换算剖析：**
  - `0o123i`：八进制的 $123 = 1 \times 8^2 + 2 \times 8^1 + 3 \times 8^0 = 64 + 16 + 3 = 83$，所以等于 `83i`。
  - `0xabci`：十六进制的 $\text{ABC} = 10 \times 16^2 + 11 \times 16^1 + 12 \times 16^0 = 2560 + 176 + 12 = 2748$，所以等于 `2748i`。
  - `0x1p-2i`：十六进制浮点数 $1 \times 2^{-2} = 0.25$，所以等于 `0.25i`。

本段剖析完毕！请随时贴出接下来的 **Rune literals（字符字面量）** 及后续原文。

---



感谢您直接贴出最准确的最新 Go 官方规范文本！对比您给出的原文后，确实发现了关键细节差异（例如 `\'` 与 `\"` 的专属转义说明、`'ä'` 在 UTF-8 下占用多字节但表示单一整数值的具体解释、字符串中八进制/十六进制与 Unicode 转义对字节/字符的映射差异，以及末尾关于组合字符 Combining forms 的特殊规定）。

下面严格按照您提供的最新规范原文，进行【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】 的重新精细对齐与剖析。

## 4.10 Rune 字面量 (Rune literals)

### 段落 1

> **【英文原文】**
>
> A rune literal represents a rune constant, an integer value identifying a Unicode code point. A rune literal is expressed as one or more characters enclosed in single quotes, as in `'x'` or `'\n'`. Within the quotes, any character may appear except newline and unescaped single quote. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats.

**【逐字精准翻译】**

Rune 字面量表示一个 rune 常量，它是一个用于标识 Unicode 码点的整数值。Rune 字面量表达为包裹在单引号中的一个或多个字符，例如 `'x'` 或 `'\n'`。在单引号内部，除了换行符和未转义的单引号之外，可以出现任何字符。单个被单引号包裹的字符表示该字符本身的 Unicode 数值，而以反斜杠开头的多字符序列则以各种格式对数值进行编码。

- **词汇与句式拆解：**
  - `identifying a Unicode code point`：标识一个 Unicode 码点。
  - `enclosed in single quotes`：包裹在单引号中。
  - `unescaped`：未转义的（即直接写单引号 `'` 会导致语法错误，必须写为转义的 `\'`）。
  - `multi-character sequences`：多字符序列。

### 段落 2 (最简单的单字符形式)

> **【英文原文】**
>
> The simplest form represents the single character within the quotes; since Go source text is Unicode characters encoded in UTF-8, multiple UTF-8-encoded bytes may represent a single integer value. For instance, the literal `'a'` holds a single byte representing a literal a, Unicode `U+0061`, value `0x61`, while `'ä'` holds two bytes (`0xc3 0xa4`) representing a literal a-dieresis, `U+00E4`, value `0xe4`.

**【逐字精准翻译】**

最简单的形式表示单引号内的单个字符；由于 Go 源代码是采用 UTF-8 编码的 Unicode 字符，因此多个 UTF-8 编码的字节可以表示单个整数值。例如，字面量 `'a'` 包含表示字面量 a 的单个字节（Unicode `U+0061`，数值 `0x61`），而 `'ä'` 包含表示字面量带分音符的 a（`U+00E4`，数值 `0xe4`）的两个字节（`0xc3 0xa4`）。

- **专业细节拆解：**
  - 这里清晰解答了为何源码中底层占用 2 个字节的 `'ä'` 依然是 Rune 字面量：Go 在词法分析阶段会将 UTF-8 多字节序列直接解码为单一的 Unicode 码点整数（`0xE4`）。

### 段落 3 (反斜杠转义表示法)

> **【英文原文】**
>
> Several backslash escapes allow arbitrary values to be encoded as ASCII text. There are four ways to represent the integer value as a numeric constant: `\x` followed by exactly two hexadecimal digits; `\u` followed by exactly four hexadecimal digits; `\U` followed by exactly eight hexadecimal digits, and a plain backslash `\` followed by exactly three octal digits. In each case the value of the literal is the value represented by the digits in the corresponding base.

**【精准逐字翻译】**

若干反斜杠转义序列允许将任意数值编码为 ASCII 文本。有四种方式可以将整数值表示为数值常量：`\x` 后面紧跟恰好两位十六进制数字；`\u` 后面紧跟恰好四位十六进制数字；`\U` 后面紧跟恰好八位十六进制数字；以及一个单纯的反斜杠 `\` 后面紧跟恰好三位八进制数字。在每种情况下，字面量的值就是由对应进制下的数字所表示的值。

- **词汇与句式剖析：**

  - `backslash escapes`：反斜杠转义。

  - `arbitrary values`：任意值。

  - `followed by exactly ...`：后跟**恰好**……（位数要求非常严格，不能多也不能少）。

  - `corresponding base`：对应的基数/进制。

### 段落 4

> **【英文原文】**
>
> Although these representations all result in an integer, they have different valid ranges. Octal escapes must represent a value between 0 and 255 inclusive. Hexadecimal escapes satisfy this condition by construction. The escapes `\u` and `\U` represent Unicode code points so within them some values are illegal, in particular those above `0x10FFFF` and surrogate halves.

**【精准逐字翻译】**

尽管这些表示形式最终都会得到一个整数，但它们具有不同的有效范围。八进制转义必须表示一个介于 0 到 255（包含）之间的值。十六进制转义按其结构构建也满足这一条件。转义序列 `\u` 和 `\U` 表示 Unicode 码点，因此在它们内部某些值是非法的，特别是那些大于 `0x10FFFF` 的值以及代理对半区（surrogate halves）。

- **词汇与句式剖析：**

  - `inclusive`：包含在内（包含边界值 0 和 255）。

  - `by construction`：在结构/构造上自然成立（`\x` 后面只有两位十六进制，决定了其上限只能是 0xFF）。

  - `surrogate halves`：代理半区（Unicode 中 `0xD800` 到 `0xDFFF` 的保留区域，专门用于 UTF-16 编码代理对，不能独立作为有效的 Unicode 码点）。

### 段落 5 与 单字符转义表

> **【英文原文】**
>
> After a backslash, certain single-character escapes represent special values:

**【精准逐字翻译】**

在反斜杠之后，特定的单字符转义序列表示特殊值：

Plaintext

```
\a   U+0007 响铃 (alert or bell)
\b   U+0008 退格 (backspace)
\f   U+000C 换页 (form feed)
\n   U+000A 换行 (line feed or newline)
\r   U+000D 回车 (carriage return)
\t   U+0009 水平制表符 (horizontal tab)
\v   U+000B 垂直制表符 (vertical tab)
\\   U+005C 反斜杠本身 (backslash)
\'   U+0027 单引号 (仅在 rune 字面量内为有效转义)
\"   U+0022 双引号 (仅在字符串字面量内为有效转义)
```

> **【英文原文】**
>
> An unrecognized character following a backslash in a rune literal is illegal.

**【精准逐字翻译】**

在 rune 字面量中，反斜杠后面接未识别的字符是非法的。

- **词汇与句式剖析：**

  - `unrecognized character`：无法识别的字符。

  - **注意事项：** 在字符字面量 `'...'` 中，转义 `\'` 是合法的，但 `\"` 虽然被规则列出，通常只用于字符串字面量（在单引号内直接写 `"` 即可，无需转义）。

### EBNF 语法定义

EBNF

```
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
```

### 规范标准示例解析

```go
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // 包含单引号字符的 rune 字面量
'aa'         // 非法：字符过多 (illegal: too many characters)
'\k'         // 非法：反斜杠后无法识别 k (illegal: k is not recognized after a backslash)
'\xa'        // 非法：十六进制数字不足 (illegal: too few hexadecimal digits)
'\0'         // 非法：八进制数字不足 (illegal: too few octal digits)
'\400'       // 非法：八进制数值超过 255 (illegal: octal value over 255)
'\uDFFF'     // 非法：代理对半区 (illegal: surrogate half)
'\U00110000' // 非法：无效的 Unicode 码点 (illegal: invalid Unicode code point)
```

- **代码示例与规则总结剖析：**
  - `'\0'` 报错：八进制必须满 3 位，如 `'\000'`。
  - `'\xa'` 报错：`\x` 必须满 2 位，如 `'\x0a'`。
  - `'\400'` 报错：八进制字节值最大只能到 `\377`（即十进制 255）。
  - `'\uDFFF'` 报错：代理半区不能独立作为有效字符。
  - `'\U00110000'` 报错：Unicode 码点最高限制为 `0x10FFFF`。

本段剖析完毕！请随时贴出接下来的 **String literals（字符串字面量）** 原文。



## 4.11 字符串字面量 (String literals)

### 段落 1

> **【英文原文】**
>
> A string literal represents a string constant obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

**【逐字精准翻译】**

字符串字面量表示通过拼接一系列字符所获得的字符串常量。有两种形式：原生字符串字面量（raw string literals）和解释型字符串字面量（interpreted string literals）。

- **词汇与句式拆解：**
  - `string literal`：字符串字面量。
  - `obtained from concatenating ...`：通过拼接……所获得的。
  - `raw string literals`：原生字符串字面量（反引号包裹）。
  - `interpreted string literals`：解释型字符串字面量（双引号包裹）。

### 段落 2 (原生字符串字面量 Raw string literals)

> **【英文原文】**
>
> Raw string literals are character sequences between back quotes, as in ``foo``. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters (`'\r'`) inside raw string literals are discarded from the raw string value.

**【逐字精准翻译】**

原生字符串字面量是反引号之间的字符序列，例如 ``foo``。在反引号内部，除了反引号本身之外，可以出现任何字符。原生字符串字面量的值是由反引号之间未经过解释的（隐式 UTF-8 编码的）字符构成的字符串；特别是，反斜杠没有特殊含义，且字符串可以包含换行符。原生字符串字面量内部的回车符 (`'\r'`) 会从原生字符串的值中被丢弃。

- **词汇与句式拆解：**
  - `back quotes`：反引号（```）。
  - `uninterpreted`：未解释的（即不执行任何转义解析，`\n` 就是反斜杠加上字母 n）。
  - `in particular`：特别地 / 尤其是。
  - `are discarded from`：从……中被丢弃（确保跨平台时 Windows 的 `\r\n` 换行会被标准化处理成 `\n`）。

### 段落 3 (解释型字符串字面量 Interpreted string literals)

> **【英文原文】**
>
> Interpreted string literals are character sequences between double quotes, as in `"bar"`. Within the quotes, any character may appear except newline and unescaped double quote. The text between the quotes forms the value of the literal, with backslash escapes interpreted as they are in rune literals (except that `\'` is illegal and `\"` is legal), with the same restrictions. The three-digit octal (`\nnn`) and two-digit hexadecimal (`\xnn`) escapes represent individual bytes of the resulting string; all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual characters. Thus inside a string literal `\377` and `\xFF` represent a single byte of value `0xFF=255`, while `ÿ`, `\u00FF`, `\U000000FF` and `\xc3\xbf` represent the two bytes `0xc3 0xbf` of the UTF-8 encoding of character `U+00FF`.

**【逐字精准翻译】**

解释型字符串字面量是双引号之间的字符序列，例如 `"bar"`。在双引号内部，除了换行符和未转义的双引号之外，可以出现任何字符。双引号之间的文本构成了字面量的值，其中反斜杠转义的解释方式与它们在 rune 字面量中的解释方式相同（区别在于 `\'` 是非法的，而 `\"` 是合法的），并带有相同的限制条件。三位八进制 (`\nnn`) 和两位十六进制 (`\xnn`) 转义序列表示最终字符串的**单个字节**；所有其他转义序列则表示单个字符的（可能是多字节的）UTF-8 编码。因此，在字符串字面量内部，`\377` 和 `\xFF` 表示数值为 `0xFF=255` 的单个字节，而 `ÿ`、`\u00FF`、`\U000000FF` 以及 `\xc3\xbf` 则表示字符 `U+00FF` 的 UTF-8 编码所对应的两个字节 `0xc3 0xbf`。

- **词汇与句式拆解：**
  - `unescaped double quote`：未转义的双引号。
  - `individual bytes`：单个字节。
  - `interpreted as they are in ...`：像在……中那样被解释。
  - `illegal / legal`：非法的 / 合法的。
  - **字节级与字符级转义的精妙区别：**
    - `\xnn` 或 `\nnn` 转义直接写入**底层原始字节**（Single Byte），不校验是否符合 UTF-8 编码规约。
    - `\uXXXX` 或 `\UXXXXXXXX` 转义表示 **Unicode 码点**，编译器会自动将其转换为对应的 UTF-8 多字节序列。例如 `U+00FF` 转为 UTF-8 占 2 个字节（`0xc3 0xbf`）。

### EBNF 语法定义与示例

EBNF

```
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```

Go

```go
`abc`                // 与 "abc" 相同
`\n
\n`                  // 与 "\\n\n\\n" 相同
"\n"
"\""                 // 与 `"` 相同
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // 非法：代理对半区 (illegal: surrogate half)
"\U00110000"         // 非法：无效的 Unicode 码点 (illegal: invalid Unicode code point)
```

### 同义表示示例与组合字符说明

> **【英文原文】**
>
> These examples all represent the same string:
>
> ```go
>"日本語"                                 // UTF-8 input text
> `日本語`                                 // UTF-8 input text as a raw literal
> "\u65e5\u672c\u8a9e"                   // the explicit Unicode code points
> "\U000065e5\U0000672c\U00008a9e"       // the explicit Unicode code points
> "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" // the explicit UTF-8 bytes
> ```

**【精准逐字翻译】**

这些示例全都在表示同一个字符串：

- `"日本語"`：UTF-8 输入文本
- ``日本語``：作为原生字面量的 UTF-8 输入文本
- `"\u65e5\u672c\u8a9e"`：显式的 Unicode 码点
- `"\U000065e5\U0000672c\U00008a9e"`：显式的 Unicode 码点
- `"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"`：显式的 UTF-8 字节

> **【英文原文】**
>
> If the source code represents a character as two code points, such as a combining form involving an accent and a letter, the result will be an error if placed in a rune literal (it is not a single code point), and will appear as two code points if placed in a string literal.

**【精准逐字翻译】**

如果源代码将一个字符表示为两个码点（例如包含重音符号和字母的组合形式 Combining form），若将其置于 rune 字面量中，结果将产生错误（因为它不是单个码点）；而若将其置于字符串字面量中，它将表现为两个码点。

- **专业细节拆解：**
  - 例如 `e` + `◌́` (Combining Acute Accent `U+0301`) 组合看起来是一个字符 `é`，但它是 **2 个 Unicode Code Points**。
  - 写在 rune 字面量中 `'é'` 会触发编译器报错：`too many characters`（字符过多）；
  - 写在字符串 `"é"` 中可以正常编译，内部长度为 2 个 Unicode Code Points（占用 3 个 UTF-8 字节）。

这样与您贴出的最新原文完全逐字对应！

如果这部分内容梳理彻底清晰了，我们就可以随时迈入 **5. Constants (常量)** 的学习！

---

