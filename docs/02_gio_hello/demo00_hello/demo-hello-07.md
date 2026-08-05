```powershell
我完成改造后，当下代码如下：
（1）代码 tree
PS D:\roc-github\gio-workspace\gorogio> goro tree -depth=2
.
├── app
│   └── app.go
├── component
│   └── button.go
├── container
│   ├── column.go
│   ├── expand.go
│   ├── flex.go
│   ├── options.go
│   ├── row.go
│   └── stack.go
├── element
│   └── element.go
├── examples
│   ├── gio
│   └── gorogio
├── go.mod
├── go.sum
├── modifier
│   └── padding.go
├── page
│   └── page.go
├── primitive
│   ├── space.go
│   └── text.go
├── state
│   └── state.go
└── style
    └── theme.go
    
（2）gorogio 代码如下
（3）examples\gorogio\hello 代码如下:
```



```ini
+---------------------+
|                     |
|     Hello, Gio      |
|                     |
|     Count: 3        |
|                     |
|   [+]   [Reset]     |
|                     |
+---------------------+

比如下一步，我是不是可以 做 窗口 布局，比如
```



