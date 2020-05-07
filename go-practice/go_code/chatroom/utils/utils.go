package utils

import (
	"fmt"
	"net"
	"encoding/binary"
	"encoding/json"
	"go-practice/go_code/chatroom/common/message"
)

func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	// 将固定长度的数据写入字节切片
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) failed", err)
		return
	}

	// 发送消息本身data
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.write(buf) failed", err)
		return
	}
	return
}

func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取对方发送的数据...")
	// conn.Read 在conn没有被关闭的情况下,才会阻塞
	// 如果对方关闭了conn则,就不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("read pkg body error")
		return
	}

	// 把pkgLen反序列成 -> message.Message
	// 注意点 &mes
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}
