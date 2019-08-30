package main

import (
	"fmt"
	"os"
)

// 判断文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func process(filename string, typestr string, project string) {
	filepath := "E:\\IdeaProjects\\workspace\\gs-i18n\\doc\\i18n\\"
	exist, err := PathExists(filepath)
	if err != nil {
		fmt.Println("get dir error![%v]", err)
		return
	}
	if !exist {
		fmt.Printf("no dir! [%v]\n", filepath)
		// 创建文件夹
		err := os.Mkdir(filepath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed! [%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}

	pathname := filepath + filename
	pathExists, err := PathExists(pathname)
	if err != nil {
		fmt.Println("get file error! [%v]", err)
		return
	}

	if !pathExists {
		os.Create(pathname)
	}

	inputFile, inputError := os.Create(pathname)
	defer inputFile.Close()
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	inputFile.WriteString("sss\n")
}

func main() {
	process("go-practice-cn.json", "CN", "go")
}
