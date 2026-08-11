我们接着保持 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的严密标准，逐字剖析 **If 语句（If statements）** 与 **Switch 语句（Switch statements）** 的开篇规范：

## 章节：If 语句 (If statements)

### 段落 1：基本定义与 EBNF 语法结构

> **【英文原文】**
>
> If statements¶
>
> "If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
>
> ```
> IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
> ```

**【逐字精准翻译】**

If 语句¶

"If" 语句根据布尔表达式的值来指定两个分支的条件执行。如果表达式求值为 true（真），则执行 "if" 分支；否则，如果存在（"else" 分支），则执行 "else" 分支。

```
If 语句 = "if" [ 简单语句 ";" ] 表达式 代码块 [ "else" ( If 语句 | 代码块 ) ] .
```

- **词汇与句式剖析：**
  - `conditional execution`：条件执行。
  - `evaluates to true`：求值为 true / 计算结果为真。
  - `if present`：如果存在（说明 `else` 部分在语法中是可选的）。
  - `SimpleStmt`：简单语句（如短变量声明 `x := f()`、赋值语句或发送语句等）。

### 段落 2：代码示例（基本形式）

> **【英文原文】**
>
> ```go
>if x > max {
> 	x = max
> }
> ```

**【逐字精准翻译】**

```go
if x > max {
	x = max
}
```

### 段落 3：包含前置简单语句的 If 形式

> **【英文原文】**
>
> The expression may be preceded by a simple statement, which executes before the expression is evaluated.
>
> ```go
>if x := f(); x < y {
> 	return x
> } else if x > z {
> 	return z
> } else {
> 	return y
> }
> ```

**【逐字精准翻译】**

表达式前面可以有一个简单语句，该语句在对表达式进行求值之前执行。

```go
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

- **规范重点解析（作用域域规则）：**
  - 在简单语句中声明的变量（如上例中的 `x`），其**作用域**一直延伸至该 `if` 结构及关联的所有 `else if` / `else` 分支代码块结束处。

## 章节：Switch 语句 (Switch statements)

### 段落 1：基本定义与类型划分

> **【英文原文】**
>
> Switch statements¶
>
> "Switch" statements provide multi-way execution. An expression or type is compared to the "cases" inside the "switch" to determine which branch to execute.
>
> ```
> SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .
> ```

**【逐字精准翻译】**

Switch 语句¶

"Switch" 语句提供多分支选择执行。将某个表达式或类型与 "switch" 内部的各 "case" 进行比较，以决定执行哪个分支。

```
Switch 语句 = 表达式 Switch 语句 | 类型 Switch 语句 .
```

- **词汇与句式剖析：**
  - `multi-way execution`：多路执行 / 多分支执行。
  - `compared to`：与……相比较。
  - `ExprSwitchStmt`：表达式 switch 语句。
  - `TypeSwitchStmt`：类型 switch 语句。

### 段落 2：两种 Switch 的核心机制与求值约定

> **【英文原文】**
>
> There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression. In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. The switch expression is evaluated exactly once in a switch statement.

**【逐字精准翻译】**

形式有两种：表达式 switch 和类型 switch。在表达式 switch 中，各 case 包含要与 switch 表达式的值进行比较的表达式。在类型 switch 中，各 case 包含要与带有特殊标注的 switch 表达式的类型进行比较的类型。在 switch 语句中，switch 表达式恰好只被求值一次。

- **词汇与句式剖析：**
  - `specially annotated switch expression`：带有特殊标注的 switch 表达式（指类型断言语法 `x.(type)`）。
  - `evaluated exactly once`：恰好只被求值一次（确保了 switch 表达式中的函数调用或表达式计算不会产生重复执行的副作用）。

你引用的文本正好停在 **Expression switches (表达式 switch)** 的标题处。

准备好后，随时发送 `Expression switches` 及接下来的原文！

我们继续保持 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的严密标准，逐字剖析 **表达式 Switch 语句（Expression switches）** 的语言规范：

## 章节：表达式 Switch 语句 (Expression switches)

### 段落 1：求值顺序与分支匹配逻辑

> **【英文原文】**
>
> Expression switches¶
>
> In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value true.

**【逐字精准翻译】**

表达式 switch 语句¶

在表达式 switch 中，首先对 switch 表达式求值，接着对各 case 表达式（其不必是常量）按照从左到右、从上到下的顺序求值；第一个与 switch 表达式相等的 case 表达式会触发执行关联 case 的语句，其余 case 将被跳过。如果没有 case 匹配且存在一个 "default"（默认）case，则执行其语句。最多只能有一个 default case，且它可以出现在 "switch" 语句中的任何位置。缺失的 switch 表达式等价于布尔值 true。

- **词汇与句式剖析：**
  - `need not be constants`：不必是常量（区别于 C/C++ 等语言，Go 的 case 后面可以跟随变量或表达式）。
  - `left-to-right and top-to-bottom`：从左到右、从上到下（短路求值原则，一旦匹配成功，后续的 case 表达式**不会**继续求值）。
  - `at most one default case`：最多只能有一个 default case。
  - `missing switch expression`：缺失的 switch 表达式（即 `switch { ... }` 形式，等价于 `switch true { ... }`）。

### 段落 2：EBNF 语法结构

> **【英文原文】**
>
> ```
> ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
> ExprCaseClause = ExprSwitchCase ":" StatementList .
> ExprSwitchCase = "case" ExpressionList | "default" .
> ```

**【逐字精准翻译】**

```
表达式 Switch 语句 = "switch" [ 简单语句 ";" ] [ 表达式 ] "{" { 表达式 Case 子句 } "}" .
表达式 Case 子句 = 表达式 Switch Case ":" 语句列表 .
表达式 Switch Case = "case" 表达式列表 | "default" .
```

### 段落 3：Switch 表达式的类型规则

> **【英文原文】**
>
> If the switch expression evaluates to an untyped constant, it is first implicitly converted to its default type. The predeclared untyped value nil cannot be used as a switch expression. The switch expression type must be comparable.

**【逐字精准翻译】**

如果 switch 表达式求值为无类型常量，则其首先被隐式转换为其默认类型。预声明的无类型值 nil 不能用作 switch 表达式。switch 表达式的类型必须是可比较的（comparable）。

- **词汇与句式剖析：**
  - `untyped constant`：无类型常量。
  - `predeclared untyped value nil`：预声明的无类型值 nil（因此 `switch nil` 是非法语法）。
  - `comparable`：可比较的（即必须支持 `==` 和 `!=` 运算符；不能对 slice、map、func 等不可比较类型直接进行 `switch x`）。

### 段落 4：Case 表达式的类型转换与隐式变量初始化

> **【英文原文】**
>
> If a case expression is untyped, it is first implicitly converted to the type of the switch expression. For each (possibly converted) case expression x and the value t of the switch expression, x == t must be a valid comparison.
>
> In other words, the switch expression is treated as if it were used to declare and initialize a temporary variable t without explicit type; it is that value of t against which each case expression x is tested for equality.

**【逐字精准翻译】**

如果某个 case 表达式是无类型的，则其首先被隐式转换为 switch 表达式的类型。对于每个（可能经过转换的）case 表达式 x 以及 switch 表达式的值 t，x == t 必须是有效的比较。

换句话说，switch 表达式的处理方式，就像是用它来声明并初始化一个没有显式类型的临时变量 t 一样；正是对 t 的该值与每个 case 表达式 x 进行相等性测试。

- **词汇与句式剖析：**
  - `valid comparison`：有效的比较（要求类型可比且满足 Go 的相等性比较规则）。
  - `temporary variable t`：临时变量 t（规范抽象出的逻辑变量，解释了隐式类型转换的基准）。

### 段落 5：fallthrough 穿透机制

> **【英文原文】**
>
> In a case or default clause, the last non-empty statement may be a (possibly labeled) "fallthrough" statement to indicate that control should flow from the end of this clause to the first statement of the next clause. Otherwise control flows to the end of the "switch" statement. A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.

**【逐字精准翻译】**

在一个 case 或 default 子句中，最后一个非空语句可以是一个（可能带有标签的）"fallthrough" 语句，以指示控制流应当从本子句的末尾流向下一个子句的第一条语句。否则，控制流将流向 "switch" 语句的末尾。除了表达式 switch 的最后一个子句之外，"fallthrough" 语句可以作为所有其他子句的最后一条语句出现。

- **规范底层关键：**
  1. Go 语言的 switch **默认不会贯穿（No implicit fallthrough）**，每个 case 执行完自动跳出，无需显式写 `break`。
  2. `fallthrough` 会**无条件强制进入下一个 case** 执行其语句，而**不会评估下一个 case 的表达式是否匹配**。
  3. `fallthrough` 只能出现在子句的末尾，且不能在最后一个 clause 中使用。

### 段落 6：前置简单语句与代码示例

> **【英文原文】**
>
> The switch expression may be preceded by a simple statement, which executes before the expression is evaluated.
>
> ```go
>switch tag {
> default: s3()
> case 0, 1, 2, 3: s1()
> case 4, 5, 6, 7: s2()
> }
> 
> switch x := f(); {  // missing switch expression means "true"
> case x < 0: return -x
> default: return x
> }
> 
> switch {
> case x < y: f1()
> case x < z: f2()
> case x == 4: f3()
> }
> ```

**【逐字精准翻译】**

switch 表达式前面可以有一个简单语句，该语句在对表达式求值之前执行。

```go
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // 缺失 switch 表达式意味着 "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```

### 段落 7：编译器实现限制（重复 Case 检测）

> **【英文原文】**
>
> Implementation restriction: A compiler may disallow multiple case expressions evaluating to the same constant. For instance, the current compilers disallow duplicate integer, floating point, or string constants in case expressions.

**【逐字精准翻译】**

实现限制：编译器可以不允许求值结果为相同常量的多个 case 表达式。例如，目前的编译器不允许在 case 表达式中出现重复的整数、浮点数或字符串常量。

- **规范解读：**

  例如 `case 1, 1:` 或两个独立的 `case 1:` 都会被官方编译器 `gc` 拒绝并报错 `duplicate case 1 in switch`。

文本停在 **Type switches (类型 switch)** 标题处。准备好后，随时发送后续原文！

继续保持 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的严密标准，逐字剖析 **类型 Switch 语句（Type switches）** 的底层规范：

## 章节：类型 Switch 语句 (Type switches)

### 段落 1：核心定义与类型约束

> **【英文原文】**
>
> Type switches¶
>
> A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a type assertion using the keyword type rather than an actual type:
>
> ```go
>switch x.(type) {
> // cases
> }
> ```
> 
> Cases then match actual types T against the dynamic type of the expression x. As with type assertions, x must be of interface type, but not a type parameter, and each non-interface type T listed in a case must implement the type of x. The types listed in the cases of a type switch must all be different.

**【逐字精准翻译】**

类型 switch 语句¶

类型 switch 比较的是类型而非数值。在其他方面，它与表达式 switch 相似。它由一个特殊的 switch 表达式标注，该表达式具有类型断言的形式，使用的是关键字 `type` 而非实际的类型：

```go
switch x.(type) {
// cases
}
```

随后，各 case 将实际类型 T 与表达式 x 的动态类型（dynamic type）进行匹配。与类型断言一样，x 必须是接口类型（interface type），但不能是类型参数（type parameter），并且 case 中列出的每个非接口类型 T 都必须实现 x 的类型。类型 switch 的各 case 中列出的类型必须互不相同。

- **词汇与句式剖析：**
  - `compares types rather than values`：比较的是类型而非数值。
  - `dynamic type`：动态类型（即接口变量在运行时具体包含的值的类型）。
  - `interface type`：接口类型（x 必须是接口，否则编译报错）。
  - `type parameter`：类型参数（Go 1.18 泛型中的类型参数不能直接用于 `x.(type)`，但可以在 case 中作为匹配目标类型）。
  - `must implement the type of x`：必须实现 x 的类型（例如若 x 为 `io.Reader`，case 中的具体类型必须实现 `Read` 方法）。

### 段落 2：EBNF 语法结构

> **【英文原文】**
>
> ```
> TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
> TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
> TypeCaseClause  = TypeSwitchCase ":" StatementList .
> TypeSwitchCase  = "case" TypeList | "default" .
> ```

**【逐字精准翻译】**

```
类型 Switch 语句 = "switch" [ 简单语句 ";" ] 类型 Switch 守卫 "{" { 类型 Case 子句 } "}" .
类型 Switch 守卫 = [ 标识符 ":=" ] 初级表达式 "." "(" "type" ")" .
类型 Case 子句   = 类型 Switch Case ":" 语句列表 .
类型 Switch Case = "case" 类型列表 | "default" .
```

### 段落 3：短变量声明与类型遮蔽（Type Shadowing / Refinement）

> **【英文原文】**
>
> The TypeSwitchGuard may include a short variable declaration. When that form is used, the variable is declared at the end of the TypeSwitchCase in the implicit block of each clause. In clauses with a case listing exactly one type, the variable has that type; otherwise, the variable has the type of the expression in the TypeSwitchGuard.
>
> Instead of a type, a case may use the predeclared identifier nil; that case is selected when the expression in the TypeSwitchGuard is a nil interface value. There may be at most one nil case.

**【逐字精准翻译】**

类型 Switch 守卫（TypeSwitchGuard）可以包含一个短变量声明。当使用该形式时，变量会在每个子句的隐式代码块中、类型 Switch Case 的末尾被声明。在那些 case 恰好列出一个类型的子句中，该变量具有该类型；否则，该变量具有类型 Switch 守卫中表达式的类型。

case 可以使用预声明标识符 nil 来代替类型；当类型 Switch 守卫中的表达式是一个 nil 接口值时，选择该 case。最多只能有一个 nil case。

- **词汇与句式剖析：**
  - `short variable declaration`：短变量声明（如 `switch i := x.(type)`）。
  - `implicit block of each clause`：每个子句的隐式代码块（这意味着变量 `i` 在每个 case 分支内部都有独立的作用域和具体类型）。
  - `listing exactly one type`：恰好列出一个类型（如 `case int:`，变量 `i` 的自动推导类型会被精确窄化为 `int`）。
  - `nil interface value`：nil 接口值（即接口变量本身的动态类型和动态值均为 nil）。

### 段落 4：类型 Switch 与 If-Else 解构的等价对照

> **【英文原文与代码例证】**
>
> Given an expression x of type interface{}, the following type switch:
>
> ```go
>switch i := x.(type) {
> case nil:
> 	printString("x is nil")                // type of i is type of x (interface{})
> case int:
> 	printInt(i)                            // type of i is int
> case float64:
> 	printFloat64(i)                        // type of i is float64
> case func(int) float64:
> 	printFunction(i)                       // type of i is func(int) float64
> case bool, string:
> 	printString("type is bool or string")  // type of i is type of x (interface{})
> default:
> 	printString("don't know the type")     // type of i is type of x (interface{})
> }
> ```
> 
> could be rewritten:
>
> ```go
>v := x  // x is evaluated exactly once
> if v == nil {
>	i := v                                 // type of i is type of x (interface{})
> 	printString("x is nil")
> } else if i, isInt := v.(int); isInt {
> 	printInt(i)                            // type of i is int
> } else if i, isFloat64 := v.(float64); isFloat64 {
> 	printFloat64(i)                        // type of i is float64
> } else if i, isFunc := v.(func(int) float64); isFunc {
> 	printFunction(i)                       // type of i is func(int) float64
> } else {
> 	_, isBool := v.(bool)
> 	_, isString := v.(string)
> 	if isBool || isString {
> 		i := v                         // type of i is type of x (interface{})
> 		printString("type is bool or string")
> 	} else {
> 		i := v                         // type of i is type of x (interface{})
> 		printString("don't know the type")
> 	}
> }
> ```

**【逐字精准翻译】**

假设给定一个类型为 `interface{}` 的表达式 x，以下类型 switch 语句：

```go
switch i := x.(type) {
case nil:
	printString("x is nil")                // i 的类型是 x 的类型 (interface{})
case int:
	printInt(i)                            // i 的类型是 int
case float64:
	printFloat64(i)                        // i 的类型是 float64
case func(int) float64:
	printFunction(i)                       // i 的类型是 func(int) float64
case bool, string:
	printString("type is bool or string")  // i 的类型是 x 的类型 (interface{})
default:
	printString("don't know the type")     // i 的类型是 x 的类型 (interface{})
}
```

可以重写为：

```go
v := x  // x 恰好只求值一次
if v == nil {
	i := v                                 // i 的类型是 x 的类型 (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // i 的类型是 int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // i 的类型是 float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // i 的类型是 func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // i 的类型是 x 的类型 (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // i 的类型是 x 的类型 (interface{})
		printString("don't know the type")
	}
}
```

- **规范重点解读：**
  - 当 case 列出多个类型（如 `case bool, string:`）时，变量 `i` **不会**自动窄化为某个具体类型，而是保留原始接口类型（`interface{}`）。

### 段落 5：泛型与实例化重复处理 (Go 1.18+)

> **【英文原文】**
>
> A type parameter or a generic type may be used as a type in a case. If upon instantiation that type turns out to duplicate another entry in the switch, the first matching case is chosen.
>
> ```go
>func f[P any](x any) int {
> 	switch x.(type) {
> 	case P:
> 		return 0
> 	case string:
> 		return 1
> 	case []P:
> 		return 2
> 	case []byte:
> 		return 3
> 	default:
> 		return 4
> 	}
> }
> 
> var v1 = f[string]("foo")   // v1 == 0
> var v2 = f[byte]([]byte{})  // v2 == 2
> ```

**【逐字精准翻译】**

类型参数或泛型类型可以作为 case 中的类型使用。如果在实例化时，该类型经证明与 switch 中的另一个条目重复，则选择第一个匹配的 case。

```go
func f[P any](x any) int {
	switch x.(type) {
	case P:
		return 0
	case string:
		return 1
	case []P:
		return 2
	case []byte:
		return 3
	default:
		return 4
	}
}

var v1 = f[string]("foo")   // v1 == 0（实例化后 P 成为 string，与 case string 重复，但优先匹配先出现的 case P）
var v2 = f[byte]([]byte{})  // v2 == 2（实例化后 P 成为 byte，[]P 即 []byte，优先匹配先出现的 case []P）
```

### 段落 6：前置语句与 fallthrough 禁令

> **【英文原文】**
>
> The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.
>
> The "fallthrough" statement is not permitted in a type switch.

**【逐字精准翻译】**

类型 switch 守卫前面可以有一个简单语句，该语句在对守卫求值之前执行。

"fallthrough" 语句在类型 switch 中是**绝对不允许**的。

- **底层逻辑：**

  因为类型 switch 会在不同的 case 分支中将变量重新绑定为不同的精确类型，一旦允许 `fallthrough` 跨分支穿透，会导致变量类型系统混乱崩溃，因此规范明确禁止在类型 switch 中使用 `fallthrough`。

文本停在 **For statements (For 循环语句)** 的标题处。准备好后，随时发送后续原文！

接着保持 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的严密标准，逐字剖析 **For 语句（For statements）前半部分：单条件循环与 For 子句循环** 的语言规范：

## 章节：For 语句 (For statements)

### 段落 1：基本定义与三元划分

> **【英文原文】**
>
> For statements¶
>
> A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.
>
> ```
> ForStmt   = "for" [ Condition | ForClause | RangeClause ] Block .
> Condition = Expression .
> ```

**【逐字精准翻译】**

For 语句¶

"for" 语句指定代码块的重复执行。形式有三种：迭代可以由单个条件（Condition）、"for" 子句（ForClause）或 "range" 子句（RangeClause）来控制。

```
For 语句 = "for" [ 条件 | For 子句 | Range 子句 ] 代码块 .
条件     = 表达式 .
```

- **词汇与句式剖析：**
  - `repeated execution`：重复执行。
  - `controlled by`：由……控制。
  - `Condition / ForClause / RangeClause`：单条件 / 三段式 For 子句 / 范围遍历 Range 子句。

## 章节：带单条件的 For 语句 (For statements with single condition)

### 段落 1：语义与缺省规则

> **【英文原文】**
>
> For statements with single condition¶
>
> In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value true.
>
> ```go
>for a < b {
> 	a *= 2
> }
> ```

**【逐字精准翻译】**

带单条件的 For 语句¶

在其最简单的形式中，只要布尔条件求值为 true，"for" 语句就会指定代码块的重复执行。在每次迭代之前均会对条件进行求值。如果条件缺失，则等价于布尔值 true（相当于 C 语言中的 `while(1)` 或 `while(cond)`）。

```go
for a < b {
	a *= 2
}
```

- **词汇与句式剖析：**
  - `as long as`：只要……。
  - `evaluated before each iteration`：在每次迭代之前求值（即前置条件循环）。
  - `if the condition is absent`：如果条件缺失（即 `for { ... }` 形式）。

## 章节：带 For 子句的 For 语句 (For statements with for clause)

### 段落 1：语法定义与简单语句限制

> **【英文原文】**
>
> For statements with for clause¶
>
> A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an init and a post statement, such as an assignment, an increment or decrement statement. The init statement may be a short variable declaration, but the post statement must not.
>
> ```
> ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
> InitStmt  = SimpleStmt .
> PostStmt  = SimpleStmt .
> ```
>
> Go
>
> ```go
> for i := 0; i < 10; i++ {
> 	f(i)
> }
> ```

**【逐字精准翻译】**

带 For 子句的 For 语句¶

带 For 子句的 "for" 语句同样由其条件控制，但此外它还可以指定一个初始化语句（init）和一个后置语句（post），例如赋值语句、自增或自减语句。初始化语句可以是短变量声明，但后置语句**绝对不能**是短变量声明。

```
For 子句   = [ 初始化语句 ] ";" [ 条件 ] ";" [ 后置语句 ] .
初始化语句 = 简单语句 .
后置语句   = 简单语句 .
```

Go

```go
for i := 0; i < 10; i++ {
	f(i)
}
```

- **规范底层约束：**
  - `post statement must not [be a short variable declaration]`：后置语句绝不能是短变量声明（例如不能在三段式的第三段写 `i := i + 1`）。

### 段落 2：执行顺序与分号规则

> **【英文原文】**
>
> If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the semicolons are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value true.
>
> Plaintext
>
> ```
> for cond { S() }    is the same as    for ; cond ; { S() }
> for      { S() }    is the same as    for true     { S() }
> ```

**【逐字精准翻译】**

如果不为空，初始化语句会在第一次迭代对条件求值之前执行一次；后置语句会在每次执行代码块之后执行（且仅当代码块被执行了时才执行）。For 子句的任何元素都可以为空，但除非只有条件，否则分号是必需的。如果条件缺失，则等价于布尔值 true。

Plaintext

```
for cond { S() }    等同于    for ; cond ; { S() }
for      { S() }    等同于    for true     { S() }
```

- **词汇与句式剖析：**
  - `and only if the block was executed`：且仅当代码块被执行了时（若因为条件评估为 false 或初始不满足而直接跳出，后置语句不执行）。

### 段落 3：Go 1.22+ 循环变量独立绑定重大变更（关键语法特性！）

> **【英文原文】**
>
> Each iteration has its own separate declared variable (or variables) [Go 1.22]. The variable used by the first iteration is declared by the init statement. The variable used by each subsequent iteration is declared implicitly before executing the post statement and initialized to the value of the previous iteration's variable at that moment.
>
> ```go
>var prints []func()
> for i := 0; i < 5; i++ {
> 	prints = append(prints, func() { println(i) })
> 	i++
> }
> for _, p := range prints {
> 	p()
> }
> ```
> 
> prints
>
> Plaintext
>
> ```
>1
> 3
> 5
> ```
> 
> Prior to [Go 1.22], iterations share one set of variables instead of having their own separate variables. In that case, the example above prints
>
> Plaintext
>
> ```
>6
> 6
> 6
> ```

**【逐字精准翻译】**

每次迭代都有其自身独立声明的变量（或多个变量）[Go 1.22 引入]。第一次迭代使用的变量由初始化语句声明。后续每次迭代使用的变量，会在执行后置语句之前隐式声明，并在那一刻初始化为上一次迭代变量的值。

```go
var prints []func()
for i := 0; i < 5; i++ {
	prints = append(prints, func() { println(i) })
	i++
}
for _, p := range prints {
	p()
}
```

打印输出：

Plaintext

```
1
3
5
```

在 [Go 1.22] 之前，所有迭代共享同一套变量，而不是拥有各自独立的变量。在那种情况下，上面的示例会打印：

Plaintext

```
6
6
6
```

- **语言规范变革与底层逻辑剖析：**
  1. **Go 1.22 之前的陷阱（共享变量）：** 以前 `i` 在整个 `for` 循环中是同一个内存地址。闭包捕获的是 `i` 的引用，循环结束后 `i` 变成了 6，因此打印全部是 `6`。
  2. **Go 1.22+ 语义（Per-iteration variable）：** 规范明确说明“每次迭代都会隐式新声明一个 `i` 并将上一次的 `i` 值拷贝过来”。
  3. **跟踪执行过程：**
     - **Iteration 1 (`i=0`)：** 闭包捕获当前的 `i_1`（值为0）。在块内执行 `i++`（`i_1` 变为 1）。打印列表存入闭包。执行 post 语句 `i++` 前，隐式创建 `i_2 = 1`，post 语句 `i++` 作用于 `i_2`（变为 2）。
     - **Iteration 2 (`i=2`)：** 闭包捕获 `i_2`（值为 2），块内 `i++` 后 `i_2` 变为 3……
     - **输出结果：** 闭包执行时打印块内 append 时刻捕获的值，即 `1, 3, 5`。

文本停在 **For statements with range clause (带 range 子句的 For 语句)** 的标题处。准备好后，随时发送后续原文！

我们需要继续保持 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的严密标准，逐字剖析 **带 Range 子句的 For 语句（For statements with range clause）**：

## 章节：带 Range 子句的 For 语句 (For statements with range clause)

### 段落 1：基本定义与迭代数据源支持

> **【英文原文】**
>
> For statements with range clause¶
>
> A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, values received on a channel, integer values from zero to an upper limit [Go 1.22], or values passed to an iterator function's yield function [Go 1.23]. For each entry it assigns iteration values to corresponding iteration variables if present and then executes the block.
>
> ```
> RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
> ```

**【逐字精准翻译】**

带 Range 子句的 For 语句¶

带 "range" 子句的 "for" 语句遍历数组（array）、切片（slice）、字符串（string）或映射（map）的所有条目，或者通道（channel）上接收到的值，或者从 0 到上限的整数值 [Go 1.22 引入]，或者传递给迭代器函数的 yield 函数的值 [Go 1.23 引入]。对于每个条目，如果存在迭代变量，它会将迭代值赋给相应的迭代变量，然后执行代码块。

```
Range 子句 = [ 表达式列表 "=" | 标识符列表 ":=" ] "range" 表达式 .
```

- **词汇与句式剖析：**
  - `iterates through`：遍历 / 迭代。
  - `iteration values / iteration variables`：迭代值 / 迭代变量。
  - `upper limit`：上限（如 `range 10` 产生 $0 \dots 9$）。
  - `yield function`：收益函数 / 迭代交出函数（Go 1.23 引入的自定义迭代器 `iter.Seq` 核心机制）。

### 段落 2：Range 表达式类型与左值变量限制

> **【英文原文】**
>
> The expression on the right in the "range" clause is called the range expression, which may be an array, pointer to an array, slice, string, map, channel permitting receive operations, an integer, or a function with specific signature (see below). As with an assignment, if present the operands on the left must be addressable or map index expressions; they denote the iteration variables. If the range expression is a function, the maximum number of iteration variables depends on the function signature. If the range expression is a channel or integer, at most one iteration variable is permitted; otherwise there may be up to two. If the last iteration variable is the blank identifier, the range clause is equivalent to the same clause without that identifier.

**【逐字精准翻译】**

"range" 子句中右侧的表达式被称为 range 表达式，它可以是数组、指向数组的指针、切片、字符串、map、允许接收操作的通道、整数，或者具有特定签名（见下文）的函数。与赋值语句一样，如果存在左侧操作数，其必须是可寻址的或 map 的索引表达式；它们表示迭代变量。如果 range 表达式是一个函数，则迭代变量的最大数量取决于该函数的签名。如果 range 表达式是通道或整数，则最多允许一个迭代变量；否则最多可以有两个。如果最后一个迭代变量是空标识符，则 range 子句等价于不带该标识符的相同子句。

- **词汇与句式剖析：**
  - `range expression`：range 表达式（即 `range` 关键字右侧的数据源）。
  - `channel permitting receive operations`：允许接收操作的通道（即双向通道或只读通道 `<-chan T`）。
  - `at most one iteration variable`：最多允许一个迭代变量。

### 段落 3：Range 表达式求值时机（优化特例）

> **【英文原文】**
>
> The range expression x is evaluated before beginning the loop, with one exception: if at most one iteration variable is present and x or len(x) is constant, the range expression is not evaluated.
>
> Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present:

**【逐字精准翻译】**

range 表达式 x 会在循环开始之前求值，但有一个例外：如果最多只存在一个迭代变量，且 x 或 len(x) 是常量，则不会对 range 表达式求值。

左侧的函数调用在每次迭代时均求值一次。对于每次迭代，如果存在相应的迭代变量，则按照如下方式产生迭代值：

- **核心优化逻辑：**

  当遍历指向固定长度数组的指针 `testdata.a` 且只获取索引 `i` 时，`len(testdata.a)` 在编译期就是常量，此时完全不需要在运行时为 `testdata` 解引用求值。

### 段落 4：迭代数据源与迭代值映射表

> **【英文原文】**
>
> Plaintext
>
> ```
> Range expression                                       1st value                2nd value
> 
> array or slice      a  [n]E, *[n]E, or []E             index    i  int          a[i]       E
> string              s  string type                     index    i  int          see below  rune
> map                 m  map[K]V                         key      k  K            m[k]       V
> channel             c  chan E, <-chan E                element  e  E
> integer value       n  integer type, or untyped int    value    i  see below
> function, 0 values  f  func(yield func() bool)
> function, 1 value   f  func(yield func(V) bool)        value    v  V                               yield cannot be variadic
> function, 2 values  f  func(yield func(K, V) bool)     key      k  K            v          V       yield cannot be variadic
> ```

**【逐字精准翻译】**

Plaintext

```
Range 表达式                                            第 1 个值                第 2 个值

数组或切片          a  [n]E, *[n]E, 或 []E              索引     i  int          a[i]       E
字符串              s  字符串类型                       索引     i  int          见下文     rune
映射 (map)          m  map[K]V                         键       k  K            m[k]       V
通道 (channel)      c  chan E, <-chan E                元素     e  E
整数值              n  整数类型, 或无类型整型            数值     i  见下文
函数 (0 个值)       f  func(yield func() bool)
函数 (1 个值)       f  func(yield func(V) bool)        数值     v  V                               yield 不能是可变参数
函数 (2 个值)       f  func(yield func(K, V) bool)     键       k  K            v          V       yield 不能是可变参数
```

### 段落 5：各种数据类型的 Range 详细机制

> **【英文原文】**
>
> For an array, pointer to array, or slice value a, the index iteration values are produced in increasing order, starting at element index 0. If at most one iteration variable is present, the range loop produces iteration values from 0 up to len(a)-1 and does not index into the array or slice itself. For a nil slice, the number of iterations is 0.
>
> For a string value, the "range" clause iterates over the Unicode code points in the string starting at byte index 0. On successive iterations, the index value will be the index of the first byte of successive UTF-8-encoded code points in the string, and the second value, of type rune, will be the value of the corresponding code point. If the iteration encounters an invalid UTF-8 sequence, the second value will be 0xFFFD, the Unicode replacement character, and the next iteration will advance a single byte in the string.
>
> The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next. If a map entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced. If a map entry is created during iteration, that entry may be produced during the iteration or may be skipped. The choice may vary for each entry created and from one iteration to the next. If the map is nil, the number of iterations is 0.
>
> For channels, the iteration values produced are the successive values sent on the channel until the channel is closed. If the channel is nil, the range expression blocks forever.

**【逐字精准翻译】**

对于数组、数组指针或切片值 a，索引迭代值从元素索引 0 开始按递增顺序产生。如果最多只存在一个迭代变量，range 循环会产生从 0 到 len(a)-1 的迭代值，且不会对数组或切片本身进行索引访问。对于 nil 切片，迭代次数为 0。

对于字符串值，"range" 子句从字节索引 0 开始遍历字符串中的 Unicode 码点（Unicode code points）。在后续迭代中，索引值将是字符串中连续 UTF-8 编码码点的首字节索引，而类型为 rune 的第二个值将是对应码点的值。如果迭代遇到无效的 UTF-8 序列，则第二个值将是 0xFFFD（Unicode 替换字符），并且下一次迭代将在字符串中前进一步（单字节）。

map 的迭代顺序未作指定，并且不能保证两次迭代之间的顺序相同。如果在迭代过程中删除了尚未到达的 map 条目，则不会产生对应的迭代值。如果在迭代过程中创建了 map 条目，则该条目可能会在迭代中被产生，也可能会被跳过。对每个创建的条目以及每次迭代，其选择可能有所不同。如果 map 为 nil，则迭代次数为 0。

对于通道，产生的迭代值是发送到通道上的连续值，直到通道被关闭。如果通道为 nil，则 range 表达式将永远阻塞。

- **关键概念解析：**
  - `Unicode replacement character (0xFFFD)`：替换字符 ``，Go 在解析到无效 UTF-8 字节时会自动用它顶替，并将索引递增 1 字节。
  - `map 迭代无序性`：Go 在底层对 map 遍历施加了随机种子，即使内容不变，多次遍历的顺序也故意设计为随机的。

### 段落 6：整数 Range 机制 (Go 1.22+)

> **【英文原文】**
>
> For an integer value n, where n is of integer type or an untyped integer constant, the iteration values 0 through n-1 are produced in increasing order. If n is of integer type, the iteration values have that same type. Otherwise, the type of n is determined as if it were assigned to the iteration variable. Specifically: if the iteration variable is preexisting, the type of the iteration values is the type of the iteration variable, which must be of integer type. Otherwise, if the iteration variable is declared by the "range" clause or is absent, the type of the iteration values is the default type for n. If n <= 0, the loop does not run any iterations.

**【逐字精准翻译】**

对于整数值 n（其中 n 为整数类型或无类型整数常量），按递增顺序产生从 0 到 n-1 的迭代值。如果 n 是整数类型，则迭代值具有相同的类型。否则，n 的类型判定如同将其赋给迭代变量一样。具体而言：如果迭代变量是已存在的，则迭代值的类型就是该迭代变量的类型（其必须是整数类型）。否则，如果迭代变量是由 "range" 子句声明的或者是缺失的，则迭代值的类型是 n 的默认类型（即 `int`）。如果 n <= 0，则循环不执行任何迭代。

### 段落 7：自定义迭代器函数 Range 机制 (Go 1.23+)

> **【英文原文】**
>
> For a function f, the iteration proceeds by calling f with a new, synthesized yield function as its argument. If yield is called before f returns, the arguments to yield become the iteration values for executing the loop body once. After each successive loop iteration, yield returns true and may be called again to continue the loop. As long as the loop body does not terminate, the "range" clause will continue to generate iteration values this way for each yield call until f returns. If the loop body terminates (such as by a break statement), yield returns false and must not be called again.

**【逐字精准翻译】**

对于函数 f，迭代是通过以一个新的、合成的 yield 函数作为参数调用 f 来进行的。如果在 f 返回之前调用了 yield，则传递给 yield 的参数将成为执行一次循环体的迭代值。在每次成功的循环迭代之后，yield 返回 true，并且可以被再次调用以继续循环。只要循环体没有终止，"range" 子句就会通过这种方式为每次 yield 调用继续生成迭代值，直到 f 返回。如果循环体终止（例如通过 break 语句），yield 会返回 false，且绝不能被再次调用。

- **底层流控（Push Iterator）剖析：**

  Go 1.23 采用的是推模式迭代器（Push-based Iterator）。当循环体内遇到 `break` 或 `return` 时，闭包 `yield` 会返回 `false`，迭代器函数 `f` 捕获到 `false` 后应当立即停止生产数据并清理资源返回。

### 段落 8：泛型类型参数下的 Range 约束

> **【英文原文】**
>
> If the type of the range expression is a type parameter, all types in its type set must have the same underlying type and the range expression must be valid for that type, or, if the type set contains channel types, it must only contain channel types with identical element types, and all channel types must permit receive operations.

**【逐字精准翻译】**

如果 range 表达式的类型是类型参数（泛型），则其类型集中的所有类型都必须具有相同的底层类型（underlying type），且 range 表达式对该类型必须是有效的；或者，如果类型集包含通道类型，则它必须仅包含具有相同元素类型的通道类型，并且所有通道类型都必须允许接收操作。

### 段落 9：迭代变量的作用域与生命周期 (Go 1.22 独立变量绑定)

> **【英文原文】**
>
> The iteration variables may be declared by the "range" clause using a form of short variable declaration (:=). In this case their scope is the block of the "for" statement and each iteration has its own new variables [Go 1.22] (see also "for" statements with a ForClause). The variables have the types of their respective iteration values.
>
> If the iteration variables are not explicitly declared by the "range" clause, they must be preexisting. In this case, the iteration values are assigned to the respective variables as in an assignment statement.

**【逐字精准翻译】**

迭代变量可以通过 "range" 子句使用短变量声明（:=）的形式来声明。在这种情况下，它们的作用域是 "for" 语句的代码块，并且每次迭代都拥有其自身全新的变量 [Go 1.22 引入]（亦参见带 For 子句的 "for" 语句）。这些变量具有其各自迭代值的类型。

如果迭代变量没有由 "range" 子句显式声明，则它们必须是已存在的。在这种情况下，迭代值如在赋值语句中那样被赋给对应的变量。

### 代码示例与逐行解析

> **【英文原文与代码】**
>
> ```go
>var testdata *struct {
> 	a *[7]int
> }
> for i, _ := range testdata.a {
> 	// testdata.a is never evaluated; len(testdata.a) is constant
> 	// i ranges from 0 to 6
> 	f(i)
> }
> 
> var a [10]string
> for i, s := range a {
> 	// type of i is int
> 	// type of s is string
> 	// s == a[i]
> 	g(i, s)
> }
> 
> var key string
> var val interface{}  // element type of m is assignable to val
> m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
> for key, val = range m {
> 	h(key, val)
> }
> // key == last map key encountered in iteration
> // val == map[key]
> 
> var ch chan Work = producer()
> for w := range ch {
> 	doWork(w)
> }
> 
> // empty a channel
> for range ch {}
> 
> // call f(0), f(1), ... f(9)
> for i := range 10 {
> 	// type of i is int (default type for untyped constant 10)
> 	f(i)
> }
> 
> // invalid: 256 cannot be assigned to uint8
> var u uint8
> for u = range 256 {
> }
> 
> // invalid: 1e3 is a floating-point constant
> for range 1e3 {
> }
> 
> // fibo generates the Fibonacci sequence
> fibo := func(yield func(x int) bool) {
> 	f0, f1 := 0, 1
> 	for yield(f0) {
> 		f0, f1 = f1, f0+f1
> 	}
> }
> 
> // print the Fibonacci numbers below 1000:
> for x := range fibo {
> 	if x >= 1000 {
> 		break
> 	}
> 	fmt.Printf("%d ", x)
> }
> // output: 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987
> 
> // iteration support for a recursive tree data structure
> type Tree[K cmp.Ordered, V any] struct {
> 	left, right *Tree[K, V]
> 	key         K
> 	value       V
> }
> 
> func (t *Tree[K, V]) walk(yield func(key K, val V) bool) bool {
> 	return t == nil || t.left.walk(yield) && yield(t.key, t.value) && t.right.walk(yield)
> }
> 
> func (t *Tree[K, V]) Walk(yield func(key K, val V) bool) {
> 	t.walk(yield)
> }
> 
> // walk tree t in-order
> var t Tree[string, int]
> for k, v := range t.Walk {
> 	// process k, v
> }
> 
> // xor returns the xor-ed bytes of S
> func xor[S ~[]byte](s S) byte {
> 	var r byte
> 	for _, b := range s {
> 		r ^= b
> 	}
> 	return r
> }
> ```

**【逐字精准翻译与语法注释】**

```go
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a 从未被求值；len(testdata.a) 是常量
	// i 的范围是从 0 到 6
	f(i)
}

var a [10]string
for i, s := range a {
	// i 的类型是 int
	// s 的类型是 string
	// s == a[i]
	g(i, s)
}

var key string
var val interface{}  // m 的元素类型可以赋值给 val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == 迭代中遇到的最后一个 map 键
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// 清空一个通道
for range ch {}

// 依次调用 f(0), f(1), ... f(9)
for i := range 10 {
	// i 的类型是 int（无类型常量 10 的默认类型）
	f(i)
}

// 无效：256 无法赋值给 uint8（溢出）
var u uint8
for u = range 256 {
}

// 无效：1e3 是浮点数常量，range 整数遍历不支持浮点数
for range 1e3 {
}

// fibo 生成斐波那契数列（Go 1.23 自定义迭代器）
fibo := func(yield func(x int) bool) {
	f0, f1 := 0, 1
	for yield(f0) {
		f0, f1 = f1, f0+f1
	}
}

// 打印 1000 以内的斐波那契数：
for x := range fibo {
	if x >= 1000 {
		break
	}
	fmt.Printf("%d ", x)
}
// 输出: 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987

// 针对递归树数据结构的迭代支持
type Tree[K cmp.Ordered, V any] struct {
	left, right *Tree[K, V]
	key         K
	value       V
}

func (t *Tree[K, V]) walk(yield func(key K, val V) bool) bool {
	return t == nil || t.left.walk(yield) && yield(t.key, t.value) && t.right.walk(yield)
}

func (t *Tree[K, V]) Walk(yield func(key K, val V) bool) {
	t.walk(yield)
}

// 中序遍历树 t
var t Tree[string, int]
for k, v := range t.Walk {
	// 处理 k, v
}

// xor 返回 S 的所有字节按位异或的结果（泛型底层类型切片示例）
func xor[S ~[]byte](s S) byte {
	var r byte
	for _, b := range s {
		r ^= b
	}
	return r
}
```

文本停在 **Go statements (Go 协程并发语句)** 的标题处。

准备好后，随时发送 `Go statements` 及后续的原文！
