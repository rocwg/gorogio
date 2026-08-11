package main // 关键字: package, 标识符: main

import ( // 关键字: import
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// 常量 (true, false, iota) 与 类型 (int, rune, float64, string, byte)
const (
	StatusOK int = iota // 关键字: const
	StatusErr
	MaxLimit = 100
	Enabled  = true  // 常量: true
	Disabled = false // 常量: false
	Tag      = 'G'   // 类型: rune; 修正：Rune 字面量只能是单个字符
)

// Storable
// 泛型接口与约束 (comparable, any)
type Storable interface { // 关键字: type, interface
	comparable
}

type Cache[K Storable, V any] struct { // 泛型类型定义
	mu    sync.RWMutex
	items map[K]V // 关键字: map
}

func NewCache[K Storable, V any]() *Cache[K, V] { // 关键字: func
	return &Cache[K, V]{items: make(map[K]V)} // 内置函数: make
}

func (c *Cache[K, V]) Set(k K, v V) {
	c.mu.Lock()
	defer c.mu.Unlock() // 关键字: defer
	c.items[k] = v
}

func (c *Cache[K, V]) Get(k K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.items[k]
	return v, ok
}

// Metrics
// 结构体、接口实现与复合类型 (uint64, uintptr, float32, complex128, struct)
type Metrics struct {
	Count   uint64
	Addr    uintptr
	Ratio   float32
	Complex complex128
	Data    []byte // 类型: byte
}

type Processor interface {
	Process(ctx context.Context) error // 类型: error
}

func specProbe() {
	// --- 1. 变量声明、零值 (nil) 与类型推导 ---
	// 关键字: var
	var (
		b    bool       = Enabled
		i8   int8       = 1
		i16  int16      = 2
		i32  int32      = 3
		i64  int64      = 4
		u    uint       = 5
		u8   uint8      = 6
		u16  uint16     = 7
		u32  uint32     = 8
		u64  uint64     = 9
		f32  float32    = 3.14
		f64  float64    = 6.28
		c64  complex64  = complex(f32, 1.0) // 内置函数: complex
		c128 complex128 = complex(f64, 2.0)
		e    error      = nil // 零值: nil
	)

	// 修正：使用空白标识符 _ 消费未使用的局部变量，防止编译器报错
	_, _, _, _, _, _, _, _, _, _ = b, i16, i32, i64, u, u16, u32, c64, Tag, e

	// --- 2. 内置函数操作 (append, len, cap, clear, delete, new, min, max) ---
	slice := make([]int, 0, 5)        // 内置函数: make
	slice = append(slice, 10, 20, 30) // 内置函数: append
	_ = len(slice)                    // 内置函数: len
	_ = cap(slice)                    // 内置函数: cap
	_ = min(slice[0], 15)             // 内置函数: min
	_ = max(slice[1], 25)             // 内置函数: max

	ptr := new(Metrics) // 内置函数: new
	ptr.Complex = c128
	ptr.Count = u64 + uint64(i8+int8(u8)) // 类型转换与计算

	realPart := real(c128) // 内置函数: real
	imagPart := imag(c128) // 内置函数: imag
	_ = complex(realPart, imagPart)

	cache := NewCache[string, any]()
	cache.Set("metrics", ptr)
	cache.Set("temp", "to_be_deleted")

	delete(cache.items, "temp") // 内置函数: delete
	clear(slice)                // 内置函数: clear (Go 1.21+)

	// --- 3. 控制流、通道 (chan) 与并发 ---
	ch := make(chan string, 2) // 关键字: chan
	done := make(chan struct{})
	_ = done

	go func() { // 关键字: go, 匿名函数
		defer close(ch) // 内置函数: close
		ch <- "task1"
		ch <- "task2"
	}()

	// for 循环与 select 多路复用
	for { // 关键字: for
		select { // 关键字: select, case
		case msg, ok := <-ch:
			if !ok { // 关键字: if, else
				goto Finished // 关键字: goto
			}
			println("Received:", msg) // 内置函数: println
		default: // 关键字: default
			time.Sleep(10 * time.Millisecond)
		}
	}

Finished: // 标签声明
	println("Channel processed successfully.") // 内置函数: print/println

	// --- 4. 接口断言与异常处理 ---
	defer func() {
		if r := recover(); r != nil { // 内置函数: recover
			_, _ = fmt.Fprintf(os.Stderr, "Recovered from panic: %v\n", r)
		}
	}()

	var i any = cache
	switch v := i.(type) { // 关键字: switch, break, fallthrough (结构展示)
	case *Cache[string, any]:
		_ = v                    // 修正：消费 switch 绑定的局部变量 v
		print("Type is Cache\n") // 内置函数: print
	default:
		break // 关键字: break
	}

	// 切片拷贝 (copy) 与 JSON 序列化 (标准库 encoding/json)
	src := []byte("hello")
	dst := make([]byte, len(src))
	copy(dst, src) // 内置函数: copy

	jsonBytes, _ := json.Marshal(map[string]string{"status": "ok"})
	fmt.Println(string(jsonBytes))

	// 触发 panic 测试 recover
	if false { // 关键字: false
		panic("trigger test panic") // 内置函数: panic
	}
}
