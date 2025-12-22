// Package internal/utils.go
// 目标：封装 gorogio 实验场的通用工具函数
// 说明：
//   - 提供常用函数，如布局助手、颜色处理、事件辅助等
//   - 供所有案例共享，避免重复编写
//   - 适合 Immediate Mode UI 风格
//
// TODO: 实现 FillBackground、Clamp、MapRange 等常用工具函数
package internal

// FillBackground TODO: 填充背景颜色
func FillBackground() {
	// TODO: 使用 op.Rect 或 paint.Fill 实现
}

// Clamp TODO: 限制数值范围
func Clamp(value, min, max float32) float32 {
	// TODO: 实现返回 value 限制在 [min, max] 之间
	return value
}
