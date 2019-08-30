package main

import (
	"github.com/urfave/negroni"
	"io"
	"net/http"
)

/**
基于它开发出我们自己的中间件，并且可以集成到Negroni中。

Negroni还兼容原生的http.Handler,你完全可以把自己的http.Handler加入到Negroni的中间件链中，Negroni会自动调用他们处理我们的HTTP Request的。
*/
func main() {
	// negroni.Classic() 返回一个Negroni 实例,然后通过这个实例,我们就可以添加一些中间件了
	n := negroni.Classic()
	// 因为Negroni完全兼容http.Handler，所以我们自己对于HTTP Request的真实业务处理也可以作为Negroni的一个中间件。
	n.UseHandler(handler())
	// 通过Negroni的Run方法，就可以启动一个服务了，这个Run方法和http.ListenAndServe是等价的
	n.Run(":1234")
}

func handler() http.Handler {
	return http.HandlerFunc(myHandler)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "hello world")
}
