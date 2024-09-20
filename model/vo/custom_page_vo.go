package vo

// CustomPageVO
//
// # 自定义分页 VO
//
// 用于接收自定义分页数据。
//
// # 请求
//   - Page: int64 分页页码
//   - Limit: int64 分页大小
//   - Search: string 搜索关键字
type CustomPageVO struct {
	Page   int    `json:"page" form:"page" binding:"required" example:"1"`
	Limit  int    `json:"limit" form:"limit" binding:"required" example:"20"`
	Search string `json:"search" form:"search" example:""`
}
