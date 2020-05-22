package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-practice/go_code/chatroom/client/utils"
	"go-practice/go_code/chatroom/common/message"
	"net"
)

type UserProcess struct {
}

// 给关联一个用户登录的方法
// 写一个函数,完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// 下一步就要开始定协议
	// fmt.Printf("userId=%d userPwd=%s\n", userId, userPwd)
	// return nil

	// 1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3. 创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4. 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 5. 把data赋给mes.Data字段
	mes.Data = string(data)

	// 6. 将mes进行序列化
	mesdata, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 7. 到这个时候, data就是我们要发送的消息
	// 7.1 先把data的长度发送给服务器
	// 先获取到data的长度=>转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(mesdata))
	var buf [4]byte
	// 将固定长度的数据写入字节切片
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) failed", err)
		return
	}

	fmt.Printf("客户端, 发送消息的长度=%d 内容是: %s\n", len(mesdata), string(mesdata))

	// 发送消息本身
	_, err = conn.Write(mesdata)
	if err != nil {
		fmt.Println("conn.Write(data) failed", err)
		return
	}

	// 休眠10秒
	// time.Sleep(time.Second*10)
	// fmt.Println("休眠了10秒...")
	// 这里还需要处理服务器端返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("tf.ReadPkg() err=", err)
		return
	}
	// 将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// fmt.Println("登录成功")
		// 1. 显示登录成功的菜单[循环显示]
		for {
			ShowMenu()
		}

	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}
