package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`
}

// 输出　{}
// 知识点：结构体访问控制，因为name首字母是小写，导致其他包不能访问，所以输出为空结构体
func main() {
	js := `{
		"name":"seekload"
	}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}
