package websocket

// 定义转发器事件
const (
	// 连接注册
	EVENT_HUB_REGISTE = iota
	// 消息发送
	EVENT_HUB_SENDMSG
	// 断开
	EVENT_HUB_UNREGISTE
)

// 定义客户端事件
const (
	// 消息发送
	EVENT_CLIENT_SENDMSG = iota
	// 断开
	EVENT_CLIENT_LEAVE
)

// Event 事件 struct
type Event struct {
	// 类型
	Type int
	// 消息体
	Message []byte
	// id
	Id string
}

// NewEvent 新建事件
func NewEvent(tp int, msg []byte, id string) Event {
	return Event{tp, msg, id}
}
