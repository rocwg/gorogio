package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	var childCount int

	// 1. 实例化业务组件 (Component)
	mainCounter := NewCounter(10,
		WithStep(5),
		WithRange(0, 50),
		WithButtonColor(color.NRGBA{R: 46, G: 139, B: 87, A: 255}),
	)

	// 2. 组装主窗口节点 (WindowNode)
	rootNode := NewWindowNode("Main Window (Root)", 450, 350, mainCounter)

	// 3. 业务回调：点击主窗口按钮，弹出子窗口
	mainCounter.OnSpawnWindow = func() {
		childCount++

		// 我们可以让子窗口显示 Counter，也可以让它显示 ProgressBar！
		var subComponent Component
		if childCount%2 == 1 {
			subComponent = NewCounter(0,
				WithStep(1),
				WithButtonColor(color.NRGBA{R: 220, G: 80, B: 50, A: 255}),
			)
		} else {
			subComponent = NewProgressBar(5 * 1000 * 1000 * 1000) // 5s 动画
		}

		// 创建平行 OS 子窗口 (Layer 2)
		childNode := NewWindowNode(
			fmt.Sprintf("Child Window #%d", childCount),
			300, 200,
			subComponent,
		)

		SpawnChild(rootNode, childNode)
	}

	// 4. 启动应用
	go func() {
		if err := ServeWindowNode(rootNode); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
