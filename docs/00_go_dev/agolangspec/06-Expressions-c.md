我们现在跳转到了 Specification 中非常核心的 **索引表达式（Index expressions）** 章节（位于表达式 Expressions 部分）。这一章极为严谨地界定了数组、切片、字符串、Map 以及泛型类型参数（Type Parameter）在执行 `a[x]` 时的全部语法与语义规则。

我们将严格延续 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与句式拆解】** 的标准，一字一句进行透彻解读。

## 章节：索引表达式 (Index expressions)

### 段落 1：基本定义

> **【英文原文】**
>
> A primary expression of the form
>
> ```
> a[x]
> ```
>
> denotes the element of the array, pointer to array, slice, string or map `a` indexed by `x`. The value `x` is called the index or map key, respectively. The following rules apply:

**【逐字精准翻译】**

形式为 `a[x]` 的初级表达式（primary expression）表示由 `x` 索引的数组、数组指针、切片、字符串或字典 `a` 中的元素。值 `x` 分别被称为索引（index）或字典键（map key）。以下规则适用：

- **词汇与句式剖析：**
  - `primary expression`：初级表达式（语法树中最基础的表达式单元，如变量名、字面量、括号表达式或带操作符的简单调用）。
  - `denotes`：表示 / 意味着。
  - `indexed by x`：被 `x` 所索引。
  - `respectively`：分别（对应前面的顺序：对于数组/切片/字符串叫 index，对于 map 叫 map key）。

### 段落 2：通用规则（当 a 不是 map 且不是类型参数时）

> **【英文原文】**
>
> If `a` is neither a map nor a type parameter:
>
> - the index `x` must be an untyped constant, or its type must be an integer or a type parameter whose type set contains only integer types
>
> - a constant index must be non-negative and representable by a value of type `int`
>
> - a constant index that is untyped is given type `int`
>
> - the index `x` is in range if `0 <= x < len(a)`, otherwise it is out of range

**【逐字精准翻译】**

如果 `a` 既不是 map 也不是类型参数（type parameter）：

- 索引 `x` 必须是一个无类型常量（untyped constant），或者其类型必须是整数，或者是类型集（type set）仅包含整数类型的类型参数
- 常量索引必须是非负的，并且可以由 `int` 类型的值表示
- 无类型的常量索引会被赋予 `int` 类型
- 如果 `0 <= x < len(a)`，则索引 `x` 在范围内（in range），否则超出范围（out of range）

- **词汇与句式剖析：**
  - `neither ... nor ...`：既不……也不……。
  - `untyped constant`：无类型常量（例如直接写 `a[0]` 中的数字字面量 `0`，在未确定类型前是无类型的）。
  - `type set`：类型集（Go 泛型接口中定义的类型集合）。
  - `non-negative`：非负的（$\ge 0$）。
  - `representable by`：可由……表示的（指数值不能溢出 `int` 的范围）。
  - `out of range`：超出范围 / 越界。

### 段落 3：数组（Array）规则

> **【英文原文】**
>
> For `a` of array type `A`:
>
> - a constant index must be in range
> - if `x` is out of range at run time, a run-time panic occurs
> - `a[x]` is the array element at index `x` and the type of `a[x]` is the element type of `A`

**【逐字精准翻译】**

对于数组类型 `A` 的 `a`：

- 常量索引必须在范围内（*注：如果常量索引越界，编译期直接报错*）
- 如果 `x` 在运行时超出范围，会发生运行时 panic（运行时恐慌）
- `a[x]` 是索引 `x` 处的数组元素，并且 `a[x]` 的类型是 `A` 的元素类型

### 段落 4：数组指针（Pointer to array）规则

> **【英文原文】**
>
> For `a` of pointer to array type:
>
> - `a[x]` is shorthand for `(*a)[x]`

**【逐字精准翻译】**

对于数组指针类型的 `a`：

- `a[x]` 是 `(*a)[x]` 的简写形式

- **词汇剖析：**
  - `shorthand for`：……的简写 / 缩写。
  - **深度解读：** 语法糖。Go 会自动对数组指针进行隐式解引用，无需显式写 `(*a)[x]`。

### 段落 5：切片（Slice）规则

> **【英文原文】**
>
> For `a` of slice type `S`:
>
> - if `x` is out of range at run time, a run-time panic occurs
> - `a[x]` is the slice element at index `x` and the type of `a[x]` is the element type of `S`

**【逐字精准翻译】**

对于切片类型 `S` 的 `a`：

- 如果 `x` 在运行时超出范围，会发生运行时 panic
- `a[x]` 是索引 `x` 处的切片元素，并且 `a[x]` 的类型是 `S` 的元素类型

### 段落 6：字符串（String）规则

> **【英文原文】**
>
> For `a` of string type:
>
> - a constant index must be in range if the string `a` is also constant
> - if `x` is out of range at run time, a run-time panic occurs
> - `a[x]` is the non-constant byte value at index `x` and the type of `a[x]` is `byte`
> - `a[x]` may not be assigned to

**【逐字精准翻译】**

对于字符串类型的 `a`：

- 如果字符串 `a` 本身也是常量，则常量索引必须在范围内
- 如果 `x` 在运行时超出范围，会发生运行时 panic
- `a[x]` 是索引 `x` 处的非常量字节值（byte value），且 `a[x]` 的类型是 `byte`（即 `uint8`）
- `a[x]` 不能被赋值

- **词汇与句式剖析：**
  - `byte value`：字节值（注意：字符串索引获得的是单字节 `byte`，而不是 Unicode 字符 `rune`！）。
  - `may not be assigned to`：不能被赋值（因为 Go 语言中的字符串是不可变 immutable 的）。

### 段落 7：字典（Map）规则

> **【英文原文】**
>
> For `a` of map type `M`:
>
> - `x`'s type must be assignable to the key type of `M`
>
> - if the map contains an entry with key `x`, `a[x]` is the map element with key `x` and the type of `a[x]` is the element type of `M`
>
> - if the map is `nil` or does not contain such an entry, `a[x]` is the zero value for the element type of `M`
>

**【逐字精准翻译】**

对于字典（map）类型 `M` 的 `a`：

- `x` 的类型必须可以赋值给 `M` 的键（key）类型
- 如果 map 包含键为 `x` 的条目（entry），则 `a[x]` 是键为 `x` 的 map 元素，且 `a[x]` 的类型是 `M` 的元素类型
- 如果 map 为 `nil` 或者不包含该条目，则 `a[x]` 是 `M` 的元素类型的零值（zero value）

- **词汇剖析：**
  - `assignable to`：可赋值给……（符合 Go 的赋值兼容规则）。
  - `entry`：条目 / 键值对。
  - `zero value`：零值（类型的默认初始值，如 `0`、`""`、`nil` 等）。

### 段落 8：泛型类型参数（Type parameter）规则

> **【英文原文】**
>
> For `a` of type parameter type `P`:
>
> - The index expression `a[x]` must be valid for values of all types in `P`'s type set.
> - The element types of all types in `P`'s type set must be identical. In this context, the element type of a string type is `byte`.
> - If there is a map type in the type set of `P`, all types in that type set must be map types, and the respective key types must be all identical.
> - `a[x]` is the array, slice, or string element at index `x`, or the map element with key `x` of the type argument that `P` is instantiated with, and the type of `a[x]` is the type of the (identical) element types.
> - `a[x]` may not be assigned to if`P`'s type set includes string types.

**【逐字精准翻译】**

对于类型参数类型 `P` 的 `a`：

- 索引表达式 `a[x]` 对 `P` 的类型集中所有类型的组合值都必须有效。
- `P` 的类型集中所有类型的元素类型必须完全相同（identical）。在此语境下，字符串类型的元素类型是 `byte`。
- 如果 `P` 的类型集中包含字典（map）类型，则该类型集中的所有类型都必须是 map 类型，且它们各自的键类型必须完全相同。
- `a[x]` 是 `P` 被实例化时的实际类型参数（type argument）在索引 `x` 处的数组、切片或字符串元素，或是键为 `x` 的 map 元素，且 `a[x]` 的类型是（完全相同的）元素类型。
- 如果 `P` 的类型集包含字符串类型，则 `a[x]` 不能被赋值。

- **词汇与句式剖析：**
  - `identical`：完全相同的 / 恒等的（Go 规范对类型相同有严格定义）。
  - `type argument`：类型参数 / 实际传入的类型实参（如 `List[int]` 中的 `int`）。
  - `instantiated with`：被……所实例化。

### 段落 9：非法情况总结

> **【英文原文】**
>
> Otherwise `a[x]` is illegal.

**【逐字精准翻译】**

否则，`a[x]` 是不合法的（编译错误）。

### 段落 10：Map 的逗号-ok 赋值模式（Comma-ok idiom）

> **【英文原文】**
>
> An index expression on a map `a` of type `map[K]V` used in an assignment statement or initialization of the special form
>
> ```go
> v, ok = a[x]
> v, ok := a[x]
> var v, ok = a[x]
> ```
>
> yields an additional untyped boolean value. The value of `ok` is `true` if the key `x` is present in the map, and `false` otherwise.

**【逐字精准翻译】**

在赋值语句或特殊形式的初始化中，用于类型为 `map[K]V` 的字典 `a` 上的索引表达式：

```go
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
```

会产生一个额外的不带类型布尔值（untyped boolean value）。如果键 `x` 存在于 map 中，`ok` 的值为 `true`；否则为 `false`。

- **词汇与句式剖析：**
  - `yields`：产生 / 产出。
  - `untyped boolean value`：无类型布尔值（可以自动赋值给任何自定义的布尔别名类型）。
  - `present in`：存在于……中。

### 段落 11：对 nil map 赋值的异常

> **【英文原文】**
>
> Assigning to an element of a `nil` map causes a run-time panic.

**【逐字精准翻译】**

给 `nil` map 的元素赋值会导致运行时 panic（运行时恐慌）。

- **核心踩坑点：** 从 `nil` map 中**读取** `a[x]` 不会报错，只会返回零值；但是**写** `nil` map（如 `a[x] = 1`）会直接触发 panic！

本章节逐字对照翻译与剖析已完成。结尾提示的 `Slice expressions¶` 标志着紧接着的下一节是 **切片表达式（Slice expressions，如 `a[low:high:max]`）**。我们可以直接按相同方式继续推进！

