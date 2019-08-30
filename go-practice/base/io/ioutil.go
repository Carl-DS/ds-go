package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "E:\\GoProjects\\go-practice\\io\\data"
	// 字符串写入文件
	data1 := []byte("hello byte")
	err := ioutil.WriteFile(filename, data1, 0644)
	checkErr(err)

	// 为了实现细颗粒度的写入，打开文件后再写入
	file, e := os.Create(filename)
	checkErr(e)

	// 在打开文件后通常应该立刻使用defer来调用打开文件的close 方法，
	// 以保证main函数结束后，文件关闭
	defer file.Close()

	// 写入字节切片
	data2 := []byte{115, 111, 109, 101, 10}
	n, err2 := file.Write(data2)
	checkErr(err2)
	fmt.Printf("wrote %d bytes\n", n)

	// 也可以使用`WriteString`直接写入字符串
	n3, err3 := file.WriteString("writes\n")
	checkErr(err3)
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用Sync方法来将缓冲区数据写入磁盘
	file.Sync()

	// bufio 除了提供上面的缓冲读取数据外，还提供了缓冲写入数据的方法
	w := bufio.NewWriter(file)
	n4, err4 := w.WriteString("buffered\n")
	checkErr(err4)
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用Flush方法确保所有缓冲区的数据写入底层writer
	w.Flush()
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
