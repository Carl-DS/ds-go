package main

import (
	"fmt"
)

func hello() []string {
	return nil
}

func GetValue() int {
	return 1
}

func main() {
	h := hello()
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	i:=GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}