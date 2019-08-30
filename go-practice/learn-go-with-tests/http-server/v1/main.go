package v1

import (
	"log"
	"net/http"
)

func main() {
	// HandlerFunc 类型是一个允许将普通函数用作 Http Handler 的适配器。
	// 如果 f 是具有适当签名的函数，则 HandlerFunc 是一个调用 f 的 Handler
	handler := http.HandlerFunc(PlayerServer)
	// ListenAndServe 会在 Handler 上监听一个端口。如果端口已被占用，它会返回一个 error，所以我们在一个 if 语句中捕获出错的场景并记录下来。
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
