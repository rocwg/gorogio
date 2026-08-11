# gorox

Personal Go workspace.

## Modules

|           | 内容：                        |                 |
|-----------|----------------------------|-----------------|
| cli       |                            |                 |
| gio       |                            |                 |
| duckdb    |                            |                 |
| grpc      |                            |                 |
| stdlib    |                            |                 |
| benchmark |                            |                 |
| scratch   | 临时代码、AI 生成代码、一次性实验。【里面很混乱】 | 不要加入 workspace。 |

## Philosophy

- one project = one module
- managed by go.work
- long-term playground for Go engineering

## Go Workspace（go work）

Workspace 本质是：“告诉 Go：这些 module 在本地是一个开发集合。”

核心文件：`go.work`.

```go
go 1.26.2

use (
	./cli
	./gio
)
```

##### 项目建立

```shell
## 第一步：创建 gorox
mkdir gorox
cd gorox
git init

## 创建 go.work（非常关键）
go work init   # 在根目录，此时会生成 go.work 文件

## 创建 module（cli）
mkdir cli
cd cli                                 # 1.创建目录
go mod init github.com/rocwg/gorox/app/cli # 2.初始化 module (使用完整路径)

## 加入 workspace
cd ..                                  # 1.回到根目录
go work use ./cli                      # 2.将 module 加入 workspace
```



###### go run

###### go build

###### go install

```powershell
> go install .\cmd\goro\
```

> 在 Go 世界里：“目录名就是 binary 名”。
> 这不是限制，而是：Go 故意设计出来的工程约束。
>
> 所以：`cmd/<binary-name>`，这个结构非常关键。

