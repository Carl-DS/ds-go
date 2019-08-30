package gorilla

import (
	"github.com/gorilla/context"
	"net/http"
	"strconv"
)

/**
使用 context 存储数据
*/
func main() {
	// 启动一个 web 服务
	//http.Handle("/", http.HandlerFunc(myHandler))
	http.Handle("/", context.ClearHandler(http.HandlerFunc(myHandler)))
	http.ListenAndServe(":1234", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// 模拟为Request 附加值,这里附加了2个
	context.Set(r, "user", "张三")
	context.Set(r, "age", 18)

	// 这里模拟一个方法或者函数的调用,大部分情况下可能不在一个包里
	doHandler(w, r)
}

func doHandler(w http.ResponseWriter, r *http.Request) {
	// 我们从这个Request 里取出对应的值
	user := context.Get(r, "user").(string)
	age := context.Get(r, "age").(int)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}

func doHandler2(w http.ResponseWriter, r *http.Request) {
	// 我们从这个Request 里取出对应的值
	allParams := context.GetAll(r)
	user := allParams["user"].(string)
	age := allParams["age"].(int)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}
