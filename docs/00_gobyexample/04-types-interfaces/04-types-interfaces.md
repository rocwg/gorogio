当然，下面给你一份可以直接放进 `04-types-interfaces.md` 的 Markdown 模板，覆盖 **Pointers / Strings and Runes / Structs / Methods / Interfaces / Enums / Struct Embedding / Generics / Range over Iterators**。



# Go 类型系统笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者到初级进阶  
> 目标：理解 Go 的类型系统、面向对象风格和泛型基础

---

## 1. Pointers

1.1 一句话理解：指针保存的是变量的地址，不是值本身。

1.2 示例代码

```go
// Go supports <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// allowing you to pass references to values and records
// within your program.

package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)

	// A new pointer to a value can be created with the
	// builtin function `new`.
	p := new(42)
	fmt.Println("value at *p:", *p)
	zeroptr(p)
	fmt.Println("value at *p:", *p)
}
```

```bash
# `zeroval` doesn't change the `i` in `main`, but
# `zeroptr` does because it has a reference to
# the memory address for that variable.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x1178e638e050
value at *p: 42
value at *p: 0
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

2.1 一句话理解：字符串是字节序列，`rune` 表示 Unicode 字符。

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

3.1 一句话理解：struct 是把多个字段组合成一个类型。

3.2 示例代码

```go
// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

// `newPerson` constructs a new person struct with the given name.
func newPerson(name string) *person {
	// Go is a garbage collected language; you can safely
	// return a pointer to a local variable - it will only
	// be cleaned up by the garbage collector when there
	// are no active references to it.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value, we don't
	// have to give it a name. The value can have an anonymous
	// struct type. This technique is commonly used for
	// [table-driven tests](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
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

4.1 一句话理解：方法是绑定到某个类型上的函数。

4.2 示例代码

```go
// Go supports _methods_ defined on struct types.

package main

import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
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

5.1 一句话理解：***==接口定义“行为”，不是定义数据==***。

5.2 示例代码

```go
// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"fmt"
	"math"
)

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Sometimes it's useful to know the runtime type of an
// interface value. One option is using a *type assertion*
// as shown here; another is a [type `switch`](switch).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
```

```bash
$ go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
circle with radius 5

# To understand how Go's interfaces work under the hood,
# check out this [blog post](https://research.swtch.com/interfaces).
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

6.1 一句话理解：Go 没有传统 enum，但常用 `const` + `iota` 模拟枚举。

6.2 示例代码

```go
// _Enumerated types_ (enums) are a special case of
// [sum types](https://en.wikipedia.org/wiki/Algebraic_data_type).
// An enum is a type that has a fixed number of possible
// values, each with a distinct name. Go doesn't have an
// enum type as a distinct language feature, but enums
// are simple to implement using existing language idioms.

package main

import "fmt"

// Our enum type `ServerState` has an underlying `int` type.
type ServerState int

// The possible values for `ServerState` are defined as
// constants. The special keyword [iota](https://go.dev/ref/spec#Iota)
// generates successive constant values automatically; in this
// case 0, 1, 2 and so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// interface, values of `ServerState` can be printed out or converted
// to strings.
//
// This can get cumbersome if there are many possible values. In such
// cases the [stringer tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// can be used in conjunction with `go:generate` to automate the
// process. See [this post](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// for a longer explanation.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// If we have a value of type `int`, we cannot pass it to `transition` - the
	// compiler will complain about type mismatch. This provides some degree of
	// compile-time type safety for enums.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition emulates a state transition for a
// server; it takes the existing state and returns
// a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some predicates here to
		// determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
```

```bash
$ go run enums.go
connected
idle
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

7.1 一句话理解：结构体嵌入是组合复用的一种方式。

7.2 示例代码

```go
// Go supports _embedding_ of structs and interfaces
// to express a more seamless _composition_ of types.
// This is not to be confused with [`//go:embed`](embed-directive) which is
// a go directive introduced in Go version 1.16+ to embed
// files and folders into the application binary.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A `container` _embeds_ a `base`. An embedding looks
// like a field without a name.
type container struct {
	base
	str string
}

func main() {

	// When creating structs with literals, we have to
	// initialize the embedding explicitly; here the
	// embedded type serves as the field name.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base's fields directly on `co`,
	// e.g. `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternatively, we can spell out the full path using
	// the embedded type name.
	fmt.Println("also num:", co.base.num)

	// Since `container` embeds `base`, the methods of
	// `base` also become methods of a `container`. Here
	// we invoke a method that was embedded from `base`
	// directly on `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedding structs with methods may be used to bestow
	// interface implementations onto other structs. Here
	// we see that a `container` now implements the
	// `describer` interface because it embeds `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
```

```bash
$ go run struct-embedding.go
co={num: 1, str: some name}
also num: 1
describe: base with num=1
describer: base with num=1
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

8.1 一句话理解：泛型让函数和类型可以处理多种类型。

8.2 示例代码

```go
// Starting with version 1.18, Go has added support for
// _generics_, also known as _type parameters_.

package main

import "fmt"

// As an example of a generic function, `SlicesIndex` takes
// a slice of any `comparable` type and an element of that
// type and returns the index of the first occurrence of
// v in s, or -1 if not present. The `comparable` constraint
// means that we can compare values of this type with the
// `==` and `!=` operators. For a more thorough explanation
// of this type signature, see [this blog post](https://go.dev/blog/deconstructing-type-parameters).
// Note that this function exists in the standard library
// as [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// As an example of a generic type, `List` is a
// singly-linked list with values of any type.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// We can define methods on generic types just like we
// do on regular types, but we have to keep the type
// parameters in place. The type is `List[T]`, not `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements returns all the List elements as a slice.
// In the next example we'll see a more idiomatic way
// of iterating over all elements of custom types.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// When invoking generic functions, we can often rely
	// on _type inference_. Note that we don't have to
	// specify the types for `S` and `E` when
	// calling `SlicesIndex` - the compiler infers them
	// automatically.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... though we could also specify them explicitly.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
```

```bash
$ go run generics.go
index of zoo: 2
list: [10 13 23]
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

9.1 一句话理解：`range` 不只可以遍历切片和 map，也可以用于迭代器风格的数据来源。

9.2 示例代码

```go
// Starting with version 1.23, Go has added support for
// [iterators](https://go.dev/blog/range-functions),
// which lets us range over pretty much anything!

package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// Let's look at the `List` type from the
// [previous example](generics) again. In that example
// we had an `AllElements` method that returned a slice
// of all elements in the list. With Go iterators, we
// can do it better - as shown below.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All returns an _iterator_, which in Go is a function
// with a [special signature](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// The iterator function takes another function as
		// a parameter, called `yield` by convention (but
		// the name can be arbitrary). It will call `yield` for
		// every element we want to iterate over, and note `yield`'s
		// return value for a potential early termination.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteration doesn't require an underlying data structure,
// and doesn't even have to be finite! Here's a function
// returning an iterator over Fibonacci numbers: it keeps
// running as long as `yield` keeps returning `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Since `List.All` returns an iterator, we can use it
	// in a regular `range` loop.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Packages like [slices](https://pkg.go.dev/slices) have
	// a number of useful functions to work with iterators.
	// For example, `Collect` takes any iterator and collects
	// all its values into a slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// Standard library packages now expose iterator helpers
	// too. For example, `strings.SplitSeq` iterates over parts
	// of a byte slice without first building a result slice.
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {

		// Once the loop hits `break` or an early return, the `yield` function
		// passed to the iterator will return `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
```

```bash
$ go run range-over-iterators.go
10
13
23
all: [10 13 23]
part: go
part: by
part: example
0
1
1
2
3
5
8
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