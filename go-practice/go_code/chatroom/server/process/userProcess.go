package process2

import (
	"fmt"
	"net"
	"encoding/json"
	"go-practice/go_code/chatroom/common/message"
	"go-practice/go_code/chatroom/server/utils"
)

type UserProcess struct {
	Conn net.Conn

}

// 编写一个函数serverProcessLogin函数,专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	// 1. 先从mes中取出mes.data, 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	// 1 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2 再声明一个 LoginResMes
	var loginResMes message.LoginResMes

	// 如果用户id=100, 密码=123456,认为合法,否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500 // 500 状态码,表示该用户不存在
		loginResMes.Error = "该用户不存在,请注册后再使用..."
	}

	// 3 将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	// 4 将data赋值给resMes
	resMes.Data = string(data)

	// 5 对resMes 进行序列化 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 6 发送data, 我们将其封装到writePkg函数中
	// 因为使用分层模式(mvc), 我们先创建一个Transfer实例,然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}