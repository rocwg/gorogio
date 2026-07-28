当然，下面给你一份可以直接放进 `05-concurrency.md` 的 Markdown 模板，覆盖 **Goroutines / Channels / Channel Buffering / Channel Synchronization / Channel Directions / Select / Timeouts / Non-Blocking Channel Operations / Closing Channels / Range over Channels / Timers / Tickers / Worker Pools / WaitGroups / Rate Limiting / Atomic Counters / Mutexes / Stateful Goroutines**。



# Go 并发笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者到初级进阶  
> 目标：理解 Go 的并发模型、通道通信和常见同步工具

---

## 1. Goroutines

1.1 一句话理解：goroutine 是 Go 里轻量级的并发执行单元。

1.2 示例代码

```go
// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now. Wait for them to finish
	// (for a more robust approach, use a [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
```

```bash
# When we run this program, we see the output of the
# blocking call first, then the output of the two
# goroutines. The goroutines' output may be interleaved,
# because goroutines are being run concurrently by the
# Go runtime.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Next we'll look at a complement to goroutines in
# concurrent Go programs: channels.
```

| 1.3 重点                                       | 1.4 常见坑                              |
| ---------------------------------------------- | --------------------------------------- |
| 用 `go` 关键字启动 goroutine。                 | 忘记等待 goroutine 完成。               |
| goroutine 非常轻量。                           | 在循环里启动 goroutine 时捕获变量出错。 |
| 主程序结束太快时，goroutine 可能来不及执行完。 |                                         |

1.5 我的理解

goroutine 是 Go 并发的基础单位。

---

## 2. Channels

2.1 一句话理解：channel 是 goroutine 之间传递数据的通道。

2.2 示例代码

```go
package main

import "fmt"

func main() {
    messages := make(chan string)
    go func() { messages <- "ping" }()
    msg := <-messages
    fmt.Println(msg)
}
```

| 2.3 重点                         | 2.4 常见坑                        |
| -------------------------------- | --------------------------------- |
| channel 用来通信，不只是“传值”。 | 忘记接收导致发送阻塞。            |
| 发送和接收默认是同步阻塞的。     | 不理解无缓冲 channel 的同步语义。 |
| 通过 `make(chan T)` 创建。       |                                   |

2.5 我的理解

channel 是 Go 并发模型的核心。

---

## 3. Channel Buffering

3.1 一句话理解

带缓冲的 channel 可以暂存一定数量的数据。

3.2 示例代码

```go
package main

import "fmt"

func main() {
    messages := make(chan string, 2)
    messages <- "buffered"
    messages <- "channel"
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

| 3.3 重点                   | 3.4 常见坑                 |
| -------------------------- | -------------------------- |
| 缓冲区允许一定程度的异步。 | 以为缓冲越大越好。         |
| 缓冲大小会影响阻塞行为。   | 不了解缓冲满了之后会阻塞。 |
| 不是无限队列。             |                            |

3.5 我的理解

缓冲 channel 是“平衡速度差”的工具。

---

## 4. Channel Synchronization

4.1 一句话理解

channel 可以用来等待某个 goroutine 完成任务。

4.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}

func main() {
    done := make(chan bool, 1)
    go worker(done)
    <-done
}
```

| 4.3 重点                               | 4.4 常见坑                  |
| -------------------------------------- | --------------------------- |
| channel 不只是传数据，也能做同步信号。 | 用错 channel 容量导致死锁。 |
| 常见于“任务完成通知”。                 | 同步信号设计不清晰。        |

4.5 我的理解

channel 也可以当“通知灯”。

---

## 5. Channel Directions

5.1 一句话理解：channel 可以限定为单向发送或单向接收。

2.2 示例代码

```go
package main

import "fmt"

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
```

| 5.3 重点                    | 5.4 常见坑                |
| --------------------------- | ------------------------- |
| `chan<- T` 只能发送。       | 看不懂单向 channel 类型。 |
| `<-chan T` 只能接收。       | 以为只是语法装饰。        |
| 能提升 API 可读性和安全性。 |                           |

5.5 我的理解

单向 channel 是一种接口约束。

---

## 6. Select

6.1 一句话理解：`select` 用于同时等待多个 channel 操作。

6.2 示例代码

```go
package main

import "fmt"

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        c1 <- "one"
    }()
    go func() {
        c2 <- "two"
    }()

    select {
    case msg1 := <-c1:
        fmt.Println("received", msg1)
    case msg2 := <-c2:
        fmt.Println("received", msg2)
    }
}
```

| 6.3 重点                             | 6.4 常见坑                          |
| ------------------------------------ | ----------------------------------- |
| `select` 类似 channel 版的多路复用。 | 以为 `select` 会按顺序执行。        |
| 哪个先就绪就执行哪个。               | 不理解 `default` 会让它变成非阻塞。 |
| 可以和 `default` 搭配。              |                                     |

6.5 我的理解

`select` 是 Go 并发里非常重要的调度工具。

---

## 7. Timeouts

7.1 一句话理解：超时用于避免 goroutine 永远等待。

7.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    c := make(chan string, 1)

    go func() {
        time.Sleep(2 * time.Second)
        c <- "result"
    }()

    select {
    case res := <-c:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout")
    }
}
```

| 7.3 重点                       | 7.4 常见坑           |
| ------------------------------ | -------------------- |
| `time.After` 常用于超时控制。  | 没有超时导致卡死。   |
| 超时是并发编程的基本保护机制。 | 超时时间设置不合理。 |

7.5 我的理解

超时是并发系统必须考虑的安全阀。

---

## 8. Non-Blocking Channel Operations

8.1 一句话理解：通过 `select` + `default` 可以做非阻塞操作。

8.2 示例代码

```go
package main

import "fmt"

func main() {
    messages := make(chan string)

    select {
    case msg := <-messages:
        fmt.Println(msg)
    default:
        fmt.Println("no message")
    }
}
```

| 8.3 重点                                 | 8.4 常见坑                 |
| ---------------------------------------- | -------------------------- |
| `default` 表示当前没准备好就走备用逻辑。 | 误把非阻塞当成“异步完成”。 |
| 常用于轮询、探测状态、避免卡住。         | 过度使用导致代码难懂。     |

8.5 我的理解

非阻塞操作适合特定场景，不要滥用。

---

## 9. Closing Channels

9.1 一句话理解：关闭 channel 表示不再发送数据了。

9.2 示例代码

```go
package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    jobs <- 1
    jobs <- 2
    close(jobs)

    for j := range jobs {
        fmt.Println(j)
    }
}
```

| 9.3 重点                            | 9.4 常见坑                         |
| ----------------------------------- | ---------------------------------- |
| 关闭后不能再发送。                  | 关闭已关闭的 channel。             |
| `range` 可以读取直到 channel 关闭。 | 多个发送方同时关闭同一个 channel。 |
| 关闭通常由发送方负责。              | 不清楚何时该 close。               |

9.5 我的理解

close 是生命周期控制，不是“清空”操作。

---

## 10. Range over Channels

10.1 一句话理解：`range` 可以持续接收 channel 数据，直到它被关闭。

10.2 示例代码

```go
// In a [previous](range-over-built-in-types) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
```

```bash
$ go run range-over-channels.go
one
two

# This example also showed that it's possible to close
# a non-empty channel but still have the remaining
# values be received.
```

| 10.3 重点               | 10.4 常见坑                     |
| ----------------------- | ------------------------------- |
| 很适合消费消息流。      | 不关闭 channel 会导致循环卡住。 |
| 依赖 channel 正确关闭。 | 误以为 `range` 会自己结束。     |

10.5 我的理解

这是 channel 消费的经典写法。

---

## 11. Timers

11.1 一句话理解：timer 用于在未来某个时间点触发一次事件。

11.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    timer1 := time.NewTimer(2 * time.Second)
    <-timer1.C
    fmt.Println("Timer fired")
}
```

| 11.3 重点              | 11.4 常见坑                   |
| ---------------------- | ----------------------------- |
| 只触发一次。           | 不清楚 Timer 和 Ticker 区别。 |
| 常用于延迟执行、超时。 | 不管理资源导致泄漏。          |

11.5 我的理解

timer 是“一次性定时器”。

---

## 12. Tickers

12.1 一句话理解：ticker 会按固定间隔重复触发。

12.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        time.Sleep(1600 * time.Millisecond)
        done <- true
    }()

    for {
        select {
        case <-done:
            fmt.Println("ticker stopped")
            return
        case t := <-ticker.C:
            fmt.Println("tick at", t)
        }
    }
}
```

| 12.3 重点             | 12.4 常见坑             |
| --------------------- | ----------------------- |
| 适合周期任务。        | 忘记停止 ticker。       |
| 退出时要停止 ticker。 | 把 ticker 当 timer 用。 |

12.5 我的理解

ticker 是“周期闹钟”。

---

## 13. Worker Pools

13.1 一句话理解：worker pool 是一组工作协程处理任务队列。

13.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        time.Sleep(500 * time.Millisecond)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 5; a++ {
        fmt.Println(<-results)
    }
}
```

| 13.3 重点                    | 13.4 常见坑                    |
| ---------------------------- | ------------------------------ |
| 用固定数量 worker 控制并发。 | jobs 和 results 的收尾不清晰。 |
| 适合批量任务处理。           | worker 数量设定不合理。        |

13.5 我的理解

worker pool 是并发的“限流执行模式”。

---

## 14. WaitGroups

14.1 一句话理解：WaitGroup 用来等待一组 goroutine 完成。

14.2 示例代码

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(time.Second)
    fmt.Println("worker", id, "done")
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
}
```

| 14.3 重点                        | 14.4 常见坑                  |
| -------------------------------- | ---------------------------- |
| `Add` / `Done` / `Wait` 是核心。 | `Add` 和 `Done` 数量不匹配。 |
| 常用于等待多个任务结束。         | 把 `WaitGroup` 复制传递。    |

14.5 我的理解

WaitGroup 是 goroutine 的“计数闸门”。

---

## 15. Rate Limiting

15.1 一句话理解：限流用来控制操作发生的频率。

15.2 示例代码

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)
    for req := range requests {
        <-limiter
        fmt.Println("request", req, "processed")
    }
}
```

| 15.3 重点                        | 15.4 常见坑                |
| -------------------------------- | -------------------------- |
| 控制速率，避免资源被打爆。       | 限流粒度不清晰。           |
| 常见于接口调用、日志、任务调度。 | 误把 `Tick` 用成永久资源。 |

15.5 我的理解

限流是系统稳定性的保护手段。

---

## 16. Atomic Counters

16.1 一句话理解：原子操作用于在并发下安全地修改计数值。

16.2 示例代码

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var ops atomic.Uint64
    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            ops.Add(1)
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("ops:", ops.Load())
}
```

| 16.3 重点                  | 16.4 常见坑                  |
| -------------------------- | ---------------------------- |
| 适合简单计数场景。         | 把原子操作当成万能同步工具。 |
| 比锁更轻量，但能力更单一。 | 对复杂共享状态仍然只用原子。 |

16.5 我的理解

原子操作适合“单个数字”的并发更新。

---

## 17. Mutexes

17.1 一句话理解：互斥锁用于保护共享数据，避免并发冲突。

17.2 示例代码

```go
package main

import (
    "fmt"
    "sync"
)

type safeCounter struct {
    mu sync.Mutex
    v  map[string]int
}

func (c *safeCounter) Inc(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.v[key]++
}

func (c *safeCounter) Value(key string) int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.v[key]
}

func main() {
    c := safeCounter{v: make(map[string]int)}
    c.Inc("a")
    fmt.Println(c.Value("a"))
}
```

| 17.3 重点          | 17.4 常见坑            |
| ------------------ | ---------------------- |
| 锁保护的是临界区。 | 忘记解锁。             |
| 加锁后要记得解锁。 | 锁范围过大导致性能差。 |
| 适合复杂共享状态。 | 死锁。                 |

17.5 我的理解

锁是处理共享状态的传统工具。

---

## 18. Stateful Goroutines

18.1 一句话理解：用 goroutine 自己管理状态，避免多个协程直接争抢数据。

18.2 示例代码

```go
package main

import "fmt"

type readOp struct {
    key  int
    resp chan int
}

func main() {
    reads := make(chan readOp)
    go func() {
        state := make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            }
        }
    }()

    resp := make(chan int)
    reads <- readOp{key: 1, resp: resp}
    fmt.Println(<-resp)
}
```

| 18.3 重点                        | 18.4 常见坑                       |
| -------------------------------- | --------------------------------- |
| 这是 Go 很典型的并发设计思路。   | 结构设计不清晰。                  |
| 用“消息传递”代替“共享内存争锁”。 | 状态 goroutine 生命周期管理不足。 |

18.5 我的理解

这是 Go 并发模型很优雅的一种写法。

---

## 总结

### 建议记忆顺序
1. Goroutines
2. Channels
3. Channel Buffering
4. Channel Synchronization
5. Select
6. Timeouts
7. Closing Channels
8. WaitGroups
9. Mutexes
10. Worker Pools
11. Rate Limiting
12. Atomic Counters
13. Stateful Goroutines

### 你要重点记住的三件事
- goroutine 负责并发执行。
- channel 负责通信和同步。
- WaitGroup、Mutex、Atomic 是不同层级的同步工具。

## 学习建议

并发这一组内容，建议你一定要边看边跑。
如果只看文字，很容易觉得懂了；但真正自己写时，才会发现“等待、关闭、同步、阻塞、竞态”这些问题都很关键。

如果你愿意，我下一步可以继续帮你写 **`06-standard-library.md`**。