package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var jsonMap = map[string]string{}

func readJson(filename string) (map[string]string, error) {
	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(e.Error())
	}

	if err := json.Unmarshal(bytes, &jsonMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return jsonMap, nil
}

func main() {
	jsonMap, err := readJson("E:\\GoProjects\\go-practice\\json\\zh-CN.json")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(jsonMap)
}
