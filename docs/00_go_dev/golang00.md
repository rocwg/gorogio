有讲究，核心不是“哪种能不能用”，而是**Go 的惯用写法更偏向第一种 `if err := ...; err != nil {}`**，因为它把“获取错误”和“检查错误”放在一起，作用域也更小，代码更紧凑。[[books.studygolang](http://books.studygolang.com/the-way-to-go_ZH_CN/05.2.html)]

## 两种写法的区别

你写的这两种，本质上都能工作：

```go
//1)
err := run(window)
if err != nil {
    log.Fatal(err)
}
//2)
if err := loop(w); err != nil {
    log.Fatal(err)
}
```

第一种是先声明一个 `err` 变量，再在下一行判断。第二种是把 `err` 直接写在 `if` 的初始化语句里，只在这个 `if` 语句内部有效。[[books.studygolang](http://books.studygolang.com/the-way-to-go_ZH_CN/05.2.html)]

## 为什么 Go 更常用第二种

第二种写法更符合 Go 的风格，主要有几个好处：

- **作用域更小**：`err` 只在这个 `if` 里有效，不会污染外层变量。[[books.studygolang](http://books.studygolang.com/the-way-to-go_ZH_CN/05.2.html)]
- **更简洁**：少一行代码，阅读时更容易把“调用”和“错误处理”看成一个整体。[[cnblogs](https://www.cnblogs.com/brynchen/p/18041827)]
- **更符合习惯**：Go 社区很常见这种 `if err := ...; err != nil` 形式。[[cnblogs](https://www.cnblogs.com/brynchen/p/18041827)]

例如：

```go
if err := loop(w); err != nil {
    log.Fatal(err)
}
```

这表示“先调用 `loop(w)`，如果出错就立刻处理”，语义很直接。[[books.studygolang](http://books.studygolang.com/the-way-to-go_ZH_CN/05.2.html)]

## 第一种什么时候也会用

第一种不是错，只是更适合这些情况：

- 你后面还要继续用这个 `err`。
- 你想先把调用结果保存下来，后面再统一判断。
- 代码逻辑比较复杂，不想把所有东西挤进 `if` 里。

比如：

```go
err := run(window)
if err != nil {
    log.Fatal(err)
}
```

如果后面还有别的地方要看 `err`，这种写法就更方便。[[books.studygolang](http://books.studygolang.com/the-way-to-go_ZH_CN/05.2.html)]

## 你的 `main` 代码怎么看

你这里两段代码的差别，**从 Go 风格上看，第二段更推荐**：

```go
go func() {
    w := new(app.Window)
    w.Option(
        app.Decorated(false),
    )
    if err := loop(w); err != nil {
        log.Fatal(err)
    }
    os.Exit(0)
}()
app.Main()
```

比：

```go
go func() {
    window := new(app.Window)
    err := run(window)
    if err != nil {
        log.Fatal(err)
    }
    os.Exit(0)
}()
app.Main()
```

更“Go 味”一点，因为错误检查离调用更近，而且 `err` 没必要提前声明。[[cnblogs](https://www.cnblogs.com/brynchen/p/18041827)]

## 额外提醒

你这里还有一个更重要的问题：`log.Fatal` 会调用 `os.Exit(1)`，而 `os.Exit(0)` 也会直接退出进程，所以要想清楚你到底是“正常退出”还是“报错退出”。另外，在 goroutine 里直接 `os.Exit` 往往会让程序立刻终止，别的 goroutine 可能还没来得及收尾。[[books.studygolang](http://books.studygolang.com/the-way-to-go_ZH_CN/05.2.html)]

如果你愿意，我可以继续帮你把这段 `app.Main()` + goroutine + `os.Exit` 的结构，按更稳妥的 Go 写法重写一版。



---



有，Go 里除了 `if err := ...; err != nil` 之外，还有一批很有代表性的**简洁、惯用**写法。它们的共同特点是：少样板代码、作用域更小、流程更直线。

## 常见简洁写法

| 说明                           |                                                              |                                                              |
| ------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `:=` 短变量声明。              | 用来在函数内同时“声明 + 赋值”，比 `var x = ...` 更常见。     | 例如：`name := "tom"`。                                      |
| `if / for / switch` 省略括号。 | Go 不像 C/Java 那样写成 `if (x > 0)`，而是直接写 `if x > 0 {}`。 | 这让控制结构更短，也更统一。                                 |
| `if` 带初始化语句。            | 除了错误处理，也常见于先创建局部变量再判断。                 | 例如：`if n := len(s); n > 0 { ... }`。                      |
| `range` 遍历。                 | 遍历切片、数组、map、字符串时很常用，写法比手动下标更简洁。  | 例如：`for i, v := range items { ... }`。                    |
| `defer` 延迟执行。             | 很适合资源释放，比如关闭文件、解锁。                         | 例如：`defer f.Close()`。<br />这使“申请资源”和“释放资源”靠得很近。 |
| `_` 空白标识符。               | 当你不需要某个返回值时，用 `_` 丢弃它。                      | 例如：`_, err := f.Read(buf)`。                              |
| 多返回值。                     | Go 很常见的风格是“返回结果 + error”，比异常更显式。          | 例如：`v, err := parse(data)`。                              |

1. 早返回（early return）。
   Go 很喜欢“先处理异常，再继续主流程”，所以经常省略多余 `else`。
   例如：

   ```go
   f, err := os.Open(name)
   if err != nil {
       return err
   }
   ```

   这种写法比层层嵌套更清晰。

2. `switch` 更轻量。
   Go 的 `switch` 通常不需要 `break`，默认只匹配一个分支，这比很多语言少很多样板。
   例如：

   ```go
   switch x {
   case 1:
       ...
   case 2:
       ...
   }
   ```

---

## 特别值得记住的几条

最有 Go 味的，通常是这三类：`:=`、`if err := ...; err != nil`、`defer`。[[cnblogs](https://www.cnblogs.com/michaelshen/p/19782916)]
它们分别解决了“少写声明”、“错误处理贴近调用点”、“资源释放不容易忘”这三个问题。[[golang.ac](https://golang.ac.cn/doc/effective_go)]

## 你可以怎么理解

Go 的简洁不是“语法花哨”，而是**把复杂性压到更少的语言规则里**。
所以你会看到它刻意减少括号、减少嵌套、减少作用域外泄，让代码更像顺序叙述。[[scribd](https://www.scribd.com/document/748828037/effective-go-zh-en)]

如果你现在是初学者，最值得优先掌握的顺序是：

1. `:=` 
2. `if err := ...; err != nil` 
3. `defer` 
4. `range` 
5. `_, err := ...` 

这些一旦熟了，读 Go 代码会顺很多。