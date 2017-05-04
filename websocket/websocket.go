package websocket

import (
	"container/list"
)

// Message 消息发送格式
type Message struct {
	Code    int
	Message string
}

var (
	GHub       *Hub
	ClientList = list.New()
)

func init() {
	GHub = NewHub()
}

// Run 启动服务
func Run() {
	GHub.run()
}
