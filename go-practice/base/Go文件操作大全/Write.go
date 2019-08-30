package main

import (
	"log"
	"os"
)

// 写文件
// 可以使用os包写入一个打开的文件。
// 因为Go可执行包是静态链接的可执行文件，你import的每一个包都会增加你的可执行文件的大小。其它的包如io、｀ioutil｀、｀bufio｀提供了一些方法，但是它们不是必须的。
func main() {
	// 可写文件打开文件
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写字节到文件中
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
