package main

import (
	"fmt"
	"go-practice/go_code/chatroom/client/process"
	"os"
)

// 定义两个变量,一个表示用户ID,一个表示用户密码
var userId int
var userPwd string

func main() {

	// 接收用户的选择
	var key int
	// 判断是否还继续显示菜单
	// var loop = true

	for true {
		fmt.Println("----------欢迎登录多人聊天系统---------")
		fmt.Println("\t\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t\t 2 注册用户")
		fmt.Println("\t\t\t\t 3 退出系统")
		fmt.Println("\t\t\t\t 4 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// loop = false
			// 完成登录
			// 1. 创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			// loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
			// loop = false
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}

	// if key == 1 {
	// 	// 说明用户要登录
	// 	fmt.Println("请输入用户的ID")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("请输入用户的密码")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	// 选择登录的函数,写到另一个文件
	// 	// todo 这里我们会需要重新调用
	// 	// login.Login(userId, userPwd)
	// 	// err := login.Login(userId, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败")
	// 	// } else {
	// 	// 	fmt.Println("登录成功")
	// 	// }
	// } else if key == 2 {
	// 	fmt.Println("进行用户注册的逻辑....")
	// }
}
