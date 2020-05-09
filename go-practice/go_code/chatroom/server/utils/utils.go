package utils

import (
	"fmt"
	"net"
	"encoding/binary"
	"encoding/json"
	"go-practice/go_code/chatroom/common/message"
)

// 这里将这些方法关联到结构中
type Transfer struct {
	// 分析它应该有哪些字段
	Conn net.Conn
	Buf [8096]byte // 这是传输时,使用的缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	// conn.Read 在conn没有被关闭的情况下,才会阻塞
	// 如果客户端关闭了conn则,就不会阻塞
	_, err = this.Conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("read pkg body error")
		return
	}

	// 把pkgLen反序列成 -> message.Message
	// 注意点 &mes
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	// 将固定长度的数据写入字节切片
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) failed", err)
		return
	}

	// 发送data消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.write(buf) failed", err)
		return
	}
	return
}
