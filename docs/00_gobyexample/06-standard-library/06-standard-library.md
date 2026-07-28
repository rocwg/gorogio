当然，下面给你一份可以直接放进 `06-standard-library.md` 的 Markdown 模板，覆盖 **Sorting / Sorting by Functions / Panic / Defer / Recover / String Functions / String Formatting / Text Templates / Regular Expressions / JSON / XML / Time / Epoch / Time Formatting / Random Numbers / Number Parsing / URL Parsing / SHA256 Hashes / Base64 Encoding**。



# Go 标准库笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者到初级进阶  
> 目标：掌握 Go 常用标准库能力，并形成基本工具箱

---

## 1. Sorting

1.1 一句话理解：排序用于把数据按一定顺序排列。

1.2 示例代码

```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    nums := []int{7, 2, 4}
    slices.Sort(nums)
    fmt.Println(nums)
}
```

| 1.3 重点                    | 1.4 常见坑                     |
| --------------------------- | ------------------------------ |
| 排序是常见基础能力。        | 不清楚原地排序会修改原切片。   |
| Go 现在常用 `slices.Sort`。 | 忘记数据类型必须支持排序逻辑。 |

1.5 我的理解

排序是处理列表数据的基础操作。

---

## 2. Sorting by Functions

2.1 一句话理解：按自定义规则排序时，可以用比较函数。

2.2 示例代码

```go
package main

import (
    "cmp"
    "fmt"
    "slices"
)

func main() {
    people := []string{"Alice", "Bob", "Eve"}
    slices.SortFunc(people, func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    })
    fmt.Println(people)
}
```

| 2.3 重点                     | 2.4 常见坑                   |
| ---------------------------- | ---------------------------- |
| 自定义排序适合复杂规则。     | 比较函数返回值规则理解不清。 |
| 可以按长度、字段、权重排序。 | 自定义排序逻辑写反。         |

2.5 我的理解

排序函数让规则更灵活，但也更容易写错。

---

## 3. Panic

3.1 一句话理解：panic 表示程序遇到严重错误并中止当前流程。

3.2 示例代码

```go
// A `panic` typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.

package main

import (
	"os"
	"path/filepath"
)

func main() {

	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.
	panic("a problem")

	// A common use of panic is to abort if a function
	// returns an error value that we don't know how to
	// (or want to) handle. Here's an example of
	// `panic`king if we get an unexpected error when creating a new file.
	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}
```

| 3.3 重点                             | 3.4 常见坑               |
| ------------------------------------ | ------------------------ |
| panic 是异常终止机制，不是普通错误。 | 把 panic 当正常错误用。  |
| 常用于不可恢复问题。                 | 在业务代码里滥用 panic。 |

3.5 我的理解

panic 适合“程序已经不能继续”的场景。

---

## 4. Defer

4.1 一句话理解：defer 用于延迟执行收尾动作。

4.2 示例代码

```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

| 4.3 重点                    | 4.4 常见坑                            |
| --------------------------- | ------------------------------------- |
| defer 会在函数结束前执行。  | 不理解执行顺序。                      |
| 常用于资源释放。            | 在 defer 中依赖已变化的变量值时出错。 |
| 多个 defer 按后进先出执行。 |                                       |

4.5 我的理解

defer 是 Go 里非常优雅的收尾工具。

---

## 5. Recover

5.1 一句话理解：recover 可以从 panic 中恢复。

5.2 示例代码

```go
package main

import "fmt"

func mayPanic() {
    panic("panic happened")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    mayPanic()
}
```

| 5.3 重点                            | 5.4 常见坑                      |
| ----------------------------------- | ------------------------------- |
| recover 只能在 defer 中使用才有效。 | 以为 recover 能随便捕获 panic。 |
| 适合兜底保护，不适合日常流程控制。  | 过度依赖 recover 掩盖真正问题。 |

5.5 我的理解

recover 是最后一道保护网。

---

## 6. String Functions

6.1 一句话理解：strings 包提供常用字符串处理能力。

6.2 示例代码

```go
// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

package main

import (
	"fmt"
	s "strings"
)

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
var p = fmt.Println

func main() {

	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need to pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](https://pkg.go.dev/strings)
	// package docs.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
```

```bash
$ go run string-functions.go
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST
```

| 6.3 重点                                       | 6.4 常见坑               |
| ---------------------------------------------- | ------------------------ |
| 常见操作：查找、计数、前后缀判断、替换、切分。 | 混淆字节长度和字符长度。 |
| 字符串处理非常常用。                           | 忘记字符串不可变。       |

6.5 我的理解

字符串处理是 Go 的高频基础能力。

---

## 7. String Formatting

7.1 一句话理解：格式化用于把数据转成可读字符串。

2.2 示例代码

```go
// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.
	fmt.Printf("struct2: %+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.
	fmt.Printf("struct3: %#v\n", p)

	// To print the type of a value, use `%T`.
	fmt.Printf("type: %T\n", p)

	// Formatting booleans is straight-forward.
	fmt.Printf("bool: %t\n", true)

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.
	fmt.Printf("int: %d\n", 123)

	// This prints a binary representation.
	fmt.Printf("bin: %b\n", 14)

	// This prints the character corresponding to the
	// given integer.
	fmt.Printf("char: %c\n", 33)

	// `%x` provides hex encoding.
	fmt.Printf("hex: %x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// For basic string printing use `%s`.
	fmt.Printf("str1: %s\n", "\"string\"")

	// To double-quote strings as in Go source, use `%q`.
	fmt.Printf("str2: %q\n", "\"string\"")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.
	fmt.Printf("str3: %x\n", "hex this")

	// To print a representation of a pointer, use `%p`.
	fmt.Printf("pointer: %p\n", &p)

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// To left-justify use the `-` flag as with numbers.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
```

```bash
$ go run string-formatting.go
struct1: {1 2}
struct2: {x:1 y:2}
struct3: main.point{x:1, y:2}
type: main.point
bool: true
int: 123
bin: 1110
char: !
hex: 1c8
float1: 78.900000
float2: 1.234000e+08
float3: 1.234000E+08
str1: "string"
str2: "\"string\""
str3: 6865782074686973
pointer: 0xc0000ba000
width1: |    12|   345|
width2: |  1.20|  3.45|
width3: |1.20  |3.45  |
width4: |   foo|     b|
width5: |foo   |b     |
sprintf: a string
io: an error
```

| 7.3 重点                     | 7.4 常见坑           |
| ---------------------------- | -------------------- |
| `Printf`、`Sprintf` 很常用。 | 格式占位符用错。     |
| 格式化适合日志和输出。       | 输出结果不符合预期。 |

7.5 我的理解

格式化是 Go 输出表达能力的核心之一。

---

## 8. Text Templates

8.1 一句话理解：模板用于按规则生成文本内容。

8.2 示例代码

```go
package main

import (
    "os"
    "text/template"
)

func main() {
    t := template.Must(template.New("t").Parse("Hello {{.Name}}"))
    t.Execute(os.Stdout, map[string]string{"Name": "Go"})
}
```

| 8.3 重点                        | 8.4 常见坑                 |
| ------------------------------- | -------------------------- |
| 模板适合生成 HTML、配置、文本。 | 模板语法和 Go 语法混淆。   |
| 数据和展示逻辑分离。            | 数据结构与模板字段不匹配。 |

8.5 我的理解

模板适合“内容拼装”，不是复杂逻辑处理。

---

## 9. Regular Expressions

9.1 一句话理解：正则用于模式匹配和文本提取。

9.2 示例代码

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    matched, _ := regexp.MatchString(`p([a-z]+)ch`, "peach")
    fmt.Println(matched)
}
```

| 9.3 重点                   | 9.4 常见坑         |
| -------------------------- | ------------------ |
| 正则适合查找、提取、验证。 | 正则过度复杂。     |
| 语法强大但也容易写复杂。   | 忽略性能和可读性。 |

9.5 我的理解

正则要“够用就好”，别过度设计。

---

## 10. JSON

10.1 一句话理解：JSON 是 Go 中最常用的数据交换格式之一。

2.2 示例代码

```go
package main

import (
    "encoding/json"
    "fmt"
)

type response struct {
    Page int    `json:"page"`
    Name string `json:"name"`
}

func main() {
    data := response{Page: 1, Name: "go"}
    b, _ := json.Marshal(data)
    fmt.Println(string(b))
}
```

| 10.3 重点                         | 10.4 常见坑                    |
| --------------------------------- | ------------------------------ |
| `Marshal` 和 `Unmarshal` 是核心。 | 标签写错。                     |
| struct tag 决定 JSON 字段名。     | 字段导出性不够导致序列化失败。 |
| 常用于接口、配置、存储。          | 类型不匹配。                   |

10.5 我的理解

JSON 是 Go 后端开发必须掌握的基础。

---

## 11. XML

11.1 一句话理解：XML 也是一种数据交换格式，结构更传统。

2.2 示例代码

```go
package main

import (
    "encoding/xml"
    "fmt"
)

type plant struct {
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
}

func main() {
    p := plant{Id: 27, Name: "coffee"}
    b, _ := xml.MarshalIndent(p, "", "  ")
    fmt.Println(string(b))
}
```

| 11.3 重点                    | 11.4 常见坑                  |
| ---------------------------- | ---------------------------- |
| Go 也支持 XML 编解码。       | 字段标签规则复杂。           |
| 适合遗留系统、部分协议场景。 | 与 JSON 混用时格式理解混乱。 |

11.5 我的理解

XML 更多是兼容性工具，而不是日常首选。

---

## 12. Time

12.1 一句话理解：time 包用于处理时间和日期。

2.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)
}
```

| 12.3 重点                    | 12.4 常见坑            |
| ---------------------------- | ---------------------- |
| 处理时间、日期、间隔都靠它。 | 时区问题。             |
| 常与定时、超时、日志配合。   | 时间格式和解析不一致。 |

12.5 我的理解

时间处理是任何业务系统都绕不开的内容。

---

## 13. Epoch

13.1 一句话理解：Epoch 通常指 Unix 时间戳。

2.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println(time.Now().Unix())
}
```

| 13.3 重点                        | 13.4 常见坑          |
| -------------------------------- | -------------------- |
| 时间戳是很多系统交互的标准格式。 | 秒、毫秒、纳秒混淆。 |
| 常用于存储、比较、日志。         | 与本地时间转换出错。 |

13.5 我的理解

时间戳很实用，但单位一定要特别清楚。

---

## 14. Time Formatting / Parsing

14.1 一句话理解：可以把时间转成字符串，也可以把字符串解析成时间。

2.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    s := t.Format("2006-01-02 15:04:05")
    fmt.Println(s)
}
```

| 14.3 重点                                 | 14.4 常见坑                |
| ----------------------------------------- | -------------------------- |
| Go 的时间格式很特别，模板是固定参考时间。 | 记错参考时间格式。         |
| 格式化和解析必须一致。                    | 格式串不一致导致解析失败。 |

14.5 我的理解

这部分很重要，尤其做日志和配置时。

---

## 15. Random Numbers

15.1 一句话理解：随机数用于生成不可预测的值。

2.2 示例代码

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println(rand.Intn(100))
}
```

| 15.3 重点                        | 15.4 常见坑            |
| -------------------------------- | ---------------------- |
| `math/rand` 适合普通随机。       | 没有初始化随机种子。   |
| 安全敏感场景要用更强的随机来源。 | 把普通随机当安全随机。 |

15.5 我的理解

随机数常见，但用途要分清。

---

## 16. Number Parsing

16.1 一句话理解：把字符串转成数字是常见输入处理任务。

2.2 示例代码

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    n, _ := strconv.Atoi("123")
    fmt.Println(n)
}
```

| 16.3 重点                               | 16.4 常见坑    |
| --------------------------------------- | -------------- |
| `strconv` 是数字转换常用包。            | 忽略转换错误。 |
| 处理命令行输入、配置、HTTP 参数很常见。 | 不同进制混淆。 |

16.5 我的理解

输入处理里数字解析非常常见。

---

## 17. URL Parsing

17.1 一句话理解：URL parsing 用于拆解和处理链接。

2.2 示例代码

```go
package main

import (
    "fmt"
    "net/url"
)

func main() {
    u, _ := url.Parse("https://example.com/path?q=go")
    fmt.Println(u.Scheme)
    fmt.Println(u.Host)
    fmt.Println(u.Query().Get("q"))
}
```

| 17.3 重点                              | 17.4 常见坑          |
| -------------------------------------- | -------------------- |
| URL 由协议、主机、路径、查询参数组成。 | 手动字符串拼接 URL。 |
| 常用于网络请求和路由处理。             | 忽略编码问题。       |

17.5 我的理解

URL 解析是做网络相关功能的基础。

---

## 18. SHA256 Hashes

18.1 一句话理解：SHA256 用于生成固定长度摘要。

2.2 示例代码

```go
package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    s := "hello"
    sum := sha256.Sum256([]byte(s))
    fmt.Printf("%x\n", sum)
}
```

| 18.3 重点                      | 18.4 常见坑            |
| ------------------------------ | ---------------------- |
| 常用于校验、签名、指纹。       | 以为哈希可以逆向还原。 |
| 结果是摘要，不是加密后的原文。 | 把哈希和加密混淆。     |

18.5 我的理解

哈希是“固定指纹”，不是“保密工具”。

---

## 19. Base64 Encoding

19.1 一句话理解：Base64 用于把二进制数据编码成文本。

2.2 示例代码

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    data := "abc123!?$*&()'-=@~"
    s := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(s)
}
```

| 19.3 重点                    | 19.4 常见坑                    |
| ---------------------------- | ------------------------------ |
| Base64 是编码，不是加密。    | 把 Base64 当安全措施。         |
| 适合在文本协议中传输二进制。 | 混淆标准编码和 URL-safe 编码。 |

19.5 我的理解

Base64 只是“转文本”，不是安全保护。

---

## 总结

### 建议记忆顺序
1. Sorting
2. Defer
3. Panic / Recover
4. Strings / Formatting
5. JSON / XML
6. Time / Epoch / Parsing
7. URL / Hash / Base64

### 你要重点记住的三件事
- `defer` 很常用，`panic/recover` 了解边界即可。
- JSON、time、strings 是日常高频。
- Base64 和 hash 都不是加密。

## 学习建议

这组内容建议你按“高频优先”来学：
先学 `strings`、`fmt`、`json`、`time`、`strconv`、`url`，这些在真实项目里最常见。
像 XML、正则、SHA256、Base64 则可以放在后面按需深入。

如果你愿意，我下一步可以继续帮你写 **`07-files-network.md`**。