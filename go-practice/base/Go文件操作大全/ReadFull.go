package main

import (
	"io"
	"log"
	"os"
)

// 读取正好N个字节
func main() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// file.Read()可以读取一个小文件到大的byte slice中,
	// 但是io.ReadFull()在文件的字节数小于byte slice 字节数的时候会返回错误
	byteSlice := make([]byte, 1)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
