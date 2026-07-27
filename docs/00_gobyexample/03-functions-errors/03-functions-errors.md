当然，下面给你一份可以直接放进 `03-functions-errors.md` 的 Markdown 模板，专门覆盖 **Functions / Multiple Return Values / Variadic Functions / Closures / Recursion / Errors / Custom Errors**。



# Go 函数与错误笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者到初级进阶  
> 目标：理解 Go 的函数写法、返回值习惯以及错误处理方式

---

## 1. Functions

1.1 一句话理解

函数是可复用的代码块，用来封装一段逻辑。

1.2 示例代码

```go
package main

import "fmt"

func plus(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println(plus(1, 2))
}
```

| 1.3 重点                     | 1.4 常见坑                           |
| ---------------------------- | ------------------------------------ |
| 使用 `func` 定义函数。       | 参数类型顺序写错。                   |
| 参数类型写在变量名后面。     | 忘记写返回类型。                     |
| 返回值类型写在参数列表后面。 | 函数名首字母小写，导致外部包不可见。 |

1.5 我的理解

Go 的函数语法简洁，但类型写法和很多语言不一样。

---

## 2. Multiple Return Values

2.1 一句话理解

Go 的函数可以返回多个值，这是它很重要的特点之一。

2.2 示例代码

```go
package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
}
```

| 2.3 重点                           | 2.4 常见坑                       |
| ---------------------------------- | -------------------------------- |
| Go 非常常见“多个返回值”。          | 返回值数量和接收变量数量不一致。 |
| 典型形式是 `result, err`。         | 不习惯 Go 把错误也作为返回值。   |
| 解构赋值时，接收变量的个数要匹配。 |                                  |

2.5 我的理解

多返回值是 Go 风格的核心特征之一，尤其在错误处理里非常重要。

---

## 3. Variadic Functions

3.1 一句话理解

可变参数函数可以接收任意数量的参数。

3.2 示例代码

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))
    fmt.Println(sum(4, 5))
}
```

| 3.3 重点                     | 3.4 常见坑                               |
| ---------------------------- | ---------------------------------------- |
| `...` 表示可变参数。         | 不理解 `...` 和切片的关系。              |
| 函数内部把它当成切片使用。   | 以为只能传多个单独参数，不能传切片展开。 |
| 很适合处理数量不固定的输入。 |                                          |

3.5 我的理解

可变参数让函数更灵活，但也不要滥用。

---

## 4. Closures

4.1 一句话理解

闭包是“函数 + 它引用的外部变量”。

4.2 示例代码

```go
package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
}
```

| 4.3 重点                         | 4.4 常见坑                           |
| -------------------------------- | ------------------------------------ |
| 内部函数可以访问外部函数的变量。 | 不理解闭包为什么能保存状态。         |
| 外部变量会被“捕获”。             | 在循环里使用闭包时产生变量捕获问题。 |
| 闭包常用于状态保持。             |                                      |

4.5 我的理解

闭包让函数可以带着自己的“小状态”一起运行。

---

## 5. Recursion

5.1 一句话理解

递归是函数调用自己，用来处理可分解的问题。

5.2 示例代码

```go
package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}
```

| 5.3 重点                       | 5.4 常见坑               |
| ------------------------------ | ------------------------ |
| 递归必须有终止条件。           | 忘记终止条件。           |
| 常用于树、分治、数学定义。     | 递归层数太深导致栈问题。 |
| 逻辑要清晰，不然容易无限递归。 | 逻辑写得比循环更难读。   |

5.5 我的理解

递归是思想工具，不是默认首选方案。

---

## 6. Errors

6.1 一句话理解

Go 用显式返回 `error` 来表示错误。

6.2 示例代码

```go
package main

import (
    "errors"
    "fmt"
)

func f(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}

func main() {
    if _, err := f(42); err != nil {
        fmt.Println("error:", err)
    }
}
```

| 6.3 重点                         | 6.4 常见坑                             |
| -------------------------------- | -------------------------------------- |
| 错误是正常返回值，不是异常机制。 | 忽略错误。                             |
| 常见模式是 `value, err := ...`。 | 把 `panic` 当成普通错误处理。          |
| `err != nil` 时立刻处理。        | 没有区分“可恢复错误”和“程序崩溃错误”。 |

6.5 我的理解

Go 的错误处理强调显式和及时。

---

## 7. Custom Errors

7.1 一句话理解

自定义错误可以携带更有意义的信息。

7.2 示例代码

```go
package main

import "fmt"

type argError struct {
    arg int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {
    _, err := f(42)
    if err != nil {
        fmt.Println(err)
    }
}
```

| 7.3 重点                                | 7.4 常见坑                       |
| --------------------------------------- | -------------------------------- |
| 自定义错误类型通常实现 `Error()` 方法。 | 自定义错误只顾结构，不顾可读性。 |
| 可以携带字段，方便调试和判断。          | 过度设计错误类型，反而更复杂。   |
| 适合复杂业务场景。                      |                                  |

7.5 我的理解

自定义错误的核心不是“炫技”，而是让错误更有上下文。

---

## 总结

### 建议记忆顺序
1. Functions
2. Multiple Return Values
3. Variadic Functions
4. Closures
5. Recursion
6. Errors
7. Custom Errors

### 你要重点记住的三件事
- Go 很常见多返回值。
- `error` 是普通返回值。
- 闭包和递归是理解 Go 进阶代码的关键。

## 学习建议

这组内容里，最重要的是先把下面三件事吃透：

- 函数怎么定义。
- 多返回值怎么接。
- `error` 为什么是“返回值”而不是“异常”。

你后面学 Go 的很多代码，都会反复见到这些模式。

如果你愿意，我下一步可以继续帮你写 **`04-types-interfaces.md`**。