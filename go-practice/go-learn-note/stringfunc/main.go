package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度,按字节 len(str)
	// golang 的统一编码为utf-8(ascii的字符(字母和数字)占一个字节, 汉字占用3个字节)
	str := "hello北"
	fmt.Println("str len: ", len(str)) // 8

	str2 := "hello北京"
	// 字符串遍历,同时处理有中文的问题 r := []rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误: ", err)
	} else {
		fmt.Println("转换的结果是: ", n)
	}
	// 整数转字符串
	str3 := strconv.Itoa(1234)
	fmt.Printf("str=%v, str=%T \n", str3, str3)
	// 字符串转 []byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)
	// []byte 转字符串
	byteStr := string([]byte{97, 98, 99})
	fmt.Printf("byteStr=%v\n", byteStr)

	// 统计一个字符串有几个指定的子串
	num := strings.Count("hello eee", "e")
	fmt.Printf("共有子串: %v\n", num)
	// 不区分大小写的字符串比较(==是区分字母大小写的)
	fmt.Println(strings.EqualFold("abc", "ABC"))
}
