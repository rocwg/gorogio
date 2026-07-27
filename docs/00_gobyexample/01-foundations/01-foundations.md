# Go 基础语法笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者  
> 目标：建立 Go 基础语法直觉，并能自己写出最小示例

---

## 1. Hello World

1.1 一句话理解

第一个 Go 程序，主要认识 `package main` 和 `func main()`。

1.2 示例代码

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

| 1.3 重点                            | 1.4 常见坑               |
| ----------------------------------- | ------------------------ |
| `package main` 表示这是可执行程序。 | 忘记写 `main` 包。       |
| `func main()` 是程序入口。          | 忘记导入 `fmt`。         |
| `fmt.Println` 用来输出。            | 花括号换行导致风格问题。 |

1.5 我的理解

这里最重要的是先接受 Go 的程序结构。

---

## 2. Values

2.1 一句话理解

Go 里的值可以是字符串、数字、布尔值等基础字面量。

2.2 示例代码

```go
// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```

| 2.3 重点             | 2.4 常见坑                         |
| -------------------- | ---------------------------------- |
| 字符串可以拼接。     | 混淆字符串和数字。                 |
| 数字可以直接运算。   | 以为 Go 会自动帮你做很多隐式转换。 |
| 布尔值用于条件判断。 |                                    |

2.5 我的理解

Go 很强调类型明确，不能随便混着用。

---

## 3. Variables

3.1 一句话理解

变量用于保存会变化的数据。

3.2 示例代码

```go
// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	// This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)
}
```

| 3.3 重点                         | 3.4 常见坑                       |
| -------------------------------- | -------------------------------- |
| `var` 是标准声明方式。           | `:=` 只能在函数内部使用。        |
| `:=` 是短变量声明，Go 里很常用。 | 同一作用域里变量名不能重复定义。 |
| 变量类型可以显式写，也可以推断。 |                                  |

3.5 我的理解

这是 Go 最常见的简洁语法之一。

---

## 4. Constants

4.1 一句话理解

常量是不能被修改的值。

4.2 示例代码

```go
// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A `const` statement can also appear inside a
	// function body.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	fmt.Println(math.Sin(n))
}
```

| 4.3 重点                     | 4.4 常见坑           |
| ---------------------------- | -------------------- |
| 常量用 `const` 声明。        | 把常量当变量改掉。   |
| 常量一旦定义，不能重新赋值。 | 不理解 iota 的作用。 |

4.5 我的理解

常量适合固定配置、枚举值、不会变化的标识。

---

## 5. For

5.1 一句话理解

Go 只有 `for`，它可以表示循环、条件循环、无限循环。

5.2 示例代码

```go
// `for` is Go's only looping construct. Here are
// some basic types of `for` loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic "do this
	// N times" iteration is `range` over an integer.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```

| 5.3 重点                         | 5.4 常见坑                  |
| -------------------------------- | --------------------------- |
| Go 没有单独的 `while`。          | 忘记循环变量作用域。        |
| `for` 是唯一的循环关键字。       | `range` 和普通 `for` 搞混。 |
| 可以写成三段式、条件式、死循环。 |                             |

5.5 我的理解

Go 的循环语法少，但表达能力足够。

---

## 6. If/Else

6.1 一句话理解

`if/else` 用来判断条件并分支执行。

6.2 示例代码

```go
package main

import "fmt"

func main() {
    if 7%2 == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }
}
```

| 6.3 重点                     | 6.4 常见坑                       |
| ---------------------------- | -------------------------------- |
| 条件表达式后面直接写花括号。 | 忘记花括号。                     |
| 不需要括号包住条件。         | 习惯性写成其他语言那种括号风格。 |
| 可以带初始化语句。           | 把 `=` 和 `==` 搞混。            |

6.5 我的理解

Go 的 `if` 重点是简洁和清晰。

---

## 7. Switch

7.1 一句话理解

`switch` 是多分支判断，比一串 `if/else` 更清晰。

7.2 示例代码

```go
package main

import "fmt"

func main() {
    i := 2
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("other")
    }
}
```

| 7.3 重点                                                | 7.4 常见坑                       |
| ------------------------------------------------------- | -------------------------------- |
| Go 的 `switch` 默认只执行匹配分支，不需要手写 `break`。 | 以为它会像某些语言一样自动贯穿。 |
| `default` 是兜底分支。                                  | 不理解“无条件 switch”的写法。    |
| `switch` 也可以没有条件，像更高级的 `if/else`。         |                                  |

7.5 我的理解

Go 的 `switch` 很适合写清晰的分支逻辑。

---

## 总结

这一组内容适合放在一个文件里，因为它们都属于 Go 的第一层基础语法。

### 建议记忆顺序
1. Hello World
2. Values
3. Variables
4. Constants
5. For
6. If/Else
7. Switch

### 你可以重点记住的三个习惯
- `:=` 很常用。
- Go 里 `for` 最核心。
- `switch` 通常比多层 `if/else` 更清楚。