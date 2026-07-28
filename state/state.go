package state

// 目标：封装 gorogio 实验场的状态管理
// 说明：
//   - 每个 UI 案例可能有自己的状态（计数器、开关、输入文本等）
//   - 可在这里统一管理共享状态或提供状态工具函数
//
// TODO: 定义 State 结构，实现初始化、更新、重置方法

// AppState TODO: 在此封装实验场通用状态
type AppState struct {
	// TODO: 添加字段，例如计数器值、开关状态等

	Nav        string
	Query      string
	Resources  []Resource
	DialogOpen bool
	EditItem   *Resource
}

func NewAppState() *AppState {
	return &AppState{
		Nav:   "resource",
		Query: "",
		Resources: []Resource{
			{ID: 1, Name: "Article API", Type: "API", Status: "Online"},
			{ID: 2, Name: "Image CDN", Type: "Storage", Status: "Offline"},
			{ID: 3, Name: "User Center", Type: "Service", Status: "Online"},
		},
	}
}
