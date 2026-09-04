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

	// 1. 创建主 Counter
	mainCounter := NewCounter(10,
		WithStep(5),
		WithRange(0, 50),
		WithButtonColor(color.NRGBA{R: 46, G: 139, B: 87, A: 255}),
	)

	// 2. 创建主窗口节点
	rootNode := NewWindowNode("Main Window (Root)", 450, 350, mainCounter)

	// 3. 注册模式 B 弹出行为：点击按钮触发 SpawnChild 生成独立子 OS 窗口
	mainCounter.OnSpawnWindow = func() {
		childCount++
		// 创建子窗口内的 Counter 组件（模式 C）
		subCounter := NewCounter(0,
			WithStep(1),
			WithButtonColor(color.NRGBA{R: 220, G: 80, B: 50, A: 255}),
		)

		// 封装为子窗口节点并挂载至根节点（模式 B）
		childNode := NewWindowNode(
			fmt.Sprintf("Child Window #%d", childCount),
			300, 200,
			subCounter,
		)

		SpawnChild(rootNode, childNode)
	}

	// 4. 运行主窗口
	go func() {
		if err := ServeWindowNode(rootNode); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
