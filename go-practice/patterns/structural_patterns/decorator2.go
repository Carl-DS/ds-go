package main

import (
	"fmt"
	"github.com/lexkong/log"
	"net/http"
)

func autoAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil || cookie.Value != "Authentic" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!"+r.URL.Path)
}

/**
装饰器模式:
装饰器模式：允许向一个现有的对象添加新的功能，同时又不改变其结构。这种类型的设计模式属于结构型模式，
它是作为现有的类的一个包装。这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，
提供了额外的功能。我们使用最为频繁的场景就是http请求的处理：对http请求做cookie校验。
*/
func main() {
	http.HandleFunc("/hello", autoAuth(hello))
	err := http.ListenAndServe(":5566", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
