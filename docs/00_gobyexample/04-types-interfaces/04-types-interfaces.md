当然，下面给你一份可以直接放进 `04-types-interfaces.md` 的 Markdown 模板，覆盖 **Pointers / Strings and Runes / Structs / Methods / Interfaces / Enums / Struct Embedding / Generics / Range over Iterators**。



# Go 类型系统笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者到初级进阶  
> 目标：理解 Go 的类型系统、面向对象风格和泛型基础

---

## 1. Pointers

1.1 一句话理解

指针保存的是变量的地址，不是值本身。

1.2 示例代码

```go
package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}
```

| 1.3 重点                        | 1.4 常见坑                         |
| ------------------------------- | ---------------------------------- |
| `*T` 表示指向 `T` 的指针。      | 分不清 `&` 和 `*`。                |
| `&x` 取地址。                   | 误以为函数参数传进来都会修改原值。 |
| `*p` 解引用，访问地址指向的值。 | 指针为空时直接解引用。             |

1.5 我的理解

指针是 Go 中理解“引用语义”的基础。

---

## 2. Strings and Runes

2.1 一句话理解

字符串是字节序列，`rune` 表示 Unicode 字符。

2.2 示例代码

```go
package main

import "fmt"

func main() {
    const s = "你好，Go"

    for i, r := range s {
        fmt.Printf("%d %c\n", i, r)
    }
}
```

| 2.3 重点                             | 2.4 常见坑                      |
| ------------------------------------ | ------------------------------- |
| 字符串底层是字节序列。               | 把字符串长度当作“字符数”。      |
| `range` 遍历字符串时会按 rune 处理。 | 中文、emoji、特殊符号处理不当。 |
| `rune` 本质上是 `int32`。            | 把 byte 和 rune 混为一谈。      |

2.5 我的理解

这部分最重要的是理解 UTF-8 和字符边界。

---

## 3. Structs

3.1 一句话理解

struct 是把多个字段组合成一个类型。

3.2 示例代码

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {
    p := person{name: "Alice", age: 25}
    fmt.Println(p)
}
```

| 3.3 重点                    | 3.4 常见坑                               |
| --------------------------- | ---------------------------------------- |
| `struct` 是自定义复合类型。 | 字段名首字母小写导致外部不可见。         |
| 字段可以显式命名初始化。    | 初始化时字段顺序和类型混乱。             |
| 常用于表示业务对象。        | 把 struct 当成“类”但忽略 Go 的组合风格。 |

3.5 我的理解

struct 是 Go 里承载数据模型的核心类型。

---

## 4. Methods

4.1 一句话理解

方法是绑定到某个类型上的函数。

4.2 示例代码

```go
package main

import "fmt"

type rect struct {
    width, height int
}

func (r rect) area() int {
    return r.width * r.height
}

func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area:", r.area())
}
```

| 4.3 重点                         | 4.4 常见坑                         |
| -------------------------------- | ---------------------------------- |
| 方法接收者写在函数名和参数之间。 | 不清楚值接收者和指针接收者的影响。 |
| 值接收者和指针接收者有区别。     | 以为方法等同于类成员函数。         |
| 方法让类型更有行为。             | 接收者类型选错导致无法修改原对象。 |

4.5 我的理解

方法是 Go 里“数据 + 行为”组合的入口。

---

## 5. Interfaces

5.1 一句话理解

接口定义“行为”，不是定义数据。

5.2 示例代码

```go
package main

import "fmt"

type geometry interface {
    area() float64
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func measure(g geometry) {
    fmt.Println(g.area())
}

func main() {
    c := circle{radius: 2}
    measure(c)
}
```

| 5.3 重点                               | 5.4 常见坑                                     |
| -------------------------------------- | ---------------------------------------------- |
| 接口是方法集合。                       | 以为必须显式声明“implements”。                 |
| 只要实现了这些方法，就“自动满足”接口。 | 接口设计太大。                                 |
| Go 倾向于隐式实现接口。                | 把接口当成“对象分类”工具，而不是行为抽象工具。 |

5.5 我的理解

Go 的接口很轻，重点是抽象行为，不是描述继承关系。

---

## 6. Enums

6.1 一句话理解

Go 没有传统 enum，但常用 `const` + `iota` 模拟枚举。

6.2 示例代码

```go
package main

import "fmt"

type serverState int

const (
    stateIdle serverState = iota
    stateConnected
    stateError
)

func main() {
    fmt.Println(stateIdle, stateConnected, stateError)
}
```

| 6.3 重点                       | 6.4 常见坑                       |
| ------------------------------ | -------------------------------- |
| `iota` 常用于连续常量。        | 不理解 `iota` 的递增规则。       |
| 常配合自定义类型使用。         | 把枚举值当成真正独立类型。       |
| 很适合状态、等级、模式等场景。 | 没有定义字符串输出，调试不直观。 |

6.5 我的理解

Go 的枚举是“约定式”的，不是语言里单独一套复杂系统。

---

## 7. Struct Embedding

7.1 一句话理解

结构体嵌入是组合复用的一种方式。

7.2 示例代码

```go
package main

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base
    str string
}

func main() {
    co := container{
        base: base{num: 1},
        str:  "hello",
    }

    fmt.Println(co.describe())
}
```

| 7.3 重点                             | 7.4 常见坑                         |
| ------------------------------------ | ---------------------------------- |
| 嵌入不是继承，但有“提升”效果。       | 把 embedding 误认为类继承。        |
| 被嵌入类型的方法会被提升到外层类型。 | 多个嵌入字段导致方法名冲突。       |
| Go 更偏向组合，不偏向传统继承。      | 以为结构体嵌入会自动复制所有语义。 |

7.5 我的理解

嵌入是 Go 组合思想的典型体现。

---

## 8. Generics

8.1 一句话理解

泛型让函数和类型可以处理多种类型。

8.2 示例代码

```go
package main

import "fmt"

func mapKeys[K comparable, V any](m map[K]V) []K {
    ks := make([]K, 0, len(m))
    for k := range m {
        ks = append(ks, k)
    }
    return ks
}

func main() {
    m := map[string]int{"a": 1, "b": 2}
    fmt.Println(mapKeys(m))
}
```

| 8.3 重点                      | 8.4 常见坑                 |
| ----------------------------- | -------------------------- |
| 泛型通过类型参数实现。        | 过早使用泛型。             |
| `any` 表示任意类型。          | 泛型约束写得太复杂。       |
| `comparable` 表示可比较类型。 | 想用泛型解决所有重复代码。 |

8.5 我的理解

泛型是工具，不是目标；能简化重复代码才值得用。

---

## 9. Range over Iterators

9.1 一句话理解

`range` 不只可以遍历切片和 map，也可以用于迭代器风格的数据来源。

9.2 示例代码

```go
package main

import "fmt"

func integers() func(func(int) bool) {
    return func(yield func(int) bool) {
        for i := 1; i <= 3; i++ {
            if !yield(i) {
                return
            }
        }
    }
}

func main() {
    for v := range integers() {
        fmt.Println(v)
    }
}
```

| 9.3 重点                            | 9.4 常见坑                 |
| ----------------------------------- | -------------------------- |
| 这是更高级的迭代方式。              | 一开始不容易直观理解。     |
| 和传统 `range` 思维类似，但更灵活。 | 容易和普通 `range` 混淆。  |
| 常用于抽象数据源。                  | 初学阶段不必深钻内部机制。 |

9.5 我的理解

这是进阶内容，先知道它能做什么就够了。

---

## 总结

### 建议记忆顺序
1. Pointers
2. Strings and Runes
3. Structs
4. Methods
5. Interfaces
6. Enums
7. Struct Embedding
8. Generics
9. Range over Iterators

### 你要重点记住的三件事
- Go 的接口是“隐式实现”的。
- Go 更强调组合，不强调传统继承。
- 泛型是增强表达能力的工具，不是必须到处用的东西。

## 学习建议

这一组内容是 Go 从“会写语法”走向“会组织代码”的关键部分。
尤其是 `struct`、`method`、`interface` 这三者，建议你反复看、反复写，因为它们会直接影响你以后写项目的方式。

如果你愿意，我下一步可以继续帮你写 **`05-concurrency.md`**。