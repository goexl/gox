package http

const (
	// MethodGet 获取
	MethodGet Method = "Get"
	// MethodPost 创建
	MethodPost Method = "Post"
	// MethodPut 更新
	MethodPut Method = "Put"
	// MethodDelete 删除
	MethodDelete Method = "Delete"
)

// 禁止IDE提示未使用
var (
	_ = MethodGet
	_ = MethodPost
	_ = MethodPut
	_ = MethodDelete
)
