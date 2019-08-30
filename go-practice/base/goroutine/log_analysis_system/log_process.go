package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	rc          chan string
	wc          chan string
	path        string // 读取文件的路径
	influxDBDsn string // influx data source
}

func (l *LogProcess) ReadFromFile() { // 需要代码优化，只能从文件中读取数据
	// 读取模块
	line := "message"
	l.rc <- line
}

func (l *LogProcess) Process() {
	// 解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

func (l *LogProcess) writeToInfluxDB() { // 需要代码优化，只能向influxDB中写入数据
	// 写入模块
	fmt.Println(<-l.wc)
}

func main() {
	lp := LogProcess{
		rc:          make(chan string),
		wc:          make(chan string),
		path:        "/tmp/access.log",
		influxDBDsn: "",
	}

	go lp.ReadFromFile()
	go lp.Process()
	go lp.writeToInfluxDB()
	time.Sleep(time.Second)
}
