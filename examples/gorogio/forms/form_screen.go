package main

import (
	"fmt"

	"gioui.org/layout"
	"github.com/rocwg/gorogio/app"

	"github.com/rocwg/gorogio/component"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/primitive"
	"github.com/rocwg/gorogio/style"
)

// 编译期检查
var _ app.Scene = (*FormScreen)(nil)

// FormScreen
//
// 一个完整页面。
// 负责组合 Component。
type FormScreen struct {
	UserName *component.FormField
	Age      *component.FormField
	City     *component.FormField

	Clear  *component.Button
	Submit *component.Button
}

// NewFormScreen 创建
func NewFormScreen() *FormScreen {
	screen := &FormScreen{
		UserName: component.NewFormField("姓名：", "请输入姓名"),
		Age:      component.NewFormField("年龄：", "请输入年龄"),
		City:     component.NewFormField("城市：", "请输入城市"),
	}
	// Button 行为绑定
	screen.Clear =
		component.NewButton("清空").
			OnClick(
				func() {
					screen.UserName.SetValue("")
					screen.Age.SetValue("")
					screen.City.SetValue("")
				},
			)
	screen.Submit =
		component.NewButton("提交").
			OnClick(
				func() {
					fmt.Println("姓名：", screen.UserName.Value())
					fmt.Println("年龄：", screen.Age.Value())
					fmt.Println("城市：", screen.City.Value())
				},
			)
	return screen
}

// Update
//
// Screen 调度 Component 更新。
func (f *FormScreen) Update(
	gtx layout.Context,
) {
	f.UserName.Update(gtx)
	f.Age.Update(gtx)
	f.City.Update(gtx)

	f.Clear.Update(gtx)
	f.Submit.Update(gtx)
}

// Element
//
// Screen 只负责组合。
func (f *FormScreen) Element(
	th *style.Theme,
) element.Element {

	return modifier.Padding(

		40,

		container.Column(

			container.Options{
				Spacing:   th.Spacing.Large,
				Alignment: layout.Middle,
			},

			primitive.H3(
				th,
				"User Profile Form",
			),

			f.UserName.Element(th),
			f.Age.Element(th),
			f.City.Element(th),

			container.Row(
				container.Options{
					Spacing: th.Spacing.Large,
				},
				f.Clear.Element(th),
				f.Submit.Element(th),
			),
		),
	)
}

func (f *FormScreen) Name() string {
	return "form"
}
