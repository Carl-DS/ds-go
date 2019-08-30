package main

import (
	"log"
	"os"
)

// 译者按： rename 和 move 原理一样
func main() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
