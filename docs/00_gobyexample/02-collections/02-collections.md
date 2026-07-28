# Go 集合类型笔记

> 来源：Go by Example  
> 适用阶段：Go 初学者  
> 目标：理解 Go 中最常用的集合类型，并能区分数组、切片和 map

---

## 1. Arrays

1.1 一句话理解

数组是长度固定、元素类型相同的一组数据。

1.2 示例代码

```go
// In Go, an _array_ is a numbered sequence of elements of a
// specific length. In typical Go code, [slices](slices) are
// much more common; arrays are useful in some special
// scenarios.

package main

import "fmt"

func main() {

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array
	// in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// You can also have the compiler count the number of
	// elements for you with `...`
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// If you specify the index with `:`, the elements in
	// between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// You can create and initialize multi-dimensional
	// arrays at once too.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
```

```bash
# Note that arrays appear in the form `[v1 v2 v3 ...]`
# when printed with `fmt.Println`.
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
dcl: [1 2 3 4 5]
idx: [100 0 0 400 500]
2d:  [[0 1 2] [1 2 3]]
2d:  [[1 2 3] [1 2 3]]
```

| 1.3 重点                 | 1.4 常见坑                   |
| ------------------------ | ---------------------------- |
| 数组长度是类型的一部分。 | 误以为数组长度可以随意变化。 |
| `len` 可以查看数组长度。 | 把数组和切片混为一谈。       |
| 数组通常不如切片常用。   |                              |

1.5 我的理解

数组更像“定长容器”，在 Go 里实际项目中不如切片常见。

---

## 2. Slices

2.1 一句话理解

切片是更灵活、更常用的动态序列。

2.2 示例代码

```go
// _Slices_ are an important data type in Go, giving
// a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has
	// length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create a slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	// By default a new slice's capacity is equal to its
	// length; if we know the slice is going to grow ahead
	// of time, it's possible to pass a capacity explicitly
	// as an additional parameter to `make`.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The `slices` package contains a number of useful
	// utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```

```bash
# Note that while slices are different types than arrays,
# they are rendered similarly by `fmt.Println`.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Check out this [great blog post](https://go.dev/blog/slices-intro)
# by the Go team for more details on the design and
# implementation of slices in Go.

# Now that we've seen arrays and slices we'll look at
# Go's other key builtin data structure: maps.
```

| 2.3 重点                         | 2.4 常见坑                     |
| -------------------------------- | ------------------------------ |
| 切片比数组更常用。               | 以为切片拷贝了所有数据。       |
| `append` 用来追加元素。          | 不理解 `len` 和 `cap` 的区别。 |
| `len` 是长度，`cap` 是容量。     | 误用切片导致共享底层数组问题。 |
| 切片本身是对底层数组的引用视图。 |                                |

2.5 我的理解

切片是 Go 最常用的数据结构之一，必须熟。

---

## 3. Maps

3.1 一句话理解

map 是键值对结构，用来根据 key 快速查找 value。

3.2 示例代码

```go
// _Maps_ are Go's built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array)
// (sometimes called _hashes_ or _dicts_ in other languages).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn't exist, the
	// [zero value](https://go.dev/ref/spec#The_zero_value) of the
	// value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs from
	// a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove *all* key/value pairs from a map, use
	// the `clear` builtin.
	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The `maps` package contains a number of useful
	// utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
```

```bash
# Note that maps appear in the form `map[k:v k:v]` when
# printed with `fmt.Println`.
$ go run maps.go 
map: map[k1:7 k2:13]
v1: 7
v3: 0
len: 2
map: map[k1:7]
map: map[]
prs: false
map: map[bar:2 foo:1]
n == n2
```

| 3.3 重点                          | 3.4 常见坑                    |
| --------------------------------- | ----------------------------- |
| map 用于键值映射。                | 忘记初始化 map 就直接写入。   |
| key 的类型必须可比较。            | 不知道缺失 key 的返回值行为。 |
| 访问不存在的 key 时，会得到零值。 | 以为 map 是有序的。           |

3.5 我的理解

map 很适合做配置、索引、统计、缓存。

---

## 4. Range over built-in types

4.1 一句话理解

`range` 是 Go 中非常常用的遍历方式。

4.2 示例代码

```go
// _range_ iterates over elements in a variety of
// built-in data structures. Let's see how to
// use `range` with some of the data structures
// we've already learned.

package main

import "fmt"

func main() {

	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// blank identifier `_`. Sometimes we actually want
	// the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	// See [Strings and Runes](strings-and-runes) for more
	// details.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
```

```bash
$ go run range-over-built-in-types.go
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
```

| 4.3 重点                                      | 4.4 常见坑                                     |
| --------------------------------------------- | ---------------------------------------------- |
| `range` 遍历切片、数组、map、字符串都很常见。 | `range` 循环里拿到的是值副本，不是原元素引用。 |
| 第一个返回值通常是索引或 key。                | 在循环中使用闭包时要特别小心变量捕获问题。     |
| 第二个返回值通常是元素值。                    |                                                |

4.5 我的理解

`range` 是 Go 里非常自然的遍历写法，几乎到处都会用。

---

## 5. 三者对比

### 5.1 一句话理解
数组、切片、map 分别适合不同场景。

### 5.2 对比表

| 类型 | 特点       | 常见用途         |
| ---- | ---------- | ---------------- |
| 数组 | 长度固定   | 较少直接使用     |
| 切片 | 灵活、常用 | 存储列表数据     |
| map  | 键值对查找 | 配置、统计、索引 |

### 5.3 我的理解
如果你不确定该用什么，优先想到切片；如果需要 key->value，就用 map。

---

## 总结

### 建议记忆顺序
1. Arrays
2. Slices
3. Maps
4. Range

### 你要重点记住的三件事
- 数组定长，切片灵活。
- 切片比数组更常用。
- map 是 Go 中最常见的键值结构之一。

## 你学习时的建议

这一组内容最好不要只记定义，还要自己补几条“什么时候用”。
比如：

- 列表数据 -> 切片。
- 键值查询 -> map。
- 遍历容器 -> `range`。

## 一个很重要的提醒

刚开始学的时候，不要急着钻进“底层数组、容量增长、map 内部实现”这些内容。
先把“会用、知道区别、知道坑点”建立起来，后面再深入会更轻松。

如果你愿意，我下一步可以继续帮你写 **`03-functions-errors.md`**。