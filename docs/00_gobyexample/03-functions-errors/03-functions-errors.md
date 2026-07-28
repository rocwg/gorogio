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
// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won't
	// automatically return the value of the last
	// expression.
	return a + b
}

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
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
// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

package main

import "fmt"

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	_, c := vals()
	fmt.Println(c)
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

3.1 一句话理解：可变参数函数可以接收任意数量的参数。

3.2 示例代码

```go
// [_Variadic functions_](https://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

import "fmt"

// Here's a function that will take an arbitrary number
// of `int`s as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// Within the function, the type of `nums` is
	// equivalent to `[]int`. We can call `len(nums)`,
	// iterate over it with `range`, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Variadic functions can be called in the usual way
	// with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

```bash
$ go run variadic-functions.go 
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10

# Another key aspect of functions in Go is their ability
# to form closures, which we'll look at next.
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

4.1 一句话理解：闭包是“函数 + 它引用的外部变量”。

4.2 示例代码

```go
// Go supports [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
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

5.1 一句话理解：递归是函数调用自己，用来处理可分解的问题。

5.2 示例代码

```go
// Go supports
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// Here's a classic example.

package main

import "fmt"

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Anonymous functions can also be recursive, but this requires
	// explicitly declaring a variable with `var` to store
	// the function before it's defined.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Since `fib` was previously declared in `main`, Go
		// knows which function to call with `fib` here.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
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

6.1 一句话理解：Go 用显式返回 `error` 来表示错误。

6.2 示例代码

```go
// In Go it's idiomatic to communicate errors via an
// explicit, separate return value. This contrasts with
// the exceptions used in languages like Java, Python and
// Ruby and the overloaded single result / error value
// sometimes used in C. Go's approach makes it easy to
// see which functions return errors and to handle them
// using the same language constructs employed for other,
// non-error tasks.
//
// See the documentation of the [errors package](https://pkg.go.dev/errors)
// and [this blog post](https://go.dev/blog/go1.13-errors) for additional
// details.

package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and
// have type `error`, a built-in interface.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` constructs a basic `error` value
		// with the given error message.
		return -1, errors.New("can't work with 42")
	}

	// A `nil` value in the error position indicates that
	// there was no error.
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to
// signify a specific error condition.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// We can wrap errors with higher-level errors to add
		// context. The simplest way to do this is with the
		// `%w` verb in `fmt.Errorf`. Wrapped errors
		// create a logical chain (A wraps B, which wraps C, etc.)
		// that can be queried with functions like `errors.Is`
		// and `errors.AsType`.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// It's idiomatic to use an inline error check in the `if`
		// line.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` checks that a given error (or any error in its chain)
			// matches a specific error value. This is especially useful with wrapped or
			// nested errors, allowing you to identify specific error types or sentinel
			// errors in a chain of errors.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
```

```bash
$ go run errors.go
f worked: 10
f failed: can't work with 42
Tea is ready!
Tea is ready!
We should buy new tea!
Tea is ready!
Now it is dark.
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

7.1 一句话理解：自定义错误可以携带更有意义的信息。

7.2 示例代码

```go
// It's possible to define custom error types by
// implementing the `Error()` method on them. Here's a
// variant on the example above that uses a custom type
// to explicitly represent an argument error.

package main

import (
	"errors"
	"fmt"
)

// A custom error type usually has the suffix "Error".
type argError struct {
	arg     int
	message string
}

// Adding this `Error` method makes `argError` implement
// the `error` interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Return our custom error.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.AsType` is a more advanced version of `errors.Is`.
	// It checks that a given error (or any error in its chain)
	// matches a specific error type and converts to a value
	// of that type, also returning `true`. If there's no match, the
	// second return value is `false`.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
```

```bash
$ go run custom-errors.go
42
can't work with it
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