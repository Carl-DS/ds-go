package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("Type=%T, val=%v\n", now, now)

	// 2. 获取年月日时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 3. 格式化日期 2006-01-02 15:04:05 参考格式
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	// Unix 和 UnixNano
	fmt.Printf("unix 时间戳: %v, nuixNano 时间戳: %v\n", now.Unix(), now.UnixNano())
}
