package main

import (
	"encoding/json"
	"fmt"
)

type Goods struct {
	Name    string
	Price   int
	Address string `json:"address2"`
	Tag     string
}

func main() {
	goods := Goods{"商品1", 100, "中国", "特价"}
	// 结构休 => json
	bytes, err := json.Marshal(goods)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", bytes)

	// json => 结构体
	jsonStr := `{"Name":"商品1","Price":100,"address2":"中国","Tag":"特价"}`

	goods2 := Goods{}
	err2 := json.Unmarshal([]byte(jsonStr), &goods2)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("%v", goods2)
}
