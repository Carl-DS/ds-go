package main

import (
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

/**
使用 context 存储数据
*/
func main() {
	// 启动一个 web 服务
	//http.Handle("/", http.HandlerFunc(myHandler))
	http.Handle("/", http.HandlerFunc(myHandler3))
	http.ListenAndServe(":1234", nil)
}

//定义一个Hander
func myHandler3(rw http.ResponseWriter, r *http.Request) {
	//模拟为Request附加值，这里附加了2个
	userContext := context.WithValue(context.Background(), "user", "张三")
	ageContext := context.WithValue(userContext, "age", 18)
	rContext := r.WithContext(ageContext)

	//这个模拟一个方法或者函数的调用，大部分情况下可能不在一个包里
	doHandler3(rw, rContext)
}

func doHandler3(rw http.ResponseWriter, r *http.Request) {
	//我们从这个Request里取出对应的值。
	user := r.Context().Value("user").(string)
	age := r.Context().Value("age").(int)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))

}
