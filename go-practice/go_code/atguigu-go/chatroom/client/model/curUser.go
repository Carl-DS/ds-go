package model

import (
	"go-practice/go_code/atguigu-go/chatroom/common/message"
	"net"
)

//因为在客户端，我们很多地方会使用到curUser,我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
}
