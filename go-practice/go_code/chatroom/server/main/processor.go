package main

import (
	"fmt"
	"net"
	"io"
	"go-practice/go_code/chatroom/common/message"
	"go-practice/go_code/chatroom/server/process"
	"go-practice/go_code/chatroom/server/utils"
)

// 先创建一个Processer的结构体
type Processer struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes函数
// 功能: 根据客户端发送消息各类的不同, 决定调用哪个函数来处理
func (this *Processer) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
	default:
		fmt.Println("消息类型不存在,无法处理...")
	}
	return
}

func (this *Processer) process2() (err error){

	// 循环读客户端发送的信息
	for {
		// 这里我们将读取数据包,直接封装成一个函数readPkg(), 返回Message, Err
		// 创建一个Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务器端也退出...")
				return err
			} else {
				fmt.Println("readPkg error=", err)
			}
			
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}