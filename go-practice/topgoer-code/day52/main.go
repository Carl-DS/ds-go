package main

import (
	"fmt"
)

// 下面代码有什么不规范的地方吗？
// 检查map是否含有某一元素，直接判断元素的值并不是一种合适的方式。最可靠的操作是使用访问map时返回的第二个值
func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}
	_, ok := x["two"]
	if !ok {
		fmt.Println("no entry")
	}
}
