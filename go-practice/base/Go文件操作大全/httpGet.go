package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// 通过HTTP下载文件
func main() {
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	url := "http://www.devdungeon.com/archive"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	// 将HTTP response body 中的内容写入到文件
	// Body 满足 reader 接口,因此我们可以使用 ioutil.Copy
	numBytesWritten, err := io.Copy(newFile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Download %d byte file.\n", numBytesWritten)

}
