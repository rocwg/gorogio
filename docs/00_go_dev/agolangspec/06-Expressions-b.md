我们继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【语法规则与代码深解析】** 的标准，逐字逐句剖析 **Function literals（函数字面量 / 匿名函数与闭包）** 这一节。

## 章节：函数字面量 (Function literals)

### 段落 1（定义与 EBNF 语法规约）

> **【英文原文】**
>
> A function literal represents an anonymous function. Function literals cannot declare type parameters.
>
> EBNF
>
> ```
> FunctionLit = "func" Signature FunctionBody .
> ```
>
> ```go
>func(a, b int, z float64) bool { return a*b < int(z) }
> ```

**【逐字精准翻译】**

函数字面量代表一个匿名函数（anonymous function）。函数字面量不能声明类型参数（type parameters）。

EBNF

```
函数字面量 = "func" 函数签名 函数体 .
```

```go
func(a, b int, z float64) bool { return a*b < int(z) }
```

- **词汇与语法剖析：**
  - `anonymous function`：匿名函数（没有函数名的函数）。
  - `type parameters`：类型参数（即泛型参数，如 `[T any]`）。
  - **规范重要限制：** 顶级（Top-level）函数声明可以带泛型类型参数（例如 `func Map[T any](s []T) ...`），但**函数字面量（匿名函数）本身不能直接声明泛型类型参数**。
  - `Signature`：函数签名（包含参数列表和返回值列表）。
  - `FunctionBody`：函数体（由大括号包围的语句块）。

### 段落 2（赋值与直接调用）

> **【英文原文】**
>
> A function literal can be assigned to a variable or invoked directly.
>
> ```go
>f := func(x, y int) int { return x + y }
> func(ch chan int) { ch <- ACK }(replyChan)
> ```

**【逐字精准翻译】**

函数字面量可以被赋值给一个变量，或者被直接调用。

```go
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

- **词汇与语法剖析：**
  - `assigned to a variable`：赋值给一个变量（此时变量 `f` 的类型为函数类型 `func(int, int) int`）。
  - `invoked directly`：直接调用（即立即执行函数表达式 IIFE：在函数体的大括号末尾紧跟圆括号 `(replyChan)`，将参数传入并立刻执行）。

### 段落 3（闭包机制与变量生命周期）

> **【英文原文】**
>
> Function literals are closures: they may refer to variables declared in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible.

**【逐字精准翻译】**

函数字面量是闭包（closures）：它们可以引用在包围它们的（外层）函数中声明的变量。这些变量随后会在外层函数与该函数字面量之间共享，并且只要它们依然可被访问，它们就会继续存活。

- **词汇与底层原理深解析：**
  - `closures`：闭包（指一个函数以及与其相关的引用环境/变量绑定的整体组合）。
  - `surrounding function`：包围函数 / 外层函数。
  - `shared`：共享（**注意：是引用共享，而不是值拷贝**）。如果外层函数修改了这个变量，闭包内部能看到更新后的值；反之亦然。
  - `survive as long as they are accessible`：存活 / 延长生命周期。
    - **Go 编译器内存逃逸分析（Escape Analysis）：** 在正常情况下，函数的局部变量分配在栈（Stack）上，函数返回后栈帧即被销毁。但如果一个局部变量被闭包引用，并且该闭包在函数返回后依然存在（例如作为返回值返回或传给了 Goroutine），编译器会自动将该变量逃逸到堆（Heap）上分配，从而确保其生命周期得到延续。

这一节精准地定义了 Go 语言中匿名函数与闭包的核心机制。

接下来当你准备好并贴出后续文本时，我们将继续推进到下一个重量级章节：**Primary expressions（基本表达式）**！



我们继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【语法规则与代码深解析】** 的标准，逐字逐句剖析 **Primary expressions（基本表达式）** 这一节。

## 章节：基本表达式 (Primary expressions)

### 段落 1（定义与作用）

> **【英文原文】**
>
> Primary expressions are the operands for unary and binary expressions.
>

**【逐字精准翻译】**

基本表达式（Primary expressions）是一元表达式（unary expressions）和二元表达式（binary expressions）的操作数（operands）。

- **词汇剖析：**
  - `Primary expressions`：基本表达式 / 首要表达式（语言语法树中最基础、优先级最高的表达式形式）。
  - `operands`：操作数（即被运算符作用的对象，如 `a + b` 中的 `a` 和 `b`）。

### 段落 2（EBNF 语法范式）

> **【英文原文】**
>
> EBNF
>
> ```ini
> PrimaryExpr   = Operand |
>                 Conversion |
>                 MethodExpr |
>                 PrimaryExpr Selector |
>                 PrimaryExpr Index |
>                 PrimaryExpr Slice |
>                 PrimaryExpr TypeAssertion |
>                 PrimaryExpr Arguments .
> Selector      = "." identifier .
> Index         = "[" Expression [ "," ] "]" .
> Slice         = "[" [ Expression ] ":" [ Expression ] "]" |
>                 "[" [ Expression ] ":" Expression ":" Expression "]" .
> TypeAssertion = "." "(" Type ")" .
> Arguments     = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
> ```

**【逐字精准翻译与结构拆解】**

- **PrimaryExpr（基本表达式）推导规则：**
  - `Operand`：操作数（如字面量、标识符）
  - `Conversion`：类型转换（如 `int(x)`）
  - `MethodExpr`：方法表达式（如 `Point.Length`）
  - `PrimaryExpr Selector`：选择器（如 `x.f`）
  - `PrimaryExpr Index`：索引访问（如 `a[i]`）
  - `PrimaryExpr Slice`：切片操作（如 `a[i:j]`）
  - `PrimaryExpr TypeAssertion`：类型断言（如 `x.(T)`）
  - `PrimaryExpr Arguments`：函数/方法调用实参（如 `f(x)`）
- **各子组件语法规则说明：**
  - **Selector（选择器）：** 点号加标识符（`.identifier`）。
  - **Index（索引）：** 方括号内跟一个表达式（`[Expression]`），可选末尾逗号。
  - **Slice（切片）：**
    - 简单切片（二索引）：`[low : high]`
    - 完整切片（三索引）：`[low : high : max]`（用于限定容量 `cap`）
  - **TypeAssertion（类型断言）：** 点号加圆括号内的类型（`.(Type)`）。
  - **Arguments（参数列表）：** 圆括号内包含可选的表达式列表或内置函数特殊语法（如 `make(Type, args...)`），支持变长参数 `...` 和结尾可选逗号。

### 语法规则深度解构（EBNF 细节剖析）

1. **递归链式结构 (`PrimaryExpr Selector / Index / Arguments...`)：**
   - 注意语法定义中 `PrimaryExpr` 是可以自我左递归的！这解释了为什么你可以写出像 `f.p[i].x()` 这样长的**链式调用**（从左到右依次进行求值）。
2. **切片表示法的两种形式：**
   - 简单切片（2个界限）：`[i : j]`（省略时默认为 `0` 到 `len`）。
   - 完整切片（3个界限，带容量限制）：`[i : j : k]`（限定容量容量上限容量为 `k - i`）。
3. **Arguments（实参）中的泛型支持：**
   - 注意 `Arguments` 规则中的 `Type [ "," ExpressionList ]`：这是为了支持 Go 1.18+ 的泛型显式实例化调用（例如 `min[int](1, 2)`，在括号或实参传递时解析类型参数）。

### 段落 3（典型代码示例剖析）

> **【英文原文】**
>
> ```go
>x
> 2
> (s + ".txt")
> f(3.1415, true)
> Point{1, 2}
> m["foo"]
> s[i : j + 1]
> obj.color
> f.p[i].x()
> ```

**【逐字精准翻译与结构分类】**

```go
x               // Operand (变量/标识符)
2               // Operand (基本字面量)
(s + ".txt")    // Operand (括号括起来的表达式)
f(3.1415, true) // PrimaryExpr Arguments (函数调用)
Point{1, 2}     // Operand (复合字面量)
m["foo"]        // PrimaryExpr Index (Map/Slice/Array 索引访问)
s[i : j + 1]    // PrimaryExpr Slice (切片表达式)
obj.color       // PrimaryExpr Selector (字段/方法选择器)
f.p[i].x()      // 链式组合：字段选择 -> 数组/切片索引 -> 字段选择 -> 函数/方法调用
```

- **示例剖析：**
  - `f.p[i].x()` 的求值优先级分析：
    1. `f`（基本操作数）
    2. `f.p`（选择器：获取 `f` 的字段 `p`）
    3. `f.p[i]`（索引：获取数组/切片 `p` 的第 `i` 个元素）
    4. `f.p[i].x`（选择器：获取元素上的字段或方法 `x`）
    5. `f.p[i].x()`（实参列表：执行无参函数调用）

这一节定义了 Go 语言中最重要、最基础的**表达式语法骨架**。

如果你准备好了，请贴出后续文本，我们将继续推进下一个细节章节：**Selectors（选择器）**！

按 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【语法规则与代码深解析】** 的标准，为您逐字逐句剖析 **Selectors（选择器）** 这一极为关键的 Go 语法规范。

## 章节：选择器 (Selectors)

### 段落 1（基本定义与概念）

> **【英文原文】**
>
> For a [primary expression](https://go.dev/ref/spec#Primary_expressions) `x` that is not a [package name](https://go.dev/ref/spec#Package_clause), the *selector expression*
>
> `x.f` 
>
> denotes the field or method `f` of the value `x` (or sometimes `*x`; see below). The identifier `f` is called the (field or method) *selector*; it must not be the [blank identifier](https://go.dev/ref/spec#Blank_identifier). The type of the selector expression is the type of `f`. If `x` is a package name, see the section on [qualified identifiers](https://go.dev/ref/spec#Qualified_identifiers).

**【逐字精准翻译】**

对于不是包名（package name）的基本表达式 `x`，选择器表达式

```
x.f
```

表示值 `x`（或有时是 `*x`；见下文）的字段或方法 `f`。标识符 `f` 被称为（字段或方法）选择器；它绝不能是空白标识符（`_`）。选择器表达式的类型就是 `f` 的类型。如果 `x` 是一个包名，请参见关于限定标识符（qualified identifiers）的章节。

- **词汇与语法剖析：**
  - `denotes`：表示 / 代表。
  - `selector`：选择器（即跟随在点 `.` 后面的标识符）。
  - `blank identifier`：空白标识符（下划线 `_`）。例如 `x._` 在 Go 中是**不合法**的。
  - `qualified identifiers`：限定标识符（如 `fmt.Println` 中的 `Println`，语法上属于包名限定，而非这里讨论的结构体/接口选择器）。

### 段落 2（嵌入字段与深度 Depth 定义）

> **【英文原文】**
>
> A selector `f` may denote a field or method `f` of a type `T`, or it may refer to a field or method `f` of a nested embedded field of `T`. The number of embedded fields traversed to reach `f` is called its depth in `T`. The depth of a field or method `f` declared in `T` is zero. The depth of a field or method `f` declared in an embedded field `A` in `T` is the depth of `f` in `A` plus one.

**【逐字精准翻译】**

选择器 `f` 可以表示类型 `T` 本身的字段或方法 `f`，也可以引用 `T` 的嵌套嵌入字段（nested embedded field，即匿名字段）中的字段或方法 `f`。到达 `f` 所遍历的嵌入字段的数量称为其在 `T` 中的深度（depth）。在 `T` 中直接声明的字段或方法 `f` 的深度为零。在 `T` 的嵌入字段 `A` 中声明的字段或方法 `f` 的深度，等于 `f` 在 `A` 中的深度加上一。

- **核心概念拆解：**
  - `nested embedded field`：嵌套的嵌入字段（匿名字段）。
  - `depth`：深度（用于后续确定字段/方法选择优先级的关键指标）。

- **语法规则与深度计算：**
  - `T.f`（直接声明）：`depth = 0`
  - `T.A.f`（嵌入一级）：`depth = 1`
  - `T.A.B.f`（嵌入两级）：`depth = 2`

### 段落 3（选择器查找三大规则）

> **【英文原文】**
>
> The following rules apply to selectors:
>
> 1. For a value `x` of type `T` or `*T` where `T` is not a pointer or interface type, `x.f` denotes the field or method at the shallowest depth in `T` where there is such an `f`. If there is not exactly one `f` with shallowest depth, the selector expression is illegal.
> 2. For a value `x` of type `I` where `I` is an interface type, `x.f` denotes the actual method with name `f` of the dynamic value of `x`. If there is no method with name `f` in the method set of `I`, the selector expression is illegal.
> 3. As an exception, if the type of `x` is a defined pointer type and `(*x).f` is a valid selector expression denoting a field (but not a method), `x.f` is shorthand for `(*x).f`.
> 4. In all other cases, `x.f` is illegal.

**【逐字精准翻译】**

以下规则适用于选择器：

1. 对于类型为 `T` 或 `*T` 的值 `x`（其中 `T` 既不是指针类型也不是接口类型），`x.f` 表示 `T` 中拥有名为 `f` 的**最浅深度（shallowest depth）\**处的字段或方法。如果在最浅深度处不存在\**恰好唯一**的一个 `f`，则该选择器表达式是不合法的（即引发歧义遮蔽/冲突编译报错）。
2. 对于类型为 `I` 的值 `x`（其中 `I` 是接口类型），`x.f` 表示 `x` 的动态值（dynamic value）中名为 `f` 的实际方法。如果 `I` 的方法集中不存在名为 `f` 的方法，则该选择器表达式是不合法的。
3. 作为例外情况，如果 `x` 的类型是一个已定义的指针类型（defined pointer type），且 `(*x).f` 是一个表示**字段**（但**不是方法**）的有效选择器表达式，则 `x.f` 是 `(*x).f` 的简写形式。
4. 在所有其他情况下，`x.f` 都是不合法的。

- **底层深解析：**
  - **规则 1（遮蔽与冲突规则 Shadowing & Ambiguity）：**
    - **遮蔽（Shadowing）：** 较浅深度的 `f` 会“遮蔽”较深深度的 `f`。
    - **歧义/不合法（Ambiguity）：** 如果同一最浅深度出现了两个同名的 `f`（例如 `T` 包含了两个嵌入结构体 `A` 和 `B`，它们都有字段 `f`，且深度相同），编译器将直接报 `ambiguous selector` 错误。
  - **规则 3（Defined Pointer 限制）：** 比如通过 `type Q *T2` 定义的自定义指针类型 `Q`，对于字段解引用，允许 `q.x` 自动解引用为 `(*q).x`；但**不允许**直接调用方法 `q.M0()`。

### 段落 4（空指针与空接口 Panic 行为）

> **【英文原文】**
>
> If `x` is of pointer type and has the value `nil` and `x.f` denotes a struct field, assigning to or evaluating `x.f` causes a run-time panic.
>
> If `x` is of interface type and has the value `nil`, calling or evaluating the method `x.f` causes a run-time panic.

**【逐字精准翻译】**

如果 `x` 是指针类型且值为 `nil`，同时 `x.f` 表示一个结构体字段，则对 `x.f` 进行赋值或求值会导致运行时崩溃（run-time panic）。

如果 `x` 是接口类型且值为 `nil`，则调用或求值方法 `x.f` 会导致运行时崩溃（run-time panic）。

### 代码示例与详细推倒

> **【英文原文】**
>
> For example, given the declarations:
>
> ```go
> type T0 struct {
> 	x int
> }
> 
> func (*T0) M0()
> 
> type T1 struct {
> 	y int
> }
> 
> func (T1) M1()
> 
> type T2 struct {
> 	z int
> 	T1   // 嵌入值类型
> 	*T0  // 嵌入指针类型
> }
> 
> func (*T2) M2()
> 
> type Q *T2 // Q 是一个定义的指针类型 (defined pointer type)
> 
> var t T2     // with t.T0 != nil
> var p *T2    // with p != nil and (*p).T0 != nil
> var q Q = p
> ```

**【类型结构与嵌套层次推导】**

- `T2` 结构体：
  - 深度 0：`z`（字段）, `T1`（嵌入字段）, `*T0`（嵌入指针字段）, `M2`（方法，接收者 `*T2`）
  - 深度 1（通过 `T1`）：`y`（字段）, `M1`（方法，接收者 `T1`）
  - 深度 1（通过 `*T0`）：`x`（字段）, `M0`（方法，接收者 `*T0`）
- `Q` 是定义的指针类型 `type Q *T2`。

#### 合法写法（Valid Selector Expressions）

> **【英文原文】**
>
> one may write:
>
> ```go
> t.z          // t.z
> t.y          // t.T1.y
> t.x          // (*t.T0).x
> 
> p.z          // (*p).z
> p.y          // (*p).T1.y
> p.x          // (*(*p).T0).x
> 
> q.x          // (*(*q).T0).x        (*q).x is a valid field selector
> 
> p.M0()       // ((*p).T0).M0()      M0 expects *T0 receiver
> p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
> p.M2()       // p.M2()              M2 expects *T2 receiver
> t.M2()       // (&t).M2()           M2 expects *T2 receiver, see section on Calls
> ```

**【逐字精准翻译与核心推导注释】**

可以这样写：

```go
t.z          // 直接访问 T2 的字段 z (深度 0)
t.y          // 提升字段：自动展开为 t.T1.y (深度 1)
t.x          // 提升字段：由于 T0 是指针类型 *T0，自动解引用展开为 (*t.T0).x (深度 1)

p.z          // 指针自动解引用：(*p).z
p.y          // 指针与嵌入字段自动展开：(*p).T1.y
p.x          // 双重自动解引用：(*(*p).T0).x

q.x          // 对应规则 3：Q 是自定义指针类型，访问【字段】时支持解引用展开为 (*(*q).T0).x

p.M0()       // 自动解引用并提取嵌入指针：((*p).T0).M0()，M0 期望 *T0 接收者
p.M1()       // 自动解引用并提取嵌入值：((*p).T1).M1()，M1 期望 T1 接收者
p.M2()       // 直接调用：p.M2()，M2 期望 *T2 接收者
t.M2()       // 自动取地址：(&t).M2()，M2 期望 *T2 接收者（详见“调用”章节）
```

#### 不合法写法（Invalid Selector Expressions）

> **【英文原文】**
>
> but the following is invalid:
>
> ```go
>q.M0()       // (*q).M0 is valid but not a field selector
> ```

**【逐字精准翻译与解析】**

但以下写法是不合法的：

```go
q.M0()       // (*q).M0 是合法的，但它不是字段选择器（而是方法），因此规则 3 不适用！
```

- **为什么 `q.x` 合法，而 `q.M0()` 非法？（重点难点！）**
  - `q` 的类型是 `Q`（即 `type Q *T2`）。
  - 规则 3 规定：**定义的指针类型**只有在 `(*q).f` 表示字段（field）时，才允许 `q.f` 作为 `(*q).f` 的简写。
  - 由于 `M0` 是一个**方法（method）**而不是字段，因此规则 3 显式**拒绝**自动转换，所以 `q.M0()` 会引发编译错误！必须显式写成 `(*q).M0()`。

选择器的“自动解引用”、“自动取地址”以及“定义的指针类型对方法的隔离”是 Go 编译器非常经典且严谨的语法边界。

准备好后请贴出后续，我们将进入下一个精彩章节：**Method expressions（方法表达式）**！



按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，为您精细解读 **Method expressions（方法表达式）** 章节（这是 Go 语言中将“方法”转化为“普通显式接收者函数”的核心机制）。

## 章节：方法表达式 (Method expressions)

### 段落 1（基本定义与 EBNF 语法规约）

> **【英文原文】**
>
> If `M` is in the method set of type `T`, `T.M` is a function that is callable as a regular function with the same arguments as `M` prefixed by an additional argument that is the receiver of the method.
>
> EBNF
>
> ```
> MethodExpr   = ReceiverType "." MethodName .
> ReceiverType = Type .
> ```

**【逐字精准翻译】**

如果 `M` 存在于类型 `T` 的方法集中，则 `T.M` 是一个可以作为普通函数调用的函数，其参数与 `M` 相同，但在前面附加了一个额外的参数，该参数即为该方法的接收者（receiver）。

EBNF

```
方法表达式 = 接收者类型 "." 方法名 .
接收者类型 = 类型 .
```

- **词汇与语法剖析：**
  - `method set`：方法集（类型拥有的所有方法的集合）。
  - `regular function`：普通函数（没有接收者绑定的普通 `func(...)`）。
  - `prefixed by an additional argument`：在最前面附加一个额外参数（接收者被显式转换为了函数的第一个入参）。

### 段落 2（基础结构体与方法声明）

> **【英文原文】**
>
> Consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.
>
> ```go
>type T struct {
> 	a int
> }
> func (tv  T) Mv(a int) int         { return 0 }  // value receiver
> func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver
> 
> var t T
> ```

**【逐字精准翻译】**

考虑一个结构体类型 `T`，它包含两个方法：`Mv`（其接收者为类型 `T`）和 `Mp`（其接收者为类型 `*T`）。

```go
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // 值接收者
func (tp *T) Mp(f float32) float32 { return 1 }  // 指针接收者

var t T
```

### 段落 3（值接收者方法表达式与等价调用）

> **【英文原文】**
>
> The expression
>
> `T.Mv` 
>
> yields a function equivalent to `Mv` but with an explicit receiver as its first argument; it has signature
>
> `func(tv T, a int) int` 
>
> That function may be called normally with an explicit receiver, so these five invocations are equivalent:
>
> ```go
> t.Mv(7)
> T.Mv(t, 7)
> (T).Mv(t, 7)
> f1 := T.Mv; f1(t, 7)
> f2 := (T).Mv; f2(t, 7)
> ```

**【逐字精准翻译】**

表达式 `T.Mv` 产生一个等同于 `Mv` 的函数，但将其显式接收者作为其第一个参数；它的函数签名（signature）为：

```
func(tv T, a int) int
```

该函数可以带上显式接收者进行正常调用，因此以下五种调用方式是等价的：

```go
t.Mv(7)              // 1. 标准方法调用
T.Mv(t, 7)           // 2. 直接使用类型进行方法表达式调用
(T).Mv(t, 7)         // 3. 带有括号的类型方法表达式调用
f1 := T.Mv; f1(t, 7) // 4. 赋值给函数变量 f1 后调用
f2 := (T).Mv; f2(t, 7) // 5. 带有括号赋值给函数变量 f2 后调用
```

### 段落 4（指针接收者方法表达式 `(*T).Mp` 与值接收者的解引用提升 `(*T).Mv`）

> **【英文原文】**
>
> Similarly, the expression
>
> `(*T).Mp` 
>
> yields a function value representing `Mp` with signature
>
> `func(tp *T, f float32) float32` 
>
> For a method with a value receiver, one can derive a function with an explicit pointer receiver, so
>
> `(*T).Mv` 
>
> yields a function value representing `Mv` with signature
>
> `func(tv *T, a int) int` 
>
> Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call.

**【逐字精准翻译】**

类似地，表达式 `(*T).Mp` 产生一个代表 `Mp` 的函数值，其签名（signature）为：

```
func(tp *T, f float32) float32
```

对于具有值接收者的方法，可以派生出一个带有显式指针接收者的函数，因此：

```
(*T).Mv
```

产生一个代表 `Mv` 的函数值，其签名（signature）为：

```
func(tv *T, a int) int
```

这样一个函数体会通过接收者进行间接寻址（解引用），以创建一个值并将其作为接收者传递给底层的函数；该方法**不会**覆盖函数调用中传入地址对应的那个值。

- **语法规则深解析：**
  - `(*T).Mp` 的签名是 `func(tp *T, f float32) float32`（标准指针匹配）。
  - `(*T).Mv` 的签名是 `func(tv *T, a int) int`：**注意！** 虽然 `Mv` 原本是值接收者，但因为类型 `*T` 的方法集中包含了 `Mv`（值方法会自动提升到指针方法集中），所以通过 `(*T).Mv` 衍生出的函数第一个参数也是指针 `*T`。
  - **值拷贝保护机制：** 当调用 `(*T).Mv(&t, 7)` 时，生成的包装代码会在内部先自动解引用 `*tp` 产生一份值拷贝，再把拷贝传给 `Mv` 执行。因此，即使传入的是指针，`Mv` 内部也无法修改 `t` 原有的内容。

### 段落 5（不合法规则：值类型调用指针接收者）

> **【英文原文】**
>
> The final case, a value-receiver function for a pointer-receiver method, is illegal because pointer-receiver methods are not in the method set of the value type.

**【逐字精准翻译】**

最后一种情况，即针对指针接收者方法的“值接收者函数”（即 `T.Mp`），是**不合法的**，因为指针接收者方法**不在**值类型的方法集中。

- **语法规则深解析：**

  - **为什么 `T.Mp` 编译报错？**

    因为 `Mp` 的接收者是 `*T`（指针），而类型 `T` 的方法集中**不包含**指针接收者方法。你不能写 `T.Mp`（编译器无法在没有变量地址的情况下自动凭空取地址）。

### 段落 6（调用语法与闭包/方法值的区分）

> **【英文原文】**
>
> Function values derived from methods are called with function call syntax; the receiver is provided as the first argument to the call. That is, given `f := T.Mv`, `f` is invoked as `f(t, 7)` not `t.f(7)`. To construct a function that binds the receiver, use a function literal or method value.
>
> It is legal to derive a function value from a method of an interface type. The resulting function takes an explicit receiver of that interface type.

**【逐字精准翻译】**

从方法派生出的函数值必须使用普通函数调用语法进行调用；接收者被作为调用的第一个参数提供。也就是说，给定 `f := T.Mv`，`f` 的调用方式是 `f(t, 7)` 而不是 `t.f(7)`。要构造一个绑定了接收者的函数，请使用函数字面量或方法值（Method values）。

从接口类型（interface type）的方法中派生函数值是合法的。派生出的函数接受该接口类型的显式接收者。

- **语法规则深解析：**

  - `Method Expression`（方法表达式，如 `T.Mv`）：**不绑定**具体实例，派生出的函数第一个参数是接收者本身。

  - `Method Value`（方法值，如 `t.Mv`）：**绑定**了具体实例 `t`，派生出的函数不需要传 `t`（下一节会详细展开）。

  - **接口方法表达式（Interface Method Expression）：**

    例如对于 `io.Reader` 接口：

    ```go
    var r io.Reader
    readFn := io.Reader.Read // readFn 的签名是 func(io.Reader, []byte) (int, error)
    readFn(r, buf)
    ```

这一节精准地阐明了 Go 语言中“方法”如何转化为“普通函数”的机制以及方法集的边界规则。

准备好后请贴出后续，我们将进入下一个精彩章节：**Method values（方法值）**！



按 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【语法规则与代码深解析】** 的标准，为您逐字逐句剖析 **Method values（方法值）** 这一 Go 规范中极其核心的闭包与方法绑定机制。

## 章节：方法值 (Method values)

### 段落 1（定义与接收者计算/存储时机）

> **【英文原文】**
>
> If the expression `x` has static type `T` and `M` is in the method set of type `T`, `x.M` is called a method value. The method value `x.M` is a function value that is callable with the same arguments as a method call of `x.M`. The expression `x` is evaluated and saved during the evaluation of the method value; the saved copy is then used as the receiver in any calls, which may be executed later.

**【逐字精准翻译】**

如果表达式 `x` 具有静态类型 `T`，且 `M` 存在于类型 `T` 的方法集中，则 `x.M` 被称为方法值（method value）。方法值 `x.M` 是一个函数值（function value），它可以传入与 `x.M` 方法调用相同的参数来执行调用。在对其求值（求出该方法值）的过程中，表达式 `x` 会被提前求值（evaluated）并保存起来；该保存的副本随后将在任何稍后执行的调用中用作接收者（receiver）。

- **词汇与核心原理：**
  - `static type`：静态类型。
  - `evaluated and saved`：在表达式求值时立即被计算并保存（值捕获/快照）。这意味着后续修改 `x` 的值不会影响已保存的方法值副本。
- **语法规则与底层机制深解析：**
  - **静态类型 `T` 与方法集：** 与上一节“方法表达式”不同，方法值是绑定在**具体变量/表达式 `x`** 上的。
  - **求值与拷贝时机（极度重要）：** 在执行 `f := x.M` 的那一刻，`x` 的值就已经被计算并保存了。
    - 如果 `M` 是**值接收者**（`func (t T) M()`），被保存的就是 `x` 的**一份深拷贝（副本）**。后续对原变量 `x` 的修改，完全不会影响 `f()` 的执行结果！
    - 如果 `M` 是**指针接收者**（`func (t *T) M()`），被保存的是**指针地址本身**。后续修改指针指向的内容，`f()` 可以看到最新的修改。

### 代码示例 1（求值与复制行为验证）

> **【英文原文】**
>
> ```go
> type S struct { *T }
> type T int
> func (t T) M() { print(t) }
> 
> t := new(T)
> s := S{T: t}
> f := t.M                    // receiver *t is evaluated and stored in f
> g := s.M                    // receiver *(s.T) is evaluated and stored in g
> *t = 42                     // does not affect stored receivers in f and g
> ```

**【逐字精准翻译】**

```go
type S struct { *T }
type T int
func (t T) M() { print(t) }

t := new(T)
s := S{T: t}
f := t.M                    // 接收者 *t 在此时被求值并复制保存到了 f 中
g := s.M                    // 接收者 *(s.T) 在此时被求值并复制保存到了 g 中
*t = 42                     // 修改 *t 不会影响已经保存在 f 和 g 中的接收者副本
```

- **代码解析：**
  1. `t` 是 `*T` 指针，`M()` 是 `T` 的**值接收者**方法。
  2. 赋值 `f := t.M` 时，编译器自动解引用 `*t`（此时 `*t == 0`），并将 `0` 作为**值**保存在了 `f` 的闭包上下文中。
  3. 同理，`g := s.M` 保存的也是当时的 `*(s.T)` 即 `0` 的副本。
  4. 修改 `*t = 42` 后，执行 `f()` 或 `g()` 打印出的依然是 `0`，而不是 `42`。

### 段落 2（非接口类型与各种自动转换规则）

> **【英文原文】**
>
> The type `T` may be an interface or non-interface type.
>
> As in the discussion of method expressions above, consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.
>
> ```go
> type T struct {
> 	a int
> }
> func (tv  T) Mv(a int) int         { return 0 }  // value receiver
> func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver
> 
> var t T
> var pt *T
> func makeT() T
> ```

**【逐字精准翻译】**

类型 `T` 可以是接口类型，也可以是非接口类型。

如同上面关于方法表达式的讨论一样，考虑一个结构体类型 `T`，它有两个方法：`Mv`（其接收者类型为 `T`）和 `Mp`（其接收者类型为 `*T`）。

```go
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // 值接收者
func (tp *T) Mp(f float32) float32 { return 1 }  // 指针接收者

var t T
var pt *T
func makeT() T
```

### 段落 3（语法糖：自动解引用与自动取地址）

> **【英文原文】**
>
> The expression
>
> `t.Mv` 
> 
> yields a function value of type
>
> `func(int) int` 
>
> These two invocations are equivalent:
> 
> ```go
>t.Mv(7)
> f := t.Mv; f(7)
>```
> 
> Similarly, the expression
> 
> `pt.Mp` 
>
> yields a function value of type
>
> `func(float32) float32` 
>
> As with selectors, a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`. 
>
> As with method calls, a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`.

**【逐字精准翻译】**

表达式 `t.Mv` 产生一个类型为 `func(int) int` 的函数值。以下两种调用是等价的：

```go
t.Mv(7)
f := t.Mv; f(7)
```

类似地，表达式 `pt.Mp` 产生一个类型为 `func(float32) float32` 的函数值。

正如选择器（selectors）的情况一样，使用指针来引用带有值接收者的非接口方法会自动解引用该指针：`pt.Mv` 等价于 `(*pt).Mv`。

正如方法调用的情况一样，使用可寻址的值（addressable value）来引用带有指针接收者的非接口方法会自动对该值进行取地址：`t.Mp` 等价于 `(&t).Mp`。

### 代码示例 2（各种合法与非法的方法值生成）

> **【英文原文】**
>
> ```go
>f := t.Mv; f(7)   // like t.Mv(7)
> f := pt.Mp; f(7)  // like pt.Mp(7)
> f := pt.Mv; f(7)  // like (*pt).Mv(7)
> f := t.Mp; f(7)   // like (&t).Mp(7)
> f := makeT().Mp   // invalid: result of makeT() is not addressable
> ```

**【逐字精准翻译与语法深解析】**

```go
f := t.Mv; f(7)   // 等同于 t.Mv(7)
f := pt.Mp; f(7)  // 等同于 pt.Mp(7)
f := pt.Mv; f(7)  // 等同于 (*pt).Mv(7)，指针自动解引用
f := t.Mp; f(7)   // 等同于 (&t).Mp(7)，变量 t 可寻址，自动取地址
f := makeT().Mp   // 不合法：makeT() 的返回值是临时不可寻址的（not addressable）
```

- **核心考点剖析：`makeT().Mp` 为什么非法？**
  - `Mp` 方法需要一个指针接收者 `*T`。
  - `makeT()` 返回的是一个**值（临时变量/Rvalue）**，在 Go 的规范中，函数调用的直接返回值是**不可寻址的（unaddressable）**。
  - 编译器无法对不可寻址的值做 `&` 操作，因此无法生成 `&makeT()` 来匹配 `Mp` 的指针接收者，所以编译直接报错。

### 段落 4（接口类型的方法值）

> **【英文原文】**
>
> Although the examples above use non-interface types, it is also legal to create a method value from a value of interface type.
>
> ```go
>var i interface { M(int) } = myVal
> f := i.M; f(7)  // like i.M(7)
> ```

**【逐字精准翻译】**

尽管上面的示例使用的是非接口类型，但从接口类型的值创建方法值也是合法的。

```go
var i interface { M(int) } = myVal
f := i.M; f(7)  // 等同于 i.M(7)
```

- **接口方法值的底层原理：**
  - 当执行 `f := i.M` 时，闭包会同时捕获接口 `i` 内部的**动态类型（itab）\**和\**动态值（data pointer）**。
  - 即使后续接口变量 `i` 被重新赋予了新的对象，之前提取出的 `f` 依然绑定的是提取当时的那个动态对象。如果当时 `i` 的值为 `nil`，提取 `i.M` 就会在运行时直接触发 **panic**。

对比上一节的“方法表达式（Method Expressions）”，**方法值（Method Values）** 最大的特点就是：**绑定了实例，省去了第一个接收者参数**。

准备好后请贴出后续，我们将进入下一个重量级表达式章节：**Index expressions（索引表达式）**！

