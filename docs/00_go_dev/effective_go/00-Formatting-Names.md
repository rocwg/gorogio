## Commentary（注释规范）

### 【英文原文】

> Go provides C-style `/* */` block comments and C++-style `//` line comments. Line comments are the norm; block comments appear mostly as package comments, but are useful within an expression or to disable large swaths of code.
>
> Comments that aid package documentation—and every package should have a package comment—should precede the package clause. For multi-file packages, the package comment only needs to be in one file, and any one will do. The package comment should introduce the package and provide information relevant to the package as a whole. It will appear first on the godoc page and should set up the detailed documentation that follows.
>
> ```Go
> /*
> Package regexp implements a simple library for regular expressions.
> 
> The syntax of the regular expressions accepted is:
> 
>     regexp:
>         concat '||' concat
> ...
> */
> package regexp
> ```
>
> If the package is simple, the package comment can be brief.
>
> ```Go
> // Package path implements utility routines for manipulating slash-separated
> // paths.
> package path
> ```
>
> Comments do not need extra formatting such as banners of stars. The output may not even be presented in a fixed-width font, so don't depend on spacing for alignment—`gofmt`, like `godoc`, takes care of that. The comments are uninterpreted text, so HTML and other annotations such as `_this_` will reproduce *verbatim* and should not be used. One quirk that `gofmt` does do is to reformat comments that look like Go code, particularly when the lines start with a tab.
>
> Inside a package, any comment immediately preceding a top-level declaration serves as a doc comment for that declaration. Every exported (capitalized) name in a program should have a doc comment.
>
> Doc comments work best as complete sentences, which allow a wide variety of automated presentations. The first sentence should be a one-sentence summary that starts with the name being declared.
>
> ```Go
> // Compile parses a regular expression and returns, if successful, a Regexp
> // object that can be used to match against text.
> func Compile(str string) (*Regexp, error) {
> ```
>
> If every doc comment begins with the name of the item it describes, the output of `godoc` can usefully be run through `grep`. Imagine you couldn't remember the name "Compile" but were looking for the parsing function for regular expressions, so you ran the command,
>
> ```sh
> $ godoc regexp | grep parse
> ```
>
> If all the doc comments in the package began, "This function...", `grep` wouldn't help you recall the name. But because the package starts each doc comment with the name, you'd see something like this, which recalls the word you're looking for.
>
> ```sh
> $ godoc regexp | grep parse
>     Compile parses a regular expression and returns, if successful, a Regexp
>     MustCompile is like Compile but panics if the expression cannot be parsed.
> $
> ```
>
> Go's declaration syntax allows grouping of declarations. A single doc comment can introduce a group of related constants or variables. Since the whole declaration is presented, such a comment can often be perfunctory.
>
> ```go
> // Error codes returned by failures to parse an expression.
> var (
>     ErrInternal      = errors.New("regexp: internal error")
>     ErrUnmatchedLpar = errors.New("regexp: unclosed '('")
>     ErrBadRune       = errors.New("regexp: bad unicode char")
> )
> ```
>
> Grouping can also indicate relationships between items, such as the fact that a set of variables is protected by a mutex.
>
> ```go
> var (
>     count LockableInt
>     mu    sync.Mutex
> )
> ```

### 【精准逐字翻译】

Go 提供了 C 风格的 `/* */` 块注释和 C++ 风格的 `//` 行注释。行注释是常规规范；块注释主要作为包注释出现，但在表达式内部或禁用大片代码时很有用。

助于包文档的注释——并且每个包都应该有一个包注释——应当置于 `package` 子句之前。对于由多个文件组成的包，包注释只需要出现在其中一个文件中，任何一个文件都可以。包注释应当介绍该包并提供与包整体相关的信息。它将首先出现在 `godoc` 页面上，并应为随后的详细文档做铺垫。

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concat '||' concat
...
*/
package regexp
```

如果包比较简单，包注释可以很简短。

```go
// Package path implements utility routines for manipulating slash-separated
// paths.
package path
```

注释不需要额外的格式化，比如由星号组成的横幅。输出甚至可能不会以等宽字体展示，因此不要依赖空格来进行对齐——`gofmt` 和 `godoc` 一样会处理好这些。注释是未经解释的纯文本，因此 HTML 以及类似 `_this_` 的其他标注将会被原样（verbatim）呈现，不应当使用。`gofmt` 确实会做的一个特殊处理是重新格式化看起来像 Go 代码的注释，特别是当行以 Tab 开头时。

在一个包内部，任何紧邻顶级声明之前的注释都会作为该声明的文档注释（doc comment）。程序中每一个导出（大写字母开头）的名称都应当有一个文档注释。

文档注释最好写成完整的句子，这允许各种各样的自动化展示。第一句话应当是以被声明的名称开头的单句总结。

```go
// Compile parses a regular expression and returns, if successful, a Regexp
// object that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

如果每个文档注释都以它所描述的项的名称开头，那么 `godoc` 的输出就可以方便地通过 `grep` 过滤。假设你想不起“Compile”这个名字，但正在寻找正则表达式的解析函数，于是你运行了命令：

```sh
$ godoc regexp | grep parse
```

如果包里的所有文档注释都以“This function...”开头，`grep` 就无法帮助你回忆起这个名字。但因为包中的每个文档注释都以名称开头，你会看到类似这样的结果，这能帮你回忆起你在寻找的单词。

```sh
$ godoc regexp | grep parse
    Compile parses a regular expression and returns, if successful, a Regexp
    MustCompile is like Compile but panics if the expression cannot be parsed.
$
```

Go 的声明语法允许对声明进行分组。单个文档注释可以引出一组相关的常量或变量。由于展示的是整个声明，这样的注释通常可以写得比较简略。

```go
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unclosed '('")
    ErrBadRune       = errors.New("regexp: bad unicode char")
)
```

分组还可以表明项之间的关联，例如一组变量受互斥锁保护这一事实。

```go
var (
    count LockableInt
    mu    sync.Mutex
)
```

### 【专业术语与句式拆解】

1. **Doc Comment（文档注释）**：
   - **含义**：紧邻顶级声明（包、函数、类型、变量、常量）上方且无空行的注释，专门用于提取生成 API 文档。
2. **Top-level declaration（顶级声明）**：
   - **含义**：直接声明在包层级（在任何函数体外部）的全局结构，包括 `func`、`type`、`const`、`var`。
3. **Exported (capitalized) name（导出的/首字母大写的名称）**：
   - **含义**：Go 语言中的可见性规则，首字母大写即对包外可见（相当于其他语言的 `public`）。
4. **Verbatim (adv. / adj.)**：
   - **含义**：逐字地、原封不动地。
5. **Perfunctory (adj.)**：
   - **含义**：敷衍的、例行公事的。此处指因为分组代码本身语义很清晰，注释无需长篇大论，简要描述即可。

### 【4. 补充：Go 1.27 演变与现代 Go 实践】

原文编写时 Go 的文档引擎还非常原始，而在现代 Go（Go 1.19+ 引入并延续至 Go 1.27）中， Go 的 **Doc Comments 语法规则进行了重大升级**：

1. **结构化文档 Markdown 增强（Go 1.19+ 标准化）：**
   - **标题（Headings）：** 支持使用 `#` 或 `##` 标注标题（需独占一行且上下有空行），例如 `# Performance`。
   - **列表（Lists）：** 支持使用数字 `1.` 或星号 `-` / `*` 编写列表。
   - **代码块（Code Blocks）：** 以 Tab 或 4 个空格缩进的文本行会被渲染为代码块，`gofmt` 会自动格式化代码块内的合法 Go 代码。
   - **自动链接与文档引用（Doc Links）：** 可以直接使用 `[Package.Item]` 格式引用当前包或其他包的标识符，例如 `[bytes.Buffer]` 或 `[Compile]`，文档生成工具（如 `pkgsite`）会自动生成超链接。
2. **包注释规范化（`doc.go` 模式）：**
   - 对于大型或复杂的 package，现代 Go 的惯用写法是创建一个专门的 `doc.go` 文件，在该文件中只写 `/* ... */` 包注释和 `package xxx` 声明，从而将庞大的包级文档与业务代码剥离。
3. **工具链迁移（`godoc` $\rightarrow$ `go doc` / `pkgsite`）：**
   - 原文频繁提到的本地命令行工具 `godoc` 已经被轻量级的内置工具 **`go doc`**（如 `go doc regexp.Compile`）取代；而线上网页端则被官方的 **`pkgsite`** 架构取代。

---



这是因为 Go 官方在 **Go 1.19（2022 年 8 月）重构文档注释系统（Doc Comments）** 时，对 Effective Go 的 `Commentary` 章节进行了**大幅精简**。

官方将原来冗长的 `grep` 示例、格式细节等内容大幅删减，转移到了全新的专题文档 **[Go Doc Comments](https://go.dev/doc/comment?utm_source=gemini)** 中，因此官方页面只保留了您看到的这两段精炼概述。

## Commentary（注释规范 - 官方最新精简版）

### 【英文原文】

> Go provides C-style `/* */` block comments and C++-style `//` line comments. Line comments are the norm; block comments appear mostly as package comments, but are useful within an expression or to disable large swaths of code.
>
> Comments that appear before top-level declarations, with no intervening newlines, are considered to document the declaration itself. These “doc comments” are the primary documentation for a given Go package or command. For more about doc comments, see “Go Doc Comments”.

### 【精准逐字翻译】

Go 提供了 C 风格的 `/* */` 块注释和 C++ 风格的 `//` 行注释。行注释是常规规范；块注释主要作为包注释出现，但在表达式内部或禁用大片代码时很有用。

出现在顶级声明之前且中间没有空行的注释，被视为该声明本身的文档。这些“文档注释”（doc comments）是给定 Go 包或命令的主要文档。有关文档注释的更多信息，请参阅“Go Doc Comments”。

### 【专业术语与句式拆解】

1. **Intervening newlines**：
   - **含义**：介于中间的换行符（空行）。在 Go 中，如果注释和声明之间隔了一个空行，编译器/文档工具就不会将其识别为该声明的 Doc Comment。
2. **Top-level declarations**：
   - **含义**：顶级声明。指包层级的 `func`、`type`、`const`、`var`。

---



## Names（命名规范）

#### 【英文原文】

> Names are as important in Go as in any other language. They even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case. It's therefore worth spending a little time discussing naming conventions in Go programs.

#### 【精准逐字翻译】

名称在 Go 中与其他任何语言一样重要。它们甚至具有语义效应：一个名称在包外的可见性取决于它的第一个字符是否是大写。因此，花费一点时间讨论 Go 程序中的命名约定是值得的。

#### 【专业术语与句式拆解】

1. **Semantic effect（语义效应）**：
   - **含义**：名字不仅代表标识符，其大小写本身就是 Go 语法和作用域规则（`Public` vs `Private`）的一部分。

### Package names

#### 【英文原文】

> When a package is imported, the package name becomes the accessor for the contents. After
>
> ```go
> import "bytes"
> ```
>
> the importing file can talk about `bytes.Buffer`. It's helpful if everyone using the package can use the same name to refer to its contents, which implies that the package name should be good: short, concise, evocative. By convention, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps. Err on the side of brevity, since anyone using your package will be typing that name. And don't worry about collisions *a priori*. The package name is only the default name for imports; it does not need to be unique across all source code, and in the rare case of a collision the importing package can choose an alias to be used locally. In any case, confusion is rare because the file name in the import determines just which package is being used.
>
> Another convention is that the package name is the base name of its source directory; the package in `src/encoding/base64` is imported as `"encoding/base64"` but has name `base64`, not `encoding_base64` and not `encodingBase64`.
>
> The importer of a package will use the name to reference its contents, so exported names in the package can use that fact to avoid stutter. (Don't use the `import .` notation, which can simplify tests that must run outside the package they are testing, but should otherwise be avoided.) For instance, the buffered reader type in the `bufio` package is called `Reader`, not `BufReader`, because users see it as `bufio.Reader`, which is a clear, concise name. Moreover, because imported entities are always addressed with their package name, `bufio.Reader` does not collide with `io.Reader`. Similarly, the function to make new instances of `ring.Ring`—which is the definition of a ring in Go—would normally be called `NewRing`, but since `Ring` is the only type exported by the package, and since the package is called `ring`, it's called just `New`, which clients of the package see as `ring.New`. That's the framework for thinking about names.
>
> Another short example is `once.Do`; `once.Do(setup)` reads well and would not be improved by writing `once.DoOrWaitUntilDone(setup)`. Long names don't automatically make things more readable. A helpful doc comment can often be more valuable than an extra long name.

#### 【精准逐字翻译】

当一个包被导入时，包名就成为了访问其内容的入口。在

```go
import "bytes"
```

之后，导入该文件的代码就可以讨论 `bytes.Buffer`。如果每个使用该包的人都能用相同的名字来引用其内容，那将大有裨益，这意味着包名应当优秀：短小、简洁、富有启发性。按照约定，包名赋予小写、单单词的名称；不应当需要下划线或大小写混排（mixedCaps）。宁可偏向于简短，因为使用你包的每个人都会输入该名称。并且不必事先（a priori）担心命名冲突。包名只是导入时的默认名称；它不需要在所有源码中保持唯一，在极少数发生冲突的情况下，导入包可以选择一个本地使用的别名。无论如何，混乱都很罕见，因为导入中的文件名明确决定了使用的是哪个包。

另一个约定是包名是其源目录的基准名称（base name）；位于 `src/encoding/base64` 的包被导入为 `"encoding/base64"`，但名称为 `base64`，而不是 `encoding_base64`，也不是 `encodingBase64`。

包的导入者将使用该名称来引用其内容，因此包中导出的名称可以利用这一事实来避免口吃（stutter，指名称重复）。（不要使用 `import .` 记法，它可以简化必须在所测试包之外运行的测试，但在其他情况下应当避免。）例如，`bufio` 包中的缓冲读取器类型被称为 `Reader`，而不是 `BufReader`，因为使用者将其视为 `bufio.Reader`，这是一个清晰且简洁的名称。此外，因为导入的实体总是通过它们的包名来寻址，`bufio.Reader` 不会与 `io.Reader` 冲突。类似地，用于创建 `ring.Ring` 新实例的函数——这是 Go 中环形链表的定义——通常会被称为 `NewRing`，但由于 `Ring` 是该包导出的唯一类型，且由于该包被称为 `ring`，它就被简单地称为 `New`，包的客户端将其视为 `ring.New`。这就是思考命名的框架。

另一个简短的例子是 `once.Do`；`once.Do(setup)` 读起来很通顺，写成 `once.DoOrWaitUntilDone(setup)` 不会有任何改进。长名字不会自动让事物变得更具可读性。一份有帮助的文档注释往往比一个超长的名字更有价值。

#### 【专业术语与句式拆解】

1. **A priori (adv.)**：
   - **含义**：先验地、事先。
2. **Stutter (n. / v.)**：
   - **含义**：口吃、结巴。在 Go 命名学中专指**冗余的重复命名**（例如 `http.HTTPClient` 或 `user.UserService`，地道的写法应为 `http.Client` 或 `user.Service`）。
3. **`import .` (Dot import)**：
   - **含义**：点导入（把被导入包的所有导出符号直接注入当前文件作用域）。Go 官方强烈反对在正规业务代码中使用。

#### 【4. 补充：Go 1.27 演变与现代 Go 实践】

1. **包路径与 Go Modules 标准化：**
   - 原文提到的 `src/encoding/base64` 是旧 `GOPATH` 时代的目录结构。在现代 Go 中，包路径由 `go.mod` 声明的模块名（如 `[github.com/foo/bar/pkg/user](https://github.com/foo/bar/pkg/user)`）决定，但**包名（Package Name）依然恪守最后一个目录名（`user`）且为全小写单单词的铁律**。
2. **构造函数命名标准：**
   - 单一类型包用 `New()`（如 `ring.New()`）。
   - 多类型包用 `NewType()`（如 `http.NewRequest()`）。

---



没问题，那我们采用**一次解读一小段**的节奏，让信息更清晰、更容易消化。

下面是 **Names** 章节中的 **Getters（Getter 方法）** 小节：

### Getters（Getter 方法）

#### 【英文原文】

> Go doesn't provide automatic support for getters and setters. There's nothing wrong with providing getters and setters yourself, and it's often appropriate to do so, but it's neither idiomatic nor necessary to put `Get` into the getter's name. If you have a field called `owner` (lower case, unexported), the getter method should be called `Owner` (upper case, exported), not `GetOwner`. The use of upper-case names for export provides the hook to discriminate the method from the field. A setter function, if needed, will likely be called `SetOwner`. Both names read well in practice:
>
> ```go
>owner := obj.Owner()
> if owner != user {
>  obj.SetOwner(user)
> }
>    ```

#### 【精准逐字翻译】

Go 不提供对 getter 和 setter 的自动支持。你自己提供 getter 和 setter 并没有错，而且这样做通常是恰当的，但在 getter 的名称中加上 `Get` 既不符合惯用法（idiomatic），也是没有必要的。如果你有一个名为 `owner`（小写，未导出）的字段，其 getter 方法应该被称为 `Owner`（大写，已导出），而不是 `GetOwner`。使用大写字母进行导出，提供了区分方法与字段的依据（挂钩）。如果需要 setter 函数，它可能会被称为 `SetOwner`。这两个名称在实践中读起来都很顺畅：

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

#### 【专业术语与句式拆解】

1. **Idiomatic (adj.)**：
   - **含义**：符合 Go 惯例的。在 Java 中几乎强制写 `getOwner()`，但在 Go 里面写 `GetOwner()` 会被视作非地道的代码。
2. **Hook (n.)**：
   - **含义**：依据、挂钩、机制。指 Go 语言利用首字母大小写实现可见性控制的机制，刚好完美隔开了同名的私有字段（`owner`）和公有 Getter（`Owner()`）。

#### 【4. 补充：Go 1.27 演变与现代 Go 实践】

1. **绝对禁忌与例外情况：**
   - **标准规范：** Getter 绝不带 `Get` 前缀（例如 `user.Name()` 而非 `user.GetName()`）。
   - **唯一例外：** 如果涉及与 **Protobuf** / **gRPC** 或某些特定 API 契约交互时，生成的 Go 代码会自动带有 `GetXxx()` 方法。在手写业务代码时，依然坚守无 `Get` 前缀规则。

---

### Interface names（接口命名规范）

#### 【英文原文】

> By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: `Reader`, `Writer`, `Formatter`, `CloseNotifier` etc.
>
> There are a number of such names and it's honoring to respect them and the function names they capture. `Read`, `Write`, `Close`, `Flush`, `String` and so on have canonical signatures and meanings. To avoid confusion, don't give your method one of those names unless it has the same signature and meaning. Conversely, if your type implements a method with the same meaning as a method on a well-known type, give it the same name and signature; call your string-converter method `String` not `ToString`.

#### 【精准逐字翻译】

按照约定，单方法接口通过方法名加上 `-er` 后缀或类似的修改来命名，以构建施动者名词：`Reader`、`Writer`、`Formatter`、`CloseNotifier` 等。

有许多这样的名称，尊重它们以及它们所涵盖的函数名称是值得遵循的。`Read`、`Write`、`Close`、`Flush`、`String` 等等都有规范的签名和含义。为避免混淆，除非你的方法具有相同的签名和含义，否则不要给你的方法赋予这些名称之一。反过来，如果你的类型实现了一个与已知类型上的方法具有相同含义的方法，请赋予它相同的名称和签名；将你的字符串转换方法称为 `String`，而不是 `ToString`。

#### 【专业术语与句式拆解】

1. **Agent noun（施动者名词）**：
   - **含义**：表示“执行动作的人或工具”的名词（如执行 `Read` 的动作体叫 `Reader`）。
2. **Canonical signatures（规范签名）**：
   - **含义**：指社区内约定俗成的标准方法签名（如 `Read(p []byte) (n int, err error)` 或 `String() string`）。
3. **`ToString` vs `String`**：
   - **含义**：Java/C# 等语言习惯叫 `ToString()`，但在 Go 中只要实现 `fmt.Stringer` 接口，就**必须**叫 `String()`，避免把其他语言的命名习惯带入 Go。

#### 【4. 补充：Go 1.27 演变与现代 Go 实践】

1. **接口定义的小而美（Interface Segregation）：**
   - Go 官方倡导“接口越小越好”，`-er` 命名规范是 Go 隐式接口（duck typing）的精髓。
2. **泛型时代的接口命名（Go 1.18+）：**
   - 在引入泛型（Generics）后，接口除用于“行为抽象”外，还新增了“类型集定义（Type Constraints）”的作用。
   - 对于**类型约束接口**，社区通常不再使用 `-er` 命名，而是采用**名词或形容词**，例如标准库中的 `cmp.Ordered`（用于限制可比较大小的类型）。

---

### MixedCaps（驼峰命名规范）

#### 【英文原文】

> Finally, the convention in Go is to use `MixedCaps` or `mixedCaps` rather than underscores to write multi-word names.

#### 【精准逐字翻译】

最后，Go 中的约定是使用 `MixedCaps`（大驼峰）或 `mixedCaps`（小驼峰），而不是下划线来编写多单词名称。

#### 【专业术语与句式拆解】

1. **`MixedCaps` vs `mixedCaps`**：
   - **含义**：分别对应 **CamelCase**（大驼峰/PascalCase，如 `ParseURL`，首字母大写用于导出符号）和 **lowerCamelCase**（小驼峰，如 `parseURL`，首字母小写用于未导出符号）。
2. **Underscores（下划线）**：
   - **含义**：像 `snake_case`（蛇形命名，常用于 Python/C/Rust/SQL）在 Go 的变量、函数、结构体字段中是**强烈反对**使用的。下划线在 Go 中几乎只保留给文件名（如 `user_service_test.go`）或特定测试/生成代码。

#### 【4. 补充：Go 1.27 演变与现代 Go 实践】

1. **缩写词（Acronyms/Initialisms）全大写/全小写规则：**
   - 原文未细讲缩写词，但在 Go 社区（及 `golangci-lint` 中的 `revive` / `golint` 规则）中有非常硬性的规定：**连续的缩写单词必须保持大小写一致**。
   - *正确：* `URL`、`HTTP`、`API`、`ID` $\rightarrow$ `parseURL`、`FetchHTTP`、`userID`、`xmlHTTPRequest`。
   - *错误：* `parseUrl`、`FetchHttp`、`userId`、`XmlHttpRequest`。
2. **文件名例外：**
   - 尽管代码标识符绝对使用驼峰命名，但 Go 源文件（`.go`）文件名统一推荐使用 **蛇形命名（`snake_case`）**（例如 `user_repository.go`），且只包含小写字母、下划线和数字，便于在不区分大小写的文件系统（Windows/macOS）上保持跨平台兼容。

---