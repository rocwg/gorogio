# Section: 错误处理（Errors）

**【英文原文】**

> **Errors**
>
> Library routines must often return some sort of error indication to the caller. As mentioned earlier, Go's multivalue return makes it easy to return a detailed error description alongside the normal return value. It is good style to use this feature to provide detailed error information. For example, as we'll see, `os.Open` doesn't just return a `nil` pointer on failure, it also returns an error value that describes what went wrong.
>
> By convention, errors have type `error`, a simple built-in interface.
>
> ```go
> type error interface {
>     Error() string
> }
> ```
>
> A library writer is free to implement this interface with a richer model under the covers, making it possible not only to see the error but also to provide some context. As mentioned, alongside the usual `*os.File` return value, `os.Open` also returns an error value. If the file is opened successfully, the error will be `nil`, but when there is a problem, it will hold an `os.PathError`:
>
> ```go
> // PathError records an error and the operation and
> // file path that caused it.
> type PathError struct {
>     Op string    // "open", "unlink", etc.
>     Path string  // The associated file.
>     Err error    // Returned by the system call.
> }
> 
> func (e *PathError) Error() string {
>     return e.Op + " " + e.Path + ": " + e.Err.Error()
> }
> ```
>
> `PathError`'s `Error` generates a string like this:
>
> Plaintext
>
> ```
> open /etc/passwx: no such file or directory
> ```
>
> Such an error, which includes the problematic file name, the operation, and the operating system error it triggered, is useful even if printed far from the call that caused it; it is much more informative than the plain "no such file or directory".
>
> When feasible, error strings should identify their origin, such as by having a prefix naming the operation or package that generated the error. For example, in package `image`, the string representation for a decoding error due to an unknown format is "image: unknown format".
>
> Callers that care about the precise error details can use a type switch or a type assertion to look for specific errors and extract details. For `PathErrors` this might include examining the internal `Err` field for recoverable failures.
>
> ```go
> for try := 0; try < 2; try++ {
>     file, err = os.Create(filename)
>     if err == nil {
>         return
>     }
>     if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
>         deleteTempFiles()  // Recover some space.
>         continue
>     }
>     return
> }
> ```
>
> The second `if` statement here is another type assertion. If it fails, `ok` will be false, and `e` will be `nil`. If it succeeds, `ok` will be true, which means the error was of type `*os.PathError`, and then so is `e`, which we can examine for more information about the error.

**【精准逐字翻译】**

**错误处理（Errors）**

库函数通常必须向调用者返回某种错误指示。正如前文所述，Go 的多返回值特性使得在返回正常返回值的同时附带详细的错误描述变得十分容易。利用此特性提供详细的错误信息是一种良好的编程风格。例如，我们将看到，`os.Open` 在失败时不仅返回一个 `nil` 指针，它还会返回一个描述出了什么问题的 `error` 值。

按照惯例，错误具有 `error` 类型，这是一个简单的内置接口：

```go
type error interface {
    Error() string
}
```

库编写者可以自由地在底层用更丰富的模型来实现这个接口，这不仅可以看到错误，还能提供上下文信息。如前所述，除了常规的 `*os.File` 返回值外，`os.Open` 还会返回一个错误值。如果成功打开文件，错误将为 `nil`，但当出现问题时，它将保存一个 `os.PathError`：

```go
// PathError 记录了一个错误以及引发该错误的指示操作和文件路径。
type PathError struct {
    Op string    // "open", "unlink" 等。
    Path string  // 关联的文件。
    Err error    // 系统调用返回的错误。
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

`PathError` 的 `Error` 方法会生成如下形式的字符串：

Plaintext

```
open /etc/passwx: no such file or directory
```

这样一个包含了问题文件名、具体操作以及触发的操作系统底层的错误，即使在远离引发调用的地方打印出来也很有用；它比单纯的 "no such file or directory" 包含的信息丰富得多。

在可行的情况下，错误字符串应该标识其来源，例如使用前缀命名生成错误的某种操作或包名。例如，在 `image` 包中，因未知格式导致解码错误的字符串表示为 "image: unknown format"。

关心精确错误细节的调用者可以使用类型选择（type switch）或类型断言（type assertion）来寻找特定错误并提取细节。对于 `PathErrors`，这可能包括检查内部的 `Err` 字段以处理可恢复的失败：

```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // 回收一些空间。
        continue
    }
    return
}
```

这里的第二个 `if` 语句是另一个类型断言。如果断言失败，`ok` 将为 false，`e` 将为 `nil`。如果成功，`ok` 将为 true，这意味着错误的类型是 `*os.PathError`，因此 `e` 也是该类型，我们可以检查它以获取有关错误的更多信息。

**【核心解析与精髓】**

- **显式错误返回值（Errors as Values）**：Go 坚决放弃传统语言的 `try-catch` 异常抛出机制，将 Error 视作一等公民（Value）显式返回。
- **带上下文的结构化 Error**：自定义类型实现 `error` 接口（如 `os.PathError`），不仅携带字符串，还保留结构化字段（`Op`、`Path`、`Err`），支持调用方进行精细化恢复决策。

### 现代 Go 演进与工程实践对比

#### 1. 错误包装（Error Wrapping）与 `errors.Is` / `errors.As`

- **经典 Effective Go**：

  依赖强类型断言 `err.(*os.PathError)` 来提取底层错误。如果错误在调用链中被其他错误嵌套包装（如 `fmt.Errorf("do sth: %w", err)`），**直接的类型断言将彻底失效**。

- **现代 Go 标杆（Go 1.13+ 规范）**：

  引入 `%w` 动词以及 `errors.Is` 和 `errors.As` 标准库函数，支持在错误树中递归解包：

```go
// 现代 Go：解包判定值相等的 sentinel error
if errors.Is(err, syscall.ENOSPC) {
    deleteTempFiles()
}

// 现代 Go：解包提取结构化 Target Error 字段
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    if errors.Is(pathErr.Err, syscall.ENOSPC) {
        deleteTempFiles()
    }
}
```

#### 2. 自定义 Error 类型的 `Unwrap` 方法实现

为了与现代 Go 错误链兼容，结构体形式的 Error 应当实现 `Unwrap` 方法：

```go
// 现代 Go 自定义错误范式
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}

// 现代规范：实现 Unwrap() 暴露内部嵌套错误
func (e *PathError) Unwrap() error {
    return e.Err
}
```

#### 3. 错误构建与上下文追加

- **经典 Effective Go**：推荐通过格式化前缀标记模块名。
- **现代 Go 范式**：
  - **轻量级上下文**：使用 `fmt.Errorf("uam: fetch user %d: %w", id, err)` 保持调用栈的可读性。
  - **多错误聚合（Go 1.20+）**：使用 `errors.Join(err1, err2)` 将并发任务中的多个错误聚合成单一 `error` 返回。

---

# Section: 恐慌（Panic）

**【英文原文】**

> **Panic**
>
> The usual way to report an error to a caller is to return an `error` as an extra return value. The canonical `Read` method is a well-known instance; it returns a byte count and an `error`. But what if the error is unrecoverable? Sometimes the program simply cannot continue.
>
> For this purpose, there is a built-in function `panic`that in effect creates a run-time error that will stop the program (but see the next section). The function takes a single argument of arbitrary type—often a string—to be printed as the program dies. It's also a way to indicate that something impossible has happened, such as exiting an infinite loop.
>
> ```go
> // A toy implementation of cube root using Newton's method.
> func CubeRoot(x float64) float64 {
>     z := x/3   // Arbitrary initial value
>     for i := 0; i < 1e6; i++ {
>         prevz := z
>         z -= (z*z*z-x) / (3*z*z)
>         if veryClose(z, prevz) {
>             return z
>         }
>     }
>     // A million iterations has not converged; something is wrong.
>     panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
> }
> ```
>
> This is only an example but real library functions should avoid `panic`. If the problem can be masked or worked around, it's always better to let things continue to run rather than taking down the whole program. One possible counterexample is during initialization: if the library truly cannot set itself up, it might be reasonable to panic, so to speak.
>
> ```go
> var user = os.Getenv("USER")
> 
> func init() {
>     if user == "" {
>         panic("no value for $USER")
>     }
> }
> ```

**【精准逐字翻译】**

**恐慌（Panic）**

向调用者报告错误的通常方式是将 `error` 作为额外的返回值返回。规范的 `Read` 方法就是一个著名的例子；它返回一个字节数和一个 `error`。但如果错误是不可恢复的呢？有时程序根本无法继续运行。

为此，内置函数 `panic` 实际上会创建一个运行时错误，该错误将停止程序（但请参见下一节）。该函数接收一个任意类型的单参数——通常是一个字符串——在程序终止时打印。这也是一种表明发生了不可能发生之事件的方式，例如退出了一个无限循环。

```go
// 使用牛顿法求解立方根的玩具实现。
func CubeRoot(x float64) float64 {
    z := x/3   // 任意初始值
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // 循环一百万次仍未收敛；出了问题。
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
```

这只是一个示例，但真实的库函数应该避免使用 `panic`。如果问题可以被掩盖或规避，让程序继续运行总是比终止整个程序更好。一个可能的反例是在初始化期间：如果库确实无法完成自我设置，可以说，抛出 `panic` 是合理的。

```go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

**【核心解析与精髓】**

- **Panic 的边界定位**：`panic` 不是 Go 的 `throw Exception`！它仅用于不可恢复的灾难性错误（致命逻辑错误、不可越界的断言错误、致命初始化失败）。
- **Fail-Fast 原则**：在程序启动阶段（如 `init()` 或 `main()`），若缺失关键环境变量或数据库配置，立即使程序崩溃（Fail-Fast）优于带着坏状态带病运行。

### 现代 Go 演进与工程实践对比

#### 1. 命名习惯与标准库中的 `Must` 模式

- **经典 Effective Go**：示例中在 `init()` 函数内部显式使用 `panic`。

- **现代 Go 标杆范式**：

  在现代 Go 基础库与架构设计中，如果一个函数可能会 panic，标准命名惯例是以 **`Must` 为前缀**，专门用于在 package 初始化或全局变量声明时快速校验：

```go
// 现代 Go 标杆：标准库与开源库常用的 Must 范式
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
var tmpl = template.Must(template.New("tpl").Parse("..."))
```

#### 2. 边界断言与无异常设计原则

- **经典 Effective Go**：展示了计算不收敛时抛出 `panic`。
- **现代工程原则**：
  - **业务逻辑禁止滥用 Panic**：业务 API 必须返回 `error`（或包装后的业务错误），由上层中间件决定降级还是拒绝服务，决不能让单个非法请求触发 `panic` 导致进程终止。
  - **跨 Goroutine 的 Panic 绝密陷阱**：在主 Goroutine 中使用 `recover()` **无法捕获**由子 Goroutine（`go func()`）内抛出的未处理 `panic`。一旦子 Goroutine panic，整个 Go 进程（Process）会直接崩溃崩溃退出。

#### 3. 现代 Web / RPC 框架中的 Panic 隔离保护

在现代高并发服务（如 HTTP/gRPC 服务、微服务）中，全局 Panic 防护是必备的 Recovery 中间件模式：

```go
// 现代 Go Web 中间件 Recovery 防护机制
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // 记录堆栈信息，避免进程宕机
                log.Printf("panic recovered: %v\n%s", err, debug.Stack())
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

---

# Section: 恢复（Recover）

**【英文原文】**

> **Recover**
>
> When `panic` is called, including implicitly for run-time errors such as indexing a slice out of bounds or failing a type assertion, it immediately stops execution of the current function and begins unwinding the stack of the goroutine, running any deferred functions along the way. If that unwinding reaches the top of the goroutine's stack, the program dies. However, it is possible to use the built-in function `recover` to regain control of the goroutine and resume normal execution.
>
> A call to `recover` stops the unwinding and returns the argument passed to `panic`. Because the only code that runs while unwinding is inside deferred functions, `recover` is only useful inside deferred functions.
>
> One application of `recover` is to shut down a failing goroutine inside a server without killing the other executing goroutines.
>
> ```go
> func server(workChan <-chan *Work) {
>     for work := range workChan {
>         go safelyDo(work)
>     }
> }
> 
> func safelyDo(work *Work) {
>     defer func() {
>         if err := recover(); err != nil {
>             log.Println("work failed:", err)
>         }
>     }()
>     do(work)
> }
> ```
>
> In this example, if `do(work)` panics, the result will be logged and the goroutine will exit cleanly without disturbing the others. There's no need to do anything else in the deferred closure; calling `recover` handles the condition completely.
>
> Because `recover` always returns `nil` unless called directly from a deferred function, deferred code can call library routines that themselves use `panic` and `recover` without failing. As an example, the deferred function in `safelyDo` might call a logging function before calling `recover`, and that logging code would run unaffected by the panicking state.
>
> With our recovery pattern in place, the `do` function (and anything it calls) can get out of any bad situation cleanly by calling `panic`. We can use that idea to simplify error handling in complex software. Let's look at an idealized version of a `regexp` package, which reports parsing errors by calling `panic` with a local error type. Here's the definition of `Error`, an `error` method, and the `Compile` function.
>
> ```go
> // Error is the type of a parse error; it satisfies the error interface.
> type Error string
> func (e Error) Error() string {
>     return string(e)
> }
> 
> // error is a method of *Regexp that reports parsing errors by
> // panicking with an Error.
> func (regexp *Regexp) error(err string) {
>     panic(Error(err))
> }
> 
> // Compile returns a parsed representation of the regular expression.
> func Compile(str string) (regexp *Regexp, err error) {
>     regexp = new(Regexp)
>     // doParse will panic if there is a parse error.
>     defer func() {
>         if e := recover(); e != nil {
>             regexp = nil    // Clear return value.
>             err = e.(Error) // Will re-panic if not a parse error.
>         }
>     }()
>     return regexp.doParse(str), nil
> }
> ```
>
> If `doParse` panics, the recovery block will set the return value to `nil`—deferred functions can modify named return values. It will then check, in the assignment to `err`, that the problem was a parse error by asserting that it has the local type `Error`. If it does not, the type assertion will fail, causing a run-time error that continues the stack unwinding as though nothing had interrupted it. This check means that if something unexpected happens, such as an index out of bounds, the code will fail even though we are using `panic` and `recover` to handle `parse` errors.
>
> With error handling in place, the `error` method (because it's a method bound to a type, it's fine, even natural, for it to have the same name as the builtin `error` type) makes it easy to report parse errors without worrying about unwinding the parse stack by hand:
>
> ```go
> if pos == 0 {
>     re.error("'*' illegal at start of expression")
> }
> ```
>
> Useful though this pattern is, it should be used only within a package. Parse turns its internal panic calls into error values; it does not expose panics to its client. That is a good rule to follow.
>
> By the way, this re-panic idiom changes the panic value if an actual error occurs. However, both the original and new failures will be presented in the crash report, so the root cause of the problem will still be visible. Thus this simple re-panic approach is usually sufficient—it's a crash after all—but if you want to display only the original value, you can write a little more code to filter unexpected problems and re-panic with the original error. That's left as an exercise for the reader.

**【精准逐字翻译】**

**恢复（Recover）**

当 `panic` 被调用时（包括因切片越界访问或类型断言失败等引发的隐式运行时错误），它会立即停止当前函数的执行，并开始解开（unwind）该 goroutine 的调用栈，在此过程中运行任何被延迟的函数（`defer`）。如果这种解栈过程到达了该 goroutine 调用栈的顶部，程序就会终止崩溃。然而，可以使用内置函数 `recover` 来重新夺回对 goroutine 的控制权并恢复正常执行。

对 `recover` 的调用会停止栈解开过程，并返回传递给 `panic` 的参数。因为在解栈过程中唯一能运行的代码是在延迟函数内部，所以 `recover` 仅在延迟函数内部有效。

`recover` 的一个应用是在服务器内部停止一个发生故障的 goroutine，而不会杀死其他正在执行的 goroutine。

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

在这个例子中，如果 `do(work)` 发生 panic，其结果将被记录，且该 goroutine 将干净地退出，而不会干扰其他 goroutine。无需在延迟闭包中做其他任何事情；调用 `recover` 就能完全处理这种情况。

由于 `recover` 除非直接从延迟函数中调用，否则总是返回 `nil`，因此延迟代码可以调用其本身使用了 `panic` 和 `recover` 的库例程而不会出错。例如，`safelyDo` 中的延迟函数可以在调用 `recover` 之前调用日志记录函数，且该日志代码的运行不受 panic 状态的影响。

借助我们的恢复模式，`do` 函数（以及它调用的任何内容）都可以通过调用 `panic` 来干净地脱离任何糟糕的状况。我们可以使用这种思想来简化复杂软件中的错误处理。让我们来看一个理想化的 `regexp` 包版本，它通过使用本地错误类型调用 `panic` 来报告解析错误。以下是 `Error` 的定义、一个 `error` 方法以及 `Compile` 函数。

```go
// Error 是解析错误的类型；它实现了 error 接口。
type Error string
func (e Error) Error() string {
    return string(e)
}

// error 是 *Regexp 的一个方法，通过抛出带有 Error 的 panic 来报告解析错误。
func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

// Compile 返回正则表达式的解析表示。
func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    // 如果存在解析错误，doParse 将会 panic。
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // 清空返回值。
            err = e.(Error) // 如果不是解析错误，类型断言失败将重新触发 panic。
        }
    }()
    return regexp.doParse(str), nil
}
```

如果 `doParse` 发生 panic，恢复块将把返回值设置为 `nil`——延迟函数可以修改具名返回值。然后，它将在对 `err` 的赋值中，通过断言其具有本地类型 `Error` 来检查该问题是否为解析错误。如果不是，类型断言将失败，从而引发运行时错误，继续进行栈解开，就好像没有任何东西中断过它一样。这种检查意味着，如果发生了意想不到的事情（例如索引越界），代码即使在使用 `panic` 和 `recover` 来处理解析错误，也依然会发生崩溃。

有了这种错误处理机制，`error` 方法（因为它是一个绑定到类型的结构体方法，所以与内置 `error` 类型同名也是完全可以的，甚至很自然）使得报告解析错误变得非常容易，而无需手动去解开解析栈：

```go
if pos == 0 {
    re.error("'*' illegal at start of expression")
}
```

尽管这种模式非常有用，但**它应该仅在包内部使用**。`Parse` 将其内部的 `panic` 调用转化为 `error` 值返回；它绝不向客户端暴露 `panic`。这是一个应当遵循的好规则。

顺便提一句，这种“重新 panic”（re-panic）的惯用手法在实际发生错误时会更改 panic 的值。然而，原始故障和新的故障都会呈现在崩溃报告中，因此问题的根本原因依然可见。因此，这种简单的重新 panic 方法通常就足够了——毕竟这终究是一次崩溃——但如果你只想显示原始值，可以编写稍微多一点的代码来过滤未预期的错误，并使用原始错误重新 panic。这留给读者作为练习。

**【核心解析与精髓】**

- **栈解开（Stack Unwinding）与 Recover 生效域**：`recover()` 必须且只能在 `defer` 的顶层闭包中直接调用。
- **边界隔离原则（Boundary Encapsulation）**：包内部可以用 `panic/recover` 简化多层深度递归（如 AST 解析器）的错误返回，但**对外暴露的 API 必须将 Panic 捕获并转译为普通的 `error` 值**。
- **带类型的 Re-Panic 选择性恢复**：只捕获内部自定义的 `Error` 类型，如果是真正的运行时致命错误（如空指针解引用 `nil pointer` 或切片越界 `out of bounds`），通过类型断言失败重新触发崩溃，确保系统不掩盖真正的 Bug。

### 现代 Go 演进与工程实践对比

#### 1. 显式 Re-panic 范式演进

- **经典 Effective Go**：利用类型断言失败 `err = e.(Error)` 的隐式 Panic 实现 Re-panic。

- **现代 Go 标杆范式**：

  现代 Go 强调代码意图的**显式表达**（Explicit is better than implicit），不再推荐依赖“类型断言失败触发隐式 Panic”的技巧，而是显式判断并原样重新抛出 `panic(e)`：

```go
// 现代 Go 标杆：显式 Re-panic 写法
defer func() {
    if r := recover(); r != nil {
        if parseErr, ok := r.(Error); ok {
            regexp = nil
            err = parseErr
        } else {
            // 显式重新抛出原始 panic，保留原始堆栈上下文
            panic(r)
        }
    }
}()
```

#### 2. 现代 Go 堆栈打印与 `debug.Stack()`

- **经典 Effective Go**：简单在 `recover()` 处打印 `err`。

- **现代 Go 运维与诊断标杆**：

  在微服务、高并发 Web 框架（如 Gin/Fiber/Echo）的 Recovery 中间件中，单纯打印 `recover()` 返回的错误对象往往无法准确定位代码出错行。现代 Go 必定搭配 `runtime/debug` 包提取完整 Goroutine 堆栈：

```go
// 现代 Go：带 Full Stack Trace 的 Recovery 机制
defer func() {
    if r := recover(); r != nil {
        stackTrace := string(debug.Stack())
        log.Printf("[PANIC RECOVERED] %v\nStack:\n%s", r, stackTrace)
        // 转译为结构化 HTTP 500 Error 返回
        err = fmt.Errorf("internal server error")
    }
}()
```

#### 3. 跨 Goroutine 边界的 Recover 铁律

现代 Go 工程中最重要的并发安全规范之一：**`recover()` 无法跨越 Goroutine 边界**。

```go
// ❌ 错误做法：主 Goroutine 的 defer 无法捕获子 Goroutine 的 panic
func MainRunner() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in main") // 无法执行到此处！
        }
    }()

    go func() {
        panic("boom!") // 导致整个进程崩溃退出
    }()
    
    time.Sleep(time.Second)
}

// ✅ 现代 Go 正确做法：每个安全派生的 Goroutine 都必须自带 Recover
func SafeGo(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("goroutine panic: %v\n%s", r, debug.Stack())
            }
        }()
        fn()
    }()
}
```

---

