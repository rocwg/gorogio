按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的三段式架构，为您精细拆解与剖析 Go 语言规范中的 **Properties of types and values（类型与值的属性）** 核心章节之 **Representation of values（值的表示）**：

## Properties of types and values (类型与值的属性)

### Representation of values (值的表示)

#### 段落 1 (自包含/值类型值与非 nil 零值)

> **【英文原文】**
>
> Values of predeclared types (see below for the interfaces `any` and `error`), arrays, and structs are self-contained: Each such value contains a complete copy of all its data, and variables of such types store the entire value. For instance, an array variable provides the storage (the variables) for all elements of the array. The respective zero values are specific to the value's types; they are never `nil`.

**【精准逐字翻译】**

预声明类型（接口 `any` 和 `error` 的情况见下文）、数组（array）和结构体（struct）的值是**自包含的（self-contained）**：每一个这样的值都包含其所有数据的完整副本，且此类类型的变量存储的是整个值。例如，一个数组变量为其所有的元素提供实际存储空间（即变量集合）。它们各自的零值是特定于该值类型的；**它们的零值绝不会是 `nil`**。

- **专业术语与句式拆解：**
  - `self-contained`：自包含/完全独立（即通常所说的“值类型”）。在赋值或按值传递时，Go 会在内存中执行完整的字节拷贝（value copy）。
  - `never nil`：基础类型（`int`, `bool`, `string`等）、数组和结构体变量从声明起即分配有确定空间，其零值由各字段/元素的零值填充，在语法和语义层面都不能与 `nil` 比较或赋予 `nil`。

#### 段落 2 (包含引用的引用类型及其内部结构)

> **【英文原文】**
>
> Non-nil pointer, function, slice, map, and channel values contain references to underlying data which may be shared by multiple values:
>
> - A pointer value is a reference to the variable holding the pointer base type value.
> - A function value contains references to the (possibly anonymous) function and enclosed variables.
> - A slice value contains the slice length, capacity, and a reference to its underlying array.
> - A map or channel value is a reference to the implementation-specific data structure of the map or channel.

**【精准逐字翻译】**

非 `nil` 的指针（pointer）、函数（function）、切片（slice）、映射（map）和通道（channel）的值都包含对底层数据（underlying data）的引用，这些底层数据可以被多个值共享：

- **指针值**是对持有指针基类型（base type）值的变量的引用。
- **函数值**包含对（可能是匿名的）函数以及捕获的上下文变量（enclosed variables / 闭包）的引用。
- **切片值**包含切片长度（length）、容量（capacity）以及对其底层数组（underlying array）的引用。
- **Map 或 Channel 值**是对特定实现的 Map 或 Channel **内部数据结构**的引用。
- **专业细节拆解：**
  - `enclosed variables`：指闭包捕获的外部局部变量。在 Go 底层，函数值本质上是一个指向闭包对象（`runtime.funcval`）的指针。
  - `implementation-specific data structure`：映射和通道在 Go 编译器/运行时层面上分别对应着指向 `runtime.hmap` 和 `runtime.hchan` 的指针，因此它们在变量传递时只需传递指针本身。

#### 段落 3 (接口值的两种表示形式与 nil 零值)

> **【英文原文】**
>
> An interface value may be self-contained or contain references to underlying data depending on the interface's dynamic type. The predeclared identifier `nil` is the zero value for types whose values can contain references.

**【精准逐字翻译】**

接口值（interface value）根据其动态类型（dynamic type）的不同，可能是自包含的，也可能包含对底层数据的引用。预声明的标识符 `nil` 是那些值可以包含引用的类型的零值。

- **专业细节拆解：**
  - `interface value`：接口在底层由双指针结构表示（`eface` 或 `iface`，包含类型信息指针与数据指针 `data`）。如果存入接口的是一个基础类型值（如 `int`），接口的 `data` 指针会指向该值的副本；如果存入的是指针或切片，则接口持有该引用。
  - 接口判定为 `nil` 的条件：只有当接口的类型（type）**和**动态值（value）同时为 `nil` 时，接口变量才等于 `nil`（即 `var err error = nil`）。如果将一个值为 `nil` 的具体指针赋给接口，接口本身就不再等于 `nil`。

#### 段落 4 (共享底层数据的副作用/写屏障影响)

> **【英文原文】**
>
> When multiple values share underlying data, changing one value may change another. For instance, changing an element of a slice will change that element in the underlying array for all slices that share the array.

**【精准逐字翻译】**

当多个值共享底层数据时，修改其中一个值可能会改变另一个值。例如，修改切片中的某个元素，将会改变所有共享该底层数组的切片在底层数组中对应的元素。

- **专业细节拆解：**
  - `share the array`：切片进行表达式截取（如 `s2 := s1[1:3]`）时只复制 Header（`array_ptr`, `len`, `cap`），底层的数组仍然是同一块内存连续区域，因此对 `s2[0]` 的写入会导致 `s1[1]` 可见。

按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的三段式架构，为您精细拆解与剖析 Go 语言规范中的 **Underlying types（底层类型）** 章节：

## Underlying types (底层类型)

### 段落 1 (底层类型的核心递归定义)

> **【英文原文】**
>
> Each type `T` has an underlying type: If `T` is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is `T` itself. Otherwise, `T`'s underlying type is the underlying type of the type to which `T` refers in its declaration. For a type parameter that is the underlying type of its type constraint, which is always an interface.

**【精准逐字翻译】**

每一个类型 `T` 都有一个**底层类型（underlying type）**：如果 `T` 是预声明的布尔、数值或字符串类型之一，或者是一个**类型字面量（type literal）**，则对应的底层类型就是 `T` 本身。否则，`T` 的底层类型就是 `T` 在其声明中所引用的类型的底层类型。对于**类型参数（type parameter，即泛型形参）**，其底层类型就是它的类型约束（type constraint）的底层类型，而该类型约束始终是一个接口。

- **专业术语与句式拆解：**
  - `type literal`：类型字面量（如 `[]int`、`map[string]bool`、`struct{ X int }`、`func()` 等复合类型的字面量值声明）。它们天然没有类型名称，其底层类型就是字面量本身。
  - **递归追溯规则**：若 `type B A`，要找 `B` 的底层类型，就看 `A` 的底层类型；若 `A` 也是具名自定义类型，就继续追溯直到遇到预声明类型或类型字面量为止。
  - `type parameter's underlying type`：泛型类型参数 `P` 的底层类型是由其 Constraint 接口决定的。

### 代码示例与段落 2 (类型别名与自定义类型的底层类型推导)

> **【英文原文】**
>
> ```Go
> type (
> 	A1 = string
> 	A2 = A1
> )
> 
> type (
> 	B1 string
> 	B2 B1
> 	B3 []B1
> 	B4 B3
> )
> 
> func f[P any](x P) { … }
> ```
>
> The underlying type of `string`, `A1`, `A2`, `B1`, and `B2` is `string`. The underlying type of `[]B1`, `B3`, and `B4` is `[]B1`. The underlying type of `P` is `interface{}`.

**【精准逐字翻译】**

`string`、`A1`、`A2`、`B1` 和 `B2` 的底层类型都是 `string`。

`[]B1`、`B3` 和 `B4` 的底层类型都是 `[]B1`。

`P` 的底层类型是 `interface{}`。

- **专业细节拆解：**
  - **类型别名（Type Alias, `A1 = string`）**：`A1` 和 `A2` 只是 `string` 的别名，它们与 `string` 是**完全相同的类型**（Identical Type），底层类型自然是 `string`。
  - **类型定义（Type Definition, `B1 string`）**：
    - `B1` 是基于 `string` 定义的新类型，底层类型为 `string`。
    - `B2` 基于 `B1` 定义（`type B2 B1`），继续追溯 `B1` 的底层类型，因此 `B2` 的底层类型也是 `string`。
    - `B3` 是基于切片类型字面量 `[]B1` 定义的新类型（`type B3 []B1`），因为 `[]B1` 是类型字面量，所以 `B3` 的底层类型是 `[]B1`（**注意：不是 `[]string`！**）。
    - `B4` 基于 `B3` 定义（`type B4 B3`），追溯 `B3` 的底层类型，因此 `B4` 的底层类型也是 `[]B1`。
  - **泛型形参（Type Parameter, `P any`）**：`P` 的约束是 `any`（即 `interface{}`），接口的底层类型是其接口字面量本身，故 `P` 的底层类型为 `interface{}`。

按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的三段式架构，为您精细拆解与剖析 Go 语言规范中的 **Type identity（类型同一性）** 章节：

## Type identity (类型同一性)

### 段落 1 (类型同一性的总则)

> **【英文原文】**
>
> Two types are either identical ("the same") or different.
>
> A named type is always different from any other type. Otherwise, two types are identical if their underlying type literals are structurally equivalent; that is, they have the same literal structure and corresponding components have identical types. In detail:

**【精准逐字翻译】**

两个类型要么是**相同的（identical，“同一个类型”）**，要么是**不同的（different）**。

具名类型（named type / 定义的类型）总是与任何其他类型不同。否则，如果两个类型的底层类型字面量（type literals）在结构上是等价的，则这两个类型相同；也就是说，它们具有相同的字面量结构，且对应的组成部分具有相同的类型。具体如下：

- **专业术语与句式拆解：**
  - `named type`：具名类型（通过 `type T ...` 定义的非别名类型）。即使两个具名类型的底层结构完全一致（如 `type T1 int` 和 `type T2 int`），它们依然是两个完全不同的类型。
  - `structurally equivalent`：结构等价。对于无名类型字面量（如 `[]int`、`map[string]int`），只要结构与内部组件类型完全相同，它们就是相同的类型。

### 细则条款 (复合类型与泛型实例化的同一性规则)

> **【英文原文】**
>
> - Two array types are identical if they have identical element types and the same array length.
> - Two slice types are identical if they have identical element types.
> - Two struct types are identical if they have the same sequence of fields, and if corresponding pairs of fields have the same names, identical types, and identical tags, and are either both embedded or both not embedded. Non-exported field names from different packages are always different.
> - Two pointer types are identical if they have identical base types.
> - Two function types are identical if they have the same number of parameters and result values, corresponding parameter and result types are identical, and either both functions are variadic or neither is. Parameter and result names are not required to match.
> - Two interface types are identical if they define the same type set.
> - Two map types are identical if they have identical key and element types.
> - Two channel types are identical if they have identical element types and the same direction.
> - Two instantiated types are identical if their defined types and all type arguments are identical.

**【精准逐字翻译】**

- **数组（Array）**：如果两个数组类型具有相同的元素类型和相同的数组长度，则它们相同。
- **切片（Slice）**：如果两个切片类型具有相同的元素类型，则它们相同。
- **结构体（Struct）**：如果两个结构体类型具有相同的字段顺序，且对应字段对具有相同的字段名、相同的类型、相同的 Tag（标签），并且同为嵌套字段或同为非嵌套字段，则它们相同。**来自于不同包的未导出字段名（小写开头）总是不同的**。
- **指针（Pointer）**：如果两个指针类型具有相同的基类型（base type），则它们相同。
- **函数（Function）**：如果两个函数类型具有相同数量的参数和返回值，对应的参数类型和返回值类型完全相同，且同为可变参数函数（variadic）或同为非可变参数函数，则它们相同。**参数名和返回值名不要求匹配**。
- **接口（Interface）**：如果两个接口类型**定义了相同的类型集（type set）**，则它们相同。
- **映射（Map）**：如果两个 Map 类型具有相同的键类型和元素类型，则它们相同。
- **通道（Channel）**：如果两个通道类型具有相同的元素类型和相同的方向，则它们相同。
- **实例化类型（Instantiated type）**：如果两个实例化类型的被定义类型（defined type）以及所有的类型实参（type arguments）完全相同，则它们相同。

### 代码声明示例

> **【英文原文】**
>
> ```go
>type (
> 	A0 = []string
> 	A1 = A0
> 	A2 = struct{ a, b int }
> 	A3 = int
> 	A4 = func(A3, float64) *A0
> 	A5 = func(x int, _ float64) *[]string
> 
> 	B0 A0
> 	B1 []string
> 	B2 struct{ a, b int }
> 	B3 struct{ a, c int }
> 	B4 func(int, float64) *B0
> 	B5 func(x int, y float64) *A1
> 
> 	C0 = B0
> 	D0[P1, P2 any] struct{ x P1; y P2 }
> 	E0 = D0[int, string]
> )
> ```

### 相同类型与不同类型的解析说明

> **【英文原文】**
>
> these types are identical:
>
> - `A0`, `A1`, and `[]string`
> - `A2` and `struct{ a, b int }`
> - `A3` and `int`
> - `A4`, `func(int, float64) *[]string`, and `A5`
> - `B0` and `C0`
> - `D0[int, string]` and `E0`
> - `[]int` and `[]int`
> - `struct{ a, b *B5 }` and `struct{ a, b *B5 }`
> - `func(x int, y float64) *[]string`, `func(int, float64) (result *[]string)`, and `A5`

**【精准逐字翻译】**

以下这些类型是相同（identical）的：

- `A0`、`A1` 以及 `[]string` （因为 `A0` 和 `A1` 都是 `[]string` 的**类型别名**）
- `A2` 与 `struct{ a, b int }` （`A2` 是结构体字面量的别名）
- `A3` 与 `int` （`A3` 是 `int` 的别名）
- `A4`、`func(int, float64) *[]string` 以及 `A5` （函数参数名和返回值名不影响类型同一性，且 `A3` 即 `int`，`A0` 即 `[]string`）
- `B0` 与 `C0` （`C0` 是具名类型 `B0` 的别名）
- `D0[int, string]` 与 `E0` （`E0` 是泛型实例化类型 `D0[int, string]` 的别名）
- `[]int` 与 `[]int`
- `struct{ a, b *B5 }` 与 `struct{ a, b *B5 }`
- `func(x int, y float64) *[]string`、`func(int, float64) (result *[]string)` 以及 `A5`

> **【英文原文】**
>
> `B0` and `B1` are different because they are new types created by distinct type definitions; `func(int, float64) *B0` and `func(x int, y float64) *[]string` are different because `B0` is different from `[]string`; and `P1` and `P2` are different because they are different type parameters. `D0[int, string]` and `struct{ x int; y string }` are different because the former is an instantiated defined type while the latter is a type literal (but they are still assignable).

**【精准逐字翻译】**

`B0` 与 `B1` 是**不同**的，因为它们是由不同的类型定义（type definitions）创建的新类型；`func(int, float64) *B0` 与 `func(x int, y float64) *[]string` 是不同的，因为 `B0` 不同于 `[]string`；`P1` 与 `P2` 是不同的，因为它们是不同的类型参数（泛型形参）。`D0[int, string]` 与 `struct{ x int; y string }` 是不同的，因为前者是一个**实例化后的具名定义类型（instantiated defined type）**，而后者是一个类型字面量（type literal）（**但它们之间依然是可赋值的 assignable**）。

- **专业细节拆解：**
  - `defined type vs. type literal`：`D0[int, string]` 虽然实例化后的底层结构就是 `struct{ x int; y string }`，但由于 `D0` 是具名泛型类型（Defined Type），因此它与匿名结构体字面量在类型同一性上判定为“不相同（different）”。不过根据 Go 的赋值规则（Assignability），底层结构相同的具名类型与字面量类型之间可以直接隐式赋值或显式转换。

按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的三段式架构，为您精细拆解与剖析 Go 语言规范中的 **Assignability（可赋值性）** 章节：

## Assignability (可赋值性)

### 段落 1 (常规类型的赋值规则)

> **【英文原文】**
>
> A value `x` of type `V` is assignable to a variable of type `T` ("`x` is assignable to `T`") if one of the following conditions applies:
>
> - `V` and `T` are identical.
> - `V` and `T` have identical underlying types but are not type parameters and at least one of `V` or `T` is not a named type.
> - `V` and `T` are channel types with identical element types, `V` is a bidirectional channel, and at least one of `V` or `T` is not a named type.
> - `T` is an interface type, but not a type parameter, and `x` implements `T`.
> - `x` is a (possibly partially instantiated) generic function, `T` is a function type, and any type arguments not provided explicitly for `x` can be inferred such that (after full instantiation) `x` and `T` have identical underlying types [Go 1.27].
> - `x` is the predeclared identifier `nil` and `T` is a pointer, function, slice, map, channel, or interface type, but not a type parameter.
> - `x` is an untyped constant representable by a value of type `T`.

**【精准逐字翻译】**

类型为 `V` 的值 `x` 在满足以下条件之一时，可赋值给类型为 `T` 的变量（“`x` 可赋值给 `T`”）：

- `V` 和 `T` 是**相同的类型（identical）**。
- `V` 和 `T` 具有**相同的底层类型（identical underlying types）**，但它们不是类型参数（type parameters），且 `V` 或 `T` 中**至少有一个不是具名类型（named type，即至少有一个是类型字面量）**。
- `V` 和 `T` 是具有相同元素类型的通道（channel）类型，`V` 是一个双向通道，且 `V` 或 `T` 中**至少有一个不是具名类型**。
- `T` 是一个接口类型（但不能是类型参数），且 `x` 实现了 `T`。
- `x` 是一个（可能部分实例化了的）泛型函数，`T` 是一个函数类型，且 `x` 未显式提供的任何类型实参（type arguments）可以被自动推导，使得在（完全实例化后）`x` 和 `T` 具有相同的底层类型 [Go 1.27]。
- `x` 是预声明标识符 `nil`，且 `T` 是指针、函数、切片、Map、通道或接口类型之一，但不能是类型参数。
- `x` 是一个**无类型常量（untyped constant）**，且能被类型 `T` 的值所表示。
- **专业术语与句式拆解：**
  - `at least one of V or T is not a named type`：这是 Go 语言中**隐式类型转换/赋值**的关键安全阀。例如 `type Age int`（具名类型）与 `int`（预声明/具名类型）：两者都是具名类型，因此 `var a Age = 1; var b int = a` 是**非法**的（必须显式转换 `int(a)`）；但如果是 `var s MyStruct = struct{ X int }{1}`，因为右侧是**无名结构体字面量（type literal，非具名类型）**，只要底层类型相同就可以直接隐式赋值！
  - `bidirectional channel`：双向通道（`chan T`）可以隐式赋值给同元素类型的单向通道（`chan<- T` 或 `<-chan T`），前提是符合非具名约束。
  - `Go 1.27 泛型函数赋值新特性`：允许泛型函数通过类型推导，直接赋值给对应的具体函数类型变量。

### 段落 2 (类型参数/泛型形参的赋值规则)

> **【英文原文】**
>
> Additionally, if `x`'s type `V` or `T` are type parameters, `x` is assignable to a variable of type `T` if one of the following conditions applies:
>
> - `x` is the predeclared identifier `nil`, `T` is a type parameter, and `x` is assignable to each type in `T`'s type set.
> - `V` is not a named type, `T` is a type parameter, and `x` is assignable to each type in `T`'s type set.
> - `V` is a type parameter and `T` is not a named type, and values of each type in `V`'s type set are assignable to `T`.

**【精准逐字翻译】**

此外，如果 `x` 的类型 `V` 或目标类型 `T` 是**类型参数（type parameters，即泛型形参）**，在满足以下条件之一时，`x` 也可赋值给类型为 `T` 的变量：

- `x` 是预声明标识符 `nil`，`T` 是一个类型参数，且 `x` **可赋值给 `T` 的类型集（type set）中的每一个类型**（这意味着 `T` 的约束接口必须只包含指针、切片、Map、channel、func 或 interface 等可 nil 类型）。
- `V` 不是具名类型（如无名切片/结构体字面量或无类型常量），`T` 是一个类型参数，且 `x` **可赋值给 `T` 的类型集中的每一个类型**。
- `V` 是一个类型参数，`T` 不是具名类型，且 `V` 的类型集中**每一个类型的任意值均可赋值给 `T`**。
- **专业细节拆解：**
  - `T's type set`（类型参数的类型集）：Go 泛型中，类型参数 `T` 代表由其 Constraint 约束定义的一个集合。要将一个值赋给泛型变量 `T`，或者把泛型变量 `V` 赋给其他变量，必须确保在**类型集中的每一种具体可能性**下，该赋值在逻辑和语义上均成立。

按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的三段式架构，为您精细拆解与剖析 Go 语言规范中的 **Representability（可表示性）** 章节：

## Representability (可表示性)

### 段落 1 (无类型常量可表示性的核心条件)

> **【英文原文】**
>
> A constant `x` is representable by a value of type `T`, where `T` is not a type parameter, if one of the following conditions applies:
>
> - `x` is in the set of values determined by `T`.
> - `T` is a floating-point type and `x` can be rounded to `T`'s precision without overflow. Rounding uses IEEE 754 round-to-even rules but with an IEEE negative zero further simplified to an unsigned zero. Note that constant values never result in an IEEE negative zero, NaN, or infinity.
> - `T` is a complex type, and `x`'s components `real(x)` and `imag(x)` are representable by values of `T`'s component type (`float32` or `float64`).
>
> If `T` is a type parameter, `x` is representable by a value of type `T` if `x` is representable by a value of each type in `T`'s type set.

**【精准逐字翻译】**

当类型 `T` 不是类型参数（泛型形参）时，如果满足以下条件之一，则常量 `x` 可以被类型 `T` 的值**所表示（representable）**：

- `x` 在由类型 `T` 所确定的值集合（set of values）内。
- `T` 是一个浮点数类型，且 `x` 可以在不发生溢出（without overflow）的情况下舍入（round）到 `T` 的精度。舍入使用 IEEE 754 偶数舍入规则（round-to-even，即四舍六入五成双），但 IEEE 负零（`-0.0`）会被进一步简化为无符号零（`0.0`）。**注意：编译期常量值永远不会产生 IEEE 负零、NaN（非数）或正负无穷大（infinity）**。
- `T` 是一个复数类型，且 `x` 的实部 `real(x)` 和虚部 `imag(x)` 均可被 `T` 的分量类型（`float32` 或 `float64`）的值所表示。

如果 `T` 是一个类型参数（泛型形参），当且仅当 `x` **可以被 `T` 的类型集（type set）中的每一个类型的值所表示**时，`x` 才可以被类型 `T` 的值所表示。

- **专业术语与句式拆解：**
  - `set of values`：值集合。例如 `byte` 的值集合是 $[0, 255]$ 的整数，常量 `42.0`（虽然带有小数点，但数学值为整数 42）属于该集合，因此可表示。
  - `round-to-even`：IEEE 754 规定的“偶数舍入”模式（即银行家舍入法）。
  - `constant values never result in NaN or Inf`：Go 语言规定编译期常量运算如果导致溢出为 `Inf` 或产生 `NaN`，直接在编译期报语法/类型错误，而不会在编译期产生 IEEE 的非法特异值。

### 表格 1 (可表示的有效示例)

> **【英文原文】**
>
> | **x**               | **T**     | **x is representable by a value of T because**               |
> | ------------------- | --------- | ------------------------------------------------------------ |
> | `'a'`               | `byte`    | 97 is in the set of byte values                              |
> | `97`                | `rune`    | `rune` is an alias for `int32`, and 97 is in the set of 32-bit integers |
> | `"foo"`             | `string`  | `"foo"` is in the set of string values                       |
> | `1024`              | `int16`   | 1024 is in the set of 16-bit integers                        |
> | `42.0`              | `byte`    | 42 is in the set of unsigned 8-bit integers                  |
> | `1e10`              | `uint64`  | 10000000000 is in the set of unsigned 64-bit integers        |
> | `2.718281828459045` | `float32` | 2.718281828459045 rounds to 2.7182817 which is in the set of float32 values |
> | `-1e-1000`          | `float64` | -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0 |
> | `0i`                | `int`     | 0 is an integer value                                        |
> | `(42 + 0i)`         | `float32` | 42.0 (with zero imaginary part) is in the set of float32 values |

**【精准逐字翻译与解析】**

| **常量 x**          | **目标类型 T** | **x 可被类型 T 的值所表示的原因**                            |
| ------------------- | -------------- | ------------------------------------------------------------ |
| `'a'`               | `byte`         | 字符 `'a'` 的 ASCII 码是 97，在 `byte`（即 `uint8`）的值集合内 |
| `97`                | `rune`         | `rune` 是 `int32` 的类型别名，97 在 32 位有符号整数的值集合内 |
| `"foo"`             | `string`       | `"foo"` 在 `string` 类型的值集合内                           |
| `1024`              | `int16`        | 1024 在 16 位有符号整数（$-32768 \sim 32767$）的值集合内     |
| `42.0`              | `byte`         | 42.0 的小数部分为 0，数学值为整数 42，属于无符号 8 位整数集合 |
| `1e10`              | `uint64`       | $10^{10} = 10000000000$，属于无符号 64 位整数集合            |
| `2.718281828459045` | `float32`      | 舍入后为 2.7182817，在 `float32` 的精度范围内且未溢出        |
| `-1e-1000`          | `float64`      | 数值极小导致下溢，舍入为 IEEE `-0.0`，并进一步简化为 `0.0`（未溢出至无穷大，因此有效） |
| `0i`                | `int`          | 虚部为 0 的复数常量，其实部数学值为 0，属于整数集合          |
| `(42 + 0i)`         | `float32`      | 虚部为 0，实部为 42.0，属于 `float32` 的值集合               |

### 表格 2 (不可表示的无效示例)

> **【英文原文】**
>
> | **x**    | **T**     | **x is not representable by a value of T because**     |
> | -------- | --------- | ------------------------------------------------------ |
> | `0`      | `bool`    | 0 is not in the set of boolean values                  |
> | `'a'`    | `string`  | `'a'` is a rune, it is not in the set of string values |
> | `1024`   | `byte`    | 1024 is not in the set of unsigned 8-bit integers      |
> | `-1`     | `uint16`  | -1 is not in the set of unsigned 16-bit integers       |
> | `1.1`    | `int`     | 1.1 is not an integer value                            |
> | `42i`    | `float32` | (0 + 42i) is not in the set of float32 values          |
> | `1e1000` | `float64` | 1e1000 overflows to IEEE +Inf after rounding           |

**【精准逐字翻译与解析】**

| **常量 x** | **目标类型 T** | **x 不可被类型 T 的值所表示的原因**                          |
| ---------- | -------------- | ------------------------------------------------------------ |
| `0`        | `bool`         | 0 不在布尔类型（仅包含 `true` 和 `false`）的值集合内         |
| `'a'`      | `string`       | `'a'` 是字符/Rune（整数），不在字符串类型的值集合内          |
| `1024`     | `byte`         | 1024 超出了无符号 8 位整数（$0 \sim 255$）的范围             |
| `-1`       | `uint16`       | -1 不在无符号 16 位整数（$0 \sim 65535$）的范围内            |
| `1.1`      | `int`          | 1.1 含有非零小数部分，不是整数值                             |
| `42i`      | `float32`      | $(0 + 42i)$ 含有非零虚部 42，无法被纯实数浮点数 `float32` 表示 |
| `1e1000`   | `float64`      | $10^{1000}$ 远超双精度浮点数最大上限（约 $1.79 \times 10^{308}$），舍入后溢出为 `+Inf` |

按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的三段式架构，为您精细拆解与剖析 Go 语言规范中的 **Method sets（方法集）** 章节：

## Method sets (方法集)

### 段落 1 (方法集的核心定义与规则)

> **【英文原文】**
>
> The method set of a type determines the methods that can be called on an operand of that type. Every type has a (possibly empty) method set associated with it:
>
> - The method set of a defined type `T` consists of all methods declared with receiver type `T`.
> - The method set of a pointer to a defined type `T` (where `T` is neither a pointer nor an interface) is the set of all methods declared with receiver `*T` or `T`.
> - The method set of an interface type is the intersection of the method sets of each type in the interface's type set (the resulting method set is usually just the set of declared methods in the interface).
>
> Further rules apply to structs (and pointer to structs) containing embedded fields, as described in the section on struct types. Any other type has an empty method set.
>
> In a method set, each method must have a unique non-blank method name.

**【精准逐字翻译】**

一个类型的方法集（method set）决定了可以在该类型的操作数（operand）上调用的方法。每一个类型都有一个与其关联的（可能为空的）方法集：

- 具名/定义类型 `T` 的方法集，由所有以接收者类型为 `T` 声明的方法组成。
- 指向具名/定义类型 `T` 的指针 `*T` 的方法集（其中 `T` 既不是指针类型，也不是接口类型），是所有以接收者类型为 `*T` 或 `T` 声明的方法的集合。
- 接口类型的方法集，是该接口类型集中**每一个类型的方法集的交集（intersection）**（最终的方法集通常就只是该接口中显式声明的方法集合）。

包含嵌套字段（embedded fields）的结构体（以及指向结构体的指针）适用更进一步的规则，正如结构体类型章节中所描述的那样。任何其他类型的方法集均为空。

在同一个方法集中，每个方法必须具有唯一的、非下划线（non-blank，即不能为 `_`）的方法名。

- **专业术语与句式拆解：**
  - **非对称性规则（Asymmetry of Method Sets）**：
    - **类型 `T`** 的方法集仅包含接收者为 `(r T)` 的方法。
    - **指针 `\*T`** 的方法集包含接收者为 `(r T)` 和 `(r *T)` 的**所有方法**。
  - **接口满足条件（Interface Satisfaction）**：一个具体类型是否实现了某个接口，取决于**该具体类型本身的方法集**（而非调用时的语法糖）是否包含该接口的所有方法。这就是为什么将值 `T` 赋给接口时，不能调用接收者为 `*T` 的方法；而将指针 `*T` 赋给接口时，既可以调用 `*T` 也可以调用 `T` 的方法。
  - `intersection of method sets`：接口的底层类型集（Type Set）中所有类型方法集的“交集”。对于普通接口（Basic Interface），这正好等于接口直接列出的方法集合。
  - `non-blank method name`：方法名不能使用空标识符 `_`。

---

