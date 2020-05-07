package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`  // 消息类型
	Data string `json:"data"`	// 消息的内容
}

// 定义两个消息...后面需要再增加
type LoginMes struct {
	UserId int	`json:"userId"`
	UserPwd string `json:"userPwd"`
	Username string `json:"username"`
}

type LoginResMes struct {
	Code int  `json:"code"`// 返回状态码 500 表示该用户未注册 200表示登录成功
	Error string  `json:"error"`// 返回错误消息
}

type RegisterMes struct {
	//...
}

