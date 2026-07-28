当然，下面给你一份可以直接放进 `07-files-network.md` 的 Markdown 模板，覆盖 **Reading Files / Writing Files / Line Filters / File Paths / Directories / Temporary Files and Directories / Embed Directive / HTTP Client / HTTP Server / TCP Server / Context / Spawning Processes / Exec'ing Processes / Signals / Exit**。



# Go 文件、网络与系统笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者到初级进阶  
> 目标：掌握文件操作、网络通信和基本系统编程能力

---

## 1. Reading Files

1.1 一句话理解：读取文件就是把磁盘内容加载到程序里。

1.2 示例代码

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data))
}
```

| 1.3 重点                  | 1.4 常见坑                   |
| ------------------------- | ---------------------------- |
| `os.ReadFile` 很常用。    | 路径不对。                   |
| 读取结果通常是 `[]byte`。 | 忘记处理错误。               |
| 需要时再转字符串。        | 二进制文件误当文本文件处理。 |

1.5 我的理解

读文件是最基础的 I/O 操作之一。

---

## 2. Writing Files

2.1 一句话理解：写文件就是把程序里的内容保存到磁盘。

2.2 示例代码

```go
package main

import (
    "os"
)

func main() {
    err := os.WriteFile("test.txt", []byte("hello\n"), 0644)
    if err != nil {
        panic(err)
    }
}
```

| 2.3 重点                  | 2.4 常见坑         |
| ------------------------- | ------------------ |
| `os.WriteFile` 简单直接。 | 误以为是追加写入。 |
| 第三个参数是权限位。      | 权限参数不理解。   |
| 会覆盖原有内容。          | 没处理写入错误。   |

2.5 我的理解

写文件是保存状态和导出数据的基础。

---

## 3. Line Filters

3.1 一句话理解：行过滤器按行读取并处理文本内容。

2.2 示例代码

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
```

| 3.3 重点                         | 3.4 常见坑                        |
| -------------------------------- | --------------------------------- |
| 适合逐行处理文本。               | 不知道 `Scanner` 默认有长度限制。 |
| 常用于日志、命令管道、文本清洗。 | 忘记处理 `Scan` 结束后的错误。    |

3.5 我的理解

行过滤器是文本处理里很实用的模式。

---

## 4. File Paths

4.1 一句话理解：path 包用于处理跨平台文件路径。

2.2 示例代码

```go
package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    p := filepath.Join("dir1", "dir2", "file.txt")
    fmt.Println(p)
}
```

| 4.3 重点                  | 4.4 常见坑                  |
| ------------------------- | --------------------------- |
| `filepath` 适合本地路径。 | 混淆 `path` 和 `filepath`。 |
| 不要手写 `/` 拼路径。     | 路径拼接方式不跨平台。      |

4.5 我的理解

路径处理要交给标准库，不要自己拼字符串。

---

## 5. Directories

5.1 一句话理解：目录操作用于创建、读取和遍历文件夹。

2.2 示例代码

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    entries, _ := os.ReadDir(".")
    for _, e := range entries {
        fmt.Println(e.Name())
    }
}
```

| 5.3 重点                             | 5.4 常见坑           |
| ------------------------------------ | -------------------- |
| 目录和文件是系统操作的一部分。       | 不区分目录项和文件。 |
| 常用于扫描资源、批量处理、构建工具。 | 遍历时忽略错误。     |

5.5 我的理解

目录操作是文件系统编程的基础。

---

## 6. Temporary Files and Directories

6.1 一句话理解：临时文件和目录用于短期存放数据。

2.2 示例代码

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, _ := os.CreateTemp("", "sample-*.txt")
    defer os.Remove(f.Name())
    fmt.Println(f.Name())
}
```

| 6.3 重点               | 6.4 常见坑         |
| ---------------------- | ------------------ |
| 临时资源适合中间产物。 | 忘记删除临时文件。 |
| 用完要清理。           | 临时路径权限问题。 |

6.5 我的理解

临时文件适合测试、缓存和中间处理。

---

## 7. Embed Directive

7.1 一句话理解：embed 可以把静态资源编译进程序里。

2.2 示例代码

```go
package main

import (
    "embed"
    "fmt"
)

//go:embed hello.txt
var content embed.FS

func main() {
    data, _ := content.ReadFile("hello.txt")
    fmt.Println(string(data))
}
```

| 7.3 重点                         | 7.4 常见坑                     |
| -------------------------------- | ------------------------------ |
| 适合内置模板、静态文件、配置。   | 忘记导入 `embed`。             |
| 资源会进二进制，不依赖外部文件。 | 路径写错。                     |
|                                  | 以为运行时还能随便改嵌入资源。 |

7.5 我的理解

embed 很适合做单文件分发或静态资源打包。

---

## 8. HTTP Client

8.1 一句话理解：HTTP Client 用于发送网络请求。

2.2 示例代码

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    resp, err := http.Get("https://example.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

| 8.3 重点                            | 8.4 常见坑       |
| ----------------------------------- | ---------------- |
| 请求完成后要关闭 body。             | 忘记关闭响应体。 |
| `http.Client` / `http.Get` 很常用。 | 不设置超时。     |
| 处理状态码和超时很重要。            | 忽略状态码。     |

8.5 我的理解

HTTP 客户端是 Go 后端开发必学内容。

---

## 9. HTTP Server

9.1 一句话理解：HTTP Server 用来接收请求并返回响应。

2.2 示例代码

```go
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}
```

| 9.3 重点                        | 9.4 常见坑                    |
| ------------------------------- | ----------------------------- |
| `HandleFunc` 注册路由处理函数。 | 没处理启动错误。              |
| `ListenAndServe` 启动服务。     | 路由设计混乱。                |
| 是 Web 服务的入口基础。         | 写 handler 时忽略请求上下文。 |

9.5 我的理解

HTTP 服务是 Go 非常常见的应用场景。

---

## 10. TCP Server

10.1 一句话理解：TCP server 直接处理底层网络连接。

2.2 示例代码

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go func(c net.Conn) {
            fmt.Fprintln(c, "hello tcp")
            c.Close()
        }(conn)
    }
}
```

| 10.3 重点                        | 10.4 常见坑      |
| -------------------------------- | ---------------- |
| 更底层，更灵活。                 | 连接关闭不及时。 |
| 适合自定义协议或高性能网络程序。 | 并发处理不当。   |
|                                  | 错误处理不完整。 |

10.5 我的理解

TCP server 比 HTTP 更底层，也更接近网络编程本质。

---

## 11. Context

11.1 一句话理解：context 用于控制取消、超时和请求范围内的数据传递。

2.2 示例代码

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel()

    select {
    case <-time.After(100 * time.Millisecond):
        fmt.Println("done")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
}
```

| 11.3 重点                       | 11.4 常见坑                 |
| ------------------------------- | --------------------------- |
| 适合超时、取消、链路传递。      | 把 context 当普通参数仓库。 |
| 常用于 HTTP、数据库、并发任务。 | 不调用 cancel。             |
|                                 | 乱传值。                    |

11.5 我的理解

context 是 Go 服务端开发的“控制总线”。

---

## 12. Spawning Processes

12.1 一句话理解：可以从 Go 程序中启动外部进程。

2.2 示例代码

```go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    out, _ := exec.Command("date").Output()
    fmt.Println(string(out))
}
```

| 12.3 重点            | 12.4 常见坑        |
| -------------------- | ------------------ |
| 适合调用系统命令。   | 命令跨平台差异。   |
| 也可用于自动化工具。 | 没处理输出和错误。 |
|                      | 安全性问题。       |

12.5 我的理解

启动进程适合工具类程序和运维脚本。

---

## 13. Exec'ing Processes

13.1 一句话理解：exec 更强调替换当前进程或更细粒度控制命令执行。

2.2 示例代码

```go
package main

import (
    "os"
    "os/exec"
)

func main() {
    cmd := exec.Command("ls", "-a")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
```

| 13.3 重点                      | 13.4 常见坑                                       |
| ------------------------------ | ------------------------------------------------- |
| 可以控制 stdin/stdout/stderr。 | 命令输出没重定向。                                |
| 适合需要流式处理的场景。       | 不理解 `Run`、`Output`、`CombinedOutput` 的区别。 |

13.5 我的理解

exec 让 Go 程序能很好地和系统命令协作。

---

## 14. Signals

14.1 一句话理解：signals 用于接收系统发来的中断或终止信号。

2.2 示例代码

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    s := <-sigs
    fmt.Println(s)
}
```

| 14.3 重点                     | 14.4 常见坑          |
| ----------------------------- | -------------------- |
| 常用于优雅退出。              | 不监听必要信号。     |
| 处理 Ctrl+C、服务停止等场景。 | 收到信号后不做清理。 |

14.5 我的理解

信号处理是服务稳定运行的重要部分。

---

## 15. Exit

15.1 一句话理解：exit 用于直接结束程序并返回状态码。

2.2 示例代码

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("bye")
    os.Exit(3)
}
```

| 15.3 重点                  | 15.4 常见坑                              |
| -------------------------- | ---------------------------------------- |
| `os.Exit` 会立即结束程序。 | 不知道 `defer` 不会在 `os.Exit` 后执行。 |
| 状态码用于表示退出原因。   | 在 goroutine 里随意 `Exit`。             |

15.5 我的理解

exit 是最后手段，要谨慎使用。

---

## 总结

### 建议记忆顺序
1. 文件读写
2. 路径与目录
3. 临时文件
4. embed
5. HTTP client / server
6. TCP server
7. context
8. 进程、信号、退出

### 你要重点记住的三件事
- 文件和路径操作优先用标准库。
- HTTP 是 Go 最常见的网络入口。
- context、signal、exit 都和程序生命周期有关。

## 学习建议

这一组内容非常适合“学一点就动手做一点”。
比如你可以自己试着做：

- 一个读写配置文件的小工具。
- 一个简单 HTTP 服务。
- 一个带超时的请求客户端。
- 一个优雅退出的小程序。

这样会比纯看例子更牢。

如果你愿意，我下一步可以继续帮你写 **`08-testing-tools.md`**。