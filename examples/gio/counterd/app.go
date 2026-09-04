package main

import (
	"fmt"
	"time"

	"gioui.org/app"
)

type AppEngine struct {
	childCount int
}

func NewAppEngine() *AppEngine {
	return &AppEngine{}
}

// CreateMainPage 组装主页面及其业务逻辑
func (e *AppEngine) CreateMainPage() *MainPage {
	page := NewMainPage()

	// 绑定页面发起的业务请求
	page.OnRequestSpawnWindow = func() {
		e.childCount++

		var subComponent Component
		if e.childCount%2 == 1 {
			subComponent = NewCounter(e.childCount)
		} else {
			subComponent = NewProgressBar(5 * time.Second)
		}

		// 创建并启动子窗口
		childNode := NewWindowNode(
			fmt.Sprintf("Child Window #%d", e.childCount),
			300, 200,
			subComponent,
		)
		go ServeWindowNode(childNode)
	}

	return page
}

func (e *AppEngine) Run() {
	// 启动主窗口
	mainPage := e.CreateMainPage()
	rootNode := NewWindowNode("Main Window (Root)", 450, 350, mainPage)

	go func() {
		_ = ServeWindowNode(rootNode)
	}()

	app.Main()
}
