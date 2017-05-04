# Websocket 封装
自己之前业务用到的

- 依赖包 [gorilla/websocket](https://github.com/gorilla/websocket)
- 对象及方法
1. Run 包方法, 启动服务
```go
import github.com/github.com/yl10/kit/websocket
// 启动socket
go websocket.Run()
```

2. Message 发送消息体格式
```go
type Message struct {
	Code    int
	Message string
}
```

3. Client 客户端
```go
// NewClient 新建客户端
// user 是包含固定顺序值的客户端 slice
// user[0] 客户端ID
// user[1] 客户端SessionID
// user[2] 客户端描述或者名称
func NewClient(user []string, w http.ResponseWriter, r *http.Request) (*Client, error)
// Run 启动客户端并连接
func (c *Client) Run()
// WriteJson 手动发送 json 数据
func (c *Client) WriteJson(msg Message)
```

4. Event 通讯事件
```go
const (
	EVENT_HUB_REGISTE = iota
	EVENT_HUB_SENDMSG
	EVENT_HUB_UNREGISTE
)

const (
	EVENT_CLIENT_SENDMSG = iota
	EVENT_CLIENT_LEAVE
)

// NewEvent 定义一个新事件
func NewEvent(tp int, msg []byte, id string) Event
```

5. Hub  通讯器
```go
type Hub struct {
    // 客户端注册事件
	Register   chan *Client
    // 客户端断开事件
	Unregister chan Event
    // 广播事件
	Broadcast  chan Event
}

// NewHub 初始化一个通讯器
// 一般情况不需要用到
func NewHub() *Hub
```


- Demo
1. 客户端初始化
```go
import github.com/github.com/yl10/kit/websocket

// 定义发送内容
// msg 为 []byte 格式 消息体
// clientId 为你需要发送给的对象ID
cli, err := websocket.NewClient(user, c.Ctx.ResponseWriter, c.Ctx.Request)
if err != nil {
    log.error(err)
    return
}
cli.Run()
```

2. 向客户端发送数据
```go
import github.com/github.com/yl10/kit/websocket

// 新建发送事件
send := websocket.NewEvent(websocket.EVENT_HUB_SENDMSG, msg, userId)
// 发送
websocket.GHub.Broadcast <- send
```
