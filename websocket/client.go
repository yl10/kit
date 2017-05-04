package websocket

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

// Client 客户端 struct
type Client struct {
	// 客户端ID
	Id string
	// SessionID 扩展用
	SesnId string
	// 名称或者描述
	Name string
	// github.com/gorilla/websocket
	Ws *websocket.Conn

	// 需要发送的消息
	SendMsg chan []byte
}

// NewClient 新建客户端
// user 与客户端struct 字段对应的 string slice
// 0 - id, 1 - session id, 2 - name
func NewClient(user []string, w http.ResponseWriter, r *http.Request) (*Client, error) {

	cli := new(Client)
	cli.Id = user[0]
	cli.SesnId = user[1]
	cli.Name = user[2]
	cli.SendMsg = make(chan []byte)

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	cli.Ws = ws

	return cli, nil
}

// Run 监听并建立连接
func (c *Client) Run() {

	// 第一次运行注册到 hub
	GHub.Register <- c

	// 持续监听事件
	go func() {
		for {
			select {
			case msg := <-c.SendMsg:
				err := c.Ws.WriteJSON(msg)
				if err != nil {
					// 从 hub 注销掉
					GHub.Unregister <- NewEvent(EVENT_HUB_UNREGISTE, nil, c.Id)
				}
			}
		}
	}()
}

// WriteJson 手动发送消息
func (c *Client) WriteJson(msg Message) {
	send, _ := json.Marshal(msg)
	c.Ws.WriteJSON(send)
}
