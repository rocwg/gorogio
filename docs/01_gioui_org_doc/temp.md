```go
我有个疑问，如下两个写法，我运行后效果一致。请您指导
（1）ops.Reset() 是不是必须的
（2）gtx := app.NewContext(&ops, e) 是否必须的
（3）gtx.Ops 和 &ops 的疑惑

var window app.Window
window.Option(app.Title(title))

var ops op.Ops
for {
    switch e := window.Event().(type) {
    case app.DestroyEvent:
        return e.Err
    case app.FrameEvent:
        ops.Reset()
        gtx := app.NewContext(&ops, e)

        drawProgressBar(gtx.Ops, e.Source, time.Now())

        e.Frame(gtx.Ops)
    }
}


var window app.Window
window.Option(app.Title(title))

var ops op.Ops
for {
    switch e := window.Event().(type) {
    case app.DestroyEvent:
        return e.Err
    case app.FrameEvent:
        ops.Reset()

        drawProgressBar(&ops, e.Source, time.Now())

        e.Frame(&ops)
    }
}
```

