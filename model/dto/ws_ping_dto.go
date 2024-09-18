package dto

// WsPingDTO
//
// # WebSocket心跳包
//
// WebSocket心跳包数据传输对象，用于接收客户端发送的心跳包数据。
//
// # 结构
//   - ClientIP: 客户端IP
//   - ClientUserAgent: 客户端UserAgent
//   - Ping: 心跳包
type WsPingDTO struct {
	ClientIP        string `json:"client_ip"`
	ClientUserAgent string `json:"client_user_agent"`
	Ping            string `json:"ping"`
}
