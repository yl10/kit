package websocket

// Hub 转发器 struct
type Hub struct {
	// 注册客户端列表
	Register chan *Client
	// 断开事件
	Unregister chan Event
	// 广播事件
	Broadcast chan Event
}

// NewHub 新建转发器
// 一般情况用不到
func NewHub() *Hub {
	return &Hub{
		make(chan *Client),
		make(chan Event),
		make(chan Event),
	}
}

func (h *Hub) run() {
	for {
		select {
		case cli := <-h.Register:

			// 有新用户接入
			ClientList.PushBack(cli)

		case event := <-h.Broadcast:

			// 遍历当前所有的连接客户
			// 找到用户ID相同的所有客户端
			// 用意是如果有消息过来，则发送给当前用户打开的所有客户端

			// 如果 event id 是 session id
			// var cliId string
			// for cli := ClientList.Front(); cli != nil; cli = cli.Next() {
			// 	if event.Id == cli.Value.(Client).SesnId {
			// 		cliId = cli.Value.(Client).Id
			// 		break
			// 	}
			// }

			for cli := ClientList.Front(); cli != nil; cli = cli.Next() {
				if event.Id == cli.Value.(*Client).Id {
					cli.Value.(*Client).SendMsg <- event.Message
				}
			}

		case event := <-h.Unregister:

			// 删除用户
			for cli := ClientList.Front(); cli != nil; cli = cli.Next() {
				if event.Id == cli.Value.(*Client).Id {
					cli.Value.(Client).Ws.Close()
					ClientList.Remove(cli)
				}
			}
		}
	}
}
