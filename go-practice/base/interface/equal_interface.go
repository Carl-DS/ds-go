package main

func main() {
	var t1, t2 interface{}
	println(t1 == nil, t1 == t2)

	t1, t2 = 100, 100
	println(t1 == t2)

	// 如果实现接口的类型支持,可做相等运算,故此处 map 无法做相等运算
	t1, t2 = map[string]int{}, map[string]int{}
	println(t1 == t2)
}
