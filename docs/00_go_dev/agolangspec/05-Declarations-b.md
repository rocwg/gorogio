继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，为您逐字逐句精准剖析 **变量声明 (Variable declarations)** 这一章节。

这一节定义了 Go 语言中最基础的 `var` 变量声明语法、初始化规则、无类型常量（Untyped constant）的隐式类型转换，以及编译器的“未引用变量检查”限制。

## 章节：变量声明 (Variable declarations)

### 段落 1 (定义与作用)

> **【英文原文】**
>
> A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value.

**【逐字精准翻译】**

变量声明创建一个或多个变量，将相应的标识符绑定到这些变量上，并为每个变量指定一个类型和一个初始值。

- **词汇与句式拆解：**
  - `creates one or more variables`：创建一个或多个变量（支持单变量与批量变量声明）。
  - `binds corresponding identifiers to them`：将相应的标识符（变量名）绑定到它们（内存中的变量）之上。
  - `initial value`：初始值（即使代码中不写显式赋值，也会给予零值作为初始值）。

### EBNF 语法规约

> **【英文原文】**
>
> EBNF
>
> ```
> VarDecl = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
> VarSpec = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
> ```

**【逐字精准翻译】**

EBNF

```
变量声明 = "var" ( 变量规范 | "(" { 变量规范 ";" } ")" ) .
变量规范 = 标识符列表 ( 类型 [ "=" 表达式列表 ] | "=" 表达式列表 ) .
```

- **语法结构深挖：**
  - 语法清楚地展现了两种形式：
    1. **单行声明：** `var` 后面直接接单个 `VarSpec`（变量规范）。
    2. **括号分组声明：** `var (...)` 内部包含零个或多个由分号分隔的 `VarSpec`。
  - `VarSpec` 揭示了两种合法的核心形式：
    - **显式指定类型：** `IdentifierList Type [ = ExpressionList ]`（类型必填，初始表达式可选）。
    - **类型推导：** `IdentifierList = ExpressionList`（省略类型，直接用 `=` 接初始表达式）。

### 代码示例 1 (各种 var 声明形态)

> **【英文原文】**
>
> ```go
>var i int
> var U, V, W float64
> var k = 0
> var x, y float32 = -1, -2
> var (
> 	i       int
> 	u, v, s = 2.0, 3.0, "bar"
> )
> var re, im = complexSqrt(-1)
> var _, found = entries[name]  // map lookup; only interested in "found"
> ```

**【逐字精准翻译】**

```go
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map 查找；仅对 "found" 感兴趣
```

- **关键语法特性说明：**
  - `var U, V, W float64`：一次性声明同类型的多个变量。
  - `var re, im = complexSqrt(-1)`：利用函数的多返回值进行**多重赋值/解构**。
  - `var _, found = ...`：使用空白标识符（Blank identifier `_`）丢弃不需要的返回值。

### 段落 2 (初始化值赋值规则)

> **【英文原文】**
>
> If a list of expressions is given, the variables are initialized with the expressions following the rules for assignment statements. Otherwise, each variable is initialized to its zero value.

**【逐字精准翻译】**

如果给出了表达式列表，则变量会按照赋值语句的规则使用这些表达式进行初始化。否则，每个变量都被初始化为其零值（zero value）。

- **词汇与句式拆解：**
  - `initialized with ...`：使用……进行初始化。
  - `following the rules for assignment statements`：遵循赋值语句的规则（例如类型匹配、可赋值性检查）。
  - `zero value`：**零值**（Go 语言极其重要的安全特性：数值型为 `0`，布尔型为 `false`，字符串为 `""`，指针/切片/接口/Map/Channel 等引用类型为 `nil`）。

### 段落 3 (类型推导与无类型常量转换)

> **【英文原文】**
>
> If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment. If that value is an untyped constant, it is first implicitly converted to its default type; if it is an untyped boolean value, it is first implicitly converted to type `bool`. The predeclared identifier `nil` cannot be used to initialize a variable with no explicit type.

**【逐字精准翻译】**

如果显式指定了类型，则每个变量都会被赋予该类型。否则，每个变量都会被赋予赋值式中对应初始化值的类型。如果该初始值是一个**无类型常量（untyped constant）**，它会首先被隐式转换为它的**默认类型（default type）**；如果它是一个无类型布尔值，它会首先被隐式转换为 `bool` 类型。预声明的标识符 `nil` 不能用于初始化没有显式指定类型的变量。

- **词汇与核心概念拆解：**

  - `untyped constant`：无类型常量（例如字面量 `0`、`3.14` 或 `"hello"` 在未显式声明类型前均属于无类型常量）。

  - `implicitly converted to its default type`：隐式转换为其默认类型（例如无整型常量 `0` 的默认类型是 `int`，无浮点型常量 `2.0` 的默认类型是 `float64`）。

  - `predeclared identifier nil`：预声明标识符 `nil`。

  - **为什么 `var n = nil` 是不法的？**

    因为 `nil` 没有默认类型！编译器无法从 `nil` 中推导出一个确切的底层类型，因此必须显式指定类型，如 `var n *int = nil` 或 `var n any = nil`。

### 代码示例 2 (类型推导示例)

> **【英文原文】**
>
> ```go
>var d = math.Sin(0.5)  // d is float64
> var i = 42             // i is int
> var t, ok = x.(T)      // t is T, ok is bool
> var n = nil            // illegal
> ```

**【逐字精准翻译】**

```go
var d = math.Sin(0.5)  // d 的类型是 float64
var i = 42             // i 的类型是 int（字面量 42 被转换为其默认类型 int）
var t, ok = x.(T)      // t 的类型是 T，ok 的类型是 bool（类型断言表达式的推导）
var n = nil            // 非法（nil 没有默认类型，无法推导）
```

### 段落 4 (编译器实现限制：未引用变量)

> **【英文原文】**
>
> Implementation restriction: A compiler may make it illegal to declare a variable inside a function body if the variable is never used.

**【逐字精准翻译】**

实现限制：如果在一个**函数体内部**声明了一个变量且该变量从未被使用过，编译器可以宣布这是非法的（禁止编译）。

- **词汇与句式拆解：**
  - `inside a function body`：在函数体内部（注意：这个限制**仅适用于局部变量**；包级别的全局变量即使未被使用，编译也是合法的）。
  - `never used`：从未被使用过。
  - **工程设计哲学：** 这是 Go 语言追求精简和防止废弃代码残留的标志性设计（Unused variable error）。

本章节逐字翻译与剖析完毕！您列出的最后一行：

> **Short variable declarations¶**（短变量声明）

是下一个小节的标题。准备好后，我们可以随时继续推进 **短变量声明 (`:=`)** 的原文！

继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，为您逐字逐句精准剖析 **短变量声明 (Short variable declarations)** 这一小节。

这一节定义了 Go 中最常用的短变量声明语法（`:=`）、它与普通 `var` 声明的等价关系、**变量重声明（Redeclaration）** 的苛刻条件，以及适用作用域。

## 章节：短变量声明 (Short variable declarations)

### 段落 1 (语法定义与等价关系)

> **【英文原文】**
>
> A short variable declaration uses the syntax:
>
> ```
> ShortVarDecl = IdentifierList ":=" ExpressionList .
> ```
>
> It is shorthand for a regular variable declaration with initializer expressions but no types:
>
> ```
> "var" IdentifierList "=" ExpressionList .
> ```

**【逐字精准翻译】**

短变量声明使用如下语法：

```
短变量声明 = 标识符列表 ":=" 表达式列表 .
```

它是带有初始化表达式但没有类型的常规变量声明的简写形式：

```
"var" 标识符列表 "=" 表达式列表 .
```

- **词汇与句式拆解：**
  - `shorthand for ...`：……的简写 / 缩写形式。
  - `regular variable declaration`：常规变量声明（即使用 `var` 关键字的声明）。
  - `initializer expressions`：初始化表达式（赋值号右边的表达式）。

### 代码示例 1 (常见用法)

> **【英文原文】**
>
> ```go
>i, j := 0, 10
> f := func() int { return 7 }
> ch := make(chan int)
> r, w, _ := os.Pipe()  // os.Pipe() returns a connected pair of Files and an error, if any
> _, y, _ := coord(p)   // coord() returns three values; only interested in y coordinate
> ```

**【逐字精准翻译】**

```go
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()  // os.Pipe() 返回一对连接好的 File 指针以及一个错误（如果有的话）
_, y, _ := coord(p)   // coord() 返回三个值；仅对 y 坐标感兴趣
```

- **语法细节剖析：**
  - `f := func() ...`：直接将匿名函数赋予变量 `f`，自动推导 `f` 为函数类型。
  - `ch := make(chan int)`：内建函数 `make` 返回的类型也是自动推导的。
  - `r, w, _ := ...`：结合多返回值与空白标识符 `_`，抛弃不需要的错误或变量。

### 段落 2 (重声明规则 —— 本节最核心、最易错的语法规则)

> **【英文原文】**
>
> Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original. The non-blank variable names on the left side of `:=` must be unique.

**【逐字精准翻译】**

与常规变量声明不同，短变量声明可以**重声明（redeclare）\**变量，条件是：这些变量最初是在\**同一个作用域块**（或者如果该作用域块是函数体，则为参数列表）中早先声明的，且具有**相同的类型**，并且**至少有一个非空白变量是新声明的**。由此产生的结果是，重声明只能出现在**多变量短声明**中。重声明并不会引入一个新的变量；它只是将一个新的值赋给原有的变量。`:=` 左侧的非空白变量名称必须是唯一的。

- **词汇与核心规则深度拆解：**
  - `provided (that)`：条件是 / 只要……（引导前提条件）。
  - `in the same block`：在同一个块/作用域中（如果是在外层作用域声明的，在内层块中使用 `:=` **不会**引发重声明，而是会造成**变量遮蔽 Shadowing**！）。
  - `at least one of the non-blank variables is new`：**至少有一个非空白变量是全新的**（例如 `err` 被重复使用，但必须伴随一个全新的变量 `v`）。
  - `Redeclaration does not introduce a new variable`：重声明不会在内存中分配新变量，它本质上退化为**纯赋值**操作。
  - `must be unique`：左侧不能出现重复名称，例如 `x, y, x := 1, 2, 3` 中 `x` 出现了两次，这是极具歧义且非法的。

### 代码示例 2 (合法重声明与非法重复)

> **【英文原文】**
>
> ```go
>field1, offset := nextField(str, 0)
> field2, offset := nextField(str, offset)  // redeclares offset
> x, y, x := 1, 2, 3                        // illegal: x repeated on left side of :=
> ```

**【逐字精准翻译】**

```go
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // 重声明 offset（offset 被赋予新值，field2 为新变量）
x, y, x := 1, 2, 3                        // 非法：x 在 := 的左侧重复出现
```

- **工程实践典型场景：**

  在 Go 处理错误时，这种重声明机制极为普遍：

  ```go
  v1, err := func1() // 声明了 v1 和 err
  v2, err := func2() // 声明了全新变量 v2，重声明了旧变量 err（给 err 赋予新值）
  ```

### 段落 3 (短变量声明的使用位置)

> **【英文原文】**
>
> Short variable declarations may appear only inside functions. In some contexts such as the initializers for "if", "for", or "switch" statements, they can be used to declare local temporary variables.

**【逐字精准翻译】**

短变量声明**只能出现在函数内部**。在某些特定语境中，例如 `"if"`、`"for"` 或 `"switch"` 语句的初始化语句中，它们可用于声明局部的临时变量。

- **词汇与句式拆解：**
  - `only inside functions`：**只能在函数内部**（在包级别/全局作用域中，必须使用 `var`，绝对不能使用 `:=`）。
  - `local temporary variables`：局部临时变量（其作用域仅限于该 `if`/`for`/`switch` 语句块内部，出了语句块即失效）。

本小节逐字翻译与剖析完毕！您列出的最后一行：

> **Function declarations¶**（函数声明）

是下一个大章节的标题。准备好后，我们可以随时继续推进 **函数声明 (Function declarations)** 的原文！

继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，为您逐字逐句精准剖析 **函数声明 (Function declarations)** 这一章节。

这一节详细定义了 Go 语言中的函数声明语法，涵盖普通函数、**泛型函数（Generic function）**、包含返回值的函数的终止语句要求，以及无函数体（外部实现）的函数声明。

## 章节：函数声明 (Function declarations)

### 段落 1 (定义与 EBNF 语法规约)

> **【英文原文】**
>
> A function declaration binds an identifier, the function name, to a function.
>
> EBNF
>
> ```
> FunctionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
> FunctionName = identifier .
> FunctionBody = Block .
> ```

**【逐字精准翻译】**

函数声明将一个标识符（即函数名）绑定到一个函数上。

EBNF

```
函数声明 = "func" 函数名 [ 类型参数列表 ] 函数签名 [ 函数体 ] .
函数名   = 标识符 .
函数体   = 代码块 .
```

- **语法结构深挖：**
  - `[ TypeParameters ]`：中括号表示可选。在 Go 1.18 引入泛型后，这里用来指定泛型的类型参数（如 `[T any]`）。
  - `Signature`：函数签名，包括入参列表和出参（返回值）列表。
  - `[ FunctionBody ]`：**函数体也是可选的！**（如果省略函数体，则对应外部实现的函数，如汇编代码或 Cgo）。

### 段落 2 (返回值与终止语句规则)

> **【英文原文】**
>
> If the function's signature declares result parameters, the function body's statement list must end in a terminating statement.

**【逐字精准翻译】**

如果函数的签名声明了结果参数（返回值），则函数体的语句列表必须以一个终止语句（terminating statement）结尾。

- **词汇与句式拆解：**
  - `result parameters`：结果参数 / 返回值参数。
  - `terminating statement`：终止语句（如 `return` 语句、控制流绝对无法穿透的 `panic` 调用，或永远无法退出的死循环语句等）。
  - **设计哲学：** Go 编译器对静态控制流分析非常严格，不允许任何带有返回值的函数有潜在路径“滑落”出函数末尾而没有明确返回。

### 代码示例 1 (缺失 return 语句的错误示例)

> **【英文原文】**
>
> ```go
>func IndexRune(s string, r rune) int {
> 	for i, c := range s {
> 		if c == r {
> 			return i
> 		}
> 	}
> 	// invalid: missing return statement
> }
> ```

**【逐字精准翻译】**

```go
func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	// 无效：缺失 return 语句
}
```

- **代码分析：**

  虽然逻辑上如果 `s` 中找不到 `r` 循环就会结束，但因为函数指定了返回值类型 `int`，且循环结束后没有放置末尾的 `return` 语句（例如 `return -1`），编译器会直接抛出 `missing return at end of function` 报错。

### 段落 3 (泛型函数声明与实例化)

> **【英文原文】**
>
> If the function declaration specifies type parameters, the function name denotes a generic function [Go 1.18]. A generic function must be instantiated before it can be called or used as a value.

**【逐字精准翻译】**

如果函数声明指定了类型参数，则该函数名代表一个**泛型函数（generic function）** [Go 1.18]。泛型函数在被调用或作为值使用之前，必须先进行**实例化（instantiated）**。

- **词汇与概念拆解：**
  - `denotes a generic function`：代表/表示一个泛型函数。
  - `must be instantiated`：必须被实例化（通过显式提供类型实参或由编译器隐式类型推导，将类型参数 $T$ 替换为具体的类型）。
  - `used as a value`：作为值使用（比如将泛型函数赋值给一个函数指针/变量时，必须指定具体的实例化类型，如 `f := min[int]`）。

### 代码示例 2 (泛型函数定义)

> **【英文原文】**
>
> ```go
>func min[T ~int|~float64](x, y T) T {
> 	if x < y {
> 		return x
> 	}
> 	return y
> }
> ```

**【逐字精准翻译】**

```go
func min[T ~int|~float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}
```

- **语法细节剖析：**
  - `[T ~int|~float64]`：定义类型参数 `T`，约束为底层类型是 `int` 或 `float64` 的任何类型（联合类型约束）。

### 段落 4 (无函数体的外部实现函数)

> **【英文原文】**
>
> A function declaration without type parameters may omit the body. Such a declaration provides the signature for a function implemented outside Go, such as an assembly routine.

**【逐字精准翻译】**

一个没有类型参数的函数声明可以省略函数体。这样的声明为在 Go 外部实现的函数（例如汇编例程）提供函数签名。

- **词汇与句式拆解：**
  - `omit the body`：省略函数体（末尾没有 `{ ... }`，直接以分号或换行结尾）。
  - `assembly routine`：汇编例程 / 汇编子程序（如 Go 运行时 `runtime` 或 `math` 包中大量采用 `.s` 汇编文件实现的底层函数）。
  - **约束提示：** 只有**普通函数**（不带类型参数的函数）才允许省略函数体，泛型函数目前不允许使用无函数体的外部实现。

### 代码示例 3 (外部函数声明)

> **【英文原文】**
>
> ```go
>func flushICache(begin, end uintptr)  // implemented externally
> ```

**【逐字精准翻译】**

```go
func flushICache(begin, end uintptr)  // 在外部实现
```

本章节逐字翻译与剖析完毕！您列出的最后一行：

> **Method declarations¶**（方法声明）

是下一个大章节的标题。准备好后，我们可以随时继续推进 **方法声明 (Method declarations)** 的原文！

继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，为您逐字逐句精准剖析 **方法声明 (Method declarations)** 这一章节。

这一节详细定义了 Go 语言中的方法声明语法、接收者（Receiver）与基础类型（Base type）的绑定规则、泛型接收者的参数传递约束、别名（Alias）对方法绑定的限制，以及 **Go 1.27** 引入的泛型方法（Generic methods）特性。

## 章节：方法声明 (Method declarations)

### 段落 1 (定义与 EBNF 语法规约)

> **【英文原文】**
>
> A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.
>
> EBNF
>
> ```
> MethodDecl = "func" Receiver MethodName [ TypeParameters ] Signature [ FunctionBody ] .
> Receiver   = Parameters .
> ```

**【逐字精准翻译】**

方法是带有接收者的函数。方法声明将一个标识符（即方法名）绑定到一个方法上，并将该方法与接收者的基础类型关联起来。

EBNF

```
方法声明 = "func" 接收者 方法名 [ 类型参数列表 ] 函数签名 [ 函数体 ] .
接收者   = 参数列表 .
```

- **词汇与句式拆解：**
  - `a function with a receiver`：带有接收者的函数（Go 语言中方法的本质）。
  - `associates ... with ...`：将……与……关联。
  - `receiver's base type`：接收者的基础类型。

### 段落 2 (接收者声明规则与限制)

> **【英文原文】**
>
> The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be a defined type $T$ or a pointer to a defined type $T$, possibly followed by a list of type parameter names $[P_1, P_2, \dots]$ enclosed in square brackets. $T$ is called the receiver base type. A receiver base type cannot be a pointer or interface type and it must be declared in the same package as the method. The method is said to be bound to its receiver base type and the method name is visible only within selectors for type $T$ or $*T$.

**【逐字精准翻译】**

接收者是通过位于方法名之前的一个额外参数段落来指定的。该参数段落必须声明单个非可变参数，即接收者。它的类型必须是一个已定义的类型 $T$，或者是指向已定义类型 $T$ 的指针，其后可以紧跟一个包含在方括号中的类型参数名称列表 $[P_1, P_2, \dots]$。$T$ 被称为接收者基础类型。接收者基础类型不能是指针类型或接口类型，并且它必须与该方法声明在同一个包中。该方法被称作绑定到了它的接收者基础类型上，且方法名仅在类型 $T$ 或 $*T$ 的选择器（selector）内部可见。

- **词汇与核心规则深度拆解：**
  - `extra parameter section preceding the method name`：在方法名之前放置的括号参数段，如 `func (p *Point) Length()`。
  - `single non-variadic parameter`：单个非可变参数（不能包含 `...` 可变参数，且只能声明一个接收者变量）。
  - `defined type T`：已定义类型 $T$（用 `type T ...` 显式定义的类型）。
  - **三大硬性约束：**
    1. $T$ 不能本身就是指针类型（如 `type P *int` 不能作为接收者基础类型）或接口类型（`interface`）。
    2. **同包限制**：基础类型 $T$ 必须与方法定义在同一个包（package）中（无法为其他包的类型如 `int` 或 `time.Time` 附加方法）。
    3. **可见性**：方法绑定后，只能通过选择器语法（如 `p.Length()`）进行调用。

### 段落 3 (接收者标识符与忽略规则)

> **【英文原文】**
>
> A non-blank receiver identifier must be unique in the method signature. If the receiver's value is not referenced inside the body of the method, its identifier may be omitted in the declaration. The same applies in general to parameters of functions and methods.

**【逐字精准翻译】**

非空白的接收者标识符在方法签名中必须是唯一的。如果接收者的值在方法体内部没有被引用，它的标识符可以在声明中省略。这通常也适用于函数和方法的参数。

- **词汇与语法细节：**
  - `must be unique`：接收者变量名不能与该方法的入参或返回值变量名重名。
  - `omitted in the declaration`：可以省略标识符（例如 `func (Point) CustomString() string`，在不需要使用接收者实例内部数据时，可以只写类型不写变量名）。

### 段落 4 (唯一性约束：方法名与字段名冲突)

> **【英文原文】**
>
> For a base type, the non-blank names of methods bound to it must be unique. If the base type is a struct type, the non-blank method and field names must be distinct.

**【逐字精准翻译】**

对于一个基础类型，绑定到它的非空白方法名必须是唯一的。如果该基础类型是一个结构体类型，则非空白的方法名和字段名必须互不相同（不能重名）。

- **关键规约：**
  - 如果结构体已经有一个字段叫 `Name string`，就**绝对不能**再定义一个叫 `func (s Struct) Name() string` 的方法，字段名与方法名必须完全隔离。

### 代码示例 1 (基础类型方法绑定)

> **【英文原文】**
>
> Given defined type Point the declarations
>
> ```go
>func (p *Point) Length() float64 {
> 	return math.Sqrt(p.x * p.x + p.y * p.y)
> }
> 
> func (p *Point) Scale(factor float64) {
> 	p.x *= factor
> 	p.y *= factor
> }
> ```
> 
> bind the methods Length and Scale, with receiver type *Point, to the base type Point.

**【逐字精准翻译】**

对于给定的已定义类型 Point，如下声明：

```go
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```

将接收者类型为 `*Point` 的方法 `Length` 和 `Scale` 绑定到了基础类型 `Point` 上。

### 段落 5 (泛型接收者类型参数规范)

> **【英文原文】**
>
> If the receiver base type is a generic type, the receiver specification must declare corresponding type parameters for the method to use. This makes the receiver type parameters available to the method. Syntactically, this type parameter declaration looks like an instantiation of the receiver base type: the type arguments must be identifiers denoting the type parameters being declared, one for each type parameter of the receiver base type. The type parameter names do not need to match their corresponding parameter names in the receiver base type definition, and all non-blank parameter names must be unique in the receiver parameter section and the method signature. The receiver type parameter constraints are implied by the receiver base type definition: corresponding type parameters have corresponding constraints.

**【逐字精准翻译】**

如果接收者基础类型是一个泛型类型，则接收者规范必须声明对应的类型参数以供该方法使用。这使得接收者类型参数在方法体内可用。在语法上，这种类型参数声明看起来像是接收者基础类型的实例化：其类型实参必须是表示正在被声明的类型参数的标识符，每一个标识符对应接收者基础类型的一个类型参数。类型参数的名称不需要与接收者基础类型定义中的对应参数名称相匹配，并且所有非空白的参数名称在接收者参数段落和方法签名中必须是唯一的。接收者类型参数的约束条件是由接收者基础类型的定义隐式决定的：对应的类型参数具有对应的约束条件。

- **概念深度拆解：**
  - `looks like an instantiation`：语法形式上看起来像实例化（如 `Pair[A, B]`），但这里的 `A` 和 `B` 不是具体的类型实参，而是该方法声明的**类型参数形参**。
  - `names do not need to match`：名称无需一致（如定义时是 `Pair[A, B]`，方法上可以写 `Pair[First, Second]` 或用 `_` 忽略）。
  - `constraints are implied`：约束是隐式继承的（无须在方法声明上重复写 `[A any, B any]` 约束条件，基础类型定义里的约束会自动继承）。

### 代码示例 2 (泛型接收者)

> **【英文原文】**
>
> ```go
>type Pair[A, B any] struct {
> 	a A
> 	b B
> }
> 
> func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // receiver declares A, B
> func (p Pair[First, _]) First() First  { … }  // receiver declares First, corresponds to A in Pair
> ```

**【逐字精准翻译】**

```go
type Pair[A, B any] struct {
	a A
	b B
}

func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // 接收者声明了类型参数 A, B
func (p Pair[First, _]) First() First  { … }  // 接收者声明了 First，对应 Pair 中的第一个类型参数 A；用 _ 忽略第二个参数
```

### 段落 6 (类型别名 Alias 绑定方法的禁令)

> **【英文原文】**
>
> If the receiver type is denoted by (a pointer to) an alias, the alias must not be generic and it must not denote an instantiated generic type, neither directly nor indirectly via another alias, and irrespective of pointer indirections.

**【逐字精准翻译】**

如果接收者类型是由一个别名（或指向别名的指针）表示的，则该别名绝不能是泛型的，并且它绝不能表示一个已实例化的泛型类型——无论是直接表示，还是通过另一个别名间接表示，亦或是无论是否存在指针间接引用。

- **关键禁令解读：**
  - 不能给**泛型别名**或**已实例化的泛型别名**绑定方法！
  - 目标是保证方法绑定的基础类型具有绝对确定且未被实例化的语法结构。

### 代码示例 3 (非法的别名接收者)

> **【英文原文】**
>
> ```go
>type GPoint[P any] = Point
> type HPoint        = *GPoint[int]
> type IPair         = Pair[int, int]
> 
> func (*GPoint[P]) Draw(P)   { … }  // illegal: alias must not be generic
> func (HPoint) Draw(P)       { … }  // illegal: alias must not denote instantiated type GPoint[int]
> func (*IPair) Second() int  { … }  // illegal: alias must not denote instantiated type Pair[int, int]
> ```

**【逐字精准翻译】**

```go
type GPoint[P any] = Point
type HPoint        = *GPoint[int]
type IPair         = Pair[int, int]

func (*GPoint[P]) Draw(P)   { … }  // 非法：别名绝不能是泛型的
func (HPoint) Draw(P)       { … }  // 非法：别名绝不能表示已实例化的类型 GPoint[int]
func (*IPair) Second() int  { … }  // 非法：别名绝不能表示已实例化的类型 Pair[int, int]
```

### 段落 7 (Go 1.27 泛型方法 Generic Methods)

> **【英文原文】**
>
> If the method declaration specifies type parameters (possibly in addition to type parameters declared by the receiver specification), the method name denotes a generic method [Go 1.27]. Like a generic function, a generic method must be instantiated before it can be called or used as a value.

**【逐字精准翻译】**

如果方法声明（除了接收者规范所声明的类型参数之外）指定了额外的类型参数，则该方法名代表一个**泛型方法（generic method）** [Go 1.27]。与泛型函数类似，泛型方法在被调用或作为值使用之前，必须先进行实例化。

- **词汇与重大演进拆解：**
  - `[Go 1.27]`：**重要语法演进标志！** Go 在 1.18 引入泛型时，仅支持泛型类型和泛型函数，禁止方法自带独立的类型参数。从 Go 1.27 开始，正式支持在方法上声明独有的类型参数。
  - `generic method`：泛型方法（方法本身拥有自己独立的类型参数列表，如 `Apply[F any]`）。

### 代码示例 4 (泛型方法示例)

> **【英文原文】**
>
> ```go
> type List[E any] []E
> 
> // Apply returns the list obtained from applying f to each element of l.
> func (l List[E]) Apply[F any](f func(E) F) List[F] {
> 	r := make(List[F], len(l))
> 	for i, x := range l {
> 		r[i] = f(x)
> 	}
> 	return r
> }
> ```

**【逐字精准翻译】**

```go
type List[E any] []E

// Apply 返回将函数 f 应用于 l 的每个元素后所获得的列表。
func (l List[E]) Apply[F any](f func(E) F) List[F] {
	r := make(List[F], len(l))
	for i, x := range l {
		r[i] = f(x)
	}
	return r
}
```

- **代码分析：**
  - `List[E]` 拥有接收者类型参数 `E`。
  - 方法 `Apply[F any]` 额外声明了自己的类型参数 `F`，将类型为 `List[E]` 的切片映射转变为类型为 `List[F]` 的切片，展现了典型的函数式映射（Map）能力。

本章节逐字翻译与剖析完毕！您列出的最后一行：

> **Expressions¶**（表达式）

是规范中极其庞大的下一个大章节的标题。准备好后，我们可以随时继续推进 **表达式 (Expressions)** 章节的原文！

