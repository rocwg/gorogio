## component

这里最重要。

你的：

> 有交互和复合能力的 UI 单元。

我认为非常准确。

但是增加一句：

Component 是：

> 拥有局部状态和行为，并组合 primitive/container/modifier 的可复用 UI 单元。

例如：

Button：

```
Button Component

state:
    pressed

behavior:
    click

render:
    Text
    Background
    Padding
```

------

# 三、关于 v0.3.1 Component Tree