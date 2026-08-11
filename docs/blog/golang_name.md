下面是一份可直接保存到笔记中的 Go 命名规范速查表，覆盖项目、模块、目录、文件、包、函数、常量、变量、结构体、接口等常见场景。你可以随时翻阅，保持代码风格一致。

------

## 1. 项目名称（Git 仓库 / module 名）

- **全小写**，单词之间 **推荐用连字符 `-`**（也可合并写，不用下划线）
- 不使用空格、特殊符号或中文
- 示例：
  - `device-monitor`
  - `userauth` 或 `user-auth`
  - `enterprise-analytics`

------

## 2. 模块名称（顶层子项目、逻辑包）

- **全小写合并单词**，保持简短
- 避免下划线，除非非常必要
- 示例：
  - `auth`（认证授权）
  - `fileupload`（文件上传）
  - `remoteapi`（第三方 API 客户端）
  - `payment`、`paychannel`（支付模块）

------

## 3. 目录 & 包 名称

- **与目录同名**，全部小写、合并写

- 少用模糊词：`utils`、`helpers`；优先用具体名称：`config`、`handler`、`repo`、`service`

- 典型结构：

  ```
  ├── cmd/          // 各命令/应用入口
  ├── internal/     // 项目内部私有包
  │   ├── handler/  // HTTP/API 处理层
  │   ├── service/  // 核心业务逻辑
  │   ├── repo/     // 数据访问层
  │   └── model/    // 领域模型定义
  ├── pkg/          // 对外公用库
  └── api/          // Protobuf/OpenAPI 定义
  ```

------

## 4. 文件名称

- **全小写**，多单词用 **下划线 `_`** 连接
- 测试文件以 `_test.go` 结尾
- 生成文件以 `zz_` 或 `zz_generated` 开头，并包含 `Code generated` 注释
- 示例：
  - `user_service.go`
  - `user_service_test.go`
  - `handler_linux.go`
  - `zz_generated_mock.go`

------

## 5. 函数 & 方法 名称

- **PascalCase**（首字母大写导出）或 **camelCase**（首字母小写私有）
- 以**动词**或**动词短语**命名
- 构造函数形如 `NewTypeName`
- 示例：
  - `func CreateUser(ctx context.Context) error`
  - `func (s *orderService) calculateTotal() float64`
  - `func NewFileUploader(cfg Config) *FileUploader`

------

## 6. 常量 名称

- **PascalCase**（首字母大写）

- 枚举/分组常量可用前缀或 `iota`

- 对于全局固定值，也可用 **ALL_CAPS**（社区较少）

- 示例：

  ```go
  const (
      MaxRetries    = 5
      StatusActive  = "active"
      DefaultTimeout = 30 * time.Second
  )
  ```

------

## 7. 变量 名称

- **camelCase**（首字母小写）

- 缩写保留大写（如 `userID`、`HTTPClient`、`dbConn`）

- 临时变量可用单字母（如 `i`, `err`）

- 示例：

  ```go
  var userID int64
  configPath := "/etc/app/config.yaml"
  for i, item := range list { ... }
  ```

------

## 8. 结构体 & 接口 名称

- **PascalCase**（首字母大写导出，名词或名词短语）

- 接口：单方法接口以 `-er` 结尾（`Reader`, `Writer`），多方法接口可用 `Service`、`Manager`、`Client` 等后缀

- 示例：

  ```go
  type UserService struct { ... }
  type FileUploader struct { ... }
  
  type Reader interface {
      Read(p []byte) (n int, err error)
  }
  type UserService interface {
      Create(ctx context.Context, u *User) error
      GetByID(ctx context.Context, id int64) (*User, error)
  }
  ```

------

## 9. 缩写 & 特殊词

- 常见缩写需保留大写：
  - `ID`, `URL`, `HTTP`, `JSON`, `XML`, `SQL`
- 不要写成 `Userid`、`Url` 等

------

## 10. 导出规则 & 访问控制

- **首字母大写** → 导出（public）；**首字母小写** → 包内私有

- 示例：

  ```go
  type Config struct { // 可导出
      Port int          // 可导出
      debug bool        // 私有
  }
  ```

------

## 速查对照表

| 对象      | 格式                          | 示例                       |
| --------- | ----------------------------- | -------------------------- |
| 项目名    | lowercase + `-`               | `order-system`             |
| 模块/包名 | lowercase 合并                | `userprofile`, `remoteapi` |
| 目录名    | 同包名                        | `handler`, `service`       |
| 文件名    | lowercase + `_`               | `payment_gateway.go`       |
| 函数名    | PascalCase / camelCase        | `NewOrder`, `processOrder` |
| 常量名    | PascalCase                    | `MaxConnections`           |
| 变量名    | camelCase                     | `orderID`, `httpClient`    |
| 结构体名  | PascalCase                    | `OrderService`             |
| 接口名    | PascalCase (`-er`, `Service`) | `Reader`, `UserService`    |

------

> **Tip**：结合 [golangci-lint](https://golangci-lint.run/) 自动检测，保持团队风格统一。
